package ioe

import (
	"bufio"
	"os"
	"strings"
)

// Read input until newline,
// ignore empty line
func ReadInput() string {
	bs := bufio.NewScanner(os.Stdin)
	for bs.Scan() {
		line := bs.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		return line
	}

	return ""
}

// Allow empty
func ReadInputEmpty() string {
	bs := bufio.NewScanner(os.Stdin)
	for bs.Scan() {
		line := bs.Text()
		line = strings.TrimSpace(line)

		return line
	}

	return ""
}
