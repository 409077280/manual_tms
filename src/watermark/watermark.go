package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	// The objective picture is sam.jpg
	imgb, err1 := os.Open("sam.jpg")
	if err1 != nil {
		fmt.Println(err1)
	}
	img, err2 := jpeg.Decode(imgb)
	if err2 != nil {
		fmt.Println(err2)
	}
	defer imgb.Close()

	wmb, err3 := os.Open("text.png")
	if err3 != nil {																																																																	
		fmt.Println(err3)
	}
	watermark, err4 := png.Decode(wmb)
	if err4 != nil {
		fmt.Println(err4)
	}
	defer wmb.Close()

	// Watermark draw to lower right of objective picture
	offset := image.Pt(img.Bounds().Dx()-watermark.Bounds().Dx()-10, img.Bounds().Dy()-watermark.Bounds().Dy()-10)
	b := img.Bounds()
	m := image.NewNRGBA(b)

	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

	// Build new.jpg，and setting picture quality
	imgw, err5 := os.Create("new.jpg")
	if err5 != nil {
		fmt.Println(err5)
	}
	jpeg.Encode(imgw, m, &jpeg.Options{100})

	defer imgw.Close()

	fmt.Println("Watermark was draw in it...")
}