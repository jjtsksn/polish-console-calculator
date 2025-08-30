package stacks

import (
	"errors"
)

type Stack []any

// Добавить сущность в стек
func (s *Stack) Push(a any) {
	*s = append(*s, a)
}

// Срезать 2 сущности с конца стека
func (s *Stack) Pop() error {
	*s = (*s)[:len(*s)-2]
	return nil
}

// Очистить стека
func (s *Stack) Clear() {
	*s = nil
}

// Проверка стека перед операцией на:
// 1: В стеке не должно быть меньше чем 2 сущности
// 2: 2 последние сущности в стеке должны быть типа int
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
