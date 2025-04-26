# Game of Checkers (in Golang)
This repository contains the source code for a multiplayer checkers/draughts game that I built while learning Golang. The game is playable through the CLI: the program outputs an 8x8 grid and contains core game logic and mechanisms to play full rounds of Checkers.
## Game Layout
24 pieces (12 red, 12 black) are positioned on an 8x8 grid that has black and white alternating cells. In this program, empty boxes represent a white cell, and boxes with a dot in the middle ```[ . ]``` represent black cells.

The pieces are symbolized as follows:
- Black, regular: b
- Black, queen: B
- Red, regular: r
- Red, queen: R
## Rules of the Game
1. All pieces shall remain on the lighter squares of the board.
2. Pieces can only move diagonally.
3. Pieces may not land on an occupied cell.
4. Only one of the current player’s checkers may move per turn.
5. Regular pieces may only move diagonally away from the end of the board from which they started.
6. Queens may move in any diagonal direction.
7. For a piece to be promoted to Queen, it must first move to its opponent's farthermost edge of the game board. The piece does not obtain a Queen's privileges until it has been placed on cell 0 or 7. **Regular pieces that jump into position to become Queens are not treated as Queens until the next turn.**

In order to win, you need to block your opponent from being able to make a move. That can be accomplished via:
1. Capturing all their pieces
2. Trapping all their pieces (meaning that they have no legal moves) [Note: This feature has not been implemented yet.]