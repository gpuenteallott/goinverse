package main

import "os"
import "fmt"
import "image"
import "image/color"
import "image/png" // register the PNG format with the image package

func main() {

	args := os.Args[1:]

	// validation of arguments
	if len(args) < 1 {
		fmt.Printf("goinverse. argument input file missing\n")
		os.Exit(1)
	}
	if len(args) < 2 {
		fmt.Printf("goinverse. argument output file missing\n")
		os.Exit(1)
	}

	input := args[0]
	output := args[1]

	fmt.Printf("goinverse. input=%q output=%q\n", input, output)

	infile, err := os.Open(input)
	if err != nil {
		// replace this with real error handling
		panic(err)
	}
	defer infile.Close()

	// Decode will figure out what type of image is in the file on its own.
	// We just have to be sure all the image packages we want are imported.
	srcImage, _, err := image.Decode(infile)
	if err != nil {
		// replace this with real error handling
		panic(err)
	}

	newImage := InvertImage(srcImage)

	// Encode the new image to the output file
	outfile, err := os.Create(output)
	if err != nil {
		// replace this with real error handling
		panic(err)
	}
	defer outfile.Close()

	png.Encode(outfile, newImage)

	fmt.Printf("goinverse. done\n")
}

/**
 * Returns an inverted image.Image from the given one.
 * Alpha channel is preserved.
 */
func InvertImage(srcImage image.Image) image.Image {

	// Create a new image
	bounds := srcImage.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	rectangle := image.Rect(0, 0, w, h)
	newImage := image.NewRGBA(rectangle)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {

			oldColor := srcImage.At(x, y)
			red, green, blue, alpha := oldColor.RGBA()

			// transforming pixel into inverse.
			ired := 255 - (uint8(red >> 8))
			igreen := 255 - (uint8(green >> 8))
			iblue := 255 - (uint8(blue >> 8))
			newColor := color.RGBA{ired, igreen, iblue, uint8(alpha >> 8)}

			newImage.Set(x, y, newColor)
		}
	}

	return newImage
}
