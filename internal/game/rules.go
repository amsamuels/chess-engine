package game

import (
	pb "chess-engine/gen"
	"chess-engine/internal/game/board"
)

func isCapturable(target board.Piece, color pb.Color) bool {
	if target == board.Empty {
		return true
	}
	return isOpponentPiece(target, color)
}

func isOpponentPiece(target board.Piece, color pb.Color) bool {
	return (color == pb.Color_WHITE && isBlack(target)) || (color == pb.Color_BLACK && isWhite(target))
}

func isClearPath(b *board.Bitboard, fromIndex, toIndex int) bool {
	fromRow, fromCol := getFrom(fromIndex)
	toRow, toCol := getTo(toIndex)

	dRow := sign(toRow - fromRow)
	dCol := sign(toCol - fromCol)

	// Step through the squares between from and to
	r, c := fromRow+dRow, fromCol+dCol
	for r != toRow || c != toCol {
		square := r*8 + c
		if b.GetBitmapIndex(square) != board.Empty {
			return false
		}
		r += dRow
		c += dCol
	}
	return true
}

func isWhite(piece board.Piece) bool {
	return piece >= board.WhitePawn && piece <= board.WhiteKing
}

func isBlack(piece board.Piece) bool {
	return piece >= board.BlackPawn && piece <= board.BlackKing
}

func isValidPawnMove(b *board.Bitboard, fromIndex, toIndex int, color pb.Color) bool {
	direction := -8 // forward one rank
	startRow := 1   // 2nd row for white (index 1)
	if color == pb.Color_BLACK {
		direction = 8
		startRow = 6
	}

	// Move forward 1
	if toIndex == fromIndex+direction && b.GetBitmapIndex(toIndex) == board.Empty {
		return true
	}

	// Move forward 2 (only from starting position)
	if toIndex == fromIndex+2*direction && fromIndex/8 == startRow && b.GetBitmapIndex(toIndex) == board.Empty &&
		b.GetBitmapIndex(fromIndex+direction) == board.Empty {
		return true
	}

	// Capture diagonally
	if (toIndex == fromIndex+direction-1 && fromIndex%8 != 0) ||
		(toIndex == fromIndex+direction+1 && fromIndex%8 != 7) {
		target := b.GetBitmapIndex(toIndex)
		if isCapturable(target, color) && target != board.Empty {
			return true
		}
	}

	return false
}

func isValidRookMove(b *board.Bitboard, fromIndex, toIndex int, color pb.Color) bool {
	if fromIndex != toIndex {
		return false
	}
	return isClearPath(b, fromIndex, toIndex) && isCapturable(b.GetBitmapIndex(toIndex), color)
}

func isValidQueenMove(b *board.Bitboard, fromIndex, toIndex int, color pb.Color) bool {
	if fromIndex == toIndex {
		return false // Can't move to the same square
	}

	fromRow, fromCol := getFrom(fromIndex)
	toRow, toCol := getTo(toIndex)

	if fromRow == toRow || fromCol == toCol || abs(fromRow-toRow) == abs(fromCol-toCol) {
		return isClearPath(b, fromIndex, toIndex) && isCapturable(b.GetBitmapIndex(toIndex), color)
	}
	return false
}

func isValidKnightMove(b *board.Bitboard, fromIndex, toIndex int, color pb.Color) bool {
	fromRow, fromCol := getFrom(fromIndex)
	toRow, toCol := getTo(toIndex)

	dr := abs(toRow - fromRow)
	dc := abs(toCol - fromCol)

	if (dr == 2 && dc == 1) || (dr == 1 && dc == 2) {
		target := b.GetBitmapIndex(toIndex)
		if target == board.Empty || isOpponentPiece(target, color) {
			return true
		}
	}
	return false
}

func isValidBishopMove(b *board.Bitboard, fromIndex, toIndex int, color pb.Color) bool {
	fromRow, fromCol := getFrom(fromIndex)
	toRow, toCol := getTo(toIndex)

	if abs(fromRow-toRow) != abs(fromCol-toCol) {
		return false
	}
	return isClearPath(b, fromIndex, toIndex) && isCapturable(b.GetBitmapIndex(toIndex), color)
}

func isValidKingMove(b *board.Bitboard, fromIndex, toIndex int, color pb.Color) bool {

	fromRow, fromCol := getFrom(fromIndex)
	toRow, toCol := getTo(toIndex)

	dr := abs(fromRow - toRow)
	dc := abs(fromCol - toCol)

	if dr <= 1 && dc <= 1 {
		return isCapturable(b.GetBitmapIndex(toIndex), color)
	}
	return false
}

func getFrom(fromIndex int) (int, int) {
	return fromIndex / 8, fromIndex % 8
}

func getTo(toIndex int) (int, int) {
	return toIndex / 8, toIndex % 8
}
