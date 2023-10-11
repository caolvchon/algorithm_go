package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	var a int
	a = 1
	fmt.Println(a)
}
func Example(slice []string, str string, i int) {
	debug.PrintStack()
}
