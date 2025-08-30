package stacks

import "errors"

type Stack []int

func (s *Stack) Push(a int) {
	*s = append(*s, a)
}

func (s *Stack) Pop() (int, int, error) {
	if s.IsEmpty() {
		return 0, 0, errors.New("в стеке нет переменных")
	}
	if len(*s) < 2 {
		return 0, 0, errors.New("в стеке не хватает элементов")
	}
	a := (*s)[len(*s)-2]
	b := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-2]
	return a, b, nil
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
