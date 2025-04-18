package patterns

import (
	"fmt"
	"time"
)

func workerForWorkerPool(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Worker %d started job %d\n", id, j)
        // Simulate some work
        time.Sleep(2 * time.Second)
        results <- j * 2 
    }
}

func RunWorkerPool() {
    jobs := make(chan int, 10)
    results := make(chan int, 10)

    for w := 1; w <= 3; w++ {
        go workerForWorkerPool(w, jobs, results)
    }

    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    for a := 1; a <= 5; a++ {
        fmt.Println(<-results)
    }
}
