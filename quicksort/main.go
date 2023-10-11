package main

import (
	"fmt"
	"math/rand"
)

//
func main() {
	arr := []int{1, 33, 55, 3322, 45, 3, 55, 5, 87, 3, 6, 23, 7865, 45}
	//divide2(arr, 55)

	quickSort3(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

//快排1.0
//快排2.0
//快排3.0 随机找划分值，等于划分值的数放在中间
func quickSort3(arr []int, L, R int) {
	if L < R {
		swap(arr, R, L+rand.Int()%(R-L+1))
		p := partition(arr, L, R)
		quickSort3(arr, L, p[0]-1)
		quickSort3(arr, p[1]+1, R)
	}
}

func partition(arr []int, L, R int) []int {
	mid := arr[R]
	lt := L - 1
	gt := R
	i := L
	for i < R && i < gt {
		if arr[i] == mid {
			i++
		} else if arr[i] < mid {
			swap(arr, lt+1, i)
			lt++
			i++
		} else {
			swap(arr, gt-1, i)
			gt--
		}
	}
	swap(arr, R, gt)
	return []int{lt + 1, gt}
}

//左右指针交换
func divide(arr []int, mid int) {
	l, r := 0, len(arr)-1
	for l < r {
		for arr[l] <= mid && l < r {
			l++
		}
		for arr[r] > mid && l < r {
			r--
		}
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

//小于等于区域往右推进
func divide2(arr []int, mid int) int {
	i := 0
	lte := -1
	for i <= len(arr)-1 {
		if arr[i] > mid {
			i++
		} else {
			arr[i], arr[lte+1] = arr[lte+1], arr[i]
			lte++
			i++
		}
	}
	return lte + 1
}

//小于mid放左边，等于放中间，大于放右边
//两个区域，小于区域，大于区域向中间挤压
func divide3(arr []int, mid int) {
	i := 0
	lt := -1
	gt := len(arr)
	for i <= len(arr)-1 && i < gt {
		if arr[i] == mid {
			i++
		} else if arr[i] > mid {
			swap(arr, i, gt-1)
			gt--
		} else {
			swap(arr, i, lt+1)
			lt++
			i++
		}
	}

}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
