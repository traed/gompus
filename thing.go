package main

import (
	"math/rand"
)

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
	return contains(to.connectsTo, to.number) != -1 || to.number == t.location.number
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
