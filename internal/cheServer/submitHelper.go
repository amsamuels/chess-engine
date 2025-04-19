package cheServer

import (
	"chess-engine/internal/game"
	"fmt"
)

func (s *GameServer) getGameByID(gameID string) (*game.GameState, error) {
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
