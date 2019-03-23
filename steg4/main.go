package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cave []room
var wumpus, player, pit1, pit2, bats1, bats2 thing
var arrows int

// Room
type room struct {
	number     int
	connectsTo []int
}

// Thing
type thing struct {
	location *room
}

// Util
func contains(arr []int, nr int) int {
	for i, x := range arr {
		if x == nr {
			return i
		}
	}

	return -1
}

func readText() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.ToLower(text)
	text = text[:len(text)-1]

	return text
}

func main() {
	// Init cave and things
	// createCave()
	// createThings()
	// arrows = 5

	fmt.Println("Välkommen till grottan, o store jägare.")

	// Main loop
	for {
		// Descerible the room you're in
		// player.location.describe()

		// Is there anything nearby?

		// Read input from user
		command := ""

		// Verify commands

		// Do action
		switch command {
		case "EXIT":

		case "MOVE":

		case "SHOOT":

		default:
			fmt.Println("Va?")
		}

		// See what happened
		switch player.location.number {
		case bats1.location.number:
		case bats2.location.number:

		case wumpus.location.number:

		case pit1.location.number:
		case pit2.location.number:

		}
	}
}
