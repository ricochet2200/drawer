package drawer

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

type RectDrawer struct {
	color     color.Color
	rect      image.Rectangle
	src       draw.Image
	thickness int
	isFilled  bool
}

// Draws rect in color onto src.  rect must be well-formed (see Rectangle.Canon()).
func NewRectDrawer(src draw.Image, rect image.Rectangle, color color.Color) *RectDrawer {
	fmt.Println("Creating new RectDrawer")
	return &RectDrawer{color: color, rect: rect, src: src, thickness: 1, isFilled: false}
}

func (this *RectDrawer) SetThickness(thickness int) *RectDrawer {
	this.thickness = thickness
	return this
}

func (this *RectDrawer) SetFilled(isFilled bool) *RectDrawer {
	this.isFilled = isFilled
	return this
}

func (this *RectDrawer) Draw() *RectDrawer {

	thickX := this.thickness

	// Even if the thickness is greater than the size of the rectangle,
	// it should still stay within the rectangle.
	if this.isFilled || thickX > (this.rect.Max.X-this.rect.Min.X)+1 {
		thickX = (this.rect.Max.X-this.rect.Min.X)/2 + 1
	}

	for y := this.rect.Min.Y; y < this.rect.Max.Y; y++ {
		for t := 0; t < thickX; t++ {
			this.src.Set(this.rect.Min.X+t, y, this.color)
			this.src.Set(this.rect.Max.X-1-t, y, this.color)
		}
	}

	if this.isFilled {
		return this
	}

	thickY := this.thickness
	if thickY > (this.rect.Max.Y-this.rect.Min.Y)/2 {
		thickY = (this.rect.Max.Y-this.rect.Min.Y)/2 + 1
	}

	for x := this.rect.Min.X + thickX; x < this.rect.Max.X-thickX; x++ {
		for t := 0; t < thickY; t++ {
			this.src.Set(x, this.rect.Min.Y+t, this.color)
			this.src.Set(x, this.rect.Max.Y-1-t, this.color)
		}
	}
	return this
}
