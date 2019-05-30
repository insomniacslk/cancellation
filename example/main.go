package main

import (
	"fmt"
	"time"

	"github.com/insomniacslk/cancellation"
)

func main() {
	c := cancellation.New()
	go func() {
		c.Cancel()
	}()
	time.Sleep(time.Second)
	fmt.Println(c.DoneNonBlock())
}
