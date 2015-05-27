package main

import (
	"fmt"
	"math"
	"strings"
)

//returns i**x
func exp(base interface{}, power int) float64 {
	switch baseVal := base.(type) {
	case int:
		var ans int = baseVal
		for i := 0; i < power-1; i++ {
			ans *= baseVal
		}
		return float64(ans)
	case float64:
		var ans float64 = baseVal
		for i := 0; i < power-1; i++ {
			ans *= baseVal
		}
		return float64(ans)
	}
	return 0
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Magnitude() float64 {
	return math.Sqrt(exp(v.X, 2) + exp(v.Y, 2))
}

type intSlice []int

func (is intSlice) Len() int {
	return len(is)
}
func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func main() {
	fmt.Println(exp(3, 2))
	a := [3]int{1, 2, 3}
	b := a[:]
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", b)
	fmt.Println(a[0])
	fmt.Println(b[0])
	a[0] = 4
	fmt.Println(a[0])
	fmt.Println(b[0])
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", b)

	for range a {
		fmt.Println("hi")
	}

	gps_map := map[string]Vertex{
		"Google": {45.454, 43.3434},
		"Yahoo":  {Y: 56.934, X: 45.0003},
	}

	fmt.Println(gps_map["Yahoo"])
	fmt.Println(gps_map["Google"])

	fmt.Println(gps_map["Google"].Magnitude())

	/*	is := intSlice{3, 1, 4, 1, 5}
		shuffler.Shuffle(is, 23458)
		fmt.Println(is)*/

	helloWorld("hello world", func(msg string) {
		fmt.Println(strings.ToUpper(msg))
	})

	helloWorld("hello world", func(suffix string) Printer {
		return func(s string) {
			fmt.Println(s + suffix)
		}
	}("!!!"))

}

type Printer func(string)

func helloWorld(message string, printFn Printer) {
	printFn(message)
}
