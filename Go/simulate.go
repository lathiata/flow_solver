package main

import (
	"log"
	"time"
)

var (
	possibleNumThreads  = []int{1, 2, 3, 5}
	trialsPerThreadSize = 20
)

// TODO(tanay) allow command line flags to pass in parameters
// TODO(tanay) move test game-boards into separate directory

// This file instantiates game boards based on command line
// inputs and runs the indicated algorithm to solve it
func simulate(s state) {
	log.Printf("Start state: %s", s)
	startTime := time.Now()
	coordinator := NewCoordinator(s, 10)
	solution := coordinator.Solve()
	elapsedTime := time.Since(startTime)
	log.Printf("Solution: %s in %f sec", solution, float64(elapsedTime)/float64(1000000000))
}

func benchmark(s state) {
	var solution state
	log.Printf("Start state: %s", s)
	averageTimes := make([]float64, len(possibleNumThreads))
	for i, numThreads := range possibleNumThreads {
		totalTime := float64(0)
		for j := 0; j < trialsPerThreadSize; j++ {
			startTime := time.Now()
			coordinator := NewCoordinator(s, numThreads)
			solution = coordinator.Solve()
			elapsedTime := time.Since(startTime)
			totalTime += float64(elapsedTime) / float64(1000000000) //nanoseconds->seconds
		}
		averageTimes[i] = totalTime / float64(trialsPerThreadSize)
	}
	log.Printf("Solution: %s", solution)
	for i, numThreads := range possibleNumThreads {
		log.Printf("For %d threads, average elapsed time is: %f", numThreads, averageTimes[i])
	}
}
func main() {
	//easy
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
	easyState := NewState(easyProblem)
	simulate(easyState)
	//medium
	coords = [][]int{
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
	gridSize = 8
	numColors = 9
	mediumProblem, err := NewProblem(gridSize, numColors, coords)
	if err != nil {
		log.Fatal(err)
	}
	mediumState := NewState(mediumProblem)
	simulate(mediumState)
	//hard
	coords = [][]int{
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
	gridSize = 8
	numColors = 5
	hardProblem, err := NewProblem(gridSize, numColors, coords)
	if err != nil {
		log.Fatal(err)
	}
	hardState := NewState(hardProblem)
	simulate(hardState)
	// very hard
	coords = [][]int{
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
	gridSize = 9
	numColors = 5
	veryHardProblem, err := NewProblem(gridSize, numColors, coords)
	if err != nil {
		log.Fatal(err)
	}
	veryHardState := NewState(veryHardProblem)
	simulate(veryHardState)
}
