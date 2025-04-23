package session

import (
	pb "chess-engine/gen"
	"chess-engine/internal/game"
	"chess-engine/internal/game/board"
	"fmt"
)

type GameState struct {
	GameID        string
	PlayerID      string
	OpponentID    string
	PlayerColor   pb.Color
	OpponentColor pb.Color
	Moves         []string // placeholder for move history
	Board         *board.Bitboard
	TurnColor     pb.Color // Track whose turn it is
}
type GameSession struct {
	GameID  string
	Player1 string
	Player2 string
	State   *game.GameState
	Ready   chan struct{}
	Input   chan GameCommand
}

type GameCommand struct {
	Move     *pb.MoveRequest
	Response chan *pb.MoveResponse
}

func (gs *GameSession) Run() {
	for {
		select {
		case cmd := <-gs.Input:
			// validate move, update state
			res := processMove(gs.State, cmd.Move)
			cmd.Response <- res
		case <-gs.Done:
			return
		}
	}
}

func processMove(gameState *GameState, moveRequest *pb.MoveRequest) *pb.MoveResponse {
	panic("unimplemented")
}

func (gs *GameState) TryMove(fromIndex, toIndex int, playerID string) (*pb.MoveResponse, error) {
	// Check player validity
	if playerID != gs.PlayerID && playerID != gs.OpponentID {
		return game.Fail("Invalid player"), nil
	}

	// Validate it's the correct player's turn
	isPlayerTurn := (playerID == gs.PlayerID && gs.PlayerColor == gs.TurnColor) ||
		(playerID == gs.OpponentID && gs.OpponentColor != gs.TurnColor)

	if !isPlayerTurn {
		return game.Fail("Not your turn"), nil
	}

	piece := gs.Board.GetBitmapIndex(fromIndex)
	if piece == board.Empty {
		return game.Fail("No piece at from_square"), nil
	}

	if !game.IsCorrectTurn(piece, gs.TurnColor) {
		return game.Fail("Not your turn"), nil
	}

	if !game.IsValidMove(piece, fromIndex, toIndex, gs.Board, gs.TurnColor) {
		return game.Fail(fmt.Sprintf("Illegal move for %s", piece)), nil
	}

	ApplyMove(gs, piece, fromIndex, toIndex)

	endMessage := CheckGameStatus(gs)
	fen := game.GenerateFEN(gs.Board, gs.TurnColor, len(gs.Moves)/2+1)

	return game.BuildMoveResponse(true, endMessage, fen, playerID), nil
}

// Add the new method to switch turns
func (gs *GameState) SwitchTurn() {
	if gs.TurnColor == pb.Color_WHITE {
		gs.TurnColor = pb.Color_BLACK
	} else {
		gs.TurnColor = pb.Color_WHITE
	}
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

func CheckGameStatus(gs *GameState) string {
	opponent := gs.TurnColor
	inCheck := game.IsKingInCheck(gs.Board, opponent)
	hasMoves := game.HasAnyLegalMoves(gs.Board, opponent)

	switch {
	case inCheck && !hasMoves:
		return fmt.Sprintf("Checkmate! %s wins.", game.OppositeColor(opponent).String())
	case !inCheck && !hasMoves:
		return "Stalemate! It's a draw."
	case inCheck:
		return fmt.Sprintf("Check against %s!", opponent.String())
	default:
		return "Move applied successfully"
	}
}
