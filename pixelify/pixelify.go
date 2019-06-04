package main

import (
	"math/rand"
	"image/gif"
	"image/color"
	"image"
	"io"
	"image/color/palette"
)

const Opaque = 0xff
const PixelSize = 10
const AnimationFrames = 5

func generateGif(f io.Writer) {
	var images []*image.Paletted
	var delays []int

	for frameIndex := 0; frameIndex < AnimationFrames; frameIndex++ {
		images = append(images, generatePixels(250, 250))
		delays = append(delays, 1)
	}
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}

func generatePixels(width int, height int)*image.Paletted {
	rect := image.Rectangle{
		Min: image.Pt(0, 0),
		Max: image.Pt(width, height)}
	image := image.NewPaletted(rect, palette.Plan9)

	for x := 0; x < width / PixelSize; x++ {
		for y := 0; y < height / PixelSize; y++ {
			color := generateColor()
			for offsetX := 0; offsetX < PixelSize; offsetX++ {
				for offsetY := 0; offsetY < PixelSize; offsetY++ {
					image.Set(x * PixelSize + offsetX, y * PixelSize + offsetY, color)
				}
			}
		}
	}

	return image
}

func generateColor()color.NRGBA {
	rgb := make([]uint8, 3)
	rand.Read(rgb)

	return color.NRGBA{
		R: rgb[0],
		G: rgb[1],
		B: rgb[2],
		A: Opaque,
	}
}