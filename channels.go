//go routines and channels
package main

import (
	"fmt"
)

func makeID(idChannel chan int) {
	id := 0
	for {
		idChannel <- id
		id += 1
	}
}

func main() {
	idChannel := make(chan int)

	go makeID(idChannel)

	fmt.Printf("%d\n", <-idChannel)
	fmt.Printf("%d\n", <-idChannel)
	fmt.Printf("%d\n", <-idChannel)
	fmt.Printf("%d\n", <-idChannel)
}
