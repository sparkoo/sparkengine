package image

import (
	"github.com/sparkoo/sparkengine/core/scene"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
)

func LoadFullImage(imagePath string) []scene.Pixel {
	img := load(imagePath)
	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	return loadImage(img, 0, 0, width, height)
}

func LoadImage(imagePath string, x int, y int, width int, height int) []scene.Pixel {
	return loadImage(load(imagePath), x, y, width, height)
}

func loadImage(img image.Image, xStart int, yStart int, width int, height int) []scene.Pixel {
	xMax, yMax := xStart+width, yStart+height
	pI := 0
	pixels := make([]scene.Pixel, width*height)

	// [xi;yi] are coordinates of result pixels
	// [x;y] are coordinates in original image
	for x, xi := xStart, 0; x < xMax; x, xi = x+1, xi+1 {
		for y, yi := yStart, 0; y < yMax; y, yi = y+1, yi+1 {
			pixels[pI] = scene.NewPixel(xi, yi, getRGBA(img.At(x, y)))
			pI++
		}
	}

	return pixels
}

func load(imagePath string) image.Image {
	imgFullPath, err := filepath.Abs(imagePath)
	if err != nil {
		panic(err)
	}

	imgReader, err := os.Open(imgFullPath)
	if err != nil {
		panic(err)
	}
	defer imgReader.Close()

	img, err := png.Decode(imgReader)
	if err != nil {
		panic(err)
	}

	return img
}

func getRGBA(c color.Color) [4]uint8 {
	switch rgba := c.(type) {
	case color.RGBA:
		return [4]uint8{rgba.R, rgba.G, rgba.B, rgba.A}
	case color.NRGBA:
		return [4]uint8{rgba.R, rgba.G, rgba.B, rgba.A}
	default:
		panic("can't do this format")
	}
}
