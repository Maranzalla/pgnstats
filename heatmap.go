package main

import "github.com/malbrecht/chess"

//WB is White/Black counts
type WB struct {
	W int `json:"w"`
	B int `json:"b"`
}

//Heatsquare is a square in the Heatmap
type Heatsquare struct {
	All WB `json:"all"`
	K   WB `json:"k"`
	Q   WB `json:"q"`
	R   WB `json:"r"`
	B   WB `json:"b"`
	N   WB `json:"n"`
	P   WB `json:"p"`
}

//Heatmap is a collection of HeatSquares in the shape of a chess board
type Heatmap [64]Heatsquare

//Count increments the counts of HeatSquares depending on piece and square
func (hm *Heatmap) Count(piece chess.Piece, sq chess.Sq) {
	index := (7-sq.Rank())*8 + sq.File()

	hmPtr := &hm[index]

	switch piece {
	case chess.WP:
		hmPtr.P.W++
	case chess.BP:
		hmPtr.P.B++
	case chess.WN:
		hmPtr.N.W++
	case chess.BN:
		hmPtr.N.B++
	case chess.WB:
		hmPtr.B.W++
	case chess.BB:
		hmPtr.B.B++
	case chess.WR:
		hmPtr.R.W++
	case chess.BR:
		hmPtr.R.B++
	case chess.WQ:
		hmPtr.Q.W++
	case chess.BQ:
		hmPtr.Q.B++
	case chess.WK:
		hmPtr.K.W++
	case chess.BK:
		hmPtr.K.B++
	}

	switch piece.Color() {
	case chess.White:
		hmPtr.All.W++
	case chess.Black:
		hmPtr.All.B++
	}
}
