package database

import (
    "github.com/boltdb/bolt" 
    "time"
)

func BoltDB(name string) *bolt.DB {
    db, err := bolt.Open(name, 0600, &bolt.Options{Timeout: 1 * time.Second})
    
    if err != nil {
        panic(err)
    }
    
    //defer db.Close()

    return db
}
