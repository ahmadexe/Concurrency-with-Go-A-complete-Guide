package patterns

import (
	"fmt"
	"sync"
)

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

func RunPubSub() {
	var wg sync.WaitGroup

	numSubs := 2
	subs := make([]chan string, numSubs)

	// Create subscriber channels
	for i := range numSubs {
		subs[i] = make(chan string)
		wg.Add(1)
		go subscriber(i+1, subs[i], &wg)
	}

	wg.Add(1)
	go publisher(subs, &wg)

	wg.Wait()
}
