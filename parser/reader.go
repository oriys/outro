package parser

import (
	"encoding/binary"
)

/**
type Class struct {
	Magic             uint32
	MinorVersion      uint16
	MajorVersion      uint16
	ConstantPoolCount uint16
	ConstantPool      []ConstantInfo
	AccessFlags       uint16
	ThisClass         uint16
	SuperClass        uint16
	InterfacesCount   uint16
	Interfaces        []uint16
	FieldsCount       uint16
	Fields            []FieldInfo
	MethodsCount      uint16
	Methods           []MethodInfo
	AttributesCount   uint16
	Attributes        []AttributeInfo
}

*/

type ByteReader struct {
	data     []byte
	position int
}

func NewByteReader(data []byte) *ByteReader {
	return &ByteReader{data: data, position: 0}
}

func (r *ByteReader) ReadUx(n int) (val []byte) {
	val = r.data[r.position : r.position+n]
	r.position += n
	return val
}

func (r *ByteReader) ReadUint8() uint8 {
	return r.ReadUx(1)[0]
}

func (r *ByteReader) ReadUint16() uint16 {
	return binary.BigEndian.Uint16(r.ReadUx(2))
}

func (r *ByteReader) ReadUint32() uint32 {
	return binary.BigEndian.Uint32(r.ReadUx(4))
}

func (r *ByteReader) ReadUint64() uint64 {
	return binary.BigEndian.Uint64(r.ReadUx(8))
}

func (r *ByteReader) ReadBytes(length uint32) []byte {
	return r.ReadUx(int(length))
}
