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

	start := image.Pt(100, 100)
	end := image.Pt(0, 0)
	ld := drawer.NewLineDrawer(src, start, end, color.RGBA{255, 0, 0, 255}).Draw()
	draw(ld, src, "negative.png")

	start = image.Pt(0, 100)
	end = image.Pt(100, 0)
	ld.SetStart(end).SetEnd(start).Draw()
	draw(ld, src, "positive.png")
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
