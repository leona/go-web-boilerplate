package migrations

import (
    "github.com/boltdb/bolt"
    "fmt"
)


func BoltDB(db *bolt.DB) {
    createBucket := func(name string) {
        tx, err := db.Begin(true)
        if err != nil {
            fmt.Printf("\n", err)
            return
        }
        
        defer tx.Rollback()
        
        _, err = tx.CreateBucket([]byte(name))
        
        if err != nil {
            fmt.Printf("\n", err)
            return
        }
        
        if err := tx.Commit(); err != nil {
            fmt.Printf("\n", err)
        }
    }
    
    createBucket("short_links")
}
