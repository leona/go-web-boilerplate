package controllers

import (
    "github.com/kataras/iris"
    "fmt"
    "regexp"
    "time"
    //"github.com/neoh/go-web-boilerplate/views"
    "github.com/neoh/go-web-boilerplate/models"
)

var linkIdRegexPattern *regexp.Regexp

func init() {
    fmt.Println("Registered default controller")
    
    iris.Get("/", DefaultViewController)
    iris.Get("/i/:param1", LinkViewController)
    iris.Post("/store", LinkStoreController)
    
    linkIdRegexPattern, _ = regexp.Compile("^[a-zA-Z0-9_-]*$")
}

func DefaultViewController(c *iris.Context) {
    c.Write("home page")
}

func LinkStoreController(c *iris.Context) {
    var redirectLink = c.PostValue("redirect")
    var identifier = c.PostValue("identifier")

    c.SetContentType("text/html")
    
    if !linkIdRegexPattern.MatchString(identifier) {
        c.Write("Invalid identifier")
        return
    }

    err := models.LinkStore(identifier, models.Link{
        Created: uint64(time.Now().Unix()),
        Redirect: redirectLink,
        Create_ip: c.RemoteAddr(),
    })
    
    if err != nil {
        c.Write("Identifier already exists")
        return
    }
    
    c.Write("Success!")
}

func LinkViewController(c *iris.Context)  {
    var link = c.Param("param1")

    if !linkIdRegexPattern.MatchString(link) {
        c.SetContentType("text/html")
        c.Write("Invalid search pattern")
        return
    }

    value, err := models.LinkByKey(link)
    
    if err != nil {
        c.Write("No link found")
       return
    }
    
    c.Redirect(string(value.Redirect), 301)
}