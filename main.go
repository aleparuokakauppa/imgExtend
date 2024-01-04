package main

import (
	"fmt"
	"image"
	"strconv"
	"image/color"
	"image/png"
	"log"
	"os"
    "path/filepath"
    "strings"
)

func main() {
    if len(os.Args) < 3 {
        log.Fatal("Not enough arguments. Need 2: Filepath, new height")
    }
    // First argument is path to file
    imageFile, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal("Image not found")
        panic(err)
    }
    // Second argument is new height
    newHeight, err := strconv.Atoi(os.Args[2])
    if err != nil {
        log.Fatal("New height couldn't be converted to type Int.")
        panic(err)
    }
    imageData, err := png.Decode(imageFile)
    if err != nil {
        log.Println("Couldn't decode image data as a PNG.")
        log.Println("Provided image is probably not a PNG file or the file is broken")
        panic(err)
    }

    if newHeight < imageData.Bounds().Dy() {
        log.Printf("New height must be larger than the original (%d)\n", imageData.Bounds().Max.Y)
        panic(fmt.Errorf("Invalid height"))
    }

    // Get the color at the center top of the image
    fillColor := imageData.At((imageData.Bounds().Max.X)/2, 0)
    // Create the new image in memory
    newImg := image.NewRGBA(image.Rect(0, 0, imageData.Bounds().Dx(), newHeight))
    // Copy the image data from the previous image into the new image created
    copyImgBottom(imageData, newImg)
    // Get the Y-axis offset between the new and the previous image
    // Fill the top-portion of the image with the fillColor
    offset := newImg.Bounds().Max.Y - imageData.Bounds().Dy()
    fillTop(fillColor, offset, newImg)

    fileBasepath := filepath.Base(os.Args[1])
    fileExt := filepath.Ext(os.Args[1])
    fileWithoutExt := strings.TrimSuffix(fileBasepath, fileExt)
    newFileName := fileWithoutExt + "_extended" + fileExt

    // Create the new file and write the newImg contents into it
    newFile, err := os.Create(newFileName)
    if err != nil {
        log.Fatal(err)
        panic(err)
    }
    defer newFile.Close()
    err = png.Encode(newFile, newImg)
    if err != nil {
        log.Fatal(err)
        panic(err)
    }
    fmt.Printf("Image saved as %s\n", newFileName)
}

func copyImgBottom(src image.Image, dest *image.RGBA) error {
    destBoundsY := dest.Bounds().Max.Y
    boundsOffsetY := destBoundsY - src.Bounds().Max.Y
    for y := 0; y < src.Bounds().Dy(); y++ {
        for x := 0; x < src.Bounds().Max.X; x++ {
            col := src.At(x, y)
            dest.Set(x, y+boundsOffsetY, col)
        }
    }
    return nil
}

//func fillTop(r uint32, g uint32, b uint32, a uint32, offset int, dest *image.RGBA) error {
func fillTop(fillColor color.Color, offset int, dest *image.RGBA) error {
    for y:= 0; y < offset; y++ {
        for x := 0; x < dest.Bounds().Dx(); x++ {
            dest.Set(x, y, fillColor)
        }
    }
    return nil
}
