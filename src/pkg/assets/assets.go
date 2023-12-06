package assets

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
)

type Assets struct{}

func (a *Assets) GetDirs(path string) []string {
	return []string{
		filepath.Join(path, "assets"),
	}
}

func (a *Assets) GetFiles(path, name string) map[string]string {

	filename := filepath.Join(path, "assets", name+".png")
	a.genIcon(filename)

	return map[string]string{}
}

func (a *Assets) genIcon(filename string) error {

	// Set the size of the image (e.g., 128x128 pixels)
	const width, height = 128, 128

	// Create a new RGBA image with a transparent background
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Define the colors
	redColor := color.RGBA{255, 0, 0, 128}   // Semi-transparent red
	greenColor := color.RGBA{0, 255, 0, 128} // Semi-transparent green
	blueColor := color.RGBA{0, 0, 255, 128}  // Semi-transparent blue

	// Radius of the circles
	const radius = 30

	// Draw overlapping circles
	drawCircle(img, image.Point{width / 3, height / 3}, radius, redColor)
	drawCircle(img, image.Point{2 * width / 3, height / 3}, radius, greenColor)
	drawCircle(img, image.Point{width / 2, 2 * height / 3}, radius, blueColor)

	// Create the file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the image to PNG and write to the file
	return png.Encode(file, img)
}

func drawCircle(img *image.RGBA, center image.Point, radius int, col color.Color) {
	for y := -radius; y <= radius; y++ {
		for x := -radius; x <= radius; x++ {
			if x*x+y*y <= radius*radius {
				img.Set(center.X+x, center.Y+y, col)
			}
		}
	}
}
