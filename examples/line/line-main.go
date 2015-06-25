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
	src := image.NewRGBA(image.Rect(0, 0, 100, 100))
	drawer.Fill(src, color.RGBA{0, 255, 255, 255})

	ld := drawer.NewLineDrawer(src, image.Pt(100, 100), image.Pt(0, 0), color.RGBA{255, 0, 0, 255}).Draw()
	draw(ld, src, "negative.png")

	ld.SetStart(image.Pt(0, 100)).SetEnd(image.Pt(100, 0)).Draw()
	draw(ld, src, "positive.png")

	ld.SetStart(image.Pt(0, 50)).SetEnd(image.Pt(100, 50)).SetThickness(5).Draw()
	draw(ld, src, "thick.png")
}

func draw(drawer *drawer.LineDrawer, src image.Image, filename string) {
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
