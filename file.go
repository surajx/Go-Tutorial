package main

import (
	"fmt"
	"os"
)

func checkAndReportError(err error, errorMsg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", errorMsg, err)
		os.Exit(1)
	}
}

func main() {
	f, err := os.Open("test.txt")
	checkAndReportError(err, "Error opening file")
	b := make([]byte, 165)
	n, err := f.Read(b)
	checkAndReportError(err, "Error reading file")
	fmt.Printf("%d bytes read.\n", n)
	fmt.Printf("%s", string(b))
}
