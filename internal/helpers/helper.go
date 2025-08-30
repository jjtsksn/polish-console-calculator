package helpers

import (
	"fmt"
	"os"

	"github.com/jjtsksn/polish-console-calculator/internal/calculators"
	"github.com/jjtsksn/polish-console-calculator/pkg/clearers"
)

func StartHelp() {
	fmt.Println("---------------------------------------------------------------")
	fmt.Println("Калькулятор работает ТОЛЬКО с целыми числами!")
	fmt.Println("Для ввода операндов и операторов используйте пробел между ними.")
	fmt.Println("Для выхода из программы введите: exit")
	fmt.Println("---------------------------------------------------------------")
	fmt.Println()
}

func ExecuteCommand(command string) {
	switch command {
	case "exit":
		clearers.ClearTerminal()
		os.Exit(0)
	default:
		clearers.ClearTerminal()
		if ans, err := calculators.Calculate(command); err != nil {
			fmt.Println(err)
			os.Exit(0)
		} else {
			fmt.Println("Резльтат:", ans)
			fmt.Println("")
		}
	}
}
