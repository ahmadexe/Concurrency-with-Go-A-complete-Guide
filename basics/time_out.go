package basics

import (
    "fmt"
    "time"
)

func work(done chan<- bool) {
    time.Sleep(2 * time.Second)
    done <- true
}

func RunTimeOut() {
    done := make(chan bool)
    go work(done)

    select {
    case <-done:
        fmt.Println("Work done")
    case <-time.After(1 * time.Second):
        fmt.Println("Timeout reached")
    }
}
