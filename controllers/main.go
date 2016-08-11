package controllers

import (
    "github.com/leonharvey/go-web-boilerplate/lib/database"   
    "github.com/boltdb/bolt" 
)

var db *bolt.DB

func LoadDependencies() {
    db = database.Instance
}