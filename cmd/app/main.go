package main

import (
	"fmt"
	"os"

	"github.com/jjtsksn/polish-console-calculator/internal/calculators"
	"github.com/jjtsksn/polish-console-calculator/internal/helpers"
	"github.com/jjtsksn/polish-console-calculator/internal/stacks"
	"github.com/jjtsksn/polish-console-calculator/pkg/clearers"
	"github.com/jjtsksn/polish-console-calculator/pkg/scanners"
)

func main() {
	clearers.ClearTerminal()
	scanner := scanners.NewScanner()
	stack := stacks.Stack{}
	for {
		helpers.StartHelp()
		fmt.Print("Введите команду: ")
		scanner.Scan()
		a := scanner.Text()
		switch a {
		case "exit":
			clearers.ClearTerminal()
			os.Exit(0)
		default:
			if err := calculators.Calculate(&scanner, &stack); err != nil {
				clearers.ClearTerminal()
				fmt.Println(err)
				os.Exit(0)
			}
			clearers.ClearTerminal()
			fmt.Println("Результат: ", stack[0])
			fmt.Println("")
			stack = nil
		}
	}

}
