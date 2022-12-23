package parser

import (
	"encoding/binary"
	"outro/constant"
	"outro/model"
)

type ClassFileParser struct {
	reader *ByteReader
}

func (p *ClassFileParser) parseMagic() uint32 {
	return p.reader.ReadUint32()
}

func (p *ClassFileParser) parseMinorVersion() uint16 {
	return p.reader.ReadUint16()
}

func (p *ClassFileParser) parseMajorVersion() uint16 {
	return p.reader.ReadUint16()
}

func (p *ClassFileParser) parseConstantPoolCount() uint16 {
	return p.reader.ReadUint16()
}

func (p *ClassFileParser) parseConstantPool(count uint16) []model.ConstantInfo {
	constantPool := make([]model.ConstantInfo, count)
	for i := 1; i < int(count); i++ {
		constantPool[i] = p.parseConstantInfo()
		switch constantPool[i].Tag {
		case constant.ConstantLong, constant.ConstantDouble:
			i++
		}
	}
	return constantPool
}

func (p *ClassFileParser) parseAccessFlags() uint16 {
	return p.reader.ReadUint16()
}

func (p *ClassFileParser) parseThisClass() uint16 {
	return p.reader.ReadUint16()
}

func (p *ClassFileParser) parseSuperClass() uint16 {
	return p.reader.ReadUint16()
}

func (p *ClassFileParser) parseInterfacesCount() uint16 {
	return p.reader.ReadUint16()
}

func (p *ClassFileParser) parseFieldsCount() uint16 {
	return p.reader.ReadUint16()
}

func (p *ClassFileParser) parseFields(count uint16) []model.FieldInfo {
	fields := make([]model.FieldInfo, count)
	for i := range fields {
		fields[i] = p.parseFieldInfo()
	}
	return fields
}

func (p *ClassFileParser) parseMethodsCount() uint16 {
	return p.reader.ReadUint16()
}

func (p *ClassFileParser) parseMethods(count uint16) []model.MethodInfo {
	methods := make([]model.MethodInfo, count)
	for i := range methods {
		methods[i] = p.parseMethodInfo()
	}
	return methods
}

func (p *ClassFileParser) parseAttributesCount() uint16 {
	return p.reader.ReadUint16()
}

func (p *ClassFileParser) parseAttributes(count uint16) []model.AttributeInfo {
	attributes := make([]model.AttributeInfo, count)
	for i := range attributes {
		attributes[i] = p.parseAttributeInfo()
	}
	return attributes
}

func (p *ClassFileParser) parseInterfaces(count uint16) []uint16 {
	interfaces := make([]uint16, count)
	for i := range interfaces {
		interfaces[i] = p.reader.ReadUint16()
	}
	return interfaces
}

