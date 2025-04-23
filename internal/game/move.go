package game

import (
	pb "chess-engine/gen"
	"chess-engine/internal/game/board"
)

func IsCorrectTurn(piece board.Piece, turn pb.Color) bool {
	return (turn == pb.Color_WHITE && isWhite(piece)) || (turn == pb.Color_BLACK && isBlack(piece))
}

func IsValidMove(piece board.Piece, fromIndex, toIndex int, b *board.Bitboard, turnColor pb.Color) bool {
	switch piece.String() {
	case board.WhitePawn.String(), board.BlackPawn.String():
		return isValidPawnMove(b, fromIndex, toIndex, turnColor)
	case board.WhiteKnight.String(), board.BlackKnight.String():
		return isValidKnightMove(b, fromIndex, toIndex, turnColor)
	case board.WhiteBishop.String(), board.BlackBishop.String():
		return isValidBishopMove(b, fromIndex, toIndex, turnColor)
	case board.WhiteRook.String(), board.BlackRook.String():
		return isValidRookMove(b, fromIndex, toIndex, turnColor)
	case board.WhiteQueen.String(), board.BlackQueen.String():
		return isValidQueenMove(b, fromIndex, toIndex, turnColor)
	case board.WhiteKing.String(), board.BlackKing.String():
		return isValidKingMove(b, fromIndex, toIndex, turnColor)
	}
	return false
}
