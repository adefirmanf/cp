package sharedpkg

import (
	"bufio"
	"os"
	"strings"
)

func Scanner() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	text := scanner.Text()
	parts := strings.Fields(text)
	return parts
}
