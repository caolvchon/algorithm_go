package main

import (
	"container/heap"
	"fmt"
)

func main() {
	intHeap := &IntHeap{1, 33, 55, 3322, 45, 3, 5, 87, 3, 6, 23, 7865, 45}
	heap.Init(intHeap)

	heap.Push(intHeap, 9999)
	heap.Push(intHeap, 33)
	heap.Push(intHeap, 444)
	heap.Push(intHeap, 11)
	for len(*intHeap) > 0 {
		fmt.Printf("%v ", heap.Pop(intHeap))
	}
}

type IntHeap []int // 定义一个类型

func (h IntHeap) Len() int { return len(h) } // 绑定len方法,返回长度
func (h IntHeap) Less(i, j int) bool { // 绑定less方法
	return h[i] < h[j] // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h IntHeap) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}
