package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{44, 31, 4, 5, 7, 9, 8, 6, 5, 4, 2, 32, 3, 23, 4, 24, 2, 32, 31, 25, 25, 24}
	arr2 := make([]int, len(arr1))
	copy(arr2, arr1)
	fmt.Printf("%v \n%v \n", arr1, arr2)
	mergeSort(arr1)
	sort.Ints(arr2)
	fmt.Printf("%v \n%v ", arr1, arr2)

}

func mergeSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	process(arr, 0, len(arr)-1)
}

func process(arr []int, L, R int) {
	if L == R {
		return
	}
	mid := L + (R-L)>>1
	process(arr, L, mid)
	process(arr, mid+1, R)
	merge(arr, L, mid, R)
}

func merge(arr []int, L, M, R int) {
	//归并
	length := R - L + 1
	help := make([]int, length)
	i, j := L, M+1
	h := 0
	for i <= M && j <= R {
		if arr[i] <= arr[j] {
			help[h] = arr[i]
			i++
		} else {
			help[h] = arr[j]
			j++
		}
		h++
	}

	for i <= M {
		help[h] = arr[i]
		h++
		i++
	}
	for j <= R {
		help[h] = arr[j]
		h++
		j++
	}

	copy(arr[L:R+1], help)
	return
}
