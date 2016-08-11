package main

import (
    "fmt"
    "log"
    "os"
    "strconv"
    "syscall"
)

func main() {
    process, err := os.FindProcess(int(pid))
    if err != nil {
        fmt.Printf("Failed to find process: %s\n", err)
    } else {
        err := process.Signal(syscall.Signal(0))
        fmt.Printf("process.Signal on pid %d returned: %v\n", pid, err)
    }
}