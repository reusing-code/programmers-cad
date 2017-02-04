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
	addDrawable(line)
	return to
}

func addDrawable(d Drawable) {
	m := len(drawElements)
	n := m + 1
	if n > cap(drawElements) {
		// allocate double what's needed, for future growth.
		newSlice := make([]Drawable, (n+1)*2)
		copy(newSlice, drawElements)
		drawElements = newSlice
	}
	drawElements = drawElements[0:n]
	drawElements[m] = d;
}