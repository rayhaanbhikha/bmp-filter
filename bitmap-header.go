package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

type BitMapHeader struct {
	data            []byte
	header          []byte
	size            uint32
	reservedVal1    uint16
	reservedVal2    uint16
	startingAddress uint32
}

func NewBitMapHeader(file *os.File) *BitMapHeader {
	data := make([]byte, 14)
	_, err := file.Read(data)
	checkErr(err)
	bmh := &BitMapHeader{data: data}
	bmh.setBitMapHeader()
	return bmh
}

func (bmh *BitMapHeader) setBitMapHeader() {
	bmh.header = bmh.data[0:2]
	bmh.size = binary.LittleEndian.Uint32(bmh.data[2:6])
	bmh.reservedVal1 = binary.LittleEndian.Uint16(bmh.data[6:8])
	bmh.reservedVal2 = binary.LittleEndian.Uint16(bmh.data[8:10])
	bmh.startingAddress = binary.LittleEndian.Uint32(bmh.data[10:14])
}

func (bmh *BitMapHeader) String() string {
	return fmt.Sprintf(`
		header: %s,
		size: %v,
		reservedVal1: %v,
		reservedVal2: %v,
		startingAddress: %v
	`, string(bmh.header), bmh.size, bmh.reservedVal1, bmh.reservedVal2, bmh.startingAddress)
}
