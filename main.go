package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	input := strings.Trim(text, "\n")
	fmt.Println("Balancing string: ", input, " => ", balance(input))
}

func balance(s string) (res string) {
	return balanceRecursive(strings.Split(s, ""), 0, 0)
}

func balanceRecursive(s []string, pos int, open int) (res string) {

	// Prepend opening braces
	for i := open; i < 0; i++ {
		s = append([]string{"{"}, s...)
		open++
		pos++
	}

	// Count open braces and recurse
	if pos < len(s) {
		current := s[pos]
		if current == "{" {
			open++
		} else {
			open--
		}
		pos++
		return balanceRecursive(s, pos, open)
	}

	// Append any closing braces to balance
	for i := 0; i < open; i++ {
		s = append(s, "}")
	}

	return strings.Join(s, "")
}
