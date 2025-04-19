package game

import (
	pb "chess-engine/gen"
	"chess-engine/internal/game/board"
	"fmt"
)

func IsCorrectTurn(piece board.Piece, turn pb.Color) bool {
	return (turn == pb.Color_WHITE && isWhite(piece)) || (turn == pb.Color_BLACK && isBlack(piece))
}

func IsValidMove(piece board.Piece, fromIndex, toIndex int, g *GameState) bool {
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
	}
	return false
}

func ApplyMove(game *GameState, bitmapIndex board.Piece, fromIndex, toIndex int) {
	game.Board.MovePieceBit(int(bitmapIndex), fromIndex, toIndex)
	// Flip turn
	if game.TurnColor == pb.Color_WHITE {
		game.TurnColor = pb.Color_BLACK
	} else {
		game.TurnColor = pb.Color_WHITE
	}
}

func CheckGameStatus(game *GameState) string {
	opponent := game.TurnColor
	inCheck := isKingInCheck(game.Board, opponent)
	hasMoves := hasAnyLegalMoves(game.Board, opponent)

	switch {
	case inCheck && !hasMoves:
		return fmt.Sprintf("Checkmate! %s wins.", oppositeColor(opponent).String())
	case !inCheck && !hasMoves:
		return "Stalemate! It's a draw."
	case inCheck:
		return fmt.Sprintf("Check against %s!", opponent.String())
	default:
		return "Move applied successfully"
	}
}
