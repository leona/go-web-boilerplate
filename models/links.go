package models

import (
    "github.com/boltdb/bolt" 
    "encoding/json"
    "fmt"
    "errors"
)

type Link struct {
    Id uint32 `json:"id"`
    Created uint64 `json:"created"`
    Redirect string `json:"redirect"`
    Create_ip string `json:"create_ip"`
}

var (
    linkBucketKey = []byte("short_links") 
)


func LinkByKey(key string) (Link, error) {
    result := Link{}
    
    err := db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket(linkBucketKey)

        value := b.Get([]byte(key))
        
        if value == nil {
            return ErrorNotFound
        }
        
        json.Unmarshal(value, &result)

        return nil
    })
    
	return result, err
}

func LinkExists(key string) bool {
    result := false
    
    db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket(linkBucketKey)

        value := b.Get([]byte(key))
        
        if value != nil {
            result = true
        }
        
        return nil
    })
    
    return result
}

func LinkStore(key string, value Link) error {
    if LinkExists(key) {
        return errors.New("Already exists")
    }
    
    return db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(linkBucketKey))
        
        id, _ := b.NextSequence()
        value.Id = uint32(id)
        
        encodedValue, err := json.Marshal(value)
        
        if err != nil {
            fmt.Print(err)
            return err
        }
        
        return b.Put([]byte(key), encodedValue)
    })
}