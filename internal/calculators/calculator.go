package calculators

import (
	"errors"
	"strings"

	"github.com/jjtsksn/polish-console-calculator/internal/stacks"
	"github.com/shopspring/decimal"
)

var stack stacks.Stack

func Calculate(s string) (any, error) {
	// Очистка стека после получения результата
	defer stack.Clear()
	fields := strings.Fields(s)

	// Обработка данных, введенных пользователем
	for _, v := range fields {
		if value, err := decimal.NewFromString(v); err == nil {
			stack.Push(value)
		} else {
			switch v {
			case "+":
				if err := Add(); err != nil {
					return nil, err
				}
			case "-":
				if err := Subtract(); err != nil {
					return nil, err
				}
			case "*":
				if err := Multiply(); err != nil {
					return nil, err
				}
			case "/":
				if err := Devide(); err != nil {
					return nil, err
				}
			}
		}
	}
	return stack[0], nil
}

func Add() error {
	if err := stack.CheckStack(); err != nil {
		return err
	}
	a, b := stack[len(stack)-2], stack[len(stack)-1]
	stack.Pop()
	stack = append(stack, a.(decimal.Decimal).Add(b.(decimal.Decimal)))
	return nil
}

func Subtract() error {
	if err := stack.CheckStack(); err != nil {
		return err
	}
	a, b := stack[len(stack)-2], stack[len(stack)-1]
	stack.Pop()
	stack = append(stack, a.(decimal.Decimal).Sub(b.(decimal.Decimal)))
	return nil
}

func Multiply() error {
	if err := stack.CheckStack(); err != nil {
		return err
	}
	a, b := stack[len(stack)-2], stack[len(stack)-1]
	stack.Pop()
	stack = append(stack, a.(decimal.Decimal).Mul(b.(decimal.Decimal)))
	return nil
}

func Devide() error {
	if err := stack.CheckStack(); err != nil {
		return err
	}
	a, b := stack[len(stack)-2], stack[len(stack)-1]
	if b == 0 {
		return errors.New("нельзя делить на 0")
	}
	stack.Pop()
	stack = append(stack, a.(decimal.Decimal).Div(b.(decimal.Decimal)))
	return nil
}
