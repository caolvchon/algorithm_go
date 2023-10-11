package main

import "fmt"

func main() {
	arr := []int{1, 33, 55, 3322, 45, 3, 5, 87, 3, 6, 23, 7865, 45}
	heapSort(arr)
	fmt.Println(arr)

}

func heapSort(arr []int) {
	if len(arr) == 0 || len(arr) == 1 {
		return
	}
	for i := 0; i < len(arr); i++ {
		heapInsert(arr, i)
	}

	size := len(arr)

	for size > 1 {
		swap(arr, 0, size-1)
		heapify(arr, 0, size-1)
		size--
	}

}

//从index开始，从上到下整理堆
func heapify(arr []int, index int, size int) {
	//index  index*2+1 index*2+2
	//左孩子的下标
	left := index*2 + 1

	for left < size { //下方还有孩子
		right := left + 1
		//largest是最大值的下标
		largest := left
		if right < size && arr[right] > arr[left] {
			largest = right
		}

		if arr[index] > arr[largest] {
			largest = index
		}

		if largest == index { //父节点最大，heapify停止
			break
		}

		swap(arr, index, largest)
		index = largest
		left = 2*index + 1
	}
}

//把index
func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		swap(arr, index, (index-1)/2)
		index = (index - 1) / 2
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

//大顶堆
//type Heap struct {
//	arr []int
//}
//
//func (h *Heap) len() int {
//	return len(h.arr)
//}

//func (h *Heap) Init() {
//	for i := 0; i < len(h.arr); i++ {
//		h.heapInsert(h.arr, i)
//	}
//}
//
////拿走index=0，把index=len-1放到0
////对index=0做heapidfy（从上到下）
//func (h *Heap) Pop() int {
//	if len(h.arr) == 0 {
//		return 0
//	}
//	res := h.arr[0]
//	swap(h.arr, 0, len(h.arr)-1)
//	h.arr = h.arr[0 : len(h.arr)-1]
//	h.heapify(h.arr, 0, len(h.arr))
//	return res
//}
//
//func (h *Heap) Push(x int) {
//	h.arr = append(h.arr, x)
//	h.heapInsert(h.arr, len(h.arr)-1)
//}
