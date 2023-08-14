package main

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetColorComponents(p uint32) (a, r, g, b uint8) {
	a = uint8((p >> 24) & 0xff)
	r = uint8((p >> 16) & 0xff)
	g = uint8((p >> 8) & 0xff)
	b = uint8(p & 0xff)
	return a, r, g, b
}

func GetPixel(img image.Image, x, y int) uint32 {
	rgbaColor := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
	pixel := (uint32(rgbaColor.A) << 24) | (uint32(rgbaColor.R) << 16) | (uint32(rgbaColor.G) << 8) | uint32(rgbaColor.B)
	return pixel
}

func ReadImage(file string) (image.Image, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		log.Printf("Error decoding image %s:%s\n", file, err)
		return nil, err
	}
	return img, nil
}

func ImageToBytes(img image.Image) []byte {
	b := new(bytes.Buffer)
	if err := png.Encode(b, img); err != nil {
		log.Fatal(err)
	}
	return b.Bytes()
}

func CreatePng(img image.Image, outputFile string) error {
	//path := filepath.Join("/mnt/d/imagens/output", outputFile)
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		return err
	}
	return nil
}

func CreateJpg(img image.Image, outputFile string) error {
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	quality := jpeg.Options{
		Quality: 90,
	}
	if err := jpeg.Encode(f, img, &quality); err != nil {
		return err
	}
	return nil
}

func ResizeImage(img image.Image, weight, height int) image.Image {
	resizedImage := image.NewRGBA(image.Rect(0, 0, weight, height))
	draw.Draw(resizedImage, resizedImage.Bounds(), img, img.Bounds().Min, draw.Src)
	return resizedImage
}

func FindImageFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && (strings.HasSuffix(info.Name(), ".png") || strings.HasSuffix(info.Name(), ".jpeg") || strings.HasSuffix(info.Name(), ".jpg")) {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func ProcessImages(files []string) chan image.Image {
	results := make(chan image.Image)
	for _, file := range files {
		go func(file string) {
			img, err := ReadImage(file)
			if err != nil {
				log.Println(err)
			}
			results <- img
		}(file)

	}
	return results
}
