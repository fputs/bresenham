package bresenham

import (
	"testing"
)

func TestBresenhamInt(t *testing.T) {
	var x0, y0, x1, y1 int
	x0 = 1
	y0 = 1
	x1 = 3
	y1 = 4
	res := []Point[int]{
		{1, 1}, {2, 2}, {2, 3}, {3, 4},
	}

	line := Bresenham(x0, y0, x1, y1)
	if line[0].X != res[0].X && line[0].Y != line[0].Y {
		t.Errorf("Wanted %d,%d, got %d,%d", res[0].X, res[0].Y, line[0].X, line[0].Y)
	}
}

func TestBresenhamIntNegative(t *testing.T) {
	var x0, y0, x1, y1 int
	x0 = -1
	y0 = 2
	x1 = 3
	y1 = -4
	res := []Point[int]{
		{-1, 2}, {0, 1}, {0, 0}, {1, -1}, {2, -2}, {2, -3}, {3, -4},
	}

	line := Bresenham(x0, y0, x1, y1)
	if line[0].X != res[0].X && line[0].Y != line[0].Y {
		t.Errorf("Wanted %d,%d, got %d,%d", res[0].X, res[0].Y, line[0].X, line[0].Y)
	}
}

func TestBresenhamInt64(t *testing.T) {
	var x0, y0, x1, y1 int64
	x0 = 1
	y0 = 1
	x1 = 3
	y1 = 4
	res := []Point[int64]{
		{1, 1}, {2, 2}, {2, 3}, {3, 4},
	}

	line := Bresenham(x0, y0, x1, y1)
	if line[0].X != res[0].X && line[0].Y != line[0].Y {
		t.Errorf("Wanted %d,%d, got %d,%d", res[0].X, res[0].Y, line[0].X, line[0].Y)
	}
}
