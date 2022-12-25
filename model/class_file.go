package model

import "encoding/binary"

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

type ClassFile struct {
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

type LocalVariableTableAttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	LocalVariableTable []LocalVariableTable
}

type LocalVariableTable struct {
	StartPC         uint16
	Length          uint16
	NameIndex       uint16
	DescriptorIndex uint16
	Index           uint16
}

func (c *ClassFile) GetMethod(name string, descriptor string) (*MethodInfo, error) {
	for _, method := range c.Methods {
		if c.readConstantPoolUtf8ValueByIndex(method.NameIndex) == name && c.readConstantPoolUtf8ValueByIndex(method.DescriptorIndex) == descriptor {
			return &method, nil
		}
	}
	return nil, nil
}

func (c *ClassFile) GetField(name string, descriptor string) (*FieldInfo, error) {
	for _, field := range c.Fields {
		if c.readConstantPoolUtf8ValueByIndex(field.NameIndex) == name && c.readConstantPoolUtf8ValueByIndex(field.DescriptorIndex) == descriptor {
			return &field, nil
		}
	}
	return nil, nil
}

func (c *ClassFile) GetCodeAttribute(method *MethodInfo) (*CodeAttributeInfo, error) {
	for _, attr := range method.Attributes {
		if c.readConstantPoolUtf8ValueByIndex(attr.AttributeNameIndex) == "Code" {
			return attr.ToCodeAttributeInfo()
		}
	}
	return nil, nil
}

func (c *ClassFile) readConstantPoolUtf8ValueByIndex(index uint16) string {
	return string(c.ConstantPool[index-1].Info)
}

func (attr AttributeInfo) ToCodeAttributeInfo() (*CodeAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	maxStack := binary.BigEndian.Uint16(attr.Info[6:8])
	maxLocals := binary.BigEndian.Uint16(attr.Info[8:10])
	codeLength := binary.BigEndian.Uint32(attr.Info[10:14])
	code := attr.Info[14 : 14+codeLength]
	exceptionTableLength := binary.BigEndian.Uint16(attr.Info[14+codeLength : 16+codeLength])
	exceptionTable := make([]ExceptionTable, exceptionTableLength)
	for i := range exceptionTable {
		baseStart := codeLength + uint32(i*8)
		startPCLeft := 16 + baseStart
		startPCRight := 18 + baseStart
		handlerPCLeft := 20 + baseStart
		handlerPCRight := 22 + baseStart
		catchTypeRight := 24 + baseStart
		exceptionTable[i] = ExceptionTable{
			StartPC:   binary.BigEndian.Uint16(attr.Info[startPCLeft:startPCRight]),
			EndPC:     binary.BigEndian.Uint16(attr.Info[startPCRight:handlerPCLeft]),
			HandlerPC: binary.BigEndian.Uint16(attr.Info[handlerPCLeft:handlerPCRight]),
			CatchType: binary.BigEndian.Uint16(attr.Info[handlerPCRight:catchTypeRight]),
		}
	}
	attributesCount := binary.BigEndian.Uint16(attr.Info[16+codeLength+uint32(exceptionTableLength*8) : 18+codeLength+uint32(exceptionTableLength*8)])
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		baseStart := codeLength + uint32(exceptionTableLength*8) + uint32(i*6)
		attributeNameIndexLeft := 18 + baseStart
		attributeNameIndexRight := 20 + baseStart
		attributeLengthLeft := 20 + baseStart
		attributeLengthRight := 24 + baseStart
		attributes[i] = AttributeInfo{
			AttributeNameIndex: binary.BigEndian.Uint16(attr.Info[attributeNameIndexLeft:attributeNameIndexRight]),
			AttributeLength:    binary.BigEndian.Uint32(attr.Info[attributeLengthLeft:attributeLengthRight]),
			Info:               attr.Info[attributeLengthRight : attributeLengthRight+attributes[i].AttributeLength],
		}
	}
	return &CodeAttributeInfo{
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
	}, nil
}

func (attr *AttributeInfo) ToLocalVariableTableAttributeInfo() (*LocalVariableTableAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	localVariableTableLength := binary.BigEndian.Uint16(attr.Info[6:8])
	localVariableTable := make([]LocalVariableTable, localVariableTableLength)
	for i := range localVariableTable {
		baseStart := uint32(i * 10)
		startPCLeft := 8 + baseStart
		startPCRight := 10 + baseStart
		lengthLeft := 10 + baseStart
		lengthRight := 12 + baseStart
		nameIndexLeft := 12 + baseStart
		nameIndexRight := 14 + baseStart
		descriptorIndexLeft := 14 + baseStart
		descriptorIndexRight := 16 + baseStart
		indexLeft := 16 + baseStart
		indexRight := 18 + baseStart
		localVariableTable[i] = LocalVariableTable{
			StartPC:         binary.BigEndian.Uint16(attr.Info[startPCLeft:startPCRight]),
			Length:          binary.BigEndian.Uint16(attr.Info[lengthLeft:lengthRight]),
			NameIndex:       binary.BigEndian.Uint16(attr.Info[nameIndexLeft:nameIndexRight]),
			DescriptorIndex: binary.BigEndian.Uint16(attr.Info[descriptorIndexLeft:descriptorIndexRight]),
			Index:           binary.BigEndian.Uint16(attr.Info[indexLeft:indexRight]),
		}
	}
	return &LocalVariableTableAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		LocalVariableTable: localVariableTable,
	}, nil
}

type LineNumberTable struct {
	StartPC    uint16
	LineNumber uint16
}

type LineNumberTableAttributeInfo struct {
	AttributeNameIndex   uint16
	AttributeLength      uint32
	LineNumberTable      []LineNumberTable
	LineNumberTableCount uint16
}

func (attr *AttributeInfo) ToLineNumberTableAttributeInfo() (*LineNumberTableAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	lineNumberTableLength := binary.BigEndian.Uint16(attr.Info[6:8])
	lineNumberTable := make([]LineNumberTable, lineNumberTableLength)
	for i := range lineNumberTable {
		baseStart := uint32(i * 4)
		startPCLeft := 8 + baseStart
		startPCRight := 10 + baseStart
		lineNumberLeft := 10 + baseStart
		lineNumberRight := 12 + baseStart
		lineNumberTable[i] = LineNumberTable{
			StartPC:    binary.BigEndian.Uint16(attr.Info[startPCLeft:startPCRight]),
			LineNumber: binary.BigEndian.Uint16(attr.Info[lineNumberLeft:lineNumberRight]),
		}
	}
	return &LineNumberTableAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		LineNumberTable:    lineNumberTable,
	}, nil
}

type SourceFileAttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	SourceFileIndex    uint16
}

func (attr *AttributeInfo) ToSourceFileAttributeInfo() (*SourceFileAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	sourceFileIndex := binary.BigEndian.Uint16(attr.Info[6:8])
	return &SourceFileAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		SourceFileIndex:    sourceFileIndex,
	}, nil

}

type ConstantValueAttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	ConstantValueIndex uint16
}

func (attr *AttributeInfo) ToConstantValueAttributeInfo() (*ConstantValueAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	constantValueIndex := binary.BigEndian.Uint16(attr.Info[6:8])
	return &ConstantValueAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		ConstantValueIndex: constantValueIndex,
	}, nil
}

type ExceptionsAttributeInfo struct {
	AttributeNameIndex  uint16
	AttributeLength     uint32
	NumberOfExceptions  uint16
	ExceptionIndexTable []uint16
}

func (attr *AttributeInfo) ToExceptionsAttributeInfo() (*ExceptionsAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	numberOfExceptions := binary.BigEndian.Uint16(attr.Info[6:8])
	exceptionIndexTable := make([]uint16, numberOfExceptions)
	for i := range exceptionIndexTable {
		left := 8 + uint32(i*2)
		right := 10 + uint32(i*2)
		exceptionIndexTable[i] = binary.BigEndian.Uint16(attr.Info[left:right])
	}
	return &ExceptionsAttributeInfo{
		AttributeNameIndex:  attributeNameIndex,
		AttributeLength:     attributeLength,
		NumberOfExceptions:  numberOfExceptions,
		ExceptionIndexTable: exceptionIndexTable,
	}, nil
}

type InnerClassesAttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	NumberOfClasses    uint16
	Classes            []InnerClassInfo
}

type InnerClassInfo struct {
	InnerClassInfoIndex       uint16
	OuterClassInfoIndex       uint16
	InnerNameIndex            uint16
	InnerClassAccessFlagsInfo uint16
}

