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

type vector struct {
	x, y int
}

func getAverage(image [][]byte) [][]byte {
	newImage := make([][]byte, 0)
	for _, row := range image {
		newRow := make([]byte, 0)
		for j := 0; j < len(row); j += 3 {
			avg := (row[j] + row[j+1] + row[j+2]) / 3
			newRow = append(newRow, avg)
		}
		newImage = append(newImage, newRow)
	}
	return newImage
}

func blurFilter(height, width int, image [][]byte) []byte {

	// compute blur value of each pixel
	blurPixelImage := make([]byte, 0)
	maxI := len(image) - 1
	maxJ := len(image[0])
	f := computeAvgPixel(maxI, maxJ, image)

	for irow := range image {
		for jcol := range image[irow] {
			avgPixelValue := f(irow, jcol)
			blurPixelImage = append(blurPixelImage, avgPixelValue)
		}
	}

	return blurPixelImage
}

func computeAvgPixel(maxI, maxJ int, image [][]byte) func(int, int) byte {
	return func(i, j int) byte {
		var vectors []vector
		switch {
		// topLeft corner
		case i == 0 && j <= 2:
			vectors = []vector{{0, 0}, {3, 0}, {3, 1}, {0, 1}}
			// topRight corner
		case i == 0 && j >= maxJ-3:
			vectors = []vector{{0, 0}, {0, 1}, {-3, 1}, {-3, 0}}
			// bottom left corner
		case i == maxI && j <= 2:
			vectors = []vector{{0, 0}, {0, -1}, {3, -1}, {3, 0}}
			// bottom right corner
		case i == maxI && j >= maxJ-3:
			vectors = []vector{{0, 0}, {0, -1}, {-3, 0}, {-3, -1}}
			// top row
		case i == 0:
			vectors = []vector{{0, 0}, {3, 0}, {3, 1}, {0, 1}, {-3, 1}, {-3, 0}}
			// bottom row
		case i == maxI:
			vectors = []vector{{0, 0}, {0, -1}, {3, -1}, {3, 0}, {-3, 0}, {-3, -1}}
			// first column
		case j <= 2:
			vectors = []vector{{0, 0}, {0, -1}, {3, -1}, {3, 0}, {3, 1}, {0, 1}}
			// second column
		case j >= maxJ-3:
			vectors = []vector{{0, 0}, {0, -1}, {0, 1}, {-3, 1}, {-3, 0}, {-3, -1}}
			// center
		default:
			vectors = []vector{{0, 0}, {0, -1}, {3, -1}, {3, 0}, {3, 1}, {0, 1}, {-3, 1}, {-3, 0}, {-3, -1}}
		}
		var avg byte
		for _, v := range vectors {
			avg += image[i+v.y][j+v.x]
		}
		return avg / byte(len(vectors))
	}
}
