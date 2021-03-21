package sort

func bubble(src []int) []int {
	l := len(src)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			if src[j] > src[j+1] {
				temp := src[j]
				src[j] = src[j+1]
				src[j+1] = temp
			}
		}
	}
	return src
}

func selection(src []int) []int {
	l := len(src)
	var minIndex int
	for i := 0; i < l; i++ {
		minIndex = i
		for j := i + 1; j < l; j++ {
			if src[minIndex] > src[j] {
				minIndex = j
			}
		}
		temp := src[i]
		src[i] = src[minIndex]
		src[minIndex] = temp
	}
	return src
}
