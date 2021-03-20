package leet

import (
	"container/list"
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := list.New()
	for _, t := range tokens {
		i, err := strconv.Atoi(t)
		if err == nil {
			stack.PushBack(i)
		} else {
			e1 := stack.Back()
			stack.Remove(e1)
			e2 := stack.Back()
			stack.Remove(e2)
			i1, _ := e1.Value.(int)
			i2, _ := e2.Value.(int)
			switch t {
			case "+":
				stack.PushBack(i2 + i1)
			case "-":
				stack.PushBack(i2 - i1)
			case "*":
				stack.PushBack(i2 * i1)
			case "/":
				stack.PushBack(i2 / i1)
			default:
				fmt.Errorf("Unknown operation. ")
			}
		}
	}
	return stack.Back().Value.(int)
}
