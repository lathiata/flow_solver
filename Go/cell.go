package main

import (
	"errors"
	"strconv"
)

const (
	Empty = -1
)

// The Cell interface is designed so that the state
// doesn't need to know about the x,y coordinates
//
// The Cell should know it's own x,y coords and its value
// Interface also includes helpful methods for
// checking if the Cell is empty and also its string repr
// for printing out the board.

type Cell interface {
	Fill(val int) error
	Val() int
	Empty() bool
	String() string
	Coords() []int
}

type CellImplementation struct {
	val    int
	coords []int
}

func NewCell(x, y int) *CellImplementation {
	return &CellImplementation{
		val:    Empty,
		coords: []int{x, y},
	}
}

func (c *CellImplementation) Val() int {
	return c.val
}

func (c *CellImplementation) Empty() bool {
	return c.val == Empty
}

func (c *CellImplementation) Fill(val int) error {
	if c.val != Empty {
		return errors.New("Can't fill a non-empty cell")
	}
	c.val = val
	return nil
}

func (c *CellImplementation) String() string {
	if c.val == Empty {
		return " "
	}
	return strconv.Itoa(c.val)
}

func (c *CellImplementation) Coords() []int {
	return c.coords
}
