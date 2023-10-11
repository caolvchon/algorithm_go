package main

import (
	"sort"
)

//leetcode 39 组合总和
func CombinationSum(candidates []int, target int) [][]int {

	sort.Ints(candidates)

	ret:=dfs(candidates,target)

	return ret
}

//给备选，返回应有的结果
func dfs(candidates []int,target int)[][]int{
	ret:=[][]int{}
	for k,v:=range candidates{
		remain:=target-v
		now:=[]int{v}
		if remain==0{
			ret = append(ret,now)
		}else if remain<0{
			break
		}else {
			res := dfs(candidates[k:], remain)
			if len(res) > 0 {
				for _, j := range res {
					ret = append(ret, append(now, j...))
				}
			}
		}
	}
	return ret
}

