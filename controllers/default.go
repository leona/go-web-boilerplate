package controllers

import (
    "github.com/kataras/iris"
    "github.com/boltdb/bolt" 
    "fmt"
    "github.com/leonharvey/go-web-boilerplate/views"
)

func init() {
    fmt.Println("Registered default controller")
    
    iris.Get("/", DefaultController)("default")
    iris.Get("/views", TemplateController)("default")
    iris.Post("/", DefaultController)("default")
}

func DefaultController(c *iris.Context)  {
    db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("short_links"))
        //err := b.Put([]byte("answer"), []byte("42"))
        
        value := b.Get([]byte("answer"))
        
        c.SetContentType("application/json")
        c.Write(string(value))
        //c.SetFlash("name", "iris")
        return nil
    })
    
    //fmt.Printf("\n", err)
}

func TemplateController(c *iris.Context) {
	p := &views.ErrorPage{
		Path: c.Path(),
	}
	views.WritePageTemplate(c.GetRequestCtx(), p)
	c.EmitError(iris.StatusInternalServerError)
}