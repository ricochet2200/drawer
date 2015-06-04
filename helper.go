package drawer

import (
	"image"
	"image/color"
	"image/draw"
)

// A convinience function that sets all of the pixels in img to color
func Set(img draw.Image, color color.Color) {
	draw.Draw(img, img.Bounds(), &image.Uniform{color}, image.ZP, draw.Src)
}
