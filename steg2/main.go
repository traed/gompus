package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readText() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.ToLower(text)
	text = text[:len(text)-1]

	return text
}

// Steg 2: For och If
func main() {
	var player string
	done := false

	for !done {
		fmt.Println("Vad heter du?")
		player = readText()

		if player == "exit" {
			done = true
		} else {
			fmt.Println("Hej " + player)
		}
	}
}
