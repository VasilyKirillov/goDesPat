/*
Example of flyweight
reusing pen objects that has been created
*/
package main

import "fmt"

type Color int

const (
	RED = iota
	GREEN
	BLUE
)

type BrushSizeType string

var BrushSize = struct {
	THICK BrushSizeType
	THIN  BrushSizeType
}{
	THICK: "thick",
	THIN:  "thin",
}

type Pen interface {
	setColor(color Color)
	draw(content string)
}

type ThickPen struct {
	color Color
}

func (p *ThickPen) setColor(color Color) {
	fmt.Printf("Using THICK pen with %v ink\n", getName(color))
	p.color = color
}

func (p ThickPen) draw(content string) {
	fmt.Printf("Writing thickly %v with %v color\n", content, getName(p.color))
}

type ThinPen struct {
	color Color
}

func (p *ThinPen) setColor(color Color) {
	fmt.Printf("Using THIN pen with %v ink\n", getName(color))
	p.color = color
}

func (p ThinPen) draw(content string) {
	fmt.Printf("Writing thinly %v with %v color\n", content, getName(p.color))
}

func getName(color Color) string {
	switch color {
	case 0:
		return "RED"
	case 1:
		return "GREEN"
	case 2:
		return "BLUE"
	default:
		return ""
	}
}

type PenFactory struct {
	penMap map[string]Pen
}

func NewPenFactory() PenFactory {
	return PenFactory{
		penMap: make(map[string]Pen, 10),
	}
}

func (pf PenFactory) GetThickPen(color Color) Pen {
	key := fmt.Sprintf("%v-THICK", getName(color))
	pen, ok := pf.penMap[key]
	if !ok {
		pen = &ThickPen{color}
		pf.penMap[key] = pen
	}
	return pen
}
func (pf PenFactory) GetThinPen(color Color) Pen {
	key := fmt.Sprintf("%v-THIN", getName(color))
	pen, ok := pf.penMap[key]
	if !ok {
		pen = &ThinPen{color}
		pf.penMap[key] = pen
	}
	return pen
}
