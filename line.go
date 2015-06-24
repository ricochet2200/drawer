package drawer

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

type LineDrawer struct {
	img       draw.Image
	start     image.Point
	end       image.Point
	thickness float64
	color     color.Color
}

func NewLineDrawer(img draw.Image, start image.Point, end image.Point, color color.Color) *LineDrawer {
	return &LineDrawer{img: img, start: start, end: end, thickness: 1, color: color}
}

func (this *LineDrawer) SetStart(start image.Point) *LineDrawer {
	this.start = start
	return this
}

func (this *LineDrawer) SetEnd(end image.Point) *LineDrawer {
	this.end = end
	return this
}

func (this *LineDrawer) Draw() *LineDrawer {

	x0 := this.start.X
	x1 := this.end.X
	y0 := this.start.Y
	y1 := this.end.Y

	dx := int(math.Abs(float64(x1 - x0)))
	sx := -1
	if x0 < x1 {
		sx = 1
	}

	dy := int(math.Abs(float64(y1 - y0)))
	sy := -1
	if y0 < y1 {
		sy = 1
	}

	err := dx - dy
	ed := 1.0
	if dx+dy != 0 {
		ed = math.Sqrt(float64(dx*dx) + float64(dy*dy))
	}
	tmpColor := ColorToRGBA(this.color)

	// Anti-alias calculations assume no alpha on the original color. Offset if that is not true.
	alpha := float64(tmpColor.A)
	for wd := (this.thickness + 1) / 2; ; { /* pixel loop */
		tmpColor.A = uint8(math.Max(0, alpha-255*(math.Abs(float64(err-dx+dy))/ed-wd+1)))
		this.img.Set(x0, y0, tmpColor)

		e2 := err
		x2 := x0
		if 2*e2 >= -dx { /* x step */
			e2 += dy
			y2 := y0

			for ; float64(e2) < ed*wd && (y1 != y2 || dx > dy); e2 += dx {
				y2 += sy
				tmpColor.A = uint8(math.Max(0, alpha-255*(math.Abs(float64(e2))/ed-wd+1)))
				Blend(this.img, x0, y2, tmpColor)
			}
			if x0 == x1 {
				break
			}
			e2 = err
			err -= dy
			x0 += sx
		}

		if 2*e2 <= dy { /* y step */
			for e2 = dx - e2; float64(e2) < ed*wd && (x1 != x2 || dx < dy); e2 += dy {
				x2 += sx
				tmpColor.A = uint8(math.Max(0, alpha-255*(math.Abs(float64(e2))/ed-wd+1)))
				Blend(this.img, x2, y0, tmpColor)
			}
			if y0 == y1 {
				break
			}
			err += dx
			y0 += sy
		}
	}
	return this
}
