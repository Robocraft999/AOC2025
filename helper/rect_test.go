package helper

import "testing"

/*
1234
1243
2134
2143
3412
3421
4312
4321
*/
func coordPairCombinations(p1, p2, p3, p4 [2]int) [8][4][2]int {
	return [8][4][2]int{
		{p1, p2, p3, p4},
		{p1, p2, p4, p3},
		{p2, p1, p3, p4},
		{p2, p1, p4, p3},
		{p3, p4, p1, p2},
		{p3, p4, p2, p1},
		{p4, p3, p1, p2},
		{p4, p3, p2, p1},
	}
}

func TestRealLineIntersect2dNormal(t *testing.T) {
	l11 := [2]int{0, 6}
	l12 := [2]int{5, 6}
	l21 := [2]int{3, 0}
	l22 := [2]int{3, 10}
	for _, points := range coordPairCombinations(l11, l12, l21, l22) {
		if !RealLineIntersect2d(points[0], points[1], points[2], points[3]) {
			t.Errorf("l1 and l2 should be intersecting (%v, %v and %v, %v)", points[0], points[1], points[2], points[3])
		}
	}
}

func TestRealLineIntersect2dTIntersection(t *testing.T) {
	l11 := [2]int{0, 6}
	l12 := [2]int{5, 6}
	l21 := [2]int{3, 6}
	l22 := [2]int{3, 10}
	for _, points := range coordPairCombinations(l11, l12, l21, l22) {
		if RealLineIntersect2d(points[0], points[1], points[2], points[3]) {
			t.Errorf("l1 and l2 should not be intersecting (%v, %v and %v, %v)", points[0], points[1], points[2], points[3])
		}
	}
}

func TestRealLineIntersect2dDisjunct(t *testing.T) {
	l11 := [2]int{0, 6}
	l12 := [2]int{5, 6}
	l21 := [2]int{13, 6}
	l22 := [2]int{13, 10}
	for _, points := range coordPairCombinations(l11, l12, l21, l22) {
		if RealLineIntersect2d(points[0], points[1], points[2], points[3]) {
			t.Errorf("l1 and l2 should not be intersecting (%v, %v and %v, %v)", points[0], points[1], points[2], points[3])
		}
	}
}

/*
.......
..+..#.
.......
.----..
.......
.......
..#..+.
*/
func TestRectLineIntersect2dNormal(t *testing.T) {
	b1 := [2]int{2, 6}
	b2 := [2]int{5, 1}
	l1 := [2]int{1, 3}
	l2 := [2]int{4, 3}
	combs := coordPairCombinations(b1, b2, l1, l2)
	for _, points := range combs[:4] {
		if !RectLineIntersect2d(points[0], points[1], points[2], points[3]) {
			t.Errorf("rect and line should be intersecting (%v, %v and %v, %v)", points[0], points[1], points[2], points[3])
		}
	}
}

/*
..|....
..+..#.
..|....
..|....
..|....
..:....
..#..+.
..:....
*/
func TestRectLineIntersect2dEdge(t *testing.T) {
	b1 := [2]int{2, 6}
	b2 := [2]int{5, 1}
	l1 := [2]int{2, 0}
	l2 := [2]int{2, 4}
	l3 := [2]int{2, 7}
	combs1 := coordPairCombinations(b1, b2, l1, l2)
	for _, points := range combs1[:4] {
		if RectLineIntersect2d(points[0], points[1], points[2], points[3]) {
			t.Errorf("rect and line should not be intersecting at the edge (out/in) (%v, %v and %v, %v)", points[0], points[1], points[2], points[3])
		}
	}

	combs2 := coordPairCombinations(b1, b2, l1, l3)
	for _, points := range combs2[:4] {
		if RectLineIntersect2d(points[0], points[1], points[2], points[3]) {
			t.Errorf("rect and line should not be intersecting at the edge (out/out)(%v, %v and %v, %v)", points[0], points[1], points[2], points[3])
		}
	}
}

func TestRectLineIntersect2dDisjunct(t *testing.T) {
	b1 := [2]int{2, 6}
	b2 := [2]int{5, 1}
	l1 := [2]int{3, 2}
	l2 := [2]int{3, 5}
	l3 := [2]int{1, 2}
	l4 := [2]int{1, 5}
	combs1 := coordPairCombinations(b1, b2, l1, l2)
	for _, points := range combs1[:4] {
		if RectLineIntersect2d(points[0], points[1], points[2], points[3]) {
			t.Errorf("rect and line should not be intersecting if inside (%v, %v and %v, %v)", points[0], points[1], points[2], points[3])
		}
	}

	combs2 := coordPairCombinations(b1, b2, l3, l4)
	for _, points := range combs2[:4] {
		if RectLineIntersect2d(points[0], points[1], points[2], points[3]) {
			t.Errorf("rect and line should not be intersecting if outside (%v, %v and %v, %v)", points[0], points[1], points[2], points[3])
		}
	}
}
