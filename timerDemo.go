//Demostrates How to use timer
package main

import (
	"fmt"
	"time"
)

func timerCallback(t *time.Timer) {
	a := <-t.C
	fmt.Printf("Timeout: %s\n", time.Now())
	t.C <- time.Now()
}

func main() {

	t := time.NewTimer(5 * time.Second)
	for i := 0; i < 5; i++ {
		go timerCallback(t)
	}

	time.Sleep(10 * time.Second)
}
