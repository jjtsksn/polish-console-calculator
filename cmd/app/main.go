package main

import (
	"fmt"

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
		calculators.ExecuteCommand(command)
	}
}
