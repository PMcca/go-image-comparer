package imgCompare

import "image"

// Transformer makes transformations to images including size and colour.

// Takes an image and converts its color to grayscale
func toGrayscale(img image.Image) image.Image {

	gray := image.NewGray(img.Bounds())

	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			gray.Set(x, y, img.At(x, y))
		}
	}

	return gray
}
