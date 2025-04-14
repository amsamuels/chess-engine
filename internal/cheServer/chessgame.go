package cheServer

import (
	pb "chess-engine/gen"
	"context"
	"fmt"
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
	game, err := s.getGameByID(req.GameId)
	if err != nil {
		return fail("Game not found"), nil
	}

	fromRow, fromCol, toRow, toCol, err := validateCoordinates(req.FromSquare, req.ToSquare)
	if err != nil {
		return fail(err.Error()), nil
	}

	piece := game.Board[fromRow][fromCol]
	if piece == Empty {
		return fail("No piece at from_square"), nil
	}
	if !isCorrectTurn(piece, game.TurnColor) {
		return fail("Not your turn"), nil
	}

	if !isValidMove(piece, fromRow, fromCol, toRow, toCol, game) {
		return fail(fmt.Sprintf("Illegal move for %s", piece)), nil
	}

	applyMove(game, fromRow, fromCol, toRow, toCol)

	endMessage := checkGameStatus(game)

	fen := generateFEN(game.Board, game.TurnColor, len(game.Moves)/2+1)

	return buildMoveResponse(true, endMessage, fen, game.PlayerID), nil
}
