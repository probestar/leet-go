package l74

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	columns := len(matrix[0])
	count := rows * columns
	i := sort.Search(count, func(i int) bool {
		return matrix[i/columns][i%columns] >= target
	})
	return i < count && matrix[i/columns][i%columns] == target
}
