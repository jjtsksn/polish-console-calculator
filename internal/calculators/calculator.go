package calculators

import (
	"bufio"
	"errors"
	"strconv"
	"strings"

	"github.com/jjtsksn/polish-console-calculator/internal/stacks"
)

func Calculate(scanner *bufio.Scanner, stack *stacks.Stack) error {
	fields := strings.Fields(scanner.Text())
	if len(fields) < 3 || !IsValid(&fields) {
		return errors.New("проверка на валидность не пройдена")
	}

	for _, v := range fields {
		if value, err := strconv.Atoi(v); err == nil {
			stack.Push(value)
		} else {
			if k, err := Check(v, stack); err == nil {
				stack.Push(k)
			} else {
				return err
			}
		}
	}
	return nil
}

func Check(s string, stack *stacks.Stack) (int, error) {
	switch s {
	case "+":
		if a, b, err := stack.Pop(); err == nil {
			return Add(a, b), nil
		} else {
			return 0, err
		}
	case "-":
		if a, b, err := stack.Pop(); err == nil {
			return Remove(a, b), nil
		} else {
			return 0, err
		}
	case "*":
		if a, b, err := stack.Pop(); err == nil {
			return Multiply(a, b), nil
		} else {
			return 0, err
		}
	case "/":
		if a, b, err := stack.Pop(); err == nil {
			if x, err := Devide(a, b); err == nil {
				return x, nil
			}
			return 0, err
		} else {
			return 0, err
		}
	default:
		return 0, errors.New("неопределенный оператор")
	}
}

func Add(a, b int) int {
	return a + b
}

func Remove(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Devide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("нельзя делить на 0")
	}
	return a / b, nil
}

func IsValid(fields *[]string) bool {
	var operators int
	var operands int
	for _, v := range *fields {
		switch v {
		case "+", "-", "*", "/":
			operators++
		default:
			operands++
		}
	}
	if operands-operators == 1 {
		return true
	} else {
		return false
	}
}
