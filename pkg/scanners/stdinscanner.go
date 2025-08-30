package scanners

import (
	"bufio"
	"os"
)

func NewScanner() bufio.Scanner {
	return *bufio.NewScanner(os.Stdin)
}
