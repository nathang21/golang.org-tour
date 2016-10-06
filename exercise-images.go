// https://tour.golang.org/methods/25
package main

import (
    "golang.org/x/tour/pic"
    "image/color"
    "image"
)

type Image struct{
    Rect image.Rectangle
}

func (image Image) ColorModel() color.Model {
        return color.RGBAModel
}
func (i Image) Bounds() image.Rectangle {
    return i.Rect
}
func (image Image) At(x, y int) color.Color {
    return color.RGBA{uint8(x*y), uint8(y+x), 255, 255}
}

func main() {
    m := Image{image.Rect(0,0,500,500)}
    pic.ShowImage(m)
}
