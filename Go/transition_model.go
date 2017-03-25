package main

type TransitionModel interface {
	Explore(board Board) []Board
}

// This transition model will solve each of the colors
// in order until all the colors are solved
// Must be thread safe
type TransitionModelImplementation struct {

}