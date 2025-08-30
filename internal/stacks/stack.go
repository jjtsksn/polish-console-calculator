package stacks

import (
	"errors"

	"github.com/shopspring/decimal"
)

type stack struct {
	items []decimal.Decimal
}

// New создает новый стек
func New() *stack {
	return &stack{items: make([]decimal.Decimal, 0)}
}

// Push добавляет элемент в стек
func (s *stack) Push(value decimal.Decimal) {
	s.items = append(s.items, value)
}

// PopTwo удаляет и возвращает 2 последних элемента из стека
func (s *stack) PopTwo() (decimal.Decimal, decimal.Decimal, error) {
	if len(s.items) < 2 {
		return decimal.Decimal{}, decimal.Decimal{}, errors.New("в стеке не хватает элементов")
	}

	b := s.items[len(s.items)-1]
	a := s.items[len(s.items)-2]
	s.items = s.items[:len(s.items)-2]
	return a, b, nil
}

// Get возвращает элемент по индекску
func (s *stack) Get(index int) (decimal.Decimal, error) {
	if index < 0 || index >= len(s.items) {
		return decimal.Decimal{}, errors.New("индекс вне диапазона")
	}
	return s.items[index], nil
}

// Clear очищает стек
func (s *stack) Clear() {
	s.items = nil
}

// Size возвращает размер стека
func (s *stack) Size() int {
	return len(s.items)
}

// IsEmpty проверят, пустой ли стек
func (s *stack) IsEmpty() bool {
	return len(s.items) == 0
}

// HasFewElements проверяет, что в стеке больше 1 элемента
func (s *stack) HasFewElements() bool {
	return len(s.items) > 1
}

// CheckStack проверяет, что в стеке не менее 2 элементов
// типа decimal.Decimal
func (s *stack) CheckStack() error {
	if len(s.items) < 2 {
		return errors.New("для операции необходимы 2 операнда в стеке")
	}
	return nil
}
