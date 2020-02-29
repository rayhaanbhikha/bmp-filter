package main

import (
	"fmt"
	"os"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// revNums := make([]int, 0)
	// n := len(numbers)
	// for i := 0; i <= n-3; i += 3 {
	// 	// fmt.Println(numbers[i : i+3])
	// 	newNums := []int{numbers[i+2], numbers[i+1], numbers[i]}
	// 	revNums = append(revNums, newNums...)
	// }
	// // numbers = append([]int{9, 8, 7}, numbers...)
	// fmt.Println(numbers)
	// fmt.Println(revNums)
	if len(os.Args[1:]) < 3 {
		fmt.Println("incorrect args provided")
		os.Exit(1)
	}

	// filter := os.Args[1]
	fileName := os.Args[2]
	newFileName := os.Args[3]

	file, err := os.Open(fileName)
	checkErr(err)
	defer file.Close()

	// read header file info.
	bitMapHeader := NewBitMapHeader(file)
	fmt.Println(bitMapHeader)

	// read bmp info header
	bitMapInfoHeader := NewBitMapInfoHeader(file, 14)
	fmt.Println(bitMapInfoHeader)

	image := make([][]byte, bitMapInfoHeader.pHeight)

	// width 600pixels.
	// each pixel made of 24 bits or 3 bytes.
	// each row is 600 * 3 = 1800 bytes
	// we have 400 pixels in height. so 400 * 1800 = 720000 bytes

	var startingOffset uint32 = bitMapHeader.startingAddress

	for i := range image {
		var rowBytes uint32 = bitMapInfoHeader.pWidth * 3
		data := make([]byte, rowBytes)
		_, err := file.ReadAt(data, int64(startingOffset))
		checkErr(err)
		image[i] = append(image[i], data...)
		startingOffset += rowBytes
	}

	// grey scale
	newImageData := reflectFilter(int(bitMapInfoHeader.pHeight), int(bitMapInfoHeader.pWidth), image)

	newImageBytes := append(bitMapHeader.data, bitMapInfoHeader.data...)
	newImageBytes = append(newImageBytes, newImageData...)

	// fmt.Printf("%08b", bitMapHeaderData)

	newFile, err := os.Create(newFileName)
	checkErr(err)

	_, err = newFile.Write(newImageBytes)
	checkErr(err)

}

func reflectFilter(height, width int, image [][]byte) []byte {
	newImage := make([]byte, 0)
	for i := 0; i < height; i++ {
		bytesInRow := width * 3
		for j := 0; j < bytesInRow/2; j++ {
			image[i][j], image[i][bytesInRow-1-j] = image[i][bytesInRow-1-j], image[i][j]
		}
		for k := 0; k <= bytesInRow-3; k += 3 {
			data := []byte{image[i][k+2], image[i][k+1], image[i][k]}
			newImage = append(newImage, data...)
		}
	}
	return newImage
}

func genGreyScale(pixel []byte) []byte {
	n := byte(len(pixel))

	var sum byte = 0
	for i := range pixel {
		sum += pixel[i]
	}
	avg := sum / n

	newPixels := make([]byte, n)
	for i := range newPixels {
		newPixels[i] = avg
	}
	return newPixels
}
