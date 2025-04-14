package cheServer

import (
	pb "chess-engine/gen"
	"fmt"
)

type Piece string

const (
	Empty       Piece = ""
	WhitePawn   Piece = "P"
	WhiteRook   Piece = "R"
	WhiteKnight Piece = "N"
	WhiteBishop Piece = "B"
	WhiteQueen  Piece = "Q"
	WhiteKing   Piece = "K"
	BlackPawn   Piece = "p"
	BlackRook   Piece = "r"
	BlackKnight Piece = "n"
	BlackBishop Piece = "b"
	BlackQueen  Piece = "q"
	BlackKing   Piece = "k"
)

type Board [8][8]Piece

func InitBoard() Board {
	var board Board

	// Place black pieces (rank 8 and 7)
	board[0] = [8]Piece{BlackRook, BlackKnight, BlackBishop, BlackQueen, BlackKing, BlackBishop, BlackKnight, BlackRook}
	board[1] = [8]Piece{BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn}

	// Place white pieces (rank 1 and 2)
	board[6] = [8]Piece{WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn}
	board[7] = [8]Piece{WhiteRook, WhiteKnight, WhiteBishop, WhiteQueen, WhiteKing, WhiteBishop, WhiteKnight, WhiteRook}

	return board
}

func oppositeColor(color pb.Color) pb.Color {
	if color == pb.Color_WHITE {
		return pb.Color_BLACK
	}
	return pb.Color_WHITE
}

func copyBoard(board Board) Board {
	var newBoard Board
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			newBoard[r][c] = board[r][c]
		}
	}
	return newBoard
}
func coordsToSquare(row, col int) string {
	file := string(rune('a' + col))  // convert to column letter (a-h)
	rank := fmt.Sprintf("%d", 8-row) // "8" through "1"
	return file + rank
}

func squareToCoords(square string) (row int, col int, err error) {
	if len(square) != 2 {
		return 0, 0, fmt.Errorf("invalid square: %s", square)
	}

	file := square[0] // column letter
	rank := square[1] // row number

	col = int(file - 'a')   // 'a' = 0, 'b' = 1, ..., 'h' = 7
	row = 8 - int(rank-'0') // '8' = 0 (top row), '1' = 7 (bottom row)

	if row < 0 || row > 7 || col < 0 || col > 7 {
		return 0, 0, fmt.Errorf("invalid square: %s", square)
	}

	return row, col, nil
}

func isWhite(piece Piece) bool {
	return piece >= "A" && piece <= "Z"
}

func isBlack(piece Piece) bool {
	return piece >= "a" && piece <= "z"
}

func printBoard(board Board) {
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			piece := board[row][col]
			if piece == "" {
				fmt.Print(". ")
			} else {
				fmt.Printf("%s ", piece)
			}
		}
		fmt.Println()
	}
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isCapturable(target Piece, color pb.Color) bool {
	if target == Empty {
		return true
	}
	return isOpponentPiece(target, color)
}

func isOpponentPiece(piece Piece, color pb.Color) bool {
	if color == pb.Color_WHITE {
		return isBlack(piece)
	}
	return isWhite(piece)
}

func isClearPath(board Board, fromRow, fromCol, toRow, toCol int) bool {
	dRow := sign(toRow - fromRow)
	dCol := sign(toCol - fromCol)

	r, c := fromRow+dRow, fromCol+dCol
	for r != toRow || c != toCol {
		if board[r][c] != Empty {
			return false
		}
		r += dRow
		c += dCol
	}
	return true
}

func sign(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}

func generateFEN(board Board, turn pb.Color, fullMoveCount int) string {
	fen := ""

	for row := 0; row < 8; row++ {
		emptyCount := 0
		for col := 0; col < 8; col++ {
			piece := board[row][col]
			if piece == Empty {
				emptyCount++
			} else {
				if emptyCount > 0 {
					fen += fmt.Sprintf("%d", emptyCount)
					emptyCount = 0
				}
				fen += string(piece)
			}
		}
		if emptyCount > 0 {
			fen += fmt.Sprintf("%d", emptyCount)
		}
		if row != 7 {
			fen += "/"
		}
	}

	// Active color
	activeColor := "w"
	if turn == pb.Color_BLACK {
		activeColor = "b"
	}

	// For now:
	castling := "-"
	enPassant := "-"
	halfmoveClock := 0 // you can increment later when needed

	return fmt.Sprintf("%s %s %s %s %d %d",
		fen,
		activeColor,
		castling,
		enPassant,
		halfmoveClock,
		fullMoveCount,
	)
}
