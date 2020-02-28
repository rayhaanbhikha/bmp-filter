package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type BitMapHeader struct {
	headerSize           uint32
	pWidth               uint32
	pHeight              uint32
	colorPanes           uint16
	bitsPerPixel         uint16
	compMethod           uint32
	imageSize            uint32
	horizontalRes        uint32
	verticalRes          uint32
	numOfColors          uint32
	numOfImportantColors uint32
}

func (bmh *BitMapHeader) setBitMapHeader(data []byte) {
	bmh.headerSize = binary.LittleEndian.Uint32(data[0:4])
	bmh.pWidth = binary.LittleEndian.Uint32(data[4:8])
	bmh.pHeight = binary.LittleEndian.Uint32(data[8:12])
	bmh.colorPanes = binary.LittleEndian.Uint16(data[12:14])
	bmh.bitsPerPixel = binary.LittleEndian.Uint16(data[14:16])
	bmh.compMethod = binary.LittleEndian.Uint32(data[16:20])
	bmh.imageSize = binary.LittleEndian.Uint32(data[20:24])
	bmh.horizontalRes = binary.LittleEndian.Uint32(data[24:28])
	bmh.verticalRes = binary.LittleEndian.Uint32(data[28:32])
	bmh.numOfColors = binary.LittleEndian.Uint32(data[32:36])
	bmh.numOfImportantColors = binary.LittleEndian.Uint32(data[36:40])
}

func (bmh *BitMapHeader) String() string {
	return fmt.Sprintf(`
		bmh.headerSize %v
		bmh.pWidth %v
		bmh.pHeight %v
		bmh.colorPanes %v
		bmh.bitsPerPixel %v
		bmh.compMethod %v
		bmh.imageSize %v
		bmh.horizontalRes %v
		bmh.verticalRes %v
		bmh.numOfColors %v
		bmh.numOfImportantColors %v
	`,
		bmh.headerSize,
		bmh.pWidth,
		bmh.pHeight,
		bmh.colorPanes,
		bmh.bitsPerPixel,
		bmh.compMethod,
		bmh.imageSize,
		bmh.horizontalRes,
		bmh.verticalRes,
		bmh.numOfColors,
		bmh.numOfImportantColors,
	)
}

type FileHeader struct {
	header          []byte
	size            uint32
	reservedVal1    uint16
	reservedVal2    uint16
	startingAddress uint32
}

func (fhi *FileHeader) setFileHeader(data []byte) {
	fhi.header = data[0:2]
	fhi.size = binary.LittleEndian.Uint32(data[2:6])
	fhi.reservedVal1 = binary.LittleEndian.Uint16(data[6:8])
	fhi.reservedVal2 = binary.LittleEndian.Uint16(data[8:10])
	fhi.startingAddress = binary.LittleEndian.Uint32(data[10:14])
}

func (fhi *FileHeader) String() string {
	return fmt.Sprintf(`
		header: %s,
		size: %v,
		reservedVal1: %v,
		reservedVal2: %v,
		startingAddress: %v
	`, string(fhi.header), fhi.size, fhi.reservedVal1, fhi.reservedVal2, fhi.startingAddress)
}

func main() {
	fmt.Println("hello world")
	file, err := os.Open("../images/courtyard.bmp")
	newFileName := "../images/filtered_image.bmp"
	checkErr(err)
	defer file.Close()

	// read header file info.
	fileHeader := &FileHeader{}
	fileHeaderData := make([]byte, 14)
	_, err = file.Read(fileHeaderData)
	fileHeader.setFileHeader(fileHeaderData)

	fmt.Println(fileHeader)

	bitMapHeader := &BitMapHeader{}
	bitMapHeaderData := make([]byte, 40)
	_, err = file.ReadAt(bitMapHeaderData, 14)
	bitMapHeader.setBitMapHeader(bitMapHeaderData)

	fmt.Println(bitMapHeader)

	oldImageData := make([]byte, 0)
	var startingOffset int64 = int64(fileHeader.startingAddress)
	bytesPerPixel := int64(bitMapHeader.bitsPerPixel / 8)

	for {
		tempData := make([]byte, bytesPerPixel)
		_, err := file.ReadAt(tempData, startingOffset)
		if err == io.EOF {
			break
		}
		if err != nil {
			os.Exit(1)
		}
		newPixelData := genGreyScale(tempData)
		oldImageData = append(oldImageData, newPixelData...)
		startingOffset += bytesPerPixel
	}

	fileHeaderData = append(fileHeaderData, bitMapHeaderData...)
	fileHeaderData = append(fileHeaderData, oldImageData...)

	// fmt.Printf("%08b", fileHeaderData)

	newFile, err := os.Create(newFileName)
	checkErr(err)

	_, err = newFile.Write(fileHeaderData)
	checkErr(err)

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