func (attr *AttributeInfo) ToInnerClassesAttributeInfo() (*InnerClassesAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	numberOfClasses := binary.BigEndian.Uint16(attr.Info[6:8])
	classes := make([]InnerClassInfo, numberOfClasses)
	for i := range classes {
		baseStart := uint32(i * 8)
		innerClassInfoIndexLeft := 8 + baseStart
		innerClassInfoIndexRight := 10 + baseStart
		outerClassInfoIndexLeft := 10 + baseStart
		outerClassInfoIndexRight := 12 + baseStart
		innerNameIndexLeft := 12 + baseStart
		innerNameIndexRight := 14 + baseStart
		innerClassAccessFlagsInfoLeft := 14 + baseStart
		innerClassAccessFlagsInfoRight := 16 + baseStart
		classes[i] = InnerClassInfo{
			InnerClassInfoIndex:       binary.BigEndian.Uint16(attr.Info[innerClassInfoIndexLeft:innerClassInfoIndexRight]),
			OuterClassInfoIndex:       binary.BigEndian.Uint16(attr.Info[outerClassInfoIndexLeft:outerClassInfoIndexRight]),
			InnerNameIndex:            binary.BigEndian.Uint16(attr.Info[innerNameIndexLeft:innerNameIndexRight]),
			InnerClassAccessFlagsInfo: binary.BigEndian.Uint16(attr.Info[innerClassAccessFlagsInfoLeft:innerClassAccessFlagsInfoRight]),
		}
	}
	return &InnerClassesAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		NumberOfClasses:    numberOfClasses,
		Classes:            classes,
	}, nil
}

type EnclosingMethodAttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	ClassIndex         uint16
	MethodIndex        uint16
}

func (attr *AttributeInfo) ToEnclosingMethodAttributeInfo() (*EnclosingMethodAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	classIndex := binary.BigEndian.Uint16(attr.Info[6:8])
	methodIndex := binary.BigEndian.Uint16(attr.Info[8:10])
	return &EnclosingMethodAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		ClassIndex:         classIndex,
		MethodIndex:        methodIndex,
	}, nil
}

type SyntheticAttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
}

func (attr *AttributeInfo) ToSyntheticAttributeInfo() (*SyntheticAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	return &SyntheticAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
	}, nil
}

type SignatureAttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	SignatureIndex     uint16
}

func (attr *AttributeInfo) ToSignatureAttributeInfo() (*SignatureAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	signatureIndex := binary.BigEndian.Uint16(attr.Info[6:8])
	return &SignatureAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		SignatureIndex:     signatureIndex,
	}, nil
}

type SourceDebugExtensionAttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	DebugExtension     []byte
}

func (attr *AttributeInfo) ToSourceDebugExtensionAttributeInfo() (*SourceDebugExtensionAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	debugExtension := attr.Info[6:]
	return &SourceDebugExtensionAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		DebugExtension:     debugExtension,
	}, nil
}

type LocalVariableTypeInfo struct {
	StartPc        uint16
	Length         uint16
	NameIndex      uint16
	SignatureIndex uint16
	Index          uint16
}

type LocalVariableTypeTableAttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	LocalVariableTable []LocalVariableTypeInfo
}

func (attr *AttributeInfo) ToLocalVariableTypeTableAttributeInfo() (*LocalVariableTypeTableAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	localVariableTableLength := binary.BigEndian.Uint16(attr.Info[6:8])
	localVariableTable := make([]LocalVariableTypeInfo, localVariableTableLength)
	for i := range localVariableTable {
		baseStart := uint32(i * 10)
		startPcLeft := 8 + baseStart
		startPcRight := 10 + baseStart
		lengthLeft := 10 + baseStart
		lengthRight := 12 + baseStart
		nameIndexLeft := 12 + baseStart
		nameIndexRight := 14 + baseStart
		signatureIndexLeft := 14 + baseStart
		signatureIndexRight := 16 + baseStart
		indexLeft := 16 + baseStart
		indexRight := 18 + baseStart
		localVariableTable[i] = LocalVariableTypeInfo{
			StartPc:        binary.BigEndian.Uint16(attr.Info[startPcLeft:startPcRight]),
			Length:         binary.BigEndian.Uint16(attr.Info[lengthLeft:lengthRight]),
			NameIndex:      binary.BigEndian.Uint16(attr.Info[nameIndexLeft:nameIndexRight]),
			SignatureIndex: binary.BigEndian.Uint16(attr.Info[signatureIndexLeft:signatureIndexRight]),
			Index:          binary.BigEndian.Uint16(attr.Info[indexLeft:indexRight]),
		}
	}
	return &LocalVariableTypeTableAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		LocalVariableTable: localVariableTable,
	}, nil
}

type DeprecatedAttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
}

func (attr *AttributeInfo) ToDeprecatedAttributeInfo() (*DeprecatedAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	return &DeprecatedAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
	}, nil
}

type AnnotationInfo struct {
	TypeIndex uint16
}

type RuntimeVisibleAnnotationsAttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	Annotations        []AnnotationInfo
}

func (attr *AttributeInfo) ToRuntimeVisibleAnnotationsAttributeInfo() (*RuntimeVisibleAnnotationsAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	annotationsLength := binary.BigEndian.Uint16(attr.Info[6:8])
	annotations := make([]AnnotationInfo, annotationsLength)
	for i := range annotations {
		baseStart := uint32(i * 2)
		annotationLeft := 8 + baseStart
		annotationRight := 10 + baseStart
		annotations[i] = AnnotationInfo{
			TypeIndex: binary.BigEndian.Uint16(attr.Info[annotationLeft:annotationRight]),
		}
	}
	return &RuntimeVisibleAnnotationsAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		Annotations:        annotations,
	}, nil
}

type RuntimeInvisibleAnnotationsAttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	Annotations        []AnnotationInfo
}

func (attr *AttributeInfo) ToRuntimeInvisibleAnnotationsAttributeInfo() (*RuntimeInvisibleAnnotationsAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	annotationsLength := binary.BigEndian.Uint16(attr.Info[6:8])
	annotations := make([]AnnotationInfo, annotationsLength)
	for i := range annotations {
		baseStart := uint32(i * 2)
		annotationLeft := 8 + baseStart
		annotationRight := 10 + baseStart
		annotations[i] = AnnotationInfo{
			TypeIndex: binary.BigEndian.Uint16(attr.Info[annotationLeft:annotationRight]),
		}
	}
	return &RuntimeInvisibleAnnotationsAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		Annotations:        annotations,
	}, nil
}

type ParameterAnnotationInfo struct {
	Annotations []AnnotationInfo
}

type RuntimeVisibleParameterAnnotationsAttributeInfo struct {
	AttributeNameIndex   uint16
	AttributeLength      uint32
	ParameterAnnotations []ParameterAnnotationInfo
}

func (attr *AttributeInfo) ToRuntimeVisibleParameterAnnotationsAttributeInfo() (*RuntimeVisibleParameterAnnotationsAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	parameterAnnotationsLength := attr.Info[6:7][0]
	parameterAnnotations := make([]ParameterAnnotationInfo, parameterAnnotationsLength)
	for i := range parameterAnnotations {
		baseStart := uint32(i * 2)
		parameterAnnotationLeft := 7 + baseStart
		parameterAnnotationRight := 9 + baseStart
		parameterAnnotations[i] = ParameterAnnotationInfo{
			Annotations: []AnnotationInfo{
				{
					TypeIndex: binary.BigEndian.Uint16(attr.Info[parameterAnnotationLeft:parameterAnnotationRight]),
				},
			},
		}
	}
	return &RuntimeVisibleParameterAnnotationsAttributeInfo{
		AttributeNameIndex:   attributeNameIndex,
		AttributeLength:      attributeLength,
		ParameterAnnotations: parameterAnnotations,
	}, nil
}

type RuntimeInvisibleParameterAnnotationsAttributeInfo struct {
	AttributeNameIndex   uint16
	AttributeLength      uint32
	ParameterAnnotations []ParameterAnnotationInfo
}

