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
	src := image.NewRGBA(image.Rect(0, 0, 2, 10))
	drawer.Set(src, color.RGBA{0, 0, 255, 255})

	rect := image.Rect(1, 1, 1, 9)

	rd := drawer.NewRectDrawer(src, rect, color.RGBA{0, 255, 0, 255})
	rd.SetFilled(false)
	rd.SetThickness(20).Draw()

	filename := "./output.png"
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
