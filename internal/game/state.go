package game

import (
	pb "chess-engine/gen"
	"chess-engine/internal/game/board"
	"fmt"
)

func (gs *GameState) TryMove(fromIndex, toIndex int, playerID string) (*pb.MoveResponse, error) {
	// Check player validity
	if playerID != gs.PlayerID && playerID != gs.OpponentID {
		return Fail("Invalid player"), nil
	}

	// Validate it's the correct player's turn
	isPlayerTurn := (playerID == gs.PlayerID && gs.PlayerColor == gs.TurnColor) ||
		(playerID == gs.OpponentID && gs.OpponentColor != gs.TurnColor)

	if !isPlayerTurn {
		return Fail("Not your turn"), nil
	}

	piece := gs.Board.GetBitmapIndex(fromIndex)
	if piece == board.Empty {
		return Fail("No piece at from_square"), nil
	}

	if !IsCorrectTurn(piece, gs.TurnColor) {
		return Fail("Not your turn"), nil
	}

	if !IsValidMove(piece, fromIndex, toIndex, gs) {
		return Fail(fmt.Sprintf("Illegal move for %s", piece)), nil
	}

	ApplyMove(gs, piece, fromIndex, toIndex)

	endMessage := CheckGameStatus(gs)
	fen := GenerateFEN(gs.Board, gs.TurnColor, len(gs.Moves)/2+1)

	return BuildMoveResponse(true, endMessage, fen, playerID), nil
}

// Add the new method to switch turns
func (gs *GameState) SwitchTurn() {
	if gs.TurnColor == pb.Color_WHITE {
		gs.TurnColor = pb.Color_BLACK
	} else {
		gs.TurnColor = pb.Color_WHITE
	}
}
