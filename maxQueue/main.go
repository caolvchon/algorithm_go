package main

import "fmt"

func main() {
	m := Constructor()
	m.Push_back(1)
	fmt.Println(m.Max_value())
	m.Push_back(2)
	fmt.Println(m.Max_value())
	m.Push_back(5)
	fmt.Println(m.Max_value())
	m.Push_back(4)
	fmt.Println(m.Max_value())
	m.Push_back(3)
	fmt.Println(m.Max_value())
	m.Push_back(4)
	fmt.Println(m.Max_value())

	fmt.Printf("pop=%v max:=%v \n", m.Pop_front(), m.Max_value())
	fmt.Printf("pop=%v max:=%v \n", m.Pop_front(), m.Max_value())
	fmt.Printf("pop=%v max:=%v \n", m.Pop_front(), m.Max_value())
	fmt.Printf("pop=%v max:=%v \n", m.Pop_front(), m.Max_value())
	fmt.Printf("pop=%v max:=%v \n", m.Pop_front(), m.Max_value())
	fmt.Printf("pop=%v max:=%v \n", m.Pop_front(), m.Max_value())
	fmt.Printf("pop=%v max:=%v \n", m.Pop_front(), m.Max_value())
	fmt.Printf("pop=%v max:=%v \n", m.Pop_front(), m.Max_value())

}

//增加一个辅助双端队列tmp
//q入队x时，判断tmp尾t跟x的大小，如果t>=x,x同时入队tmp，如果t<x,tmp队尾出队，直到t>=x,x入队tmp
//q出队x时，判断tmp队头h跟x的大小，如果h==x,tmp对头出队
//max_value,直接取tmp队头h

type MaxQueue struct {
	q   []int
	tmp []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		q:   []int{},
		tmp: []int{},
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.q) == 0 {
		return -1
	}
	return this.tmp[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.q = append(this.q, value)
	tmp := this.tmp
	i := len(tmp) - 1
	for ; i >= 0; i-- {
		if value <= tmp[i] {
			break
		}
	}

	this.tmp = this.tmp[0 : i+1]
	this.tmp = append(this.tmp, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.q) == 0 {
		return -1
	}
	res := this.q[0]
	this.q = this.q[1:]
	if res == this.tmp[0] {
		this.tmp = this.tmp[1:]
	}
	return res
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
