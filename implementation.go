package lab2

import (
	"fmt"
	"strings"
)

func pop(_list *[]string) string {
	n := len(*_list)
	last := (*_list)[n-1]
	*_list = (*_list)[:n-1]
	return last
}

// PostfixToPrefix converts reverse polish notation into prefix notation.
func PostfixToPrefix(input string) (string, error) {
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
			stack = append(stack, element)
		}
	}
	return strings.Join(stack, " "), nil

}
