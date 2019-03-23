package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type thing struct {
	name     string
	location *room
}

type room struct {
	number int
}

func readText() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.ToLower(text)
	text = text[:len(text)-1]

	return text
}

// Steg 2: Typer
func main() {
	var player thing
	var room room

	for {
		fmt.Println("Vad heter du?")
		player.name = readText()

		fmt.Println("Var är du?")
		// player.location = readText() Varför funkar inte det? Typer!
		input := readText()
		nr, err := strconv.Atoi(input)

		// Felhantering?
		if err != nil {
			fmt.Println("Det där är inte ett nummer!")
		} else {
			room.number = nr
			// player.location = room Varför funkar inte det? Pointers!
			player.location = &room
			fmt.Printf("Hej %s. Du är i rum %d.\n", player.name, player.location.number)
		}
	}
}
