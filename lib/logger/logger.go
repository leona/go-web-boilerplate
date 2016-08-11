package logger

import (
    "github.com/kataras/iris"
    "fmt"
    "time"
    "strings"
    "strconv"
    "github.com/leonharvey/go-web-boilerplate/lib/settings"
    "github.com/leonharvey/go-web-boilerplate/lib/utilities"
)

type logger struct {
	logPool []string
	lastUpdate int64
}

var Logger *logger

func init() {
	Logger = &logger{}
}

func Middleware(c *iris.Context) {
	if settings.Config.LoggerEnabled == false {
		c.Next()
		return
	}
	
	var date, status, ip, method, path string
	var latency time.Duration
	var startTime, endTime time.Time
	path = c.PathString()
	method = c.MethodString()

	startTime = time.Now()

	c.Next()

	endTime = time.Now()
	date = endTime.Format("02/01 15:04:05")
	latency = endTime.Sub(startTime)

	status = strconv.Itoa(c.Response.StatusCode())

	ip = c.RemoteAddr()

	if queryString := string(c.URI().QueryString()); len(queryString) > 0 {
		path += "?" + string(c.URI().QueryString())
	}
	
	logValue := fmt.Sprintf("%s %v %4v %s %s %s", date, status, latency, ip, method, path)
	
	New(logValue)
	
	return
}

func New(value string) {
	if settings.Config.LoggerEnabled == false {
		return
	}
	
	if settings.Config.DevelopmentMode {
		fmt.Println(value)
	} else {
		Logger.logPool = append(Logger.logPool, value)
		
		if len(Logger.logPool) > 30000 {
			//fix race condition
			utilities.AppendFile(settings.Config.AbsolutePath + "/logs.txt", strings.Join(Logger.logPool[:],"\n"))

			Logger.logPool = []string{}
		}
	}
}

func Fatal(value string) {
	panic(value)
}