// File: internal/game/movevalidator.go
package game

import (
	pb "chess-engine/gen"
	"chess-engine/internal/game/board"
)

func Validate(piece board.Piece, fromIndex, toIndex int, g *GameState) bool {
	switch piece.String() {
	case board.WhitePawn.String(), board.BlackPawn.String():
		return isValidPawnMove(g.Board, fromIndex, toIndex, g.TurnColor)
	case board.WhiteKnight.String(), board.BlackKnight.String():
		return isValidKnightMove(g.Board, fromIndex, toIndex, g.TurnColor)
	case board.WhiteBishop.String(), board.BlackBishop.String():
		return isValidBishopMove(g.Board, fromIndex, toIndex, g.TurnColor)
	case board.WhiteRook.String(), board.BlackRook.String():
		return isValidRookMove(g.Board, fromIndex, toIndex, g.TurnColor)
	case board.WhiteQueen.String(), board.BlackQueen.String():
		return isValidQueenMove(g.Board, fromIndex, toIndex, g.TurnColor)
	case board.WhiteKing.String(), board.BlackKing.String():
		return isValidKingMove(g.Board, fromIndex, toIndex, g.TurnColor)
	default:
		return false
	}
}

func isCapturable(target board.Piece, color pb.Color) bool {
	if target == board.Empty {
		return true
	}
	return isOpponentPiece(target, color)
}

func isOpponentPiece(target board.Piece, color pb.Color) bool {
	return (color == pb.Color_WHITE && isBlack(target)) || (color == pb.Color_BLACK && isWhite(target))
}

func isWhite(piece board.Piece) bool {
	return piece >= board.WhitePawn && piece <= board.WhiteKing
}

func isBlack(piece board.Piece) bool {
	return piece >= board.BlackPawn && piece <= board.BlackKing
}
