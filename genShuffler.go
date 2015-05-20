package main

import (
	"fmt"
	"math/rand"
)

type shuffler interface {
	Len() int
	Swap(i, j int)
}

func shuffle(s shuffler) {
	for i := 0; i < s.Len(); i++ {
		s.Swap(i, rand.Intn(s.Len()-i))
	}
}

type intSlice []int

func (is intSlice) Len() int {
	return len(is)
}
func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

type student struct {
	roll int
	name string
}

type studentSlice []student

func (ss studentSlice) Len() int {
	return len(ss)
}
func (ss studentSlice) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func main() {
	is := intSlice{1, 2, 3, 4, 5, 6, 7}
	ss := studentSlice{
		{roll: 1, name: "Suraj"},
		{roll: 2, name: "Ramya"},
		{roll: 3, name: "Unni"},
		{roll: 4, name: "Priya"}}

	fmt.Printf("IntSlice Value: %v", is)
	fmt.Printf("StringSlice Value: %v", ss)

	shuffle(is)
	shuffle(ss)

	fmt.Printf("\nIntSlice Value: %v", is)
	fmt.Printf("StringSlice Value: %v\n", ss)
}
