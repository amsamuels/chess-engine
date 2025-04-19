package game

import (
	pb "chess-engine/gen"
	"chess-engine/internal/game/board"
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
