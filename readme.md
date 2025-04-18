# Go's Concurrency Model: Everything you need to know.

This repository everything you need to know about concurrency in terms of software engineering and Go. From the basics to advanced patterns.

---

## 1. Fan-In / Fan-Out Pattern

**Fan-Out** distributes tasks to multiple goroutines, while **Fan-In** gathers their results into a single channel.


```go
func workerFIFO(id int, ch chan int) {
    for i := 0; i < 5; i++ {
        time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
        ch <- i + id*100 
    }
}

func RunFanInFanOut() {
    ch := make(chan int, 10)
    for i := 0; i < 3; i++ {
        go workerFIFO(i, ch) 
    }

    for i := 0; i < 15; i++ {
        fmt.Println(<-ch)
    }
}
```


## 2. Pipeline Pattern

Each stage is a goroutine with input/output channels. Used for processing steps in sequence asynchronously.


```go
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n * 10
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func RunPipeline() {
	gen := generate(1, 2, 3, 4, 5)
	sq := square(gen)

	for result := range sq {
		fmt.Println(result)
	}
}
```


## 3. Worker Pool

Limit the number of goroutines by having a fixed number of workers pulling tasks from a jobs channel.


```go
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
```


## 4. Pub/Sub Pattern

Publish the data to subscriber, the subscribers can listen to messages and perform their jobs.


```go
func publisher(subs []chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	messages := []string{"message1", "message2", "message3"}
	for _, msg := range messages {
		for _, sub := range subs {
			sub <- msg 
		}
	}

	for _, sub := range subs {
		close(sub)
	}
}

func subscriber(id int, ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range ch {
		fmt.Printf("Subscriber %d received: %s\n", id, msg)
	}
}
```


## ⭐️ Star this repo if you learned something cool!
