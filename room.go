package main

import "fmt"

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
