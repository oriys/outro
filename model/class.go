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
	Attributes      []CodeAttributeInfo
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

type ExceptionTable struct {
	StartPC   uint16
	EndPC     uint16
	HandlerPC uint16
	CatchType uint16
}
type CodeAttributeInfo struct {
	AttributeNameIndex   uint16
	AttributeLength      uint32
	MaxStack             uint16
	MaxLocals            uint16
	CodeLength           uint32
	Code                 []byte
	ExceptionTableLength uint16
	ExceptionTable       []ExceptionTable
	AttributesCount      uint16
	Attributes           []AttributeInfo
}

func (a *CodeAttributeInfo) Print() {
	println("Method MaxStack:", a.MaxStack)
	println("Method MaxLocals:", a.MaxLocals)
	println("Method CodeLength:", a.CodeLength)
	for _, b := range a.Code {
		println("Method Code ", InstructDisplayNameMap[Instruct(b)])
	}
	println("Method ExceptionTableLength:", a.ExceptionTableLength)
	println("Method ExceptionTable:", a.ExceptionTable)
	println("Method AttributesCount:", a.AttributesCount)
	println("Method Attributes:", a.Attributes)
}
