package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(findNthDigit(2889 + 1))

}

func findNthDigit(n int) int {
	base, i := 9, 1
	for {
		comp := i * base
		if n > comp {
			n -= comp
			i++
			base *= 10
			continue
		} else {
			break
		}
	}
	fmt.Printf("n:%v i:%v \n", n, i)

	x := (n - 1) / i
	y := (n - 1) % i

	target := int(math.Pow10(i)) + x
	fmt.Println("target=", target)
	s := strconv.Itoa(target)
	return int(s[y] - byte('0'))
}
