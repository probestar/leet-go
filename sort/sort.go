package sort

func bubble(src []int) []int {
	n := len(src)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
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
	n := len(src)
	var minIndex int
	for i := 0; i < n; i++ {
		minIndex = i
		for j := i + 1; j < n; j++ {
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

func insertion(src []int) []int {
	n := len(src)
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if src[j] > src[j+1] {
				temp := src[j+1]
				src[j+1] = src[j]
				src[j] = temp
			} else {
				break
			}
		}
	}
	return src
}

func shell(src []int) []int {
	n := len(src)
	for i := n / 2; n > 0; n /= 2 {
		for j := 0; j < i; j++ {
			if src[j] > src[j+i] {
				temp := src[j]
				src[j] = src[j+i]
				src[j+i] = temp
			}
		}
	}
	return src
}
