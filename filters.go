package main

import (
	"image"
	"image/color"
	"math"
)

type FilterFunc func(r, g, b uint32) (uint8, uint8, uint8)

func ApplyFilter(img image.Image, filter FilterFunc) image.Image {
	bounds := img.Bounds()
	filteredImg := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			oldColor := img.At(x, y)
			r, g, b, a := oldColor.RGBA()

			newR, newG, newB := filter(r, g, b)

			filteredColor := color.RGBA{
				R: newR,
				G: newG,
				B: newB,
				A: uint8(a >> 8),
			}
			filteredImg.Set(x, y, filteredColor)
		}
	}
	return filteredImg
}

func NegativeFilter(r, g, b uint32) (uint8, uint8, uint8) {
	negativeR := uint8(255 - r>>8)
	negativeG := uint8(255 - g>>8)
	negativeB := uint8(255 - b>>8)
	return negativeR, negativeG, negativeB
}

func GrayFilter(r, g, b uint32) (uint8, uint8, uint8) {
	grayValue := uint8((r + g + b) / 3 >> 8)
	return grayValue, grayValue, grayValue
}

func SepiaFilter(img image.Image) image.Image {
	bounds := img.Bounds()
	filteredImg := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := GetPixel(img, x, y)

			// Extraindo compontentes de cor
			a, r, g, b := GetColorComponents(pixel)

			// Aplicando o filtro sepia
			sepiaR := 0.393*float64(r) + 0.769*float64(g) + 0.189*float64(b)
			sepiaG := 0.349*float64(r) + 0.686*float64(g) + 0.168*float64(b)
			sepiaB := 0.272*float64(r) + 0.534*float64(g) + 0.131*float64(b)

			// Ajustar os valores de sepia para o máximo de 255
			sepiaR = math.Min(sepiaR, 255)
			sepiaG = math.Min(sepiaG, 255)
			sepiaB = math.Min(sepiaB, 255)

			// Converter os valores de sepia para uint8
			sepiaRUint8 := uint8(sepiaR)
			sepiaGUint8 := uint8(sepiaG)
			sepiaBUint8 := uint8(sepiaB)

			// Setando os novos valores de cor
			filteredImg.SetRGBA(x, y, color.RGBA{
				R: sepiaRUint8,
				G: sepiaGUint8,
				B: sepiaBUint8,
				A: a,
			})
		}
	}

	return filteredImg
}
