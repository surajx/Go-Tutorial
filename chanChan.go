//How to use channel of channels
package main

import (
	"fmt"
	"time"
)

func emit(chanChannel chan chan string, doneChannel chan bool) {
	wordChannel := make(chan string)
	defer close(wordChannel)
	chanChannel <- wordChannel
	words := []string{"I", "Hate", "Lorem", "Ipsum"}
	i := 0
	t := time.NewTimer(3 * time.Second)
	for {
		select {
		case wordChannel <- words[i]:
			i += 1
			if i == len(words) {
				i = 0
			}
		case <-doneChannel:
			return
		case <-t.C:
			return
		}
	}
}

func main() {
	chanChannel := make(chan chan string)
	wordChannel := make(chan string)
	doneChannel := make(chan bool)

	go emit(chanChannel, doneChannel)

	wordChannel = <-chanChannel

	for word := range wordChannel {
		fmt.Printf("%s", word)
	}

}
