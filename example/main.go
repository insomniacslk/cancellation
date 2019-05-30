package main

import (
	"fmt"
	"time"

	"github.com/insomniacslk/cancellation"
)

func main() {
	// non-blocking
	c, cancel := cancellation.New()
	fmt.Println("non-blocking cancellation:", c.DoneNonBlock())
	cancel()
	fmt.Println("non-blocking cancellation:", c.DoneNonBlock())

	// blocking
	c, cancel = cancellation.New()
	go func() {
		time.Sleep(time.Second)
		cancel()
	}()
	<-c.Done()
	fmt.Println("blocking cancellation: done")
}
