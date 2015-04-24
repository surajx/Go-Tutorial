//Demostrates how closing a channel can trigger execution of
//a go routine which was othewise blocked.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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

func getPageNow(url string, startChan chan bool) {
	<-startChan
	size, err := getPageSize(url)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Size of %s: %d\n", url, size)
}

func main() {
	startChan := make(chan bool)
	urls := []string{"http://yahoo.com", "http://google.com", "http://quora.com",
		"http://facebook.com", "http://bing.com", "http://github.com"}
	for _, url := range urls {
		go getPageNow(url, startChan)
	}

	time.Sleep(5 * time.Second)
	close(startChan)
	time.Sleep(20 * time.Second)
}
