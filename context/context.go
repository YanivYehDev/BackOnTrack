package context

import (
	"context"
	"fmt"
	"time"
)


// Basic //
func CreateBasicCtx() {
	ctx := context.Background()
	// Use the context, for example, pass it to a function that requires context.Context
	doSomething(ctx)
}
func doSomething(ctx context.Context) {
	fmt.Println("doing Something...")
}
////////////////////////////////////////


// Example with 'WithCancel' //
func CancelTask() {
	// Step 1: Create a cancelable context
	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan struct{})

	// Step 2: Start a goroutine that listens for ctx.Done or timeout
	go func() {
		select {
		case <-ch:
			fmt.Println("completed normally")
		case <-ctx.Done():
			fmt.Println("context canceled")
		}
	}()

	go longRunningTask(ch)

	// Step 3: Wait 5 seconds, then cancel the context
	time.Sleep(5 * time.Second)
	cancel()

	// Give the goroutine a moment to print before main exits
	time.Sleep(1 * time.Second)

}

func longRunningTask(ch chan struct{}) {
	time.Sleep(10 * time.Second)
	ch <- struct{}{}
}

/////////////////////////////////////////////
