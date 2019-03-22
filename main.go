package main

import (
	"bufio"
	"os"
	"strings"
)

var name string
var cave []room
var wumpus, player, pit1, pit2, bats1, bats2 thing

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
	cave = make([]room, 20)

	// Create list of rooms
	for i := range cave {
		cave[i] = newRoom(i + 1)
	}
}
