package cheServer

import (
	pb "chess-engine/gen"
	"chess-engine/internal/game/board"
)

func isKingInCheck(board *board.Bitboard, color pb.Color) bool {
	var kingRow, kingCol int
	found := false

	// Step 1: Find the king
	for r := range 63 {
		for c := range 63 {

			Index := r*8 + c

			bitmapIndex := board.GetBitmapIndex(Index)
			piece := board.Symbols[bitmapIndex]

			if color == pb.Color_WHITE && piece == "w" {
				kingRow, kingCol = r, c
				found = true
			}
			if color == pb.Color_BLACK && piece == BlackKing {
				kingRow, kingCol = r, c
				found = true
			}
		}
	}
	if !found {
		return false // fail-safe
	}

	// Step 2: See if any opponent piece can move to the king
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			piece := board[r][c]
			if piece == Empty {
				continue
			}
			if color == pb.Color_WHITE && isBlack(piece) {
				if isLegalAttack(board, r, c, kingRow, kingCol, pb.Color_BLACK) {
					return true
				}
			}
			if color == pb.Color_BLACK && isWhite(piece) {
				if isLegalAttack(board, r, c, kingRow, kingCol, pb.Color_WHITE) {
					return true
				}
			}
		}
	}

	return false
}

func isLegalAttack(board Board, fromRow, fromCol, toRow, toCol int, color pb.Color) bool {
	piece := board[fromRow][fromCol]

	switch piece {
	case WhitePawn, BlackPawn:
		return isValidPawnMove(board, fromRow, fromCol, toRow, toCol, color)
	case WhiteKnight, BlackKnight:
		return isValidKnightMove(board, fromRow, fromCol, toRow, toCol, color)
	case WhiteBishop, BlackBishop:
		return isValidBishopMove(board, fromRow, fromCol, toRow, toCol, color)
	case WhiteRook, BlackRook:
		return isValidRookMove(board, fromRow, fromCol, toRow, toCol, color)
	case WhiteQueen, BlackQueen:
		return isValidQueenMove(board, fromRow, fromCol, toRow, toCol, color)
	case WhiteKing, BlackKing:
		return isValidKingMove(board, fromRow, fromCol, toRow, toCol, color)
	}
	return false
}

func hasAnyLegalMoves(board Board, color pb.Color) bool {
	for fromRow := range 8 {
		for fromCol := range 8 {
			piece := board[fromRow][fromCol]

			if piece == Empty {
				continue
			}
			if color == pb.Color_WHITE && !isWhite(piece) {
				continue
			}
			if color == pb.Color_BLACK && !isBlack(piece) {
				continue
			}

			for toRow := 0; toRow < 8; toRow++ {
				for toCol := 0; toCol < 8; toCol++ {
					if fromRow == toRow && fromCol == toCol {
						continue
					}

					if isLegalAttack(board, fromRow, fromCol, toRow, toCol, color) {
						// Make a simulated move
						simulated := copyBoard(board)
						simulated[toRow][toCol] = simulated[fromRow][fromCol]
						simulated[fromRow][fromCol] = Empty

						if !isKingInCheck(simulated, color) {
							return true
						}
					}
				}
			}
		}
	}
	return false
}
