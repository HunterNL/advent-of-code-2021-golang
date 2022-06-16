package vec2

import (
	"math"
)

type Vec2 struct {
	X float64
	Y float64
}

type vec2int struct {
	x int
	y int
}

type Line struct {
	A Vec2
	B Vec2
}

func (l *Line) Isorthogonal() bool {
	if l.A.X == l.B.X {
		return true
	}

	if l.A.Y == l.B.Y {
		return true
	}

	return false
}

func (a *Vec2) Sub(b *Vec2) Vec2 {
	return Vec2{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}
func (a *Vec2) Add(b *Vec2) Vec2 {
	return Vec2{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (vec *Vec2) Length() float64 {
	return math.Sqrt(math.Pow(float64(vec.X), 2) + math.Pow(float64(vec.Y), 2))
}

func (vec *Vec2) Scale(factor float64) Vec2 {
	return Vec2{X: vec.X * factor, Y: vec.Y * factor}
}

func (vec *Vec2) Normalized() Vec2 {
	return vec.Scale(1.0 / vec.Length())
}

func (from *Vec2) DirectionTo(to *Vec2) Vec2 {
	dir := to.Sub(from)
	return dir.Normalized()
}

// func (from *Vec2) DirectionAsInt(to *Vec2) Vec2 {
// 	dir := to.Sub(from)
// 	norm := dir.Normalized()
// 	return Vec2{
// 		x: math.Ceil()
// 	}
// }

func (from *Vec2) StepsTo(to *Vec2) []Vec2 {
	ret := make([]Vec2, 0)
	dir := from.DirectionTo(to)
	cur := *from
	for cur != *to {
		ret = append(ret, cur)
		cur = cur.Add(&dir)
	}
	ret = append(ret, *to)
	return ret
}
