package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	inDir := flag.String("input", "/mnt/d/imagens/png", "input directory")
	outDir := flag.String("output", "/mnt/d/imagens/output", "output directory")
	FilterDir(*inDir, *outDir)
}

func FilterDir(dir, dirout string) {
	files, err := FindPngFiles(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	imagensChannel := ProcessImages(files)
	for i := range files {
		image := <-imagensChannel
		negativeImage := ApplyFilter(image, NegativeFilter)
		grayImage := ApplyFilter(image, GrayFilter)
		sepiaImage := SepiaFilter(image)
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
	}
}
