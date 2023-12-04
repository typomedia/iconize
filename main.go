package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/typomedia/ico"
	"golang.org/x/image/draw"
	"image"
	"image/color"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var (
		input string
		data  []byte
		err   error
	)

	outfile := pflag.StringP("out", "o", "", "outfile ico file")
	version := pflag.BoolP("version", "V", false, "show version")
	pflag.Parse()

	if *version {
		fmt.Println("iconize 0.2.0 <philipp@typo.media>")
		os.Exit(0)
	}

	input = pflag.Arg(0)
	stat, _ := os.Stdin.Stat()

	switch true {
	case input != "":
		data, err = os.ReadFile(input)
	case (stat.Mode() & os.ModeCharDevice) == 0:
		data, err = io.ReadAll(os.Stdin)
	default:
		log.Fatal("no input data given")
	}
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	width, height := size(img)
	size := max(width, height)

	alpha := image.NewRGBA(image.Rect(0, 0, size, size))

	rgba := color.RGBA{}
	draw.Draw(alpha, alpha.Bounds(), &image.Uniform{C: rgba}, image.Point{}, draw.Src)

	draw.Draw(alpha, img.Bounds().Add(
		image.Point{
			X: (size - img.Bounds().Dx()) / 2,
			Y: (size - img.Bounds().Dy()) / 2,
		}), img, image.Point{}, draw.Over)

	icon := ico.NewIcon()
	// https://learn.microsoft.com/en-us/windows/win32/uxguide/vis-icons
	sizes := []int{256, 128, 64, 48, 32, 24, 16}
	for _, size := range sizes {
		resizedImg := scale(alpha, size, size)
		icon.AddPng(resizedImg)
	}

	enc, err := icon.Encode()
	if err != nil {
		log.Fatal(err)
	}

	// write to stdout
	if *outfile == "" {
		fmt.Println(string(enc))
		os.Exit(0)
	}

	// write to file
	icoFile, err := os.Create(*outfile)
	if err != nil {
		log.Fatal(err)
	}
	defer icoFile.Close()

	icoFile.Write(enc)
	os.Exit(0)

}

func max(width, height int) int {
	size := width
	if height > width {
		size = height
	}
	return size
}

func size(img image.Image) (int, int) {
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y
	return width, height
}

func scale(img image.Image, width, height int) image.Image {
	resizedImg := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(resizedImg, resizedImg.Bounds(), img, img.Bounds(), draw.Over, nil)

	return resizedImg
}

func name(filePath string) string {
	fileName := filepath.Base(filePath)
	extension := filepath.Ext(fileName)
	return fileName[0 : len(fileName)-len(extension)]
}
