package clearers

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// ClearTerminal очищает консоль в зависимости от операционной системы
func ClearTerminal() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default: // Unix-like системы: Linux, macOS, BSD
		cmd = exec.Command("clear")
	}

	// Устанавливаем вывод команды в stdout
	cmd.Stdout = os.Stdout

	// Выполняем команду и обрабатываем возможные ошибки
	if err := cmd.Run(); err != nil {
		// Fallback: если команда очистки не сработала,
		// используем вывод escape-последовательности
		clearWithEscapeSequence()
	}
}

// clearWithEscapeSequence резервный метод очистки через ANSI escape codes
func clearWithEscapeSequence() {
	// ANSI escape code для очистки экрана и перемещения курсора в начало
	const clearCode = "\033[H\033[2J"

	fmt.Print(clearCode)
	fmt.Fprint(os.Stdout, clearCode)
}
