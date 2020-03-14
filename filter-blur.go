package main

type vector struct {
	x, y int
}

func blurFilter(pixels [][]Pixel) (newPixels []Pixel) {
	f := computeAvgPixel(pixels)
	for i, row := range pixels {
		for j := range row {
			newPixels = append(newPixels, f(i, j))
		}
	}
	return
}

func computeAvgPixel(image [][]Pixel) func(int, int) Pixel {
	maxI := len(image) - 1
	maxJ := len(image[0]) - 1
	return func(i, j int) Pixel {
		var vectors []vector
		switch {
		// topLeft corner
		case i == 0 && j == 0:
			vectors = []vector{{0, 0}, {1, 0}, {1, 1}, {0, 1}}
			// topRight corner
		case i == 0 && j == maxJ:
			vectors = []vector{{0, 0}, {0, 1}, {-1, 1}, {-1, 0}}
			// bottom left corner
		case i == maxI && j == 0:
			vectors = []vector{{0, 0}, {0, -1}, {1, -1}, {1, 0}}
			// bottom right corner
		case i == maxI && j == maxJ:
			vectors = []vector{{0, 0}, {0, -1}, {-1, 0}, {-1, -1}}
			// top row
		case i == 0:
			vectors = []vector{{0, 0}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}}
			// bottom row
		case i == maxI:
			vectors = []vector{{0, 0}, {0, -1}, {1, -1}, {1, 0}, {-1, 0}, {-1, -1}}
			// first column
		case j == 0:
			vectors = []vector{{0, 0}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}}
			// second column
		case j == maxJ:
			vectors = []vector{{0, 0}, {0, -1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}}
			// center
		default:
			vectors = []vector{{0, 0}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}}
		}
		var avgB, avgG, avgR byte

		for _, v := range vectors {
			newI, newJ := i+v.y, j+v.x
			avgB += image[newI][newJ].b
			avgG += image[newI][newJ].g
			avgR += image[newI][newJ].r
		}

		b := avgB / byte(len(vectors))
		g := avgG / byte(len(vectors))
		r := avgR / byte(len(vectors))

		return Pixel{b, g, r}
	}
}
