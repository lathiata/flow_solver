package main

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

	return nextStates
}