func (p *ClassFileParser) parseConstantInfo() model.ConstantInfo {
	tag := p.reader.ReadUint8()
	switch tag {
	case constant.ConstantClass:
		return p.parseConstantClassInfo()
	case constant.ConstantFieldRef:
		return p.parseConstantFieldrefInfo()
	case constant.ConstantMethodRef:
		return p.parseConstantMethodrefInfo()
	case constant.ConstantInterfaceMethodRef:
		return p.parseConstantInterfaceMethodrefInfo()
	case constant.ConstantString:
		return p.parseConstantStringInfo()
	case constant.ConstantInteger:
		return p.parseConstantIntegerInfo()
	case constant.ConstantFloat:
		return p.parseConstantFloatInfo()
	case constant.ConstantLong:
		return p.parseConstantLongInfo()
	case constant.ConstantDouble:
		return p.parseConstantDoubleInfo()
	case constant.ConstantNameAndType:
		return p.parseConstantNameAndTypeInfo()
	case constant.ConstantUtf8:
		return p.parseConstantUtf8Info()
	case constant.ConstantMethodHandle:
		return p.parseConstantMethodHandleInfo()
	case constant.ConstantMethodType:
		return p.parseConstantMethodTypeInfo()
	case constant.ConstantInvokeDynamic:
		return p.parseConstantInvokeDynamicInfo()
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}

func (p *ClassFileParser) parseConstantClassInfo() model.ConstantInfo {
	readUint16 := p.reader.ReadUint16()
	info := make([]uint8, 2)
	binary.BigEndian.PutUint16(info, readUint16)
	return model.ConstantInfo{Tag: constant.ConstantClass, Info: info}
}

func (p *ClassFileParser) parseConstantFieldrefInfo() model.ConstantInfo {
	readUint16 := p.reader.ReadUint16()
	info := make([]uint8, 4)
	binary.BigEndian.PutUint16(info, readUint16)
	binary.BigEndian.PutUint16(info[2:], p.reader.ReadUint16())
	return model.ConstantInfo{Tag: constant.ConstantFieldRef, Info: info}
}

func (p *ClassFileParser) parseConstantMethodrefInfo() model.ConstantInfo {
	readUint16 := p.reader.ReadUint16()
	info := make([]uint8, 4)
	binary.BigEndian.PutUint16(info, readUint16)
	binary.BigEndian.PutUint16(info[2:], p.reader.ReadUint16())
	return model.ConstantInfo{Tag: constant.ConstantMethodRef, Info: info}
}

func (p *ClassFileParser) parseConstantInterfaceMethodrefInfo() model.ConstantInfo {
	readUint16 := p.reader.ReadUint16()
	info := make([]uint8, 4)
	binary.BigEndian.PutUint16(info, readUint16)
	binary.BigEndian.PutUint16(info[2:], p.reader.ReadUint16())
	return model.ConstantInfo{Tag: constant.ConstantInterfaceMethodRef, Info: info}
}

func (p *ClassFileParser) parseConstantStringInfo() model.ConstantInfo {
	readUint16 := p.reader.ReadUint16()
	info := make([]uint8, 2)
	binary.BigEndian.PutUint16(info, readUint16)
	return model.ConstantInfo{Tag: constant.ConstantString, Info: info}
}

func (p *ClassFileParser) parseConstantIntegerInfo() model.ConstantInfo {
	readUint32 := p.reader.ReadUint32()
	info := make([]uint8, 4)
	binary.BigEndian.PutUint32(info, readUint32)
	return model.ConstantInfo{Tag: constant.ConstantInteger, Info: info}
}

func (p *ClassFileParser) parseConstantFloatInfo() model.ConstantInfo {
	readUint32 := p.reader.ReadUint32()
	info := make([]uint8, 4)
	binary.BigEndian.PutUint32(info, readUint32)
	return model.ConstantInfo{Tag: constant.ConstantFloat, Info: info}
}

func (p *ClassFileParser) parseConstantLongInfo() model.ConstantInfo {
	readUint64 := p.reader.ReadUint64()
	info := make([]uint8, 8)
	binary.BigEndian.PutUint64(info, readUint64)
	return model.ConstantInfo{Tag: constant.ConstantLong, Info: info}
}

func (p *ClassFileParser) parseConstantDoubleInfo() model.ConstantInfo {
	readUint64 := p.reader.ReadUint64()
	info := make([]uint8, 8)
	binary.BigEndian.PutUint64(info, readUint64)
	return model.ConstantInfo{Tag: constant.ConstantDouble, Info: info}
}

func (p *ClassFileParser) parseConstantNameAndTypeInfo() model.ConstantInfo {
	readUint16 := p.reader.ReadUint16()
	info := make([]uint8, 4)
	binary.BigEndian.PutUint16(info, readUint16)
	binary.BigEndian.PutUint16(info[2:], p.reader.ReadUint16())
	return model.ConstantInfo{Tag: constant.ConstantNameAndType, Info: info}
}

func (p *ClassFileParser) parseConstantUtf8Info() model.ConstantInfo {
	length := p.reader.ReadUint16()
	bytes := p.reader.ReadBytes(uint32(length))
	return model.ConstantInfo{Tag: constant.ConstantUtf8, Info: bytes}
}

func (p *ClassFileParser) parseConstantMethodHandleInfo() model.ConstantInfo {
	info := make([]uint8, 3)
	info[0] = p.reader.ReadUint8()
	binary.BigEndian.PutUint16(info[1:], p.reader.ReadUint16())
	return model.ConstantInfo{Tag: constant.ConstantMethodHandle, Info: info}
}

func (p *ClassFileParser) parseConstantMethodTypeInfo() model.ConstantInfo {
	readUint16 := p.reader.ReadUint16()
	info := make([]uint8, 2)
	binary.BigEndian.PutUint16(info, readUint16)
	return model.ConstantInfo{Tag: constant.ConstantMethodType, Info: info}
}

func (p *ClassFileParser) parseConstantInvokeDynamicInfo() model.ConstantInfo {
	readUint16 := p.reader.ReadUint16()
	info := make([]uint8, 4)
	binary.BigEndian.PutUint16(info, readUint16)
	binary.BigEndian.PutUint16(info[2:], p.reader.ReadUint16())
	return model.ConstantInfo{Tag: constant.ConstantInvokeDynamic, Info: info}
}

func (p *ClassFileParser) parseFieldInfo() model.FieldInfo {
	accessFlags := p.reader.ReadUint16()
	nameIndex := p.reader.ReadUint16()
	descriptorIndex := p.reader.ReadUint16()
	attributesCount := p.reader.ReadUint16()
	attributes := make([]model.AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = p.parseAttributeInfo()
	}
	return model.FieldInfo{
		AccessFlags:     accessFlags,
		NameIndex:       nameIndex,
		DescriptorIndex: descriptorIndex,
		AttributesCount: attributesCount,
		Attributes:      attributes,
	}
}

func (p *ClassFileParser) parseAttributeInfo() model.AttributeInfo {
	attributeNameIndex := p.reader.ReadUint16()
	attributeLength := p.reader.ReadUint32()
	attributeInfo := p.reader.ReadBytes(attributeLength)
	return model.AttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		Info:               attributeInfo,
	}
}

func (p *ClassFileParser) parseMethodInfo() model.MethodInfo {
	accessFlags := p.reader.ReadUint16()
	nameIndex := p.reader.ReadUint16()
	descriptorIndex := p.reader.ReadUint16()
	attributesCount := p.reader.ReadUint16()
	attributes := make([]model.AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = p.parseAttributeInfo()
	}
	return model.MethodInfo{
		AccessFlags:     accessFlags,
		NameIndex:       nameIndex,
		DescriptorIndex: descriptorIndex,
		AttributesCount: attributesCount,
		Attributes:      attributes,
	}
}

func (p *ClassFileParser) Parse() *model.Class {
	class := model.Class{}
	class.Magic = p.parseMagic()
	class.MinorVersion = p.parseMinorVersion()
	class.MajorVersion = p.parseMajorVersion()
	class.ConstantPoolCount = p.parseConstantPoolCount()
	class.ConstantPool = p.parseConstantPool(class.ConstantPoolCount)
	class.AccessFlags = p.parseAccessFlags()
	class.ThisClass = p.parseThisClass()
	class.SuperClass = p.parseSuperClass()
	class.InterfacesCount = p.parseInterfacesCount()
	class.Interfaces = p.parseInterfaces(class.InterfacesCount)
	class.FieldsCount = p.parseFieldsCount()
	class.Fields = p.parseFields(class.FieldsCount)
	class.MethodsCount = p.parseMethodsCount()
	class.Methods = p.parseMethods(class.MethodsCount)
	class.AttributesCount = p.parseAttributesCount()
	class.Attributes = p.parseAttributes(class.AttributesCount)
	return &class

}

func (p *ClassFileParser) parseCodeAttributeInfo() model.CodeAttributeInfo {
	attributeNameIndex := p.reader.ReadUint16()
	attributeLength := p.reader.ReadUint32()
	maxStack := p.reader.ReadUint16()
	maxLocals := p.reader.ReadUint16()
	codeLength := p.reader.ReadUint32()
	code := p.reader.ReadBytes(codeLength)
	exceptionTableLength := p.reader.ReadUint16()
	exceptionTable := make([]model.ExceptionTable, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = p.parseExceptionTable()
	}
	attributesCount := p.reader.ReadUint16()
	attributes := make([]model.AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = p.parseAttributeInfo()
	}
	return model.CodeAttributeInfo{
		AttributeNameIndex:   attributeNameIndex,
		AttributeLength:      attributeLength,
		MaxStack:             maxStack,
		MaxLocals:            maxLocals,
		CodeLength:           codeLength,
		Code:                 code,
		ExceptionTableLength: exceptionTableLength,
		ExceptionTable:       exceptionTable,
		AttributesCount:      attributesCount,
		Attributes:           attributes,
	}

}

func (p *ClassFileParser) parseExceptionTable() model.ExceptionTable {
	return model.ExceptionTable{
		StartPC:   p.reader.ReadUint16(),
		EndPC:     p.reader.ReadUint16(),
		HandlerPC: p.reader.ReadUint16(),
		CatchType: p.reader.ReadUint16(),
	}
}

func (p *ClassFileParser) parseLocalVariableAttributeInfo() model.LocalVariableTableAttributeInfo {
	attributeNameIndex := p.reader.ReadUint16()
	attributeLength := p.reader.ReadUint32()
	localVariableTableLength := p.reader.ReadUint16()
	localVariableTable := make([]model.LocalVariableTable, localVariableTableLength)
	for i := range localVariableTable {
		localVariableTable[i] = p.parseLocalVariableTable()
	}
	return model.LocalVariableTableAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		LocalVariableTable: localVariableTable,
	}
}

func (p *ClassFileParser) parseLocalVariableTable() model.LocalVariableTable {
	startPC := p.reader.ReadUint16()
	length := p.reader.ReadUint16()
	nameIndex := p.reader.ReadUint16()
	descriptorIndex := p.reader.ReadUint16()
	index := p.reader.ReadUint16()
	return model.LocalVariableTable{
		StartPC:         startPC,
		Length:          length,
		NameIndex:       nameIndex,
		DescriptorIndex: descriptorIndex,
		Index:           index,
	}
}

func NewClassFileParser(reader *ByteReader) *ClassFileParser {
	return &ClassFileParser{reader: reader}
}
