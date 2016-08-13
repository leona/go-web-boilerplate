package middleware

import (
    "github.com/kataras/iris"
    "github.com/neoh/go-web-boilerplate/views"
    "github.com/neoh/go-web-boilerplate/lib/settings"
)

func init() {
    iris.OnError(iris.StatusInternalServerError, func(c *iris.Context) {
    	p := &views.ErrorPage{
    		Path: c.Path(),
    		Msg: settings.Errors.IntervalServer,
    	}
    	c.SetContentType("text/html")
    	views.WritePageTemplate(c.GetRequestCtx(), p)
    })

    iris.OnError(iris.StatusNotFound, func(c *iris.Context) {
    	p := &views.ErrorPage{
    		Path: c.Path(),
    		Msg: settings.Errors.NotFound,
    	}
    	
    	c.SetContentType("text/html")
    	views.WritePageTemplate(c.GetRequestCtx(), p)
    })

    // emit the errors to test them
    iris.Get("/500", func(c *iris.Context) {
        c.EmitError(iris.StatusInternalServerError) // ctx.Panic()
    })

    iris.Get("/404", func(c *iris.Context) {
        c.EmitError(iris.StatusNotFound) // ctx.NotFound()
    })
}