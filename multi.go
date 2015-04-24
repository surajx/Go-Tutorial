//Using channels to make concurrent web requests
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getPageSize(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	return len(body), nil
}

func getPageSizeRoutine(urlChannel chan string, sizeChannel chan int) {
	for {
		url := <-urlChannel
		fmt.Printf("Retrieving size of %s...\n", url)
		size, err := getPageSize(url)
		if err != nil {
			sizeChannel <- 0
		} else {
			sizeChannel <- size
		}
	}
}

func main() {
	urls := []string{"http://yahoo.com", "http://google.com", "http://quora.com",
		"http://facebook.com", "http://bing.com", "http://github.com"}
	urlChannel := make(chan string)
	sizeChannel := make(chan int)

	for i := 0; i < 3; i++ {
		go getPageSizeRoutine(urlChannel, sizeChannel)
	}

	for _, url := range urls {
		go func(url string) {
			urlChannel <- url
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Printf("Size: %d\n", <-sizeChannel)
	}
}
