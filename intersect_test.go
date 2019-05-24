package hcover_test

import (
	"fmt"
	"testing"

	"github.com/savardiego/hcover"
)

func TestLineIntersection(t *testing.T) {
	//array of x1,y1,x2,y2
	lines := map[byte][]float64{
		'A': {2.0, 1.0, 8.0, 4.0},
		'B': {1.0, 3.0, 6.0, 2.0},
		'C': {3.0, 3.0, 7.0, 3.0},
		'D': {5.0, 4.0, 5.0, 1.0},
		'E': {2.0, 4.0, 7.0, 3.0},
		'F': {-4.0, 1.0, 2.0, -3.0},
		'G': {-3.0, -2.0, 6.0, -2.0},
	}
	intersection := map[string]bool{
		"AB": true,
		"AC": true,
		"AD": true,
		"AE": true,
		"AF": true,
		"AG": true,
		"BC": true,
		"BD": true,
		"BE": false,
		"BF": true,
		"CD": true,
		"CE": true,
		"CF": true,
		"CG": false,
		"DE": true,
		"DF": true,
		"DG": true,
		"EF": true,
		"EG": true,
		"FG": true,
	}
	for ll, exp := range intersection {
		a1 := hcover.Point{lines[ll[0]][0], lines[ll[0]][1]}
		b1 := hcover.Point{lines[ll[0]][2], lines[ll[0]][3]}
		l1 := hcover.LineTru2Points(a1, b1)
		fmt.Printf("Line1: %v\n", l1)
		a2 := hcover.Point{lines[ll[1]][0], lines[ll[1]][1]}
		b2 := hcover.Point{lines[ll[1]][2], lines[ll[1]][3]}
		l2 := hcover.LineTru2Points(a2, b2)
		fmt.Printf("Line2: %v\n", l2)
		intersect, intersection := l1.Intersect(l2)
		if intersect != exp {
			t.Errorf("Intersect gave the frong result for couple %v -> %t\n", ll, intersect)
		}
		fmt.Printf("Intersection for %v is: %v\n", ll, intersection)
	}
}

func TestLineAndIntersect(t *testing.T) {
	p1 := hcover.Point{X: 2, Y: 1}
	p2 := hcover.Point{X: 8, Y: 4}
	l1 := hcover.LineTru2Points(p1, p2)
	p3 := hcover.Point{X: 5, Y: 4}
	p4 := hcover.Point{X: 7, Y: 2}
	l2 := hcover.LineTru2Points(p3, p4)
	y1 := l1.EvalY(p1.X)
	x2 := l1.EvalX(p2.Y)
	if y1 != p1.Y {
		t.Errorf("Inconsitent calculation for Y, is %f and should be %f", y1, p1.Y)
	}
	if x2 != p2.X {
		t.Errorf("Inconsitent calculation for X, is %f and should be %f", x2, p2.X)
	}
	intersect, pi := l1.Intersect(l2)
	fmt.Println(pi)
	if !intersect {
		t.Errorf("Lines (l1:%v   l2:%v) should intersect but result is %t", l1, l2, intersect)

	}
	xl1 := l1.EvalX(pi.Y)
	xl2 := l2.EvalX(pi.Y)
	if xl1 != pi.X || xl2 != pi.X {
		t.Errorf("Inconsitent calculation for X: l1:%f l2:%f  intesection:%f", xl1, xl2, pi.X)
	}
	yl1 := l1.EvalY(pi.X)
	yl2 := l2.EvalY(pi.X)
	if yl1 != pi.Y || yl2 != pi.Y {
		t.Errorf("Inconsitent calculation for Y: l1:%f l2:%f  intesection:%f", yl1, yl2, pi.Y)
	}
}

func TestLineCrossSegment(t *testing.T) {
	lines := map[byte][]float64{
		'A': {2.0, 1.0, 8.0, 4.0},
		'B': {1.0, 3.0, 6.0, 2.0},
		'C': {3.0, 3.0, 7.0, 3.0},
		'D': {5.0, 4.0, 5.0, 1.0},
		'E': {2.0, 4.0, 7.0, 3.0},
		'F': {-4.0, 1.0, 2.0, -3.0},
		'G': {-3.0, -2.0, 6.0, -2.0},
	}
	//First letter is line, second is segment
	intersection := map[string]bool{
		"AB": true,
		"AC": true,
		"AD": true,
		"AE": true,
		"AF": true,
		"AG": false,
		"BC": false,
		"BD": true,
		"BE": false,
		"BF": false,
		"BG": false,
		"CD": true,
		"CE": true,
		"CF": false,
		"CG": false,
		"DE": true,
		"DF": false,
		"DG": true,
		"EF": false,
		"EG": false,
		"FG": true,
	}
	for ll, exp := range intersection {
		a1 := hcover.Point{lines[ll[0]][0], lines[ll[0]][1]}
		b1 := hcover.Point{lines[ll[0]][2], lines[ll[0]][3]}
		line := hcover.LineTru2Points(a1, b1)
		fmt.Printf("Line: %v\n", line)
		a2 := hcover.Point{lines[ll[1]][0], lines[ll[1]][1]}
		b2 := hcover.Point{lines[ll[1]][2], lines[ll[1]][3]}
		seg := hcover.NewSegment(a2, b2)
		fmt.Printf("Segment: %v\n", seg)
		cross, intersection := line.Cross(seg)
		if cross != exp {
			t.Errorf("Cross gave the wrong result for line %v and segment %v -> %t\n", line, seg, cross)
		}
		t.Logf("Cross is:%t && instersection: %v", cross, intersection)
	}

}
