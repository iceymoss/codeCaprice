package main

import "fmt"

func flipImage(width int, pixels []int) []int {
	h := len(pixels) / (4 * width)
	for i := 0; i < h/2; i++ {
		for j := 0; j < width*4; j++ {
			cur := i*width*4 + j
			syIndex := (h-i-1)*width*4 + j
			pixels[cur], pixels[syIndex] = pixels[syIndex], pixels[cur]
		}
	}
	return pixels
}

func main() {
	fmt.Println(flipImage(2, []int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4}))
}
