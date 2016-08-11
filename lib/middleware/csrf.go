package middleware

import (
    "github.com/kataras/iris"
    "github.com/valyala/fasthttp"
    "bytes"
    "time"
    "github.com/leonharvey/go-web-boilerplate/lib/logger"
    "github.com/leonharvey/go-web-boilerplate/lib/settings"
    "github.com/leonharvey/go-web-boilerplate/lib/utilities"
    "fmt"
)

func init() {
    fmt.Println("Loading CSRF middleware")
    
    triggerMethod := []byte("POST")

    iris.UseFunc(func(c *iris.Context) {
        if bytes.Compare(c.Method(), triggerMethod) != 0 {
            //Is GET method
            
            if len(c.GetCookie("csrf_token")) < 64 {
            	cookie := new(fasthttp.Cookie)
            	cookie.SetKey(settings.Keys.Csrf)
            	cookie.SetValue(utilities.RandomString(64))
            	cookie.SetPath("/")
            	cookie.SetDomain(settings.Config.Host)
            	cookie.SetExpire(time.Now().Add(time.Duration(120) * time.Minute))
            	cookie.SetSecure(settings.Config.ForceSSL)
            	cookie.SetHTTPOnly(true)

            	c.SetCookie(cookie)
            }
            
            c.Next()
            return
        }
        
        originHeader := c.RequestHeader("Origin")
        refererHeader := c.RequestHeader("referer")
        
        errorSwitch := func(msg string) {
            if settings.Config.DevelopmentMode {
                c.Write(msg)
            } else {
                c.EmitError(iris.StatusInternalServerError)
            }
            
            logger.New(msg + " - " + c.RemoteAddr())
        }
        
        if (len(originHeader) > 0 && originHeader != settings.Config.Host) || 
           (len(refererHeader) > 0 && refererHeader != settings.Config.Host) {
               
            errorSwitch("Origin/referer header mismatch")
            return
        }
    
        cookieToken := c.GetCookie(settings.Keys.Csrf)
        requestToken := c.GetString(settings.Keys.Csrf)
        
        if len(cookieToken) < 64 || len(requestToken) < 64 || cookieToken != requestToken  {
            //Verify cookie is still HTTP_ONLY to stop client from overwriting
            errorSwitch("Invalid CSRF token")
            return
        }
        
        c.Next()
    })
}