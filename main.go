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

// create the board
func NewBoard() *Board {
    board := &Board{
        Turn: Black, // black is always the staring piece
    }

    // fill up board with black pieces
    for y:= 0; y < 3; y++ {
        for x:= 0; x < 8; x++ {
            if ((x+y) % 2) != 0 {
                board.Grid[y][x] = &Piece{
                    Type: Normal,
                    Color: Black,
                }
            }
        }
    }

    // fill up board with red pieces
    for y:= 0; y < 3; y++ {
        for x:= 0; x < 8; x++ {
            if ((x+y) % 2) != 0 {
                board.Grid[y][x] = &Piece{
                    Type: Normal,
                    Color: Red,
                }
            }
        }
    }

    return board
}


func (b *Board) VisualizeBoard() {
    for y:= 0; y < 8; y++ {
        fmt.Printf("%d |", y)
        for x:= 0; x < 8; x++ {
            // check if it is empty
            // what color is that space
            if b.Grid[y][x] == nil {
                if (x+y) % 2 == 0 {
                    fmt.Print(" . |")
                } else {
                    fmt.Print("   |")
                }
            } else {
                // if the space is not empty:
                // check what color piece it is
                // check whether piece is normal or queen

                // this space is NOT empty
                piece := b.Grid[y][x]

                if piece.Color == Black {
                    if piece.Type == Queen {
                        // black ,queen
                        fmt.Print(" B |")
                    } else {
                        // black, regular
                        fmt.Print(" b |")
                    }
                } else {
                    if piece.Type == Queen {
                        // red, queen
                        fmt.Print(" R |")
                    } else {
                        // red, regular
                        fmt.Print(" r |")
                    }
                }
            }
        }

        fmt.Println("\n  +---+---+---+---+---+---+---+---+")
    }
}


func main() {
    board := NewBoard()
    board.VisualizeBoard()

}
