package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

type BitMapInfoHeader struct {
	data                 []byte
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

func NewBitMapInfoHeader(file *os.File, offset int64) *BitMapInfoHeader {
	data := make([]byte, 40)
	_, err := file.ReadAt(data, offset)
	checkErr(err)
	bmih := &BitMapInfoHeader{data: data}
	bmih.setBitMapInfoHeader()
	return bmih
}

func (bmih *BitMapInfoHeader) setBitMapInfoHeader() {
	bmih.headerSize = binary.LittleEndian.Uint32(bmih.data[0:4])
	bmih.pWidth = binary.LittleEndian.Uint32(bmih.data[4:8])
	bmih.pHeight = binary.LittleEndian.Uint32(bmih.data[8:12])
	bmih.colorPanes = binary.LittleEndian.Uint16(bmih.data[12:14])
	bmih.bitsPerPixel = binary.LittleEndian.Uint16(bmih.data[14:16])
	bmih.compMethod = binary.LittleEndian.Uint32(bmih.data[16:20])
	bmih.imageSize = binary.LittleEndian.Uint32(bmih.data[20:24])
	bmih.horizontalRes = binary.LittleEndian.Uint32(bmih.data[24:28])
	bmih.verticalRes = binary.LittleEndian.Uint32(bmih.data[28:32])
	bmih.numOfColors = binary.LittleEndian.Uint32(bmih.data[32:36])
	bmih.numOfImportantColors = binary.LittleEndian.Uint32(bmih.data[36:40])
}

func (bmih *BitMapInfoHeader) String() string {
	return fmt.Sprintf(`
		bmih.headerSize %v
		bmih.pWidth %v
		bmih.pHeight %v
		bmih.colorPanes %v
		bmih.bitsPerPixel %v
		bmih.compMethod %v
		bmih.imageSize %v
		bmih.horizontalRes %v
		bmih.verticalRes %v
		bmih.numOfColors %v
		bmih.numOfImportantColors %v
	`,
		bmih.headerSize,
		bmih.pWidth,
		bmih.pHeight,
		bmih.colorPanes,
		bmih.bitsPerPixel,
		bmih.compMethod,
		bmih.imageSize,
		bmih.horizontalRes,
		bmih.verticalRes,
		bmih.numOfColors,
		bmih.numOfImportantColors,
	)
}
