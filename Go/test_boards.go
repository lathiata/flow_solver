package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func benchmark(s state, threads, trials int) {
	var solution state
	log.Printf("Start state: %s", s)
	totalTime := float64(0)
	for i := 0; i < trials; i++ {
		startTime := time.Now()
		coordinator := NewCoordinator(s, threads)
		solution = coordinator.Solve()
		elapsedTime := time.Since(startTime)
		totalTime += float64(elapsedTime) / float64(1000000000) //nanoseconds->seconds
	}
	avgTime := totalTime / float64(trials)

	log.Printf("Solution: %s", solution)
	log.Printf("For %d threads, average elapsed time is: %f over %d trials", threads, avgTime, trials)
}

func randomState(difficulty string) (state, error) {
	var states []state
	switch difficulty {
	case "easy":
		states = easyStates()
	case "medium":
		states = mediumStates()
	case "hard":
		states = hardStates()
	case "extreme":
		states = extremeStates()
	default:
		return nil, errors.New(fmt.Sprintf("invalid difficulty string %s. Must be one of"+
			" {easy, medium, hard, extreme}.", difficulty))
	}

	randomState := states[rand.Intn(len(states))]
	return randomState, nil
}

func easyStates() []state {
	easyStates := make([]state, 0)
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
	easyProblem, err := NewProblem(gridSize, numColors, coords)
	if err != nil {
		log.Fatal(err)
	}
	easyStates = append(easyStates, NewState(easyProblem))
	return easyStates
}

func mediumStates() []state {
	mediumStates := make([]state, 0)
	coords := [][]int{
		[]int{0, 0},
		[]int{3, 0},
		[]int{4, 0},
		[]int{7, 6},
		[]int{2, 1},
		[]int{3, 5},
		[]int{2, 2},
		[]int{1, 6},
		[]int{2, 3},
		[]int{3, 6},
		[]int{4, 1},
		[]int{6, 1},
		[]int{3, 1},
		[]int{4, 4},
		[]int{4, 3},
		[]int{6, 6},
		[]int{5, 3},
		[]int{7, 7},
	}
	gridSize := 8
	numColors := 9
	mediumProblem, err := NewProblem(gridSize, numColors, coords)
	if err != nil {
		log.Fatal(err)
	}
	mediumState := NewState(mediumProblem)
	mediumStates = append(mediumStates, mediumState)
	return mediumStates
}
func hardStates() []state {
	hardStates := make([]state, 0)
	coords := [][]int{
		[]int{0, 0},
		[]int{5, 2},
		[]int{1, 1},
		[]int{3, 0},
		[]int{1, 2},
		[]int{2, 6},
		[]int{1, 6},
		[]int{5, 4},
		[]int{4, 0},
		[]int{5, 3},
	}
	gridSize := 8
	numColors := 5
	hardProblem, err := NewProblem(gridSize, numColors, coords)
	if err != nil {
		log.Fatal(err)
	}
	hardState := NewState(hardProblem)
	hardStates = append(hardStates, hardState)
	return hardStates
}

func extremeStates() []state {
	extremeStates := make([]state, 0)
	coords := [][]int{
		[]int{0, 8}, //red
		[]int{3, 3},
		[]int{1, 8}, //green
		[]int{8, 0},
		[]int{2, 8}, //yellow
		[]int{5, 2},
		[]int{1, 4}, //orange
		[]int{3, 1},
		[]int{3, 2}, //blue
		[]int{7, 7},
	}
	gridSize := 9
	numColors := 5
	extremeProblem, err := NewProblem(gridSize, numColors, coords)
	if err != nil {
		log.Fatal(err)
	}
	extremeState := NewState(extremeProblem)
	extremeStates = append(extremeStates, extremeState)
	return extremeStates
}
