package main

func reflectFilter(pixels [][]Pixel) (newPixels []Pixel) {
	for _, pixelRow := range pixels {
		n := len(pixelRow)
		for i := 0; i < n/2; i++ {
			pixelRow[i], pixelRow[n-1-i] = pixelRow[n-1-i], pixelRow[i]
		}
		newPixels = append(newPixels, pixelRow...)
	}
	return
}

func greyScaleFilter(pixels [][]Pixel) (newPixels []Pixel) {
	for _, pixelRow := range pixels {
		for _, pixel := range pixelRow {
			avg := (pixel.b + pixel.g + pixel.r) / 3
			newPixels = append(newPixels, Pixel{b: avg, g: avg, r: avg})
		}
	}
	return
}

func blurFilter(pixels [][]Pixel) (newPixels []Pixel) {

}
