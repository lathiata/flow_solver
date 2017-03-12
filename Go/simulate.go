package main

import (
	"log"
)

// TODO(tanay) allow command line flags to pass in parameters

// This file instantiates game boards based on command line
// inputs and runs the indicated algorithm to solve it

func main() {
	b := NewBoard(5)
	log.Printf("Starting flow game simulation with board: %v", b)
}
