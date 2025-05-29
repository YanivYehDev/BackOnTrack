package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func HowToWorkWithSelect() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Second) // simulate work delay
		ch <- "hello"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout")
	}
}

func ParallelExecutionWithWaitGroup() {
	resultsChan := make(chan string)
	var wg sync.WaitGroup

	task := func(id int, resultsChan chan string) {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		resultsChan <- fmt.Sprintf("done %d", id)
	}

	wg.Add(3)
	go task(1, resultsChan)
	go task(2, resultsChan)
	go task(3, resultsChan)

	// one option - wait for all go routines separatly
	// <- resultsChan
	// <- resultsChan
	// <- resultsChan

	// second option - use range on channel + waitGroup in order to close the channel

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// A key point to note is that the loop will automatically exit when the channel is closed,
	// preventing any deadlocks or infinite loops.
	// This makes it a safe and idiomatic way to consume all values sent on a channel in Go.
	for res := range resultsChan {
		fmt.Printf("finished task with message %s\n", res)
	}

	fmt.Println("finished all tasks")
}


func ParallelExecutionWithSelect() {
	resultsChan := make(chan string)
	var wg sync.WaitGroup

	task := func(id int, resultsChan chan string) {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		resultsChan <- fmt.Sprintf("done %d", id)
	}

	wg.Add(3)
	go task(1, resultsChan)
	go task(2, resultsChan)
	go task(3, resultsChan)

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// A key point to note is that the loop will automatically exit when the channel is closed,
	// preventing any deadlocks or infinite loops.
	// This makes it a safe and idiomatic way to consume all values sent on a channel in Go.
	timeout := time.After(2 * time.Second)
	for {
		select {
		case res, ok := <-resultsChan:
			if !ok {
				return
			}
			fmt.Printf("finished task with message %s\n", res)
		case <-timeout:
			fmt.Println("timeout while waiting for tasks")
			return
		}
	}
}