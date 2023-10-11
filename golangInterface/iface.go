package main

type A interface {
	Jump(x int) int
}

type B struct {
}

func (b B) Jump(x int) int {
	return 1
}

func main() {
	var a A
	b := B{}
	a = b
	a.Jump()
}
