package cheServer

import (
	pb "chess-engine/gen"
	"context"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type GameServer struct {
	pb.UnimplementedChessServiceServer
	GameState map[string]*GameState
}

type GameState struct {
	GameID     string
	PlayerID   string
	OpponentID string
	Color      pb.Color
	Moves      []string // placeholder for move history
	Board      Board
	TurnColor  pb.Color // Track whose turn it is
}

func New() *GameServer {
	return &GameServer{
		GameState: make(map[string]*GameState),
	}
}

// StartGame starts a new chess game session.
func (s *GameServer) StartGame(ctx context.Context, req *pb.StartGameRequest) (*pb.StartGameResponse, error) {

	// 1. Generate UUIDs for game, player, and opponent
	gameID := uuid.NewString()
	playerID := uuid.NewString()
	opponentID := uuid.NewString() // or "AI" prefix if agent

	// 2. Randomly assign a color (WHITE or BLACK)
	color := pb.Color(rand.Intn(2))

	// 3. Create the GameState object
	gs := &GameState{
		GameID:     gameID,
		PlayerID:   playerID,
		OpponentID: opponentID,
		Color:      color,
		Moves:      []string{},
		Board:      InitBoard(),
		TurnColor:  pb.Color_WHITE,
	}
	// 4. Store it in the GameServer's map
	s.GameState[gameID] = gs

	// 5. Return a StartGameResponse
	return &pb.StartGameResponse{
		GameId:     gameID,
		PlayerId:   playerID,
		OpponentId: opponentID,
		Color:      color,
	}, nil
}

func (s *GameServer) SubmitMove(ctx context.Context, req *pb.MoveRequest) (*pb.MoveResponse, error) {
	// 1. Look up the game by req.GameId
	// 2. Validate the move (for now just basic checks: exists, turn, from != to)
	// 3. Update the move history (append "e2e4" format or similar)
	// 4. Update FEN (you can skip or use placeholder for now)
	// 5. Toggle turn
	// 6. Return success with new FEN and next player's ID

	_, ok := s.GameState[req.GameId]
	if !ok {
		return &pb.MoveResponse{
			Success: false,
			Message: "gameId not valid no session for this gameid",
		}, nil
	}

	if req.GetFromSquare() == "" || req.GetToSquare() == "" {
		return &pb.MoveResponse{
			Success: false,
			Message: "Missing move squares: both from_square and to_square are required.",
		}, nil
	}

	return &pb.MoveResponse{
		Success: true,
	}, nil
}
