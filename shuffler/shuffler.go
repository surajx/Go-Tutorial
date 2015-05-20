package shuffler

import (
	"math/rand"
	"sort"
)

type Shufflable interface {
	Len() int
	Swap(i, j int)
}

type WeightedShufflable interface {
	Shufflable
	Weight(i int) int
}

func Shuffle(s Shufflable) {
	for i := 0; i < s.Len(); i++ {
		s.Swap(i, rand.Intn(s.Len()-i))
	}
}

type sortByWeight WeightedShufflable

func (ws sortByWeight) Len() int           { return len(ws) }
func (ws sortByWeight) Swap(i, j int)      { ws[i], ws[j] = ws[j], ws[i] }
func (ws sortByWeight) Less(i, j int) bool { return ws[i].weight < ws[j].weight }

func WeightedShuffler(ws WeightedShufflable) {
	Shuffle(ws)
	sort.Sort(sortByWeight(ws))
}
