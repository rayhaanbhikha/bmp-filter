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
