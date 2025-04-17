package patterns

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i 
		fmt.Println("Produced:", i)
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func consumer(ch chan int) {
	for msg := range ch {
		fmt.Println("Consumed:", msg)
		time.Sleep(2 * time.Second)
	}
}

func RunProducerConsumer() { 
	ch := make(chan int, 3)

	go producer(ch)
	go consumer(ch)

	time.Sleep(10 * time.Second)
}