package main

import (
	"fmt"
	"os"

	"github.com/jjtsksn/polish-console-calculator/internal/calculators"
	"github.com/jjtsksn/polish-console-calculator/internal/helpers"
	"github.com/jjtsksn/polish-console-calculator/pkg/clearers"
	"github.com/jjtsksn/polish-console-calculator/pkg/scanners"
)

func main() {
	clearers.ClearTerminal()
	scanner := scanners.NewScanner()

	for {
		helpers.StartHelp()
		fmt.Print("Введите комманду: ")
		scanner.Scan()
		command := scanner.Text()
		executeCommand(command)
	}
}

// Функция принимает комманду в качестве аргумента и выполняет
// код в зависимости от того, какую команду ввел пользователь
func executeCommand(command string) {
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
