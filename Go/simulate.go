package main

import (
	"log"
	"time"
)

// TODO(tanay) allow command line flags to pass in parameters

// This file instantiates game boards based on command line
// inputs and runs the indicated algorithm to solve it
func simulate(s state) {
	log.Printf("Start state: %s", s)
	startTime := time.Now()
	coordinator := NewCoordinator(s)
	solution := coordinator.Solve()
	elapsedTime := time.Since(startTime)
	log.Printf("Solution: %s in %f sec", solution, float64(elapsedTime)/float64(1000000000))

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
	p, err := NewProblem(gridSize, numColors, coords)
	if err != nil {
		log.Fatal(err)
	}
	s := NewState(p)
	simulate(s)
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
	p, err = NewProblem(gridSize, numColors, coords)
	if err != nil {
		log.Fatal(err)
	}
	s = NewState(p)
	simulate(s)
	//hard
	//coords = [][]int{
	//	[]int{0, 0},
	//	[]int{5, 2},
	//	[]int{1, 1},
	//	[]int{3, 0},
	//	[]int{1, 2},
	//	[]int{2, 6},
	//	[]int{1, 6},
	//	[]int{5, 4},
	//	[]int{4, 0},
	//	[]int{5, 3},
	//}
	//gridSize = 8
	//numColors = 5
	//p, err = NewProblem(gridSize, numColors, coords)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//s = NewState(p)
	//simulate(s)
	//// very hard
	//coords = [][]int{
	//	[]int{0, 8}, //red
	//	[]int{3, 3},
	//	[]int{1, 8}, //green
	//	[]int{8, 0},
	//	[]int{2, 8}, //yellow
	//	[]int{5, 2},
	//	[]int{1, 4}, //orange
	//	[]int{3, 1},
	//	[]int{3, 2}, //blue
	//	[]int{7, 7},
	//}
	//gridSize = 9
	//numColors = 5
	//p, err = NewProblem(gridSize, numColors, coords)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//s = NewState(p)
	//simulate(s)
}
