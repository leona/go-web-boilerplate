package controllers

import "github.com/kataras/iris"

func init() {
    //iris.Get("/page/:test", handlerFunc)("page")
}

func handlerFunc(c *iris.Context)  {
    c.Write(c.Param("test"))
    c.Write("Hello - " + iris.Lookup("page").Name())
}