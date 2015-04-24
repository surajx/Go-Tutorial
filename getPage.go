//The program retrieves 3 webpages concurrently and displays their size
//Shows usage of types, how to execute a limited number of routines from a
//pool of workers.
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (w *webPage) get() {
	if w.url != "" {
		resp, err := http.Get(w.url)
		if err != nil {
			w.err = err
			return
		}
		defer resp.Body.Close()
		w.body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			w.err = err
		}
	} else {
		w.err = errors.New("Web Page Instance URL not set")
	}
}

func worker(aWebPage webPage, sema chan bool, sizeChan chan string, id int) {
	<-sema
	fmt.Printf("Launching worker %d\n", id)
	aWebPage.get()
	if aWebPage.err != nil {
		sizeChan <- fmt.Sprintf("Error retrieving page size for %s", aWebPage.url)
	} else {
		sizeChan <- fmt.Sprintf("(%d) Size of %s: %d", id, aWebPage.url, len(aWebPage.body))
	}
	sema <- true
}

type webPage struct {
	url  string
	body []byte
	err  error
}

func main() {
	urls := []string{"http://yahoo.com", "http://google.com", "http://quora.com",
		"http://facebook.com", "http://bing.com", "http://github.com"}

	webPages := make([]webPage, len(urls))
	for idx, url := range urls {
		webPages[idx] = webPage{url: url}
	}

	fmt.Printf("%+v\n", webPages)

	sema := make(chan bool, 3)
	sizeChan := make(chan string)
	for id, aWebPage := range webPages {
		go worker(aWebPage, sema, sizeChan, id)
	}

	for i := 0; i < cap(sema); i++ {
		sema <- true
	}

	for i := 0; i < len(webPages); i++ {
		fmt.Printf("%s\n", <-sizeChan)
	}
}
