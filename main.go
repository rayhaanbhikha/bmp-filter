package main

import (
	"fmt"
	"io"
	"os"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Pixel struct {
	b, g, r byte
}

func main() {

	if err := os.Mkdir("filtered-images", os.FileMode(0777)); os.IsExist(err) {
		fmt.Println("directory already exists")
	}
	// if len(os.Args[1:]) < 3 {
	// 	fmt.Println("incorrect args provided")
	// 	os.Exit(1)
	// }

	// filter := os.Args[1]
	// fileName := os.Args[2]
	// newFileName := os.Args[3]
	// filter := os.Args[1]
	fileName := "./images/courtyard.bmp"
	newFileName := "./filtered-images/g-courtyard.bmp"

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

	// collect bytes into pixels
	image := make([]Pixel, 0)
	offset := int64(bitMapHeader.startingAddress)

	for {
		bytesRead := make([]byte, 3)
		n, err := file.ReadAt(bytesRead, offset)
		if err == io.EOF {
			fmt.Println("image data read")
			break
		}
		checkErr(err)
		image = append(image, Pixel{bytesRead[0], bytesRead[1], bytesRead[2]})
		offset += int64(n)
	}

	newPixels := greyScaleFilter(int(bitMapInfoHeader.pHeight), int(bitMapInfoHeader.pWidth), image)
	newImageData := pixelsToBytes(newPixels)

	// image := make([][]byte, bitMapInfoHeader.pHeight)

	// // width 600pixels.
	// // each pixel made of 24 bits or 3 bytes.
	// // each row is 600 * 3 = 1800 bytes
	// // we have 400 pixels in height. so 400 * 1800 = 720000 bytes

	// var startingOffset uint32 = bitMapHeader.startingAddress

	// for i := range image {
	// 	var rowBytes uint32 = bitMapInfoHeader.pWidth * 3
	// 	data := make([]byte, rowBytes)
	// 	_, err := file.ReadAt(data, int64(startingOffset))
	// 	checkErr(err)
	// 	image[i] = append(image[i], data...)
	// 	startingOffset += rowBytes
	// }

	// var newImageData []byte

	// switch filter {
	// case "-r":
	// 	newImageData = reflectFilter(int(bitMapInfoHeader.pHeight), int(bitMapInfoHeader.pWidth), image)
	// case "-g":
	// 	newImageData = greyScaleFilter(int(bitMapInfoHeader.pHeight), int(bitMapInfoHeader.pWidth), image)
	// case "-b":
	// 	newImageData = blurFilter(int(bitMapInfoHeader.pHeight), int(bitMapInfoHeader.pWidth), image)
	// default:
	// 	fmt.Printf("filter %s does not exist\n", filter)
	// 	os.Exit(1)
	// }

	newImageBytes := append(bitMapHeader.data, bitMapInfoHeader.data...)
	newImageBytes = append(newImageBytes, newImageData...)

	newFile, err := os.Create(newFileName)
	checkErr(err)

	_, err = newFile.Write(newImageBytes)
	checkErr(err)
}

func pixelsToBytes(pixels []Pixel) (bytes []byte) {
	for _, pixel := range pixels {
		bytes = append(bytes, pixel.b, pixel.g, pixel.r)
	}
	return
}
