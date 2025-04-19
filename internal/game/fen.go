package game

import (
	pb "chess-engine/gen"
	"chess-engine/internal/game/board"
	"fmt"
)

func GenerateFEN(b *board.Bitboard, turn pb.Color, fullMoveCount int) string {
	fen := ""

	for row := 0; row < 8; row++ {
		emptyCount := 0
		for col := 0; col < 8; col++ {
			index := row*8 + col
			piece := b.GetBitmapIndex(index)
			if piece == board.Empty {
				emptyCount++
			} else {
				if emptyCount > 0 {
					fen += fmt.Sprintf("%d", emptyCount)
					emptyCount = 0
				}
				fen += piece.Symbol()
			}
		}
		if emptyCount > 0 {
			fen += fmt.Sprintf("%d", emptyCount)
		}
		if row != 7 {
			fen += "/"
		}
	}

	activeColor := "w"
	if turn == pb.Color_BLACK {
		activeColor = "b"
	}

	castling := "-"
	enPassant := "-"
	halfmoveClock := 0

	return fmt.Sprintf("%s %s %s %s %d %d",
		fen,
		activeColor,
		castling,
		enPassant,
		halfmoveClock,
		fullMoveCount,
	)
}
