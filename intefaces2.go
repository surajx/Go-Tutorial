package main

import (
	"./shuffler"
	"fmt"
)

type intSlice []int

func (is intSlice) Len() int {
	return len(is)
}

func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

type weightedInt struct {
	val    int
	weight int
}

type weightedIntSlice []weightedInt

func (wis weightedIntSlice) Len() int {
	return len(wis)
}

func (wis weightedIntSlice) Swap(i, j int) {
	wis[i], wis[j] = wis[j], wis[i]
}

func (wis weightedIntSlice) Weight(i int) int {
	return wis[i].weight
}

func main() {
	is := intSlice{1, 2, 3, 4, 5, 6}
	shuffler.Shuffle(is)
	fmt.Printf("%v\n", is)

	wis := weightedIntSlice{{val: 1, weight: 10}, {val: 2, weight: 100},
		{val: 3, weight: 1000}, {val: 4, weight: 1}, {val: 5, weight: 400},
		{val: 6, weight: 100000}}
	shuffler.WeightedShuffler(wis)
	fmt.Printf("%v\n", wis)
}
