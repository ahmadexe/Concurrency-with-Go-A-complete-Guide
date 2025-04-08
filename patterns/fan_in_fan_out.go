package patterns

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, ch chan int) {
    for i := range 5 {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
        ch <- i + id*100 
    }
}

func RunFanInFanOut() {
    ch := make(chan int, 10)
    for i := range 3 {
        go worker(i, ch) 
    }

    for range 15 {
        fmt.Println(<-ch)
    }
}
