package main

import (
	"fibonnaci-art/sigil"
	"fmt"

	"github.com/fogleman/gg"
)

func main() {
	var modulo uint
	var width, height int
	var radius, lineWidth float64
	var path string
	fmt.Printf("Generate .PNG from Pisano Period of entered modulo\nPlease input the following values:\n\n")
	fmt.Printf("Modulo: ")
	fmt.Scanf("%d", &modulo)
	fmt.Printf("Image size, one edge: ")
	fmt.Scanf("%d", &width)
	height = width
	fmt.Printf("Radius: ")
	fmt.Scanf("%f", &radius)
	fmt.Printf("Line width: ")
	fmt.Scanf("%f", &lineWidth)
	fmt.Printf("Save path: ")
	fmt.Scanf("%s", &path)
	sigil, err := sigil.MakeSigil(modulo, width, height, radius, lineWidth)
	if err != nil {
		fmt.Printf("Error: %q", err)
		fmt.Printf("s: %q", sigil)
		return
	}

	gg.SavePNG(path, sigil)
}
