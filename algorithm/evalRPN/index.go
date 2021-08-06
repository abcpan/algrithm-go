package evalRPN

import "strconv"

var Table = map[string] func(a int, b int) int {
	"+": func(a int, b int) int {
		return a + b
	},
	"-": func(a int, b int) int {
		return a -b
	},
	"/": func(a, b int) int {
		return a / b
	},
	"*": func(a, b int) int {
		return a * b
	},
}

func evalRPN(tokens []string) int {
	stack := make([]string, 0)
	for _, token := range tokens {
		if _, isIn := Table[token]; !isIn {
			stack = append(stack, token)
			continue
		}
		pair := make([]int, 0)
		for i := 0; i < 2;i++ {
			size := len(stack)
			item := stack[size - 1]
			stack = stack[:size - 1]
			num, _ := strconv.Atoi(item)
			pair = append(pair,  num)
		}
		ret := Table[token](pair[1], pair[0])
		stack = append(stack, strconv.Itoa(ret))
	}
	result, _ := strconv.Atoi(stack[0])
	return result
}