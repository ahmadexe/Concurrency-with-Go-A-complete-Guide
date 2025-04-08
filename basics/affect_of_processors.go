package basics

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(id int,  wg *sync.WaitGroup) {
	defer wg.Done() 

	for i := range 5 {
		_ = calculatePrime(i) 
		fmt.Printf("%s - Worker %d: step %d\n", time.Now().Format("15:04:05.000"), id, i)
	}
}

func RunAffectOfProcessors() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup

	wg.Add(2) 

	for i := range 2 {
		go worker(i, &wg)
	}

	wg.Wait()
}

func calculatePrime(n int) int {
	// CPU-bound task: finding prime number
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			return n
		}
	}
	return n
}