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

	if len(os.Args[1:]) < 3 {
		fmt.Println("incorrect args provided")
		os.Exit(1)
	}

	filter := os.Args[1]
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

	if v := bitMapInfoHeader.bitsPerPixel; v != 24 {
		fmt.Printf("unsupported format! %v bit bmp\n", v)
		os.Exit(1)
	}

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

	var newImageData []byte

	switch filter {
	case "-r":
		newImageData = reflectFilter(int(bitMapInfoHeader.pHeight), int(bitMapInfoHeader.pWidth), image)
	case "-g":
		newImageData = greyScaleFilter(int(bitMapInfoHeader.pHeight), int(bitMapInfoHeader.pWidth), image)
	case "-b":
		newImageData = blurFilter(int(bitMapInfoHeader.pHeight), int(bitMapInfoHeader.pWidth), image)
	default:
		fmt.Printf("filter %s does not exist\n", filter)
		os.Exit(1)
	}

	newImageBytes := append(bitMapHeader.data, bitMapInfoHeader.data...)
	newImageBytes = append(newImageBytes, newImageData...)

	newFile, err := os.Create(newFileName)
	checkErr(err)

	_, err = newFile.Write(newImageBytes)
	checkErr(err)

}
