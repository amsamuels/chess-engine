package cheServer

import pb "chess-engine/gen"

func isValidPawnMove(board Board, fromRow, fromCol, toRow, toCol int, color pb.Color) bool {
	direction := -1
	startRow := 6
	oppositeColor := isBlack

	if color == pb.Color_BLACK {
		direction = 1
		startRow = 1
		oppositeColor = isWhite
	}

	// Move forward 1
	if toCol == fromCol && toRow == fromRow+direction && board[toRow][toCol] == Empty {
		return true
	}

	// Move forward 2 (only from starting position)
	if toCol == fromCol && fromRow == startRow && toRow == fromRow+2*direction &&
		board[fromRow+direction][toCol] == Empty && board[toRow][toCol] == Empty {
		return true
	}

	// Capture diagonally
	if (toCol == fromCol+1 || toCol == fromCol-1) && toRow == fromRow+direction {
		target := board[toRow][toCol]
		if target != Empty && oppositeColor(target) {
			return true
		}
	}

	return false
}

func isValidRookMove(board Board, fromRow, fromCol, toRow, toCol int, color pb.Color) bool {
	if fromRow != toRow && fromCol != toCol {
		return false
	}
	return isClearPath(board, fromRow, fromCol, toRow, toCol) &&
		isCapturable(board[toRow][toCol], color)
}

func isValidQueenMove(board Board, fromRow, fromCol, toRow, toCol int, color pb.Color) bool {
	if fromRow == toRow || fromCol == toCol || abs(fromRow-toRow) == abs(fromCol-toCol) {
		return isClearPath(board, fromRow, fromCol, toRow, toCol) &&
			isCapturable(board[toRow][toCol], color)
	}
	return false
}

func isValidKnightMove(board Board, fromRow, fromCol, toRow, toCol int, color pb.Color) bool {
	dr := abs(toRow - fromRow)
	dc := abs(toCol - fromCol)

	if (dr == 2 && dc == 1) || (dr == 1 && dc == 2) {
		target := board[toRow][toCol]
		if target == Empty || isOpponentPiece(target, color) {
			return true
		}
	}
	return false
}

func isValidBishopMove(board Board, fromRow, fromCol, toRow, toCol int, color pb.Color) bool {
	if abs(fromRow-toRow) != abs(fromCol-toCol) {
		return false
	}
	return isClearPath(board, fromRow, fromCol, toRow, toCol) &&
		isCapturable(board[toRow][toCol], color)
}

func isValidKingMove(board Board, fromRow, fromCol, toRow, toCol int, color pb.Color) bool {
	dr := abs(fromRow - toRow)
	dc := abs(fromCol - toCol)

	if dr <= 1 && dc <= 1 {
		return isCapturable(board[toRow][toCol], color)
	}
	return false
}
