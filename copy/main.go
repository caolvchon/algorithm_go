package main

import "fmt"

type A struct {
	name *string
	age  string
}

func (a A) String() string {
	s := fmt.Sprintf("AAA %v %v", *a.name, a.age)
	return s
}

func main() {
	//slice copy
	//a := []int{1, 2, 3}
	//b := make([]int, 3)
	//copy(b, a)
	//fmt.Println(a, b)
	//b[1] = 111
	//fmt.Println(a, b)
	//
	//b = a
	//b[1] = 222
	//fmt.Println(a, b)
	name := "aaa"
	fmt.Println(A{&name, "a"})

	//n1 := "11"
	//n2 := "22"
	////n3 := "33"
	////n4 := "44"
	//
	//a := []A{{&n1, "1"}, {&n2, "2"}}
	//b := a
	//b[0].age = "33"
	//c := make([]A, 2)
	//copy(c, a)
	////*(c[1].name) = n3

	a := []int{1, 2, 3}
	b := a
	b[1] = 1
	fmt.Println(a, b)

}