func (attr *AttributeInfo) ToRuntimeInvisibleParameterAnnotationsAttributeInfo() (*RuntimeInvisibleParameterAnnotationsAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	parameterAnnotationsLength := attr.Info[6:7][0]
	parameterAnnotations := make([]ParameterAnnotationInfo, parameterAnnotationsLength)
	for i := range parameterAnnotations {
		baseStart := uint32(i * 2)
		parameterAnnotationLeft := 7 + baseStart
		parameterAnnotationRight := 9 + baseStart
		parameterAnnotations[i] = ParameterAnnotationInfo{
			Annotations: []AnnotationInfo{
				{
					TypeIndex: binary.BigEndian.Uint16(attr.Info[parameterAnnotationLeft:parameterAnnotationRight]),
				},
			},
		}
	}
	return &RuntimeInvisibleParameterAnnotationsAttributeInfo{
		AttributeNameIndex:   attributeNameIndex,
		AttributeLength:      attributeLength,
		ParameterAnnotations: parameterAnnotations,
	}, nil
}

type AnnotationDefaultAttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	Default            []byte
}

func (attr *AttributeInfo) ToAnnotationDefaultAttributeInfo() (*AnnotationDefaultAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	defaultValue := attr.Info[6:]
	return &AnnotationDefaultAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		Default:            defaultValue,
	}, nil
}

type BootstrapMethodInfo struct {
	BootstrapMethodRef uint16
	BootstrapArguments []uint16
}

type BootstrapMethodsAttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	BootstrapMethods   []BootstrapMethodInfo
}

func (attr *AttributeInfo) ToBootstrapMethodsAttributeInfo() (*BootstrapMethodsAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	bootstrapMethodsLength := binary.BigEndian.Uint16(attr.Info[6:8])
	bootstrapMethods := make([]BootstrapMethodInfo, bootstrapMethodsLength)
	for i := range bootstrapMethods {
		baseStart := uint32(i * 2)
		bootstrapMethodLeft := 8 + baseStart
		bootstrapMethodRight := 10 + baseStart
		bootstrapMethods[i] = BootstrapMethodInfo{
			BootstrapMethodRef: binary.BigEndian.Uint16(attr.Info[bootstrapMethodLeft:bootstrapMethodRight]),
		}
	}
	return &BootstrapMethodsAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		BootstrapMethods:   bootstrapMethods,
	}, nil
}

type ParameterInfo struct {
	AccessFlags uint16
	NameIndex   uint16
	Descriptor  uint16
	Attributes  []AttributeInfo
}

type MethodParametersAttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	Parameters         []ParameterInfo
}

func (attr *AttributeInfo) ToMethodParametersAttributeInfo() (*MethodParametersAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	parametersLength := attr.Info[6:7][0]
	parameters := make([]ParameterInfo, parametersLength)
	for i := range parameters {
		baseStart := uint32(i * 2)
		parameterLeft := 7 + baseStart
		parameterRight := 9 + baseStart
		parameters[i] = ParameterInfo{
			AccessFlags: binary.BigEndian.Uint16(attr.Info[parameterLeft:parameterRight]),
		}
	}
	return &MethodParametersAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		Parameters:         parameters,
	}, nil
}

type ModuleInfo struct {
	ModuleNameIndex    uint16
	ModuleFlags        uint16
	ModuleVersionIndex uint16
	Requires           []RequiresInfo
	Exports            []ExportsInfo
	Opens              []OpensInfo
	Uses               []UsesInfo
	Provides           []ProvidesInfo
}

type RequiresInfo struct {
	RequiresIndex   uint16
	RequiresFlags   uint16
	RequiresVersion uint16
}

type ExportsInfo struct {
	ExportsIndex   uint16
	ExportsFlags   uint16
	ExportsToCount uint16
	ExportsToIndex []uint16
}

type OpensInfo struct {
	OpensIndex   uint16
	OpensFlags   uint16
	OpensToCount uint16
	OpensToIndex []uint16
}

type UsesInfo struct {
	UsesIndex uint16
}

type ProvidesInfo struct {
	ProvidesIndex   uint16
	ProvidesWith    uint16
	ProvidesWithLen uint16
}

type ModuleAttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	ModuleInfo         ModuleInfo
}

func (attr *AttributeInfo) ToModuleAttributeInfo() (*ModuleAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	moduleInfo := ModuleInfo{
		ModuleNameIndex:    binary.BigEndian.Uint16(attr.Info[6:8]),
		ModuleFlags:        binary.BigEndian.Uint16(attr.Info[8:10]),
		ModuleVersionIndex: binary.BigEndian.Uint16(attr.Info[10:12]),
	}
	return &ModuleAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		ModuleInfo:         moduleInfo,
	}, nil
}

type ModulePackagesAttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	PackageCount       uint16
	PackageIndex       []uint16
}

func (attr *AttributeInfo) ToModulePackagesAttributeInfo() (*ModulePackagesAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	packageCount := binary.BigEndian.Uint16(attr.Info[6:8])
	packageIndex := make([]uint16, packageCount)
	for i := range packageIndex {
		baseStart := uint32(i * 2)
		packageLeft := 8 + baseStart
		packageRight := 10 + baseStart
		packageIndex[i] = binary.BigEndian.Uint16(attr.Info[packageLeft:packageRight])
	}
	return &ModulePackagesAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		PackageCount:       packageCount,
		PackageIndex:       packageIndex,
	}, nil
}

type ModuleMainClassAttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	MainClassIndex     uint16
}

func (attr *AttributeInfo) ToModuleMainClassAttributeInfo() (*ModuleMainClassAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	mainClassIndex := binary.BigEndian.Uint16(attr.Info[6:8])
	return &ModuleMainClassAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		MainClassIndex:     mainClassIndex,
	}, nil
}

type NestHostAttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	ClassIndex         uint16
}

func (attr *AttributeInfo) ToNestHostAttributeInfo() (*NestHostAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	classIndex := binary.BigEndian.Uint16(attr.Info[6:8])
	return &NestHostAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		ClassIndex:         classIndex,
	}, nil
}

type NestMembersAttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	NumberOfClasses    uint16
	Classes            []uint16
}

func (attr *AttributeInfo) ToNestMembersAttributeInfo() (*NestMembersAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	numberOfClasses := binary.BigEndian.Uint16(attr.Info[6:8])
	classes := make([]uint16, numberOfClasses)
	for i := range classes {
		baseStart := uint32(i * 2)
		classLeft := 8 + baseStart
		classRight := 10 + baseStart
		classes[i] = binary.BigEndian.Uint16(attr.Info[classLeft:classRight])
	}
	return &NestMembersAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		NumberOfClasses:    numberOfClasses,
		Classes:            classes,
	}, nil
}

type PermittedSubclassesAttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	NumberOfClasses    uint16
	Classes            []uint16
}

func (attr *AttributeInfo) ToPermittedSubclassesAttributeInfo() (*PermittedSubclassesAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	numberOfClasses := binary.BigEndian.Uint16(attr.Info[6:8])
	classes := make([]uint16, numberOfClasses)
	for i := range classes {
		baseStart := uint32(i * 2)
		classLeft := 8 + baseStart
		classRight := 10 + baseStart
		classes[i] = binary.BigEndian.Uint16(attr.Info[classLeft:classRight])
	}
	return &PermittedSubclassesAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		NumberOfClasses:    numberOfClasses,
		Classes:            classes,
	}, nil
}

type RecordComponentInfo struct {
	ComponentNameIndex       uint16
	ComponentDescriptorIndex uint16
	ComponentAttributesCount uint16
}

type RecordAttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	NumberOfComponents uint16
	Components         []RecordComponentInfo
}

func (attr *AttributeInfo) ToRecordAttributeInfo() (*RecordAttributeInfo, error) {
	attributeNameIndex := binary.BigEndian.Uint16(attr.Info[0:2])
	attributeLength := binary.BigEndian.Uint32(attr.Info[2:6])
	numberOfComponents := binary.BigEndian.Uint16(attr.Info[6:8])
	components := make([]RecordComponentInfo, numberOfComponents)
	for i := range components {
		baseStart := uint32(i * 4)
		componentLeft := 8 + baseStart
		componentRight := 12 + baseStart
		components[i] = RecordComponentInfo{
			ComponentNameIndex:       binary.BigEndian.Uint16(attr.Info[componentLeft:componentRight]),
			ComponentDescriptorIndex: binary.BigEndian.Uint16(attr.Info[componentLeft+2 : componentRight+2]),
			ComponentAttributesCount: binary.BigEndian.Uint16(attr.Info[componentLeft+4 : componentRight+4]),
		}
	}
	return &RecordAttributeInfo{
		AttributeNameIndex: attributeNameIndex,
		AttributeLength:    attributeLength,
		NumberOfComponents: numberOfComponents,
		Components:         components,
	}, nil
}
