package game

import (
	pb "chess-engine/gen"
	"fmt"
)

func ValidateCoordinates(from, to string) (int, int, int, int, error) {
	fromRow, fromCol, err := squareToCoords(from)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("invalid from_square")
	}
	toRow, toCol, err := squareToCoords(to)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("invalid to_square")
	}
	return fromRow, fromCol, toRow, toCol, nil
}

func OppositeColor(color pb.Color) pb.Color {
	if color == pb.Color_WHITE {
		return pb.Color_BLACK
	}
	return pb.Color_WHITE
}

func coordsToSquare(row, col int) string {
	file := string(rune('a' + col))  // convert to column letter (a-h)
	rank := fmt.Sprintf("%d", 8-row) // "8" through "1"
	return file + rank
}

func squareToCoords(square string) (row int, col int, err error) {
	if len(square) != 2 {
		return 0, 0, fmt.Errorf("invalid square: %s", square)
	}

	file := square[0] // column letter
	rank := square[1] // row number

	col = int(file - 'a')   // 'a' = 0, 'b' = 1, ..., 'h' = 7
	row = 8 - int(rank-'0') // '8' = 0 (top row), '1' = 7 (bottom row)

	if row < 0 || row > 7 || col < 0 || col > 7 {
		return 0, 0, fmt.Errorf("invalid square: %s", square)
	}

	return row, col, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}
