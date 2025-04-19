package game

import (
	pb "chess-engine/gen"
)

func BuildMoveResponse(success bool, message, fen, nextPlayer string) *pb.MoveResponse {
	return &pb.MoveResponse{
		Success:  success,
		Message:  message,
		Fen:      fen,
		NextTurn: nextPlayer,
	}
}

func Fail(msg string) *pb.MoveResponse {
	return &pb.MoveResponse{
		Success: false,
		Message: msg,
	}
}
