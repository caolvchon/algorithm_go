package main

import "sort"

//leetcode 40 组合总和2
func combinationSum2(candidates []int, target int) [][]int {

	sort.Ints(candidates)

	ret:=dfs2(candidates,target)

	return ret
}

//给备选，返回应有的结果
func dfs2(candidates []int,target int)[][]int{
	ret:=[][]int{}
	for k,v:=range candidates{
		remain:=target-v
		now:=[]int{v}
		if remain==0{
			ret = append(ret,now)
		}else if remain<0{
			break
		}else {
			if k==len(candidates)-1{
				break
			}
			res := dfs2(candidates[k+1:], remain)
			if len(res) > 0 {
				for _, j := range res {
					ret = append(ret, append(now, j...))
				}
			}
		}
	}
	return ret
}


