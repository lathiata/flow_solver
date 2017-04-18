package main

import (
	"log"
)

// TODO(tanay) allow command line flags to pass in parameters

// This file instantiates game boards based on command line
// inputs and runs the indicated algorithm to solve it

func main() {
	coords := [][]int{
		[]int{0, 2},
		[]int{3, 0},
		[]int{0, 3},
		[]int{4, 3},
		[]int{1, 3},
		[]int{2, 1},
		[]int{3, 3},
		[]int{4, 4},
		[]int{3, 1},
		[]int{4, 0},
	}
	gridSize := 5
	numColors := 5
	p, err := NewProblem(gridSize, numColors, coords)
	if err != nil {
		log.Fatal(err)
	}
	s := NewState(p)
	log.Printf("Starting flow game simulation with board: %s", s)
	nextStates := s.NextStates()
	log.Printf("Next Level States: %s", nextStates)
	log.Printf("Two levels deep: %s", nextStates[0].NextStates())
	log.Printf("Also two levels deep: %s", nextStates[1].NextStates())
}
