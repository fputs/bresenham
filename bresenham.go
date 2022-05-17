// Package bresenham provides a generic implementation of
// Bresenham's Line Algorithm
package bresenham

import "golang.org/x/exp/constraints"

// Point is a generic struct holding two generic coordinates
// with a constraint of Signed
type Point[T constraints.Signed] struct {
	X T
	Y T
}

func abs[T constraints.Signed](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

// Bresenham requires arguments of any signed type. The line is generated
// and returned as a slice of Point[T]
func Bresenham[T constraints.Signed](x0, y0, x1, y1 T) []Point[T] {
	var line []Point[T]
	var cx, cy, dx, dy, sx, sy, err T
	cx = x0
	cy = y0
	dx = abs(x1 - x0)
	dy = abs(y1 - y0)

	if cx < x1 {
		sx = 1
	} else {
		sx = -1
	}
	if cy < y1 {
		sy = 1
	} else {
		sy = -1
	}
	err = dx - dy

	for {
		line = append(line, Point[T]{cx, cy})
		if cx == x1 && cy == y1 {
			return line
		}
		e2 := 2 * err
		if e2 > 0-dy {
			err = err - dy
			cx = cx + sx
		}
		if e2 < dx {
			err = err + dx
			cy = cy + sy
		}
	}
}
