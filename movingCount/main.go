package main

import "fmt"

func movingCount(m int, n int, k int) int {

	//数位之和计算
	//可达性分析（从00开始向右/下移动）
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	var traverse func(int, int)
	traverse = func(x, y int) {
		if x > m-1 || y > n-1 || visited[x][y] != 0 || getSum(x)+getSum(y) > k {
			return
		}
		visited[x][y] = 1
		traverse(x, y+1)
		traverse(x+1, y)

	}
	traverse(0, 0)
	ret := 0
	for i := 0; i < m; i++ {
		fmt.Println(visited[i])
		for j := 0; j < n; j++ {
			if visited[i][j] != 0 {
				ret++
			}
		}
	}
	return ret

}

func getSum(x int) int {
	s := 0
	for x != 0 {
		s += x % 10
		x = x / 10
	}
	return s
}

//DFS depth first traversal
