package main

func isBipartite(graph [][]int) bool {
	//染色
	//从一个点出发，自己染为白，相连节点x染为黑，然后x的相连节点染为白，直到出现色彩不一致，或都遍历完
	//从第二个点出发，如果被染过，continue，如果未被染过，重复上一步骤

	nodeColor := make(map[int]int)

	var addColor func(i int, color int) bool
	addColor = func(i int, color int) bool {
		//有染色，判断颜色正确与否
		if c, ok := nodeColor[i]; ok {
			if c != color {
				return false
			}
			return true
		}

		//没染色，进行染色
		nodeColor[i] = color
		nexts := graph[i]
		if len(nexts) == 0 {
			return true
		}
		for _, v := range nexts {
			res := addColor(v, 1-color)
			if res == false {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(graph); i++ {
		color, ok := nodeColor[i]
		if ok {
			continue
		}
		res := addColor(i, color)
		if res == false {
			return false
		}
	}
	return true
}
