package main

import "sort"

func minimumRemoval(beans []int) int64 {
	//剩的最多等于拿的最少

	//1、排序 2、遍历，数每次剩下的豆子数

	beanNum := 0
	sort.Ints(beans)

	remainNumMax := 0
	for i, num := range beans {
		beanNum += num

		remainNum := num * (len(beans) - i)
		if remainNum > remainNumMax {
			remainNumMax = remainNum
		}
	}
	return int64(beanNum - remainNumMax)

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
