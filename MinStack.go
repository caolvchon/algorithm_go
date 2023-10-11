package main

type MinStack struct {
	arr    []int
	minArr []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		arr: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.arr = append(this.arr, x)
	if len(this.minArr) == 0 {
		this.minArr = append(this.minArr, x)
	} else {
		if this.minArr[len(this.minArr)-1] >= x {
			this.minArr = append(this.minArr, x)
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.arr) == 0 {
		return
	}
	p := this.arr[len(this.arr)-1]
	this.arr = this.arr[:len(this.arr)-1]
	if p == this.minArr[len(this.minArr)-1] {
		this.minArr = this.minArr[:len(this.minArr)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.arr) == 0 {
		return 0
	}
	return this.arr[len(this.arr)-1]
}

func (this *MinStack) Min() int {
	if len(this.minArr) == 0 {
		return 0
	}
	return this.minArr[len(this.minArr)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
