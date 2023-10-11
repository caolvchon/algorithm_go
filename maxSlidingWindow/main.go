package main

//单调队列，单调非增队列
func maxSlidingWindow(nums []int, k int) []int {

}

//单调队列，头部为最大值，单调递减
//push x时，如果尾部<x，尾部出队，直到为空或尾部>=x,则x入队
//pop x时，如果头部==x，则x出队
type queue struct {
	arr []int
}

func (q *queue) push(x int) {
	for i := len(q.arr) - 1; i >= 0; i-- {

	}
}

func (q *queue) pop() {

}
