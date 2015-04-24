//Demostrates how the channel capacity parameter can be used to trigger
//a subset of go routines from a pool of workers.
package main

import (
	"fmt"
	"time"
)

func work(id int) {
	time.Sleep(5 * time.Second)
	fmt.Printf("Work Done by %d\n", id)
}

func worker(sema chan bool, id int) {
	<-sema
	work(id)
	sema <- true
}

func main() {
	sema := make(chan bool, 10)

	for i := 0; i < 1000; i++ {
		go worker(sema, i)
	}

	for i := 0; i < cap(sema); i++ {
		sema <- true
	}

	time.Sleep(30 * time.Second)
}
