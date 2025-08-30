package calculators

import (
	"errors"
	"strings"

	"github.com/jjtsksn/polish-console-calculator/internal/stacks"
	"github.com/shopspring/decimal"
)

var stack = stacks.New()

func Calculate(s string) (decimal.Decimal, error) {
	// Очистка стека после получения результата
	defer stack.Clear()

	// Проверка на пустой ввод
	fields := strings.Fields(s)
	if len(fields) == 0 {
		return decimal.Decimal{}, errors.New("пустой ввод")
	}

	// Обработка данных, введенных пользователем
	for _, v := range fields {
		if value, err := decimal.NewFromString(v); err == nil {
			stack.Push(value)
		} else {
			if err := handleOperation(v); err != nil {
				return decimal.Decimal{}, err
			}
		}
	}

	if stack.IsEmpty() {
		return decimal.Decimal{}, errors.New("нет результата")
	}
	if stack.HasFewElements() {
		return decimal.Decimal{}, errors.New("в стеке остались лишние значения")
	}
	if result, err := stack.Get(0); err != nil {
		return decimal.Decimal{}, err
	} else {
		return result, nil
	}
}

// handleOperation определяет операцию, которую необходимо произвести
func handleOperation(op string) error {
	switch op {
	case "+":
		return add()
	case "-":
		return subtract()
	case "*":
		return multiply()
	case "/":
		return divide()
	default:
		return errors.New("неизвестный оператор: " + op)
	}
}

// executeBinaryOperation выполняет бинарную операцию
func executeBinaryOperation(operation func(a, b decimal.Decimal) decimal.Decimal) error {
	if err := stack.CheckStack(); err != nil {
		return err
	}
	if a, b, err := stack.PopTwo(); err != nil {
		return err
	} else {
		result := operation(a, b)
		stack.Push(result)
		return nil
	}
}

// executeBinaryOperationWithCheck выполняет бинарную операцию
// дополнительно проверяя, равен ли делитель 0 для операции деления
func executeBinaryOperationWithCheck(operation func(a, b decimal.Decimal) decimal.Decimal,
	check func(b decimal.Decimal) error) error {
	if err := stack.CheckStack(); err != nil {
		return err
	}
	if a, b, err := stack.PopTwo(); err != nil {
		return err
	} else {
		if check != nil {
			if err := check(b); err != nil {
				return err
			}
		}
		result := operation(a, b)
		stack.Push(result)
		return nil
	}
}

func add() error {
	return executeBinaryOperation(func(a, b decimal.Decimal) decimal.Decimal {
		return a.Add(b)
	})
}

func subtract() error {
	return executeBinaryOperation(func(a, b decimal.Decimal) decimal.Decimal {
		return a.Sub(b)
	})
}

func multiply() error {
	return executeBinaryOperation(func(a, b decimal.Decimal) decimal.Decimal {
		return a.Mul(b)
	})
}

func divide() error {
	return executeBinaryOperationWithCheck(func(a, b decimal.Decimal) decimal.Decimal {
		return a.Div(b)
	}, func(b decimal.Decimal) error {
		if b.IsZero() {
			return errors.New("нельзя делить на 0")
		}
		return nil
	})
}
