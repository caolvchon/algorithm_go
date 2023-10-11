package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Printf("panic:%v", err)
	//	}
	//}()
	//executePanic()

	//a := 1
	//defer func(x int) {
	//	fmt.Println(x)
	//}(a + 1)
	//a = 100

	//x := f()
	//fmt.Println(x)

	//	x := f2()
	//	fmt.Printf("x=%v g=%v", x, g)

	//fmt.Println(reflect.TypeOf(A))
	//fmt.Println(reflect.TypeOf(B))
	//fmt.Println(reflect.TypeOf(C))
	//fmt.Println(reflect.TypeOf(D))
	//
	//x := 1
	//fmt.Println(A * x)
	//
	//s := "GO语言"
	//for i := 0; i < len(s); i++ {
	//	fmt.Println(s[i])
	//}

	//a := 1
	//
	//go func() {
	//	for {
	//		a++
	//	}
	//
	//}()
	//
	//go func() {
	//	for {
	//		a--
	//	}
	//}()
	//time.Sleep(time.Second * 100)

	a := int32(1)
	succ := atomic.CompareAndSwapInt32(&a, 1, 2)
	succ1 := atomic.CompareAndSwapInt32(&a, 1, 2)
	fmt.Println(succ, succ1)
	m := sync.Mutex{}
	m.Lock()
	m.Unlock()

}

const A = 1
const B = 1.2
const C = 0x111
const D = 333111111

func executePanic() {
	panic("xxxxx panic xxxx")

}

func f() (r int) {
	defer func() {
		r = 100
	}()
	r = 0
	return r
}

var g = 100

func f2() (r int) {
	defer func() {
		g = 200
	}()
	return g
}
