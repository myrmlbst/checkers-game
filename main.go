package main

import "fmt"

type PieceType int
type Color int

const (
	Normal PieceType = iota
	Queen
)

const (
	Black Color = iota
	Red
	Empty
)

// piece struct
type Piece struct {
	Type  PieceType
	Color Color
}

// position
type Position struct {
	X, Y int
}

// move
type Move struct {
	Start    Position
	End      Position
	Captures Position
}

// board
type Board struct {
	Grid [8][8]*Piece
	Turn Color
}

func main() {
	fmt.Println("Hello world!")
}
