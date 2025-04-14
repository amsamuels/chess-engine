package cheServer

import (
	pb "chess-engine/gen"
	"fmt"
)

func (s *GameServer) getGameByID(gameID string) (*GameState, error) {
	game, ok := s.GameState[gameID]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return game, nil
}

func validateCoordinates(from, to string) (int, int, int, int, error) {
	fromRow, fromCol, err := squareToCoords(from)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("Invalid from_square")
	}
	toRow, toCol, err := squareToCoords(to)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("Invalid to_square")
	}
	return fromRow, fromCol, toRow, toCol, nil
}

func isCorrectTurn(piece Piece, turn pb.Color) bool {
	return (turn == pb.Color_WHITE && isWhite(piece)) || (turn == pb.Color_BLACK && isBlack(piece))
}

func isValidMove(piece Piece, fromRow, fromCol, toRow, toCol int, game *GameState) bool {
	switch piece {
	case WhitePawn, BlackPawn:
		return isValidPawnMove(game.Board, fromRow, fromCol, toRow, toCol, game.TurnColor)
	case WhiteKnight, BlackKnight:
		return isValidKnightMove(game.Board, fromRow, fromCol, toRow, toCol, game.TurnColor)
	case WhiteBishop, BlackBishop:
		return isValidBishopMove(game.Board, fromRow, fromCol, toRow, toCol, game.TurnColor)
	case WhiteRook, BlackRook:
		return isValidRookMove(game.Board, fromRow, fromCol, toRow, toCol, game.TurnColor)
	case WhiteQueen, BlackQueen:
		return isValidQueenMove(game.Board, fromRow, fromCol, toRow, toCol, game.TurnColor)
	case WhiteKing, BlackKing:
		return isValidKingMove(game.Board, fromRow, fromCol, toRow, toCol, game.TurnColor)
	}
	return false
}

func applyMove(game *GameState, fromRow, fromCol, toRow, toCol int) {
	piece := game.Board[fromRow][fromCol]
	game.Board[toRow][toCol] = piece
	game.Board[fromRow][fromCol] = Empty
	game.Moves = append(game.Moves, coordsToSquare(fromRow, fromCol)+coordsToSquare(toRow, toCol))

	// Flip turn
	if game.TurnColor == pb.Color_WHITE {
		game.TurnColor = pb.Color_BLACK
	} else {
		game.TurnColor = pb.Color_WHITE
	}
}

func checkGameStatus(game *GameState) string {
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

func buildMoveResponse(success bool, message, fen, nextPlayer string) *pb.MoveResponse {
	return &pb.MoveResponse{
		Success:  success,
		Message:  message,
		Fen:      fen,
		NextTurn: nextPlayer,
	}
}

func fail(msg string) *pb.MoveResponse {
	return &pb.MoveResponse{
		Success: false,
		Message: msg,
	}
}
