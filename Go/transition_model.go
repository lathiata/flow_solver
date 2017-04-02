package main

import (
	"log"
)

type TransitionModel interface {
	Explore(s state) []state
}

// This transition model will solve each of the colors
// in order until all the colors are solved
// Must be thread safe
type TransitionModelImplementation struct {
}

// TODO(tanay) keeping track of state here makes this a LOT easier
func (t *TransitionModelImplementation) Explore(s state) []state {
	nextStates := make([]state, 0)
	// First find the next color to solve
	// do this by checking the end spot and seeing if any of the adjacent squares are filled
	var i int
	for i = 0; i < s.Problem().NumColors(); i++ {
		solved := false
		// get the cells adjacent to the "end" coordinate for color i
		coords, err := s.Problem().ColorCoords(i)
		if err != nil {
			log.Fatal(err)
		}
		adjacentCells, err := s.AdjacentCells(coords[1][0], coords[1][1])
		if err != nil {
			log.Fatal(err)
		}
		// check if any of the cells adjacent to end cell are equal to i,
		// if so, that color is solved
		for _, cell := range adjacentCells {
			if cell.Val() == i {
				solved = true
			}
		}
		if !solved {
			break
		}
	}

	return nextStates
}
