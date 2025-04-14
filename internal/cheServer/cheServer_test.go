package cheServer_test

import (
	"context"
	"testing"

	pb "chess-engine/gen"
	"chess-engine/internal/cheServer"
)

func TestSubmitMove(t *testing.T) {
	tests := []struct {
		name        string
		setup       func() (*cheServer.GameServer, *pb.MoveRequest)
		wantSuccess bool
		wantMessage string
	}{
		{
			name: "Valid pawn move e2 to e4",
			setup: func() (*cheServer.GameServer, *pb.MoveRequest) {
				s := cheServer.New()
				gameID := "test-game"
				playerID := "p1"

				// Setup initial board
				board := cheServer.InitBoard()

				// Ensure it's white's turn and piece is white
				s.GameState[gameID] = &cheServer.GameState{
					GameID:     gameID,
					PlayerID:   playerID,
					OpponentID: "p2",
					Color:      pb.Color_WHITE, // This player's assigned color
					TurnColor:  pb.Color_WHITE, // Must match pawn's color
					Moves:      []string{},
					Board:      board,
				}

				req := &pb.MoveRequest{
					GameId:     gameID,
					PlayerId:   playerID,
					FromSquare: "e2",
					ToSquare:   "e4",
				}
				return s, req
			},
			wantSuccess: true,
			wantMessage: "Move applied successfully",
		},
		{
			name: "Invalid move from empty square",
			setup: func() (*cheServer.GameServer, *pb.MoveRequest) {
				s := cheServer.New()
				gameID := "game2"
				playerID := "p1"

				board := cheServer.InitBoard()

				s.GameState[gameID] = &cheServer.GameState{
					GameID:     gameID,
					PlayerID:   playerID,
					OpponentID: "p2",
					Color:      pb.Color_WHITE,
					TurnColor:  pb.Color_WHITE,
					Moves:      []string{},
					Board:      board,
				}

				req := &pb.MoveRequest{
					GameId:     gameID,
					PlayerId:   playerID,
					FromSquare: "e5", // no piece here in initial board
					ToSquare:   "e6",
				}
				return s, req
			},
			wantSuccess: false,
			wantMessage: "No piece at from_square",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			server, req := tc.setup()
			res, err := server.SubmitMove(context.Background(), req)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if res.Success != tc.wantSuccess {
				t.Errorf("Expected success = %v, got %v", tc.wantSuccess, res.Success)
			}
			if res.Message != tc.wantMessage {
				t.Errorf("Expected message = %q, got %q", tc.wantMessage, res.Message)
			}
		})
	}
}
