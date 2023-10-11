package main

import (
	"fmt"
	"reflect"
	"strings"
	//"syscall"
)

type A struct {
	a string
	b int
	c []int
}

func main() {
	//var a interface{}
	//fmt.Println(reflect.TypeOf(a))
	//fmt.Println("aaa")
	//name := "xxx"
	//res := getInstance2(name)
	//fmt.Printf("%v\n", res)

	a1 := A{
		a: "a",
		b: 0,
		c: []int{1, 2, 2},
	}
	a2 := A{
		a: "a",
		b: 0,
		c: []int{1, 2, 3},
	}
	//fmt.Println(a1 == a2)
	fmt.Println(reflect.DeepEqual(a1, a2))

	s := strings.Builder{}

}

type demoA struct {
	name string
}

func getInstance() *demoA {

	dd := new(demoA)
	dd.name = `aaa`
	return dd
}

func getInstance2(s string) demoA {
	dd := new(demoA)
	dd.name = s
	return *dd
}
