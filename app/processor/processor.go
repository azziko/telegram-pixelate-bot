package processor

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
)

const (
	blockSize = 8
)

type Pixel struct {
	R int
	G int
	B int
	A int
}

func Pixelate(filepath string, data io.ReadCloser) error {
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("Pixelate failed: %v", err)
	}

	pixels, err := getPixels(data)
	if err != nil {
		return fmt.Errorf("getPixels failed: %v", err)
	}

	output := pixelateGivenPixels(pixels)

	if err := writeOutput(output, file); err != nil {
		file.Close()
		return fmt.Errorf("writeOutput failed: %v", err)
	}

	return file.Close()
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	var pixel Pixel

	pixel.R = int(r / 257)
	pixel.G = int(g / 257)
	pixel.B = int(b / 257)
	pixel.A = int(a / 257)

	return pixel
}

func getPixels(file io.Reader) ([][]Pixel, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for x := 0; x < width; x++ {
		var row []Pixel
		for y := 0; y < height; y++ {
			color := img.At(x, y)
			row = append(row, rgbaToPixel(color.RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}

func averagePixels(a Pixel, b Pixel) Pixel {
	var result Pixel

	result.R = int((a.R + b.R) / 2)
	result.G = int((a.G + b.G) / 2)
	result.B = int((a.B + b.B) / 2)
	result.A = int((a.A + b.A) / 2)

	return result
}

func pixelateGivenPixels(pixels [][]Pixel) [][]Pixel {
	width, height := len(pixels), len(pixels[0])
	result := pixels

	averages := make([][]Pixel, int(width/blockSize)+1)
	for i := range averages {
		averages[i] = make([]Pixel, int(height/blockSize)+1)
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			avgX := int(x / blockSize)
			avgY := int(y / blockSize)
			averages[avgX][avgY] = averagePixels(averages[avgX][avgY], pixels[x][y])
		}
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			result[x][y] = averages[int(x/blockSize)][int(y/blockSize)]
		}
	}

	return result
}

func pixelToRGBA(pixel Pixel) color.RGBA {
	var rgba color.RGBA

	rgba.R = uint8(pixel.R)
	rgba.G = uint8(pixel.G)
	rgba.B = uint8(pixel.B)
	rgba.A = uint8(pixel.A)

	return rgba
}

func writeOutput(pixels [][]Pixel, file *os.File) error {
	width, height := len(pixels), len(pixels[0])

	img := image.NewRGBA64(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			pixel := pixels[x][y]
			img.Set(x, y, pixelToRGBA(pixel))
		}
	}

	return png.Encode(file, img)
}
