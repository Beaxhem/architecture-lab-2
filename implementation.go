package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

func pop(_list *[]string) string {
	n := len(*_list)
	last := (*_list)[n-1]
	*_list = (*_list)[:n-1]
	return last
}

// Converts postfix notation into prefix notation.
func PostfixToPrefix(input string) (string, error) {
	if len(input) == 0 {
		return "", fmt.Errorf("empty string")
	}
	elements := strings.Split(input, " ")
	stack := make([]string, 0)
	for _, element := range elements {
		if element == "+" || element == "-" || element == "*" || element == "/" {
			if len(stack) < 2 {
				return "", fmt.Errorf("not enough operands")
			}
			op2 := pop(&stack)
			op1 := pop(&stack)
			stack = append(stack, element+" "+op1+" "+op2)
		} else {
			if _, err := strconv.Atoi(element); err != nil {
				return "", fmt.Errorf("%v is not a number", element)
			}
			stack = append(stack, element)
		}
	}
	if len(stack) != 1 {
		return "", fmt.Errorf("not enough operands")
	}
	return strings.Join(stack, " "), nil
}
