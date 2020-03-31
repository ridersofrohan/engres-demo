package simplifications

import "log"

type Point struct {
	a int
	b int
}

type Piece struct {
	a int
	b int
	c Point
	d []Point
	e *Point
	f *Point
}

func j() {
	//Simplify range stmt.
	for k, _ := range []int{} {
		log.Print(k)
	}

	// Simplify slice stmt.
	var a [5]int
	_ = a[2:len(a)]

	// Simplify composite literal stmt.
	_ = []*Piece{
		&Piece{0, 0, Point{}, []Point{Point{}, Point{}}, nil, nil},
	}
}
