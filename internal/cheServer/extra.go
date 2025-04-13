package cheServer

type Piece string

const (
	Empty       Piece = ""
	WhitePawn   Piece = "P"
	WhiteRook   Piece = "R"
	WhiteKnight Piece = "N"
	WhiteBishop Piece = "B"
	WhiteQueen  Piece = "Q"
	WhiteKing   Piece = "K"
	BlackPawn   Piece = "p"
	BlackRook   Piece = "r"
	BlackKnight Piece = "n"
	BlackBishop Piece = "b"
	BlackQueen  Piece = "q"
	BlackKing   Piece = "k"
)

type Board [8][8]Piece

func InitBoard() Board {
	var board Board

	// Place white pieces
	board[0] = [8]Piece{WhiteRook, WhiteKnight, WhiteBishop, WhiteQueen, WhiteKing, WhiteBishop, WhiteKnight, WhiteRook}
	board[1] = [8]Piece{WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn}

	// Place black pieces
	board[6] = [8]Piece{BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn}
	board[7] = [8]Piece{BlackRook, BlackKnight, BlackBishop, BlackQueen, BlackKing, BlackBishop, BlackKnight, BlackRook}

	return board
}
