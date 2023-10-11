package main

import "fmt"

func main() {
	a := 1.2
	b := 1.6
	fmt.Println(int(a), int(b))
}

func reversePairs(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	pairsCount := mergeSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	return pairsCount

}

func mergeSort(nums []int, L, R int) int {
	if L == R {
		return 0
	}
	mid := L + (R-L)>>1

	count1 := mergeSort(nums, L, mid)
	count2 := mergeSort(nums, mid+1, R)
	count3 := getPairCount(nums, L, mid, R)
	merge(nums, L, mid, R)
	return count1 + count2 + count3
}

func merge(nums []int, L, M, R int) {
	help := make([]int, R-L+1)
	p1, p2 := L, M+1
	i := 0
	for ; p1 <= M && p2 <= R; i++ {
		if nums[p1] <= nums[p2] {
			help[i] = nums[p1]
			p1++
		} else {
			help[i] = nums[p2]
			p2++
		}
	}
	if p1 <= M {
		copy(help[i:], nums[p1:M+1])
	}
	if p2 <= R {
		copy(help[i:], nums[p2:R+1])
	}

	copy(nums[L:R+1], help)

}

func getPairCount(nums []int, L, M, R int) int {
	p1, p2 := L, M+1

	pairCount := 0
	for ; p1 <= M; p1++ {
		for ; p2 <= R; p2++ {
			sub := nums[p1] - nums[p2]*2
			if sub > 0 {
				p2++
			} else {
				pairCount += (p2 - M - 1)
				break
			}
		}
		if p2 == R+1 {
			//加上
			pairCount += (p2 - M - 1)
		}
	}

	return pairCount

}
