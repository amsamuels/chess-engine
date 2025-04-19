package board

import "fmt"

type Piece int

const (
	Empty Piece = iota
	WhitePawn
	WhiteRook
	WhiteKnight
	WhiteBishop
	WhiteQueen
	WhiteKing
	BlackPawn
	BlackRook
	BlackKnight
	BlackBishop
	BlackQueen
	BlackKing
)

var pieceSymbols = [...]string{
	".", "P", "R", "N", "B", "Q", "K", "p", "r", "n", "b", "q", "k",
}

var pieceNames = [...]string{
	"Empty", "WhitePawn", "WhiteRook", "WhiteKnight", "WhiteBishop", "WhiteQueen", "WhiteKing",
	"BlackPawn", "BlackRook", "BlackKnight", "BlackBishop", "BlackQueen", "BlackKing",
}

// Symbol returns the one-character symbol used to represent the piece.
func (p Piece) Symbol() string {
	if p < 0 || int(p) >= len(pieceSymbols) {
		return "?"
	}
	return pieceSymbols[p]
}

// String returns the human-readable name of the piece.
func (p Piece) String() string {
	if p < 0 || int(p) >= len(pieceNames) {
		return "Unknown"
	}
	return pieceNames[p]
}

// Construct a new Bitboard using New. There are also convenience
// functions for constructing bitboards for specific games.
type Bitboard struct {
	Bitmaps  [13]uint64 // Bitmaps for each colour/piece combination
	Occupied uint64     // Union of all bitmaps (occupied squares)
	Ranks    int        // Number of rows
	Files    int        // Number of columns
}

// NewChessBoard is a convenience function for constructing a new chess board.
func NewChessBoard() *Bitboard {
	bb := &Bitboard{Ranks: 8, Files: 8}

	bb.Bitmaps[WhiteRook] = 0x0000000000000081
	bb.Bitmaps[WhiteKnight] = 0x0000000000000042
	bb.Bitmaps[WhiteBishop] = 0x0000000000000024
	bb.Bitmaps[WhiteQueen] = 0x0000000000000008
	bb.Bitmaps[WhiteKing] = 0x0000000000000010
	bb.Bitmaps[WhitePawn] = 0x000000000000ff00

	bb.Bitmaps[BlackRook] = 0x8100000000000000
	bb.Bitmaps[BlackKnight] = 0x4200000000000000
	bb.Bitmaps[BlackBishop] = 0x2400000000000000
	bb.Bitmaps[BlackQueen] = 0x0800000000000000
	bb.Bitmaps[BlackKing] = 0x1000000000000000
	bb.Bitmaps[BlackPawn] = 0x00ff000000000000

	bb.Occupied = Union(bb.Bitmaps[:]...)
	return bb
}

// PrettyPrint pretty-prints a Bitboard using the symbols for each colour/piece
// combination. Empty squares are represented by periods.
func (b *Bitboard) PrettyPrint() {
	for r := b.Ranks - 1; r >= 0; r-- {
		for f := 0; f < b.Files; f++ {
			p := r*8 + f
			piece := b.GetBitmapIndex(p)
			fmt.Print(piece.Symbol(), " ")
		}
		fmt.Println()
	}
}

// GetBitmapIndex returns the array index of the bitmap including a particular
// square.
func (b *Bitboard) GetBitmapIndex(p int) Piece {
	if GetBit(&b.Occupied, p) == 0 {
		return Empty
	}
	for i := Piece(1); i <= BlackKing; i++ {
		if GetBit(&b.Bitmaps[i], p) != 0 {
			return i
		}
	}
	return Empty
}

// Move a piece from bit position p1 to p2.
func (b *Bitboard) MovePieceBit(m int, p1 int, p2 int) {
	b.RemovePieceBit(m, p1)
	b.PlacePieceBit(m, p2)
}

// Remove the piece at bit position p.
func (b *Bitboard) RemovePieceBit(m int, p int) {
	// Update the occupancy bitmap.
	ClearBit(&b.Occupied, p)
	ClearBit(&b.Bitmaps[m], p)
}

// Place the piece at bit position p.
func (b *Bitboard) PlacePieceBit(m int, p int) {
	// Update the occupancy bitmap.
	SetBit(&b.Occupied, p)
	SetBit(&b.Bitmaps[m], p)
}

func (b *Bitboard) Copy() *Bitboard {
	return &Bitboard{
		Bitmaps:  b.Bitmaps, // direct array copy
		Occupied: b.Occupied,
		Ranks:    b.Ranks,
		Files:    b.Files,
	}
}
