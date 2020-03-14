package main

func reflectFilter(height, width int, image [][]byte) []byte {
	newImage := make([]byte, 0)
	for i := 0; i < height; i++ {
		bytesInRow := width * 3
		for j := 0; j < bytesInRow/2; j++ {
			image[i][j], image[i][bytesInRow-1-j] = image[i][bytesInRow-1-j], image[i][j]
		}
		for k := 0; k <= bytesInRow-3; k += 3 {
			newImage = append(newImage, image[i][k+2], image[i][k+1], image[i][k])
		}
	}
	return newImage
}

func greyScaleFilter(height, width int, image [][]byte) []byte {
	newImage := make([]byte, 0)
	for i := 0; i < height; i++ {
		bytesInRow := width * 3
		for k := 0; k <= bytesInRow-3; k += 3 {
			avg := (image[i][k+2] + image[i][k+1] + image[i][k]) / 3
			newImage = append(newImage, avg, avg, avg)
		}
	}
	return newImage
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
