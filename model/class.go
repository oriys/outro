package model

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.6
type ConstantInfo struct {
	Tag  uint8
	Info []uint8
}

type AttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	Info               []uint8
}

type FieldInfo struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes      []AttributeInfo
}

type MethodInfo struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes      []AttributeInfo
}

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
