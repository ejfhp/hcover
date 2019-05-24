package hcover

import "math"

type Point struct {
	X float64
	Y float64
}

type Line struct {
	A float64
	B float64
	C float64
}

type Segment struct {
	Line
	Start Point
	End   Point
}

func LineTru2Points(p, q Point) Line {
	// fmt.Printf("Points A:%v    B:%v", a, b)
	a := q.Y - p.Y
	b := p.X - q.X
	c := p.Y*q.X - p.X*q.Y
	return Line{A: a, B: b, C: c}
}

func NewSegment(a, b Point) Segment {
	l := LineTru2Points(a, b)
	return Segment{l, a, b}
}

func (l Line) EvalY(x float64) float64 {
	return -1.0*(l.C/l.B) - (x * (l.A / l.B))
}

func (l Line) EvalX(y float64) float64 {
	return -1.0*(l.C/l.A) - (y * (l.B / l.A))
}

func (l1 Line) Intersect(l2 Line) (bool, Point) {
	determinant := l1.A*l2.B - l2.A*l1.B
	if determinant == 0 {
		return false, Point{}
	}
	x := (l1.B*l2.C - l2.B*l1.C) / determinant
	y := (l1.C*l2.A - l2.C*l1.A) / determinant
	return true, Point{x, y}
}

func (l1 Line) Cross(s Segment) (bool, Point) {
	intersect, point := l1.Intersect(s.Line)
	if intersect {
		if isPointInsideExtent(s.Start, s.End, point) {
			return true, point
		}
	}
	return false, Point{}
}

func isPointInsideExtent(vertex, opposedV, p Point) bool {
	maxX := math.Max(vertex.X, opposedV.X)
	minX := math.Min(vertex.X, opposedV.X)
	maxY := math.Max(vertex.Y, opposedV.Y)
	minY := math.Min(vertex.Y, opposedV.Y)
	if p.X <= maxX && p.X >= minX && p.Y <= maxY && p.Y >= minY {
		return true
	}
	return false
}
