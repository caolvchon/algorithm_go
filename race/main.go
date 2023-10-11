package main

import (
	"time"
)

func main() {
	conAdd()
	time.Sleep(10 * time.Second)
}

func conAdd() {
	a := 0
	go func() {
		for {
			a++
			//fmt.Printf("%v", a)
		}
	}()
	go func() {
		for {
			a += 2
			//fmt.Printf("%v", a)
		}
	}()
}
