package main

import (
	"log"
)

// TODO(tanay) allow command line flags to pass in parameters

// This file instantiates game boards based on command line
// inputs and runs the indicated algorithm to solve it

func main() {
	xCoords := []int{0, 3, 0, 4, 1, 2, 3, 4, 3, 4}
	yCoords := []int{2, 0, 3, 3, 3, 1, 3, 4, 1, 0}
	gridSize := 5
	numColors := 5
	p := NewProblem(gridSize, numColors, xCoords, yCoords)
	err := p.Validate()
	if err != nil {
		log.Fatal(err)
	}
	b := NewBoard(p)
	log.Printf("Starting flow game simulation with board:\n %s", b.String())
}
