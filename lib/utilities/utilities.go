package utilities

import (
    "os" 
    "sync"
    "math/rand"
    "time"
    //"fmt"
)

func AppendFile(filename string, contents string) error {
    file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
    
    if err != nil {
        panic(err)
    }
    
    defer file.Close()
    
    if _, err = file.WriteString(contents); err != nil {
        panic(err)
    }
    
    return nil
}


const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
    letterIdxBits = 6                    // 6 bits to represent a letter index
    letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
    letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)
var src = rand.NewSource(time.Now().UnixNano())

func RandomString(n int) string {
    b := make([]byte, n)

    for i, cache, remain := n-1, int63(), letterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
            b[i] = letterBytes[idx]
            i--
        }
        cache >>= letterIdxBits
        remain--
    }

    return string(b)
}

var mutex sync.Mutex

func int63() int64 {
	mutex.Lock()
	v := src.Int63()
	mutex.Unlock()
	return v
}
