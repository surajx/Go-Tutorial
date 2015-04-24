//How to use for-select construct to control flow in go routines.
package main

import (
	"fmt"
)

func emitWords(wordChannel chan string, doneChannel chan bool) {
	defer close(wordChannel)
	defer close(doneChannel)
	words := [9]string{"The", "Quick", "Brown", "Fox", "Jumps", "Over", "The", "Lazy", "Dog"}

	i := 0
	for {
		select {
		case wordChannel <- words[i]:
			i += 1
			if i == len(words) {
				i = 0
			}
		case <-doneChannel:
			return
		}
	}

}

func main() {

	wordChannel := make(chan string)
	doneChannel := make(chan bool)

	go emitWords(wordChannel, doneChannel)

	repeatCount := 10
	for word := range wordChannel {
		if repeatCount == 0 {
			doneChannel <- true
			break
		}
		fmt.Printf("%s\n", word)
		repeatCount -= 1
	}

}
