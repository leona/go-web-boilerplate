package middleware

import (
    "fmt"
    "github.com/kataras/iris"
    "github.com/leonharvey/go-web-boilerplate/lib/settings"
    //"github.com/leonharvey/go-web-boilerplate/lib/logger"
)

func init() {
    var devMode string
    staticDirectory := settings.Config.AbsolutePath + settings.Config.Static
    
    if settings.Config.DevelopmentMode {
        devMode = "true"
    } else {
        devMode = "false"
    }
    
    fmt.Println("Loading base config middleware")
    fmt.Println("DevelopmentMode is: " + devMode)
    fmt.Println("Static directory is: " + staticDirectory)

    //iris.UseFunc(logger.Middleware)
    iris.Config.Gzip = !settings.Config.DevelopmentMode
    iris.Config.Charset = "UTF-8"
    iris.Config.IsDevelopment = settings.Config.DevelopmentMode
    iris.Static(staticDirectory, "/assets", 1)
    iris.Favicon(staticDirectory + settings.Config.Favicon)
    /*
     * Template engine config
    iris.UseTemplate(html.New(html.Config{
        Layout: "layout.html",
    })).Directory("../../views", ".html")
    */
}

