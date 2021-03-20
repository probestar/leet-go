package l54

func spiralOrder(matrix [][]int) []int {
	var directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	rows := len(matrix)
	columns := len(matrix[0])

	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, columns)
	}

	directionsIndex := 0
	all := rows * columns
	ret := make([]int, all)
	i := 0
	j := 0
	for k := 0; k < all; k++ {
		ret[k] = matrix[i][j]
		visited[i][j] = true
		i1 := i + directions[directionsIndex][0]
		j1 := j + directions[directionsIndex][1]
		if i1 >= rows || j1 >= columns || i1 < 0 || j1 < 0 || visited[i1][j1] {
			directionsIndex = (directionsIndex + 1) % 4
		}
		i = i + directions[directionsIndex][0]
		j = j + directions[directionsIndex][1]
	}

	return ret
}
