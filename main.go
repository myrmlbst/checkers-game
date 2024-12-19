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
    board:= &Board{
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
                piece:= b.Grid[y][x]

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


// check if moves are valid
func (b *Board) IsValidMove(move Move) (bool, string) {
    // utility functions:

    if !isInBounds(move.Start) || !isInBounds(move.End) {
        return false, "This move is out of bounds"
    }

    piece:= b.Grid[move.Start.Y][move.Start.X]
    if piece == nil {
        return false, "No piece at starting position"
    }

    // player should not land on an occupied piece
    if b.Grid[move.End.Y][move.End.X] != nil {
        return false, "Destination square is already occupied"
    }

    // check turns
    if piece.Color != b.Turn {
        return false, fmt.Sprintf("%s's turn", colorToString(b.Turn))
    }

    // move distances
    dx:= move.End.X - move.Start.X
    dy:= move.End.Y - move.Start.Y

    if abs(dx) != abs(dy) {
        return false, "Move must be diagonal"
    }

    if piece.Type == Normal {
        // ensure that no piece is going backwards
        if piece.Color == Black && dy <= 0 {
            return false, "Normal black pieces cannot go upwards"
        }
        if piece.Color == Red && dy >= 0 {
            return false, "Normal red pieces cannot go downwards"
        }
    }

    if abs(dx) == 1 && abs(dy) == 1 {
        // valid, regular move
        return true, ""
    } else if abs(dx) == 2 && abs(dy) == 2 {
        // check for a valid capture
        captureX:= move.Start.X + dx/2
        captureY:= move.Start.Y + dy/2

        capturedPiece:= b.Grid[captureY][captureX]

        // check if capturedPiece actually exists
        if capturedPiece == nil {
            return false, "No piece to be captured"
        }
        if capturedPiece.Color == piece.Color {
            return false, "You may capture your own piece"
        }
        return true, "" // valid capture move
    }

    return false, "Invalid move distance"
}


// what is processed when user inputs arguments into CLI
func (b *Board) MakeMove(move Move) bool {
    // check move validity
    valid, reason:= b.IsValidMove(move)
    if !valid {
        fmt.Printf("Invalid move: %s\n", reason)
        return false
    }

    // execute move
    piece:= b.Grid[move.Start.Y][move.Start.X]
    b.Grid[move.End.Y][move.End.X] = piece
    b.Grid[move.Start.Y][move.Start.X] = nil

    // handle captures
    if abs(move.End.X - move.Start.X) == 2 {
        captureX:= (move.Start.X + move.End.X) / 2
        captureY:= (move.Start.Y + move.End.Y) / 2
        capturedPiece:= b.Grid[captureY][captureX]

        b.Grid[captureY][captureX] = nil // piece no longer on this square
        fmt.Printf("Captured %s piece at %d, %d", colorToString(capturedPiece.Color), captureX, captureY)
    }

    // handle piece promotions
    if piece.Type == Normal {
        if (piece.Color == Black && move.End.Y == 7) || (piece.Color == Red && move.End.Y == 0) {
            piece.Type = Queen
            fmt.Printf("%s piece promoted to Queen at %d, %d!\n", colorToString(piece.Color), move.End.X, move.End.Y)
        }
    }

    // switch turns
    b.Turn = opponent(b.Turn)

    return true
}


// check if move is inside the boundaries of the board
func isInBounds(position Position) bool {
    return position.X >= 0 && position.X < 8 && position.Y >= 0 && position.Y < 8
}


// helper function to check whose move it is
func colorToString(c Color) string {
    if c == Black {
        return "Black"
    } else {
        return "Red"
    }
}


// helper function to check if moves are diagonal
func abs(val int) int {
    if val < 0 {
        return -val
    } else {
        return val
    }
}


// helper function to find out who is the next opponent
func opponent(c Color) Color {
    if c == Black {
        return Red
    } else {
        return Black
    }
}


func GameLoop(board *Board) {
    scanner:= bufio.NewScanner(os.Stdin) // reads from std input (cli)

    fmt.Println("\nWelcome to Checkers!")
    fmt.Println("Instructions:")
    fmt.Println("- Black pieces (b/B) move down the board")
    fmt.Println("- Red pieces (r/R) move up the board")
    fmt.Println("- Uppercase letters (B/R) represent Queens")
    fmt.Println("- Enter moves in format: START_X, START_Y, END_X, END_Y")
    fmt.Println("- Example: '2 2 3 3' moves from {2,2} to {3,3}")

    for {
        board.VisualizeBoard() // updated tracker of what the board looks like (it updates with every move)

        fmt.Printf("Enter move (start X start Y end X end Y):")
        if !scanner.Scan() {
            return
        }

        input:= scanner.Text()
        coords:= strings.Split(input, " ")

        if len(coords) != 4 {
            fmt.Println("Invalid input format")
            continue
        }

        numbers:= make([]int, 4)
        valid:= true

        for i, coord := coords {
            num. err := strconv.Atoi(coord)
            if err != nil || num < 0 || num > 7 {
                fmt.Println("Invalid input. Please enter numbers between 0 and 7")
                valid = false
                break
            }
            numbers[i] = num
        }

        if !valid {
            continue
        }

        move:= Move{
            Start: Position{X: numbers[0], Y: numbers[1]},
            End: Position{X: numbers[2], Y: numbers[3]},
        }

        board.MakeMove(move)
    }
}


func main() {
    board:= NewBoard()
    GameLoop(board)
}
