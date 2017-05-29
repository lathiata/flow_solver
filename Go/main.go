package main

import (
	"flag"
	"log"
)

// TODO(tanay) allow command line flags to pass in parameters
// TODO(tanay) move test game-boards into separate directory

// This file instantiates game boards based on command line
// inputs and runs the indicated algorithm to solve it
func main() {
	var parallelism = flag.Int("parallelism", 1, "Number of threads (goroutines) used to solve board")
	var difficulty = flag.String("difficulty", "easy", "complexity of board {easy, medium, hard, extreme}")
	var trials = flag.Int("trials", 1, "Number of trials (output time is the mean of all trials)")
	flag.Parse()

	if *parallelism <= 0 {
		log.Fatalf("Input parallelism, %d, is not positive", *parallelism)
	}

	if *trials <= 0 {
		log.Fatalf("Input trials, %d, is not positive", *trials)
	}

	s, err := randomState(*difficulty)
	if err != nil {
		log.Fatal(err)
	}

	benchmark(s, *parallelism, *trials)
}
