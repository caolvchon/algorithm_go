package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := X{
		Name: "1",
		Age:  "2",
	}

	fmt.Printf("type:%v kind:%v", reflect.TypeOf(a).Name(), reflect.TypeOf(a).Kind())
	fmt.Printf("value :%v", reflect.ValueOf(a))

}

type X struct {
	Name string `json:"x_name"`
	Age  string `json:"x_age"`
}
