package main

var directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	r := 0
	c := 0
	index := 0
	for i := 1; i <= n*n; i++ {
		ret[r][c] = i
		r1 := r + directions[index][0]
		c1 := c + directions[index][1]
		if r1 < 0 || r1 >= n || c1 < 0 || c1 >= n || ret[r1][c1] != 0 {
			index = (index + 1) % 4
		}
		r += directions[index][0]
		c += directions[index][1]
	}
	return ret
}

func main() {
	generateMatrix(3)
}
