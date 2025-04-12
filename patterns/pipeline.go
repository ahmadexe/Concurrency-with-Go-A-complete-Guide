package patterns

import "fmt"

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
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
