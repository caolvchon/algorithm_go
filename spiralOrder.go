package main

func SpiralOrder(matrix [][]int) []int {

	rows := len(matrix)
	if rows == 0 {
		return []int{}
	}
	columns := len(matrix[0])
	ret := []int{}

	short := rows
	if columns < rows {
		short = columns
	}

	//
	for n := 0; n < (short+1)/2; n++ {
		for i := n; i < columns-n-1; i++ {
			ret = append(ret, matrix[n][i])
		}
		if rows-n-1 <= n {
			ret = append(ret, matrix[n][columns-n-1])
			break
		}

		for i := n; i < rows-n-1; i++ {
			ret = append(ret, matrix[i][columns-n-1])
		}
		if columns-n-1 <= n {
			ret = append(ret, matrix[rows-n-1][columns-n-1])
			break
		}

		for i := columns - n - 1; i > n; i-- {
			ret = append(ret, matrix[rows-n-1][i])
		}

		for i := rows - n - 1; i > n; i-- {
			ret = append(ret, matrix[i][n])
		}
	}
	return ret
}
