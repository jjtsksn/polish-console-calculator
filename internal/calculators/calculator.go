package calculators

import (
	"errors"
	"strconv"
	"strings"

	"github.com/jjtsksn/polish-console-calculator/internal/stacks"
)

var stack stacks.Stack

func Calculate(s string) (any, error) {
	defer stack.Clear()
	fields := strings.Fields(s)

	for _, v := range fields {
		if value, err := strconv.Atoi(v); err == nil {
			stack.Push(value)
		} else {
			switch v {
			case "+":
				if err := Add(); err != nil {
					return nil, err
				}
			case "-":
				if err := Remove(); err != nil {
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
	stack = append(stack, a.(int)+b.(int))
	return nil
}

func Remove() error {
	if err := stack.CheckStack(); err != nil {
		return err
	}
	a, b := stack[len(stack)-2], stack[len(stack)-1]
	stack.Pop()
	stack = append(stack, a.(int)-b.(int))
	return nil
}

func Multiply() error {
	if err := stack.CheckStack(); err != nil {
		return err
	}
	a, b := stack[len(stack)-2], stack[len(stack)-1]
	stack.Pop()
	stack = append(stack, a.(int)*b.(int))
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
	stack = append(stack, a.(int)/b.(int))
	return nil
}
