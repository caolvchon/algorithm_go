package main

func maxSlidingWindow(nums []int, k int) []int {
	ret := make([]int, len(nums)-k+1)
	maxQueue := []int{nums[0]}

	//初始化
	for i := 1; i < k; i++ {
		end := 0
		for j := len(maxQueue) - 1; j >= 0; j-- {
			if maxQueue[j] >= nums[i] {
				end = j + 1
				break
			}
		}
		maxQueue = append(maxQueue[:end], nums[i])

	}

	for i := 0; i <= len(nums)-k; i++ {
		if i == 0 {
			ret[i] = maxQueue[0]
			continue
		}
		if maxQueue[0] == nums[i-1] {
			maxQueue = maxQueue[1:]
		}
		end := 0
		for j := len(maxQueue) - 1; j >= 0; j-- {
			if nums[i+k-1] <= maxQueue[j] {
				end = j + 1
				break
			}
		}
		maxQueue = append(maxQueue[:end], nums[i+k-1])
		ret[i] = maxQueue[0]
	}
	return ret
}
