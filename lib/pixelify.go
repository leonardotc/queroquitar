package pixelify

import (
	"math/rand"
	"image/gif"
	"image/color"
	"image"
	"io"
	"image/color/palette"
)

func generateGif(f io.Writer) {
	var images []*image.Paletted
	var delays []int

	for i := 0; i < 5; i++ {
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

	for x := 0; x < width / 25; x++ {
		for y := 0; y < height / 25; y++ {
			color := generateColor()
			for i := 0; i < 25; i++ {
				for j := 0; j < 25; j++ {
					image.Set(x * 25 + i, y * 25 + j, color)
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
		A: 0xff,
	}
}