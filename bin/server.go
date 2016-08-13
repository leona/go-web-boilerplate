package main

//go:generate qtc -dir=../views

import (
    "runtime"
    "fmt"
    "github.com/kataras/iris"
    "github.com/fatih/color"
    "github.com/neoh/go-web-boilerplate/lib/settings"
    _ "github.com/neoh/go-web-boilerplate/lib/limiter"
    _ "github.com/neoh/go-web-boilerplate/lib/middleware"
    _ "github.com/neoh/go-web-boilerplate/controllers"
    "github.com/neoh/go-web-boilerplate/models"
    "github.com/neoh/go-web-boilerplate/lib/database"
)


func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    switch settings.Config.DatabaseType {
        case "bolt":
            fmt.Println("Setting up Bolt database")
            
            database.Use(database.BoltDB("bolt.db"))
            database.RunMigrations()
        default:
            fmt.Println("No database driver found")
    }

    models.LoadDependencies()
    
    displayRouteTree()
    
    iris.Listen(":8080")
    //iris.ListenTLS(":443")
    //iris.ListenTLSAuto(":443") // Auto letsencrypt configuration
}

func displayRouteTree() {
    for _, value := range iris.Lookups() {
        color.Cyan("Route: '" + value.Name() + "' at: '" + value.Path() + "' for: '" + value.Method() + "' requests")
    }
}