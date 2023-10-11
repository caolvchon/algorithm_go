package main

import "fmt"

func main() {
	f := testClosure()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func testClosure() func() int {
	x := 1
	return func() int {
		x++
		return x
	}
}
