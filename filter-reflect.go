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
