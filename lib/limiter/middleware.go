package limiter

import (
    "github.com/kataras/iris"
    "github.com/leonharvey/go-web-boilerplate/lib/logger"
    "github.com/leonharvey/go-web-boilerplate/lib/settings"
    "fmt"
    "bytes"
)

func init() {
    fmt.Println("Loading limiter middleware")
    
    triggerMethod := []byte("POST")

    iris.UseFunc(func(c *iris.Context) {
        
        if bytes.Compare(c.Method(), triggerMethod) != 0 {
            //Is GET method
            c.Next()
            return
        }
        
        errorSwitch := func(msg string) {
            if settings.Config.DevelopmentMode {
                c.Write(msg)
            } else {
                c.EmitError(iris.StatusInternalServerError)
            }
            
            logger.New(msg + " - " + c.RemoteAddr())
        }
        
        //i, _ := St.Strikes(remoteAddr, "NOT_GET")
        remoteAddr := c.RemoteAddr()
        
        if St.IsJailed(remoteAddr) {
            errorSwitch("Too many non GET requests")
            
            return
        }
  
        St.Infraction(remoteAddr, "NOT_GET")
        
        c.Next()
    })
}