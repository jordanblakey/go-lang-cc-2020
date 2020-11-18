package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// goRoutineExample()
	// waitGroupExample()
	// channelExample()
	// channelCapacityExample()
	// selectExample()
	workersExample()
}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func goRoutineExample() {
	// Run function in background as a Go Routine
	go count("sheep")
	go count("fish")

	// Waits for user input and prevents the program from exiting
	fmt.Scanln()
}

func waitGroupExample() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		count("sheep")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		count("fish")
		wg.Done()
	}()

	wg.Wait()
}

func chanCount(thing string, c chan string) {
	for i := 0; i <= 5; i++ {
		c <- fmt.Sprint(i) + ": " + thing
		time.Sleep(time.Millisecond * 500)
	}
	// closes the channel
	close(c)
}

func channelExample() {
	channel := make(chan string)
	go chanCount("sheep", channel)

	// for {
	// 	msg, open := <-channel // Execution stops to wait for value to assign from channel
	// 	if !open { // Detects if channel has been closed and breaks the for loop
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }

	// Syntactic sugar for the above
	for msg := range channel {
		fmt.Println(msg)
	}
}

func channelCapacityExample() {
	c := make(chan string, 2) // Make a channel with capacity 2
	c <- "hello"              // Capacity makes this non-blocking
	c <- "world"
	msg := <-c
	fmt.Println(msg)
	msg = <-c
	fmt.Println(msg)
}

func selectExample() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		// forks the two statements waiting for assignment from a channel
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func workersExample() {
	defer timeTrack(time.Now(), "workersExample: 4 workers")

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 45; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 45; j++ {
		fmt.Printf("%v, ", <-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("\n%s took %s", name, elapsed)
}
