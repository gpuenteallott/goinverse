package main

import "os"
import "fmt"
import "image"
import "image/color"
import "image/png" // register the PNG format with the image package

func main() {

	args := os.Args[1:]
	input := args[0]
	output := args[1]

    fmt.Printf("goinverse. Input=%q Output=%q\n", input, output )

    infile, err := os.Open( input )
    if err != nil {
        // replace this with real error handling
        panic(err)
    }
    defer infile.Close()

    // Decode will figure out what type of image is in the file on its own.
    // We just have to be sure all the image packages we want are imported.
    src, _, err := image.Decode(infile)
    if err != nil {
        // replace this with real error handling
        panic(err)
    }

    // Create a new grayscale image
    bounds := src.Bounds()
    w, h := bounds.Max.X, bounds.Max.Y
    rectangle := image.Rect(0, 0, w, h)
    grayImage := image.NewGray( rectangle )
    for x := 0; x < w; x++ {
        for y := 0; y < h; y++ {

            oldColor := src.At(x, y)
            red, green, blue, alpha := oldColor.RGBA()


            averaged := uint8( ( (red + green + blue) / 3 ) >> 8 )
            grayColor := color.RGBA{ averaged, averaged, averaged, uint8(alpha>>8) }

            // grayColor := color.GrayModel.Convert(oldColor)

            grayImage.Set(x, y, grayColor)
        }
    }

    // Encode the grayscale image to the output file
    outfile, err := os.Create( output )
    if err != nil {
        // replace this with real error handling
        panic(err)
    }
    defer outfile.Close()
    png.Encode(outfile, grayImage)
}