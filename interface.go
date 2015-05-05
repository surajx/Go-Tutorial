package main

import (
	"fmt"
	"math/rand"
)

type shuffler interface {
	blaher()
	Len() int
	Swap(i, j int)
}

func shuffle(s shuffler) {
	for i := 0; i < s.Len(); i++ {
		j := rand.Intn(s.Len() - 1)
		s.Swap(i, j)
	}
}

type marks struct {
	mark int
	name string
}

type marksSlice []marks

func (m marksSlice) Len() int {
	return len(m)
}

func (ma marksSlice) Swap(m1, m2 int) {
	ma[m1], ma[m2] = ma[m2], ma[m2]
}

func (mar marksSlice) blaher() {
	stringShuffler := func(name string) string {
		for i := 0; i < len(name); i++ {

		}
	}
	for i := 0; i < len(mar); i++ {
		mar[i].name = stringShuffler(mar[i].name)
	}
}

func main() {
	studentMarks := marksSlice{
		{mark: 12, name: "suraj"},
		{mark: 16, name: "ramya"},
		{mark: 67, name: "priya"},
		{mark: 68, name: "unni"},
	}
	shuffle(studentMarks)

	fmt.Printf("Shuffled marks: %v", studentMarks)
}
