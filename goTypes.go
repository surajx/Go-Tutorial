//Creating user defined types analogous to classes et.al.
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type webPage struct {
	url  string
	body []byte
	err  error
}

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

func main() {
	w := &webPage{url: "http://google.com"}
	w.get()
	fmt.Printf("URL: %s, BodyLength: %d, Error: %s\n", w.url, len(w.body), w.err)
}
