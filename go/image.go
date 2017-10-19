package main

import (
	"image"	
	"image/color"
	"golang.org/x/tour/pic"
)

type Image struct{
	Width, Height int
	colr uint8	
}

func (r *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, r.Width, r.Height)
}

func (r *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (r *Image) At(x, y int) color.Color {
	return color.RGBA{ r.colr + uint8(x), r.colr + uint8(y), 255, 255 }
}

func main() {
	m := Image{90, 90, 128}
	pic.ShowImage(&m)
}