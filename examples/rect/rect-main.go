package main

import (
	"../../../drawer"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	src := image.NewRGBA(image.Rect(0, 0, 10, 10))
	drawer.Set(src, color.RGBA{0, 0, 255, 255})

	rect := image.Rect(0, 0, 10, 10)
	rd := drawer.NewRectDrawer(src, rect, color.RGBA{0, 255, 0, 255})

	draw(rd, src, "full-square.png")
	rd.SetThickness(3).Draw()
	draw(rd, src, "three-pixel-border.png")
	rd.SetThickness(30).Draw()
	draw(rd, src, "oversized-border.png")
	rd.SetFilled(true)
	draw(rd, src, "filled.png")
	rd.SetColor(color.RGBA{0, 0, 0, 255})
	draw(rd, src, "black.png")

}

func draw(drawer *drawer.RectDrawer, src image.Image, filename string) {
	out, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer out.Close()
	fmt.Println("Writing output to:", filename)

	err = png.Encode(out, src)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
