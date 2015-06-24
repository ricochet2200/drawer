package drawer

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

// A convinience function that sets all of the pixels in img to color
func Fill(img draw.Image, color color.Color) {
	draw.Draw(img, img.Bounds(), &image.Uniform{color}, image.ZP, draw.Src)
}

func ColorToRGBA(c color.Color) color.RGBA {
	r, g, b, a := c.RGBA()
	return color.RGBA{uint8(r / 0x101), uint8(g / 0x101), uint8(b / 0x101), uint8(a / 0x101)}
}

func Blend(img draw.Image, x int, y int, c color.Color) {
	dst := ColorToRGBA(img.At(x, y))
	src := ColorToRGBA(c)

	fmt.Println("dst:", dst, "src:", src)

	//	dst: {0 255 255 255} src: {255 0 0 74}
	//	col: {1 0 0 147}

	alpha := src.A + dst.A*(1-src.A)
	r := (src.R + dst.R*(1-src.A))
	g := (src.G + dst.G*(1-src.A))
	b := (src.B + dst.B*(1-src.A))

	col := color.RGBA{R: r, G: g, B: b, A: alpha}
	fmt.Println("col:", col)

	img.Set(x, y, col)
}
