package stacks

import (
	"errors"
)

type Stack []any

func (s *Stack) Push(a any) {
	*s = append(*s, a)
}

func (s *Stack) Pop() error {
	*s = (*s)[:len(*s)-2]
	return nil
}

func (s *Stack) GetTop() any {
	return (*s)[0]
}

func (s *Stack) Clear() {
	*s = nil
}

func (s *Stack) CheckStack() error {
	if len(*s) < 2 {
		return errors.New("длина стека не может быть меньше 2")
	}
	switch (*s)[len(*s)-2].(type) {
	case int:
	default:
		return errors.New("перед оператором должны быть 2 операнда")
	}
	switch (*s)[len(*s)-1].(type) {
	case int:
	default:
		return errors.New("перед оператором должны быть 2 операнда")
	}
	return nil
}
