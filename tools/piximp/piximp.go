package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

func main() {
	srcImage, dstFile := handleArgs(os.Args)

	log.Println("reading image ...")

	fileReader, err := os.Open(srcImage)
	panicErr(err)
	defer fileReader.Close()

	img, err := png.Decode(fileReader)
	panicErr(err)

	resultFile, err := os.Create(dstFile)
	panicErr(err)
	defer resultFile.Close()

	writePixels(resultFile, img)

	log.Println("done")
	log.Println("function saved at ", dstFile)
}

func handleArgs(args []string) (srcImage string, dstFile string) {
	if len(os.Args) != 3 {
		log.Fatal("expected 2 args. piximp [source_image] [dest_file]")
	}

	srcImage, err := filepath.Abs(args[1])
	panicErr(err)

	dstFile, err = filepath.Abs(args[2])
	panicErr(err)

	return
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func writePixels(resFile *os.File, srcImage image.Image) {
	width, height := srcImage.Bounds().Dx(), srcImage.Bounds().Dy()
	size := width * height
	writeLine(resFile, "import \"github.com/sparkoo/sparkengine/scene\"")
	writeLine(resFile, "")
	writeLine(resFile, "func initPixels() []scene.Pixel {")
	writeLine(resFile, fmt.Sprintf("\tpixels := make([]scene.Pixel, %d)", size))

	pi := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, a := getRGBA(srcImage.At(x, y))
			writeLine(
				resFile, fmt.Sprintf("\tpixels[%v] = scene.NewPixel(%d, %d, scene.Color{%d, %d, %d, %d})",
					pi, x, y, r, g, b, a))
			pi++
		}
	}
	writeLine(resFile, "\treturn pixels")
	writeLine(resFile, "}")
}

func writeLine(file *os.File, line string) {
	_, err := file.WriteString(line + "\n")
	panicErr(err)
}

func getRGBA(c color.Color) (uint8, uint8, uint8, uint8) {
	switch rgba := c.(type) {
	case color.RGBA:
		return rgba.R, rgba.G, rgba.B, rgba.A
	case color.NRGBA:
		return rgba.R, rgba.G, rgba.B, rgba.A
	default:
		panic("can't do this format")
	}
}
