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

// Steg 1: Hello player name, variabler och funktioner
func main() {
	var player string

	player = readText()

	fmt.Println("Hej " + player)
}
