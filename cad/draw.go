package cad

var drawElements []Drawable

type Drawable interface {
	draw()
}

type Point struct {
	x float64
	y float64
}

type Line struct {
	p1 Point
	p2 Point
}

func (l Line) draw() {
	// TODO
}

type Rectangle struct {
	p1 Point
	p2 Point
}

func (r Rectangle) draw() {
	// TODO
}

type Canvas struct {
	r Rectangle
}

func (p Point) Line(x float64, y float64) Point {
	to := Point{p.x + x, p.y + y}
	line := Line{p, to}
	append(drawElements, line)
	return to
}
