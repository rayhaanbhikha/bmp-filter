package main

func greyScaleFilter(pixels [][]Pixel) (newPixels []Pixel) {
	for _, pixelRow := range pixels {
		for _, pixel := range pixelRow {
			avg := (pixel.b + pixel.g + pixel.r) / 3
			newPixels = append(newPixels, Pixel{b: avg, g: avg, r: avg})
		}
	}
	return
}
