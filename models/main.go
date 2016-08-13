package models

import (
    "github.com/neoh/go-web-boilerplate/lib/database"   
    "github.com/boltdb/bolt"
    "errors"
)

var db *bolt.DB

func LoadDependencies() {
    db = database.Instance
}

var (
	ErrorNotFound = errors.New("Not found in database")
)
