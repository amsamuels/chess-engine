package cheServer

import (
	pb "chess-engine/gen"
	"chess-engine/internal/game"
	"chess-engine/internal/game/board"
	"chess-engine/internal/game/session"
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano())) // No need to seed the global rand in Go 1.20+
}

type GameManager struct {
	Session   map[string]*session.GameState
	joinQueue chan *session.GameSession
	mu        sync.RWMutex
}

type GameServer struct {
	pb.UnimplementedChessServiceServer
	GameState map[string]*game.GameState
}

func NewGameManager() *GameManager {
	gm := &GameManager{
		Session:   make(map[string]*session.GameState),
		joinQueue: make(chan *session.GameSession, 10),
	}
	go gm.matchPlayers()
	return gm
}

func (gm *GameManager) StartNewGame(ctx context.Context, playerID string) (string, error) {
	gameID := uuid.NewString()
	ready := make(chan struct{})
	session := &session.GameState{
		GameID:  gameID,
		Player1: playerID,
		Ready:   ready,
		Input:   make(chan GameCommand),
		State: &game.GameState{
			GameID:    gameID,
			PlayerID:  playerID,
			Moves:     []string{},
			Board:     board.NewChessBoard(),
			TurnColor: pb.Color_WHITE,
		},
	}

	gm.mu.Lock()
	gm.sessions[gameID] = session
	gm.mu.Unlock()

	gm.joinQueue <- session

	return gameID, nil // The client waits for notification through polling or future streaming
}

// StartGame starts a new chess game session.
func (m *GameServer) StartGame(ctx context.Context, req *pb.StartGameRequest) (*pb.StartGameResponse, error) {

	// 1. Generate UUIDs for game, player, and opponent
	gameID := uuid.NewString()

	// 2. Randomly assign a color (WHITE or BLACK)
	color := pb.Color(rand.Intn(2))
	oppColor := game.OppositeColor(color)

	// 3. Create the GameState object
	gs := &session.GameState{
		GameID:        gameID,
		PlayerID:      playerID,
		OpponentID:    opponentID,
		PlayerColor:   color,
		OpponentColor: oppColor,
		Moves:         []string{},
		Board:         board.NewChessBoard(),
		TurnColor:     pb.Color_WHITE,
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

/*
Convert (row, col) coordinates into a linear bit index (0–63) for bitboard usage.

A bitboard is a 64-bit unsigned integer where each bit represents a square on the 8×8 chess board.
The squares are indexed from top-left (a8) to bottom-right (h1), row by row:

    0  → a8   1  → b8   2  → c8  ... 7  → h8
    8  → a7   9  → b7   ...       ...
    ...
    56 → a1   ...               63 → h1

To compute the index from row and column:
    index = row * 8 + col

Example:
    fromRow = 6, fromCol = 4 → "e2" → index = 6*8 + 4 = 52
    toRow   = 4, toCol   = 4 → "e4" → index = 4*8 + 4 = 36

This mapping lets us efficiently manipulate pieces on the board using bitwise operations.
*/

func (s *GameServer) SubmitMove(ctx context.Context, req *pb.MoveRequest) (*pb.MoveResponse, error) {
	session, err := s.getGameByID(req.GameId)
	if err != nil {
		return game.Fail("Game not found"), nil
	}

	fromRow, fromCol, toRow, toCol, err := game.ValidateCoordinates(req.FromSquare, req.ToSquare)
	if err != nil {
		return game.Fail(err.Error()), nil
	}
	fromIndex := fromRow*8 + fromCol
	toIndex := toRow*8 + toCol

	return session.TryMove(fromIndex, toIndex, req.PlayerId)
}

func (s *GameServer) getGameByID(gameID string) (*game.GameState, error) {
	game, ok := s.GameState[gameID]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return game, nil
}
