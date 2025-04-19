package game

import (
	pb "chess-engine/gen"
	"chess-engine/internal/game/board"
)

func isKingInCheck(b *board.Bitboard, color pb.Color) bool {
	var kingRow, kingCol int
	found := false

	// Step 1: Find the king
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			index := r*8 + c
			piece := b.GetBitmapIndex(index)

			if (color == pb.Color_WHITE && piece == board.WhiteKing) ||
				(color == pb.Color_BLACK && piece == board.BlackKing) {
				kingRow, kingCol = r, c
				found = true
				break
			}
		}
	}

	if !found {
		return false // Fail-safe: king not found
	}

	// Step 2: See if any opponent piece can move to the king
	opponentColor := OppositeColor(color)

	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			index := r*8 + c
			piece := b.GetBitmapIndex(index)

			if piece == board.Empty {
				continue
			}
			fromIndex := r*8 + c
			toIndex := kingRow*8 + kingCol
			if (color == pb.Color_WHITE && isBlack(piece)) || (color == pb.Color_BLACK && isWhite(piece)) {
				if isLegalAttack(b, fromIndex, toIndex, opponentColor) {
					return true
				}
			}
		}
	}

	return false
}

func isLegalAttack(b *board.Bitboard, fromIndex, toIndex int, color pb.Color) bool {
	piece := b.GetBitmapIndex(fromIndex)
	switch piece.String() {
	case board.WhitePawn.String(), board.BlackPawn.String():
		return isValidPawnMove(b, fromIndex, toIndex, color)
	case board.WhiteKnight.String(), board.BlackKnight.String():
		return isValidKnightMove(b, fromIndex, toIndex, color)
	case board.WhiteBishop.String(), board.BlackBishop.String():
		return isValidBishopMove(b, fromIndex, toIndex, color)
	case board.WhiteRook.String(), board.BlackRook.String():
		return isValidRookMove(b, fromIndex, toIndex, color)
	case board.WhiteQueen.String(), board.BlackQueen.String():
		return isValidQueenMove(b, fromIndex, toIndex, color)
	case board.WhiteKing.String(), board.BlackKing.String():
		return isValidKingMove(b, fromIndex, toIndex, color)
	}
	return false
}

func hasAnyLegalMoves(b *board.Bitboard, color pb.Color) bool {
	for fromRow := 0; fromRow < 8; fromRow++ {
		for fromCol := 0; fromCol < 8; fromCol++ {
			fromIndex := fromRow*8 + fromCol
			piece := b.GetBitmapIndex(fromIndex)

			// Skip empty or opponent pieces
			if piece == board.Empty {
				continue
			}
			if color == pb.Color_WHITE && !isWhite(piece) {
				continue
			}
			if color == pb.Color_BLACK && !isBlack(piece) {
				continue
			}

			for toRow := 0; toRow < 8; toRow++ {
				for toCol := 0; toCol < 8; toCol++ {
					toIndex := toRow*8 + toCol

					if fromIndex == toIndex {
						continue
					}

					if isLegalAttack(b, fromIndex, toIndex, color) {
						// Simulate the move
						sim := b.Copy()
						sim.MovePieceBit(int(piece), fromIndex, toIndex)

						if !isKingInCheck(sim, color) {
							return true
						}
					}
				}
			}
		}
	}
	return false
}
