package main

import (
	"fmt"
	"time"

	"github.com/insomniacslk/interruption"
)

func main() {
	// non-blocking
	c, cancel := interruption.New()
	fmt.Println("non-blocking interruption:", c.DoneNonBlock())
	cancel()
	fmt.Println("non-blocking interruption:", c.DoneNonBlock())

	// blocking
	c, cancel = interruption.New()
	go func() {
		time.Sleep(time.Second)
		cancel()
	}()
	<-c.Done()
	fmt.Println("blocking interruption: done")
}
