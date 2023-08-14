package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	inDir := flag.String("input", "/mnt/d/imagens/png", "input directory")
	outDir := flag.String("output", "/mnt/d/imagens/output", "output directory")
	FilterDir(*inDir, *outDir)
	fmt.Printf("O tempo de execução foi de %f segundos\n", time.Since(start).Seconds())
}

func FilterDir(dir, dirout string) {
	files, err := FindImageFiles(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	imagensChannel := ProcessImages(files)
	for i, file := range files {
		image := <-imagensChannel
		negativeImage := ApplyFilter(image, NegativeFilter)
		grayImage := ApplyFilter(image, GrayFilter)
		sepiaImage := SepiaFilter(image)
		extension := strings.ToLower(filepath.Ext(file))
		if extension == ".png" {
			//fileName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
			outputNegativeFile := fmt.Sprintf("%d.png", i)
			outputGrayFile := fmt.Sprintf("%d_gray.png", i)
			outputSepiaFile := fmt.Sprintf("%d_sepia.png", i)
			outputNegativePath := filepath.Join(dirout, outputNegativeFile)
			outputGrayPath := filepath.Join(dirout, outputGrayFile)
			outputSepiaPath := filepath.Join(dirout, outputSepiaFile)
			err := CreatePng(negativeImage, outputNegativePath)
			_ = CreatePng(grayImage, outputGrayPath)
			_ = CreatePng(sepiaImage, outputSepiaPath)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
		} else if extension == ".jpeg" || extension == ".jpg" {
			outputNegativeFile := fmt.Sprintf("%d.jpeg", i)
			outputGrayFile := fmt.Sprintf("%d_gray.jpeg", i)
			outputSepiaFile := fmt.Sprintf("%d_sepia.jpeg", i)
			outputNegativePath := filepath.Join(dirout, outputNegativeFile)
			outputGrayPath := filepath.Join(dirout, outputGrayFile)
			outputSepiaPath := filepath.Join(dirout, outputSepiaFile)
			err := CreateJpg(negativeImage, outputNegativePath)
			_ = CreateJpg(grayImage, outputGrayPath)
			_ = CreateJpg(sepiaImage, outputSepiaPath)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
		} else {
			log.Println("Extensão inválida")
		}
	}
}
