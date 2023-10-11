package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcd"
	indices := []int{0, 2}
	sources := []string{"a", "cd"}
	targets := []string{"eee", "ffff"}
	res := findReplaceString(s, indices, sources, targets)
	fmt.Println(res)
}

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	//linux

	if len(indices) == 0 || len(sources) == 0 || len(targets) == 0 {
		return s
	}

	r := strings.Builder{}
	r.WriteString(s[0:indices[0]])

	for i, index := range indices {
		nextFrom := index
		nextTo := len(s)
		if i != len(indices)-1 {
			nextTo = indices[i+1]
		}

		if sources[i] == s[index:index+len(sources[i])] {
			r.WriteString(targets[i])
			nextFrom += len(sources[i])
		}

		r.WriteString(s[nextFrom:nextTo])
	}
	return r.String()
}
