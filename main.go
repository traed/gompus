package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var name string
var cave []room
var wumpus, player, pit1, pit2, bats1, bats2 thing
var arrows int

// Room
type room struct {
	number      int
	name        string
	connectsTo  []int
	description string
}

func newRoom(nr int) room {
	return room{number: nr}
}

func (r *room) connect(nr int) {
	if contains(r.connectsTo, nr) == -1 {
		r.connectsTo = append(r.connectsTo, nr)
	}
}

func (r *room) describe() {
	if r.description != "" {
		fmt.Println(r.description)
	} else {
		fmt.Printf("Du är i rum %d.\nGångar leder till %v", r.number, r.connectsTo)
	}
}

// Thing
type thing struct {
	name     string
	location *room
}

func newThing(r *room) thing {
	return thing{location: r}
}

func (t *thing) move(to *room) bool {
	if t.validateMove(to) {
		t.location = to
		return true
	}

	return false
}

func (t *thing) validateMove(to *room) bool {
	return contains(t.location.connectsTo, to.number) != -1 || to.number == t.location.number
}

func (t *thing) wakeUp() {
	if rand.Intn(3) != 0 {
		i := rand.Intn(len(cave)) - 1
		t.location = &cave[i]
	}
}

func (t *thing) isHit(r *room) bool {
	return t.location.number == r.number
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

func createCave() {
	cave = make([]room, 20)

	// Create list of rooms
	for i := range cave {
		cave[i].number = i + 1
	}

	// Connect rooms
	for i := range cave {
		if i == 9 {
			cave[i].connect(cave[0].number)
		} else if i == 19 {
			cave[i].connect(cave[10].number)
		} else {
			cave[i].connect(cave[i+1].number)
		}

		if i == 0 {
			cave[i].connect(cave[9].number)
		} else if i == 10 {
			cave[i].connect(cave[19].number)
		} else {
			cave[i].connect(cave[i-1].number)
		}

		if i < 10 {
			cave[i].connect(cave[i+10].number)
			cave[i+10].connect(cave[i].number)
		}
	}
}

func createThings() {
	player = newThing(&cave[0])
	wumpus = newThing(&cave[15])
	bats1 = newThing(&cave[5])
	bats2 = newThing(&cave[11])
	pit1 = newThing(&cave[7])
	pit2 = newThing(&cave[18])
}

func main() {
	createCave()
	createThings()
	arrows = 5

	fmt.Println("Välkommen till grottan, o store jägare.")
	fmt.Println("Du är på jakt efter den mytomspunne Wumpusen.")
	fmt.Println("Du navigera mellan grottans olika rum genom att skriva GÅ <rumsnummer>, tex GÅ 12.")
	fmt.Println("Du kan också välja att skjuta en pil in i ett rum genom att skriva SKJUT <rumsnummer>.")
	fmt.Println("Skriv EXIT för att avsluta spelet.")
	fmt.Println("Lycka till!")

	// Main loop
	for {
		player.location.describe()

		// Is there anything nearby?
		for currentRoom := range player.location.connectsTo {
			if wumpus.location.number == currentRoom {
				fmt.Println("Det luktar Wumpus...")
			}

			if pit1.location.number == currentRoom || pit2.location.number == currentRoom {
				fmt.Println("Jag känner ett luftdrag...")
			}

			if bats1.location.number == currentRoom || bats2.location.number == currentRoom {
				fmt.Println("Jag trampade i något äckligt...")
			}
		}

		// Read input from user
		input := readText()
		commandList := strings.Split(input, " ")
		command := strings.ToUpper(commandList[0])
		var move *room

		if len(commandList) > 1 {
			i, err := strconv.Atoi(commandList[1])
			if err != nil {
				fmt.Println("Va?")
				continue
			}

			move = &cave[i-1]
		} else {
			move = player.location
		}

		switch command {
		case "EXIT":
			fmt.Println("... och jägaren syntes aldrig mera till.")
			os.Exit(0)
		case "MOVE":
			if player.move(move) {
				if player.location.number == wumpus.location.number {
					fmt.Println("Jag känner någe mjukt... det är en WUMPUS!")
					wumpus.wakeUp()
				}
				break
			} else {
				fmt.Println("Du kan inte gå dit från detta rum.")
				continue
			}
		case "SHOOT":
			if player.validateMove(move) {
				fmt.Println("*TWANG*")
				if wumpus.location.number == move.number {
					fmt.Println("Snyggt skjutet! Du träffade Wumpusen.")
					os.Exit(0)
				}
			} else {
				fmt.Println("Du kan inte skjuta genom väggar...")
			}

			wumpus.wakeUp()
			arrows--
			if arrows == 0 {
				fmt.Println("Du har slut på pilar. Bäst att vända hemåt innan Wumpusen hittar dig...")
				os.Exit(0)
			}
		default:
			fmt.Println("Va?")
		}

		switch player.location.number {
		case bats1.location.number:
		case bats2.location.number:
			fmt.Println("Fladdermöss! Varför måste det alltid vara fladdermöss?!")
			fmt.Println("Du lyfts iväg och flyger iväg djupare in i grottan.")
			player.location = &cave[rand.Intn(19)]
			break
		case wumpus.location.number:
			fmt.Println("Jag känner någe mjukt... Det är en WUMPUS! Nom nom nom...")
			os.Exit(0)
		case pit1.location.number:
		case pit2.location.number:
			fmt.Println("Hmm det här ser ut som ett bottenlöst HÅÅÅÅÅÅååååååå...")
			os.Exit(0)
		}
	}
}
