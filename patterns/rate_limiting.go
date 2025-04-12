package patterns

import (
    "fmt"
    "time"
)

func task(id int) {
    fmt.Printf("Task %d processed\n", id)
}

func RunRateLimiting() {
    ticker := time.NewTicker(5 * time.Second)
    defer ticker.Stop()

    for i := 1; i <= 5; i++ {
        <-ticker.C
        task(i)
    }
}
