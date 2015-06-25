package drawer

/*
This package contains several structs for drawing, all following a similar convention:

Structs have a constructor function that takes basic arguments for that type, ie:

func NewLineDrawer(draw.Image, image.Point, image.Point, color.Color)

All functions in the Drawer structs return a pointer to itself to allow chaining:

NewLineDrawer(src, start, end, color.RGBA{255, 0, 0, 255}).SetThickness(10).Draw()

This also allows you to reuse an instance:

lineDrawer := NewLineDrawer(draw.Image, image.Point, image.Point, color.Color).Draw()
lineDrawer = lineDrawer.SetStart(image.Pt(0, 50)).Draw()

*/
