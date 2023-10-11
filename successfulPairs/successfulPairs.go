package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {

	spells := []int{5, 1, 3}
	potions := []int{1, 2, 3, 4, 5}
	success := int64(7)

	arr := []int{1, 2, 2, 3, 3, 3, 4, 5, 6}
	x := sort.SearchInts(arr, 3)
	fmt.Println(x)
	os.Exit(1)

	res := successfulPairs(spells, potions, success)
	_ = res

}

func successfulPairs(spells []int, potions []int, success int64) []int {
	if len(spells) == 0 || len(potions) == 0 {
		return []int{}
	}

	sort.Ints(potions)

	index := make([]int, len(spells))
	for i := 0; i < len(spells); i++ {
		index[i] = i
	}
	t := T{
		index:  index,
		values: spells,
	}

	sort.Sort(t)

	res := make([]int, len(spells))
	j := len(potions) - 1
	for i := 0; i < len(spells); i++ {
		for j >= 0 && int64(spells[i]*potions[j]) >= success {
			j--
		}
		res[i] = len(potions) - 1 - j
	}

	t2 := T{
		index:  res,
		values: t.index,
	}
	sort.Sort(t2)
	return t2.index
}

type T struct {
	index  []int
	values []int
}

func (t T) Len() int {
	return len(t.index)
}

func (t T) Less(x, y int) bool {
	if t.values[x] <= t.values[y] {
		return true
	}
	return false
}

func (t T) Swap(x, y int) {
	t.index[x], t.index[y] = t.index[y], t.index[x]
	t.values[x], t.values[y] = t.values[y], t.values[x]
}
