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

// A convineice function that creates a color.RGBA struct from color.Color interface.  While fairly
// straitforward, tedious conversions need to be made to change between uint32 and uint8.
func ColorToRGBA(c color.Color) color.RGBA {
	r, g, b, a := c.RGBA()
	return color.RGBA{uint8(r / 0x101), uint8(g / 0x101), uint8(b / 0x101), uint8(a / 0x101)}
}

// Performs alpha blending on the pixel on img.At(x,y).  While draw.Set(x,y) *sets* the color to c,
// Blend uses the existing background to create a new color that is a mix of the existing annew color.
// This is similar to treating the two colors as stained glass windows laying over eachother.  Not that if
// the new color is completely opache, Blend() and Set() have the same result.
func Blend(img draw.Image, x int, y int, c color.Color) {
	dst := ColorToRGBA(img.At(x, y))
	src := ColorToRGBA(c)

	fmt.Println("dst:", dst, "src:", src)

	alpha := src.A + dst.A*(1-src.A)
	r := (src.R + dst.R*(1-src.A))
	g := (src.G + dst.G*(1-src.A))
	b := (src.B + dst.B*(1-src.A))

	col := color.RGBA{R: r, G: g, B: b, A: alpha}
	fmt.Println("col:", col)

	img.Set(x, y, col)
}
