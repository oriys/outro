package rtda

import (
	"encoding/binary"
	"errors"
	"math"
	"outro/constant"
	"outro/model"
)

func (c *Class) GetConstant(index uint16) interface{} {
	return c.ConstantPool[index]
}

type Method struct {
	AccessFlag uint16
	Name       string
	Descriptor string
	MaxStack   uint16
	MaxLocals  uint16
	Code       []byte
	Class      *Class
}

type Field struct {
	AccessFlag uint16
	Name       string
	Descriptor string
	Class      *Class
}

type Class struct {
	AccessFlag        uint16
	Name              string
	SuperClassName    string
	InterfaceNames    []string
	ConstantPool      []interface{}
	Fields            []*Field
	Methods           []*Method
	Loader            *ApplicationClassLoader
	SuperClass        *Class
	Interfaces        []*Class
	InstanceSlotCount uint
	StaticSlotCount   uint
	StaticVars        []*interface{}
}

func (c *Class) GetMainMethod() (*Method, error) {
	return c.GetStaticMethod("main", "([Ljava/lang/String;)V")

}

func (c *Class) GetStaticMethod(s string, s2 string) (*Method, error) {
	for _, method := range c.Methods {
		if method.Name == s && method.Descriptor == s2 {
			return method, nil
		}
	}
	return nil, errors.New("method not found")
}

func NewClass(classFile *model.ClassFile) *Class {
	class := &Class{}
	class.AccessFlag = classFile.AccessFlags
	class.Name = string(classFile.ConstantPool[classFile.ThisClass].Info)
	class.SuperClassName = string(classFile.ConstantPool[classFile.SuperClass].Info)
	class.InterfaceNames = make([]string, len(classFile.Interfaces))
	for i, interfaceIndex := range classFile.Interfaces {
		class.InterfaceNames[i] = string(classFile.ConstantPool[interfaceIndex].Info)
	}
	class.ConstantPool = make([]interface{}, len(classFile.ConstantPool))
	for i, constantInfo := range classFile.ConstantPool {
		class.ConstantPool[i] = newConstant(constantInfo, class, classFile)
	}
	class.Fields = make([]*Field, len(classFile.Fields))
	for i, fieldInfo := range classFile.Fields {
		class.Fields[i] = newField(fieldInfo, class, classFile)
	}
	class.Methods = make([]*Method, len(classFile.Methods))
	for i, methodInfo := range classFile.Methods {
		class.Methods[i] = newMethod(methodInfo, class, classFile)
	}
	return class
}

func newMethod(info model.MethodInfo, class *Class, file *model.ClassFile) *Method {
	m := &Method{
		AccessFlag: info.AccessFlags,
		Name:       string(file.ConstantPool[info.NameIndex].Info),
		Descriptor: string(file.ConstantPool[info.DescriptorIndex].Info),
		Class:      class,
	}
	for _, attr := range info.Attributes {
		if string(file.ConstantPool[attr.AttributeNameIndex].Info) == "Code" {
			m.MaxStack = binary.BigEndian.Uint16(attr.Info[0:2])
			m.MaxLocals = binary.BigEndian.Uint16(attr.Info[2:4])
			m.Code = attr.Info[8:]
			break
		}
	}
	return m

}

func newField(info model.FieldInfo, class *Class, file *model.ClassFile) *Field {
	return &Field{
		AccessFlag: info.AccessFlags,
		Name:       string(file.ConstantPool[info.NameIndex].Info),
		Descriptor: string(file.ConstantPool[info.DescriptorIndex].Info),
		Class:      class,
	}
}

func newConstant(info model.ConstantInfo, class *Class, classFile *model.ClassFile) interface{} {
	switch info.Tag {
	case constant.ConstantInteger:
		return binary.BigEndian.Uint32(info.Info)
	case constant.ConstantFloat:
		return math.Float32frombits(binary.BigEndian.Uint32(info.Info))
	case constant.ConstantLong:
		return binary.BigEndian.Uint64(info.Info)
	case constant.ConstantDouble:
		return math.Float64frombits(binary.BigEndian.Uint64(info.Info))
	case constant.ConstantString:
		return string(info.Info)
	case constant.ConstantClass:
		return string(info.Info)
	case constant.ConstantFieldRef:
		return newFieldRef(info, class, classFile)
	case constant.ConstantMethodRef:
		return newMethodRef(info, class, classFile)
	case constant.ConstantInterfaceMethodRef:
		return newInterfaceMethodRef(info, class, classFile)
	case constant.ConstantNameAndType:
		return newNameAndType(info, class, classFile)
	case constant.ConstantMethodHandle:
		return newMethodHandle(info, class, classFile)
	case constant.ConstantMethodType:
		return newMethodType(info, class, classFile)
	case constant.ConstantInvokeDynamic:
		return newInvokeDynamic(info, class, classFile)
	}
	return nil

}

func newInvokeDynamic(info model.ConstantInfo, class *Class, file *model.ClassFile) interface{} {
	return nil
}

func newMethodType(info model.ConstantInfo, class *Class, file *model.ClassFile) interface{} {
	return string(file.ConstantPool[info.Info[0]].Info)

}

type MethodHandle struct {
	ReferenceKind  uint8
	ReferenceIndex uint16
}

func newMethodHandle(info model.ConstantInfo, class *Class, file *model.ClassFile) *MethodHandle {
	return &MethodHandle{
		ReferenceKind:  info.Info[0],
		ReferenceIndex: binary.BigEndian.Uint16(info.Info[1:]),
	}
}

type NameAndType struct {
	Name       string
	Descriptor string
}

func newNameAndType(info model.ConstantInfo, class *Class, file *model.ClassFile) *NameAndType {
	return &NameAndType{
		Name:       string(file.ConstantPool[binary.BigEndian.Uint16(info.Info[0:2])].Info),
		Descriptor: string(file.ConstantPool[binary.BigEndian.Uint16(info.Info[2:4])].Info),
	}
}

type InterfaceMethodRef struct {
	ClassName       string
	NameAndTypeName string
}

func newInterfaceMethodRef(info model.ConstantInfo, class *Class, file *model.ClassFile) *InterfaceMethodRef {
	return &InterfaceMethodRef{
		ClassName:       string(file.ConstantPool[binary.BigEndian.Uint16(info.Info[0:2])].Info),
		NameAndTypeName: string(file.ConstantPool[binary.BigEndian.Uint16(info.Info[2:4])].Info),
	}
}

type MethodRef struct {
	ClassName       string
	NameAndTypeName string
	Descriptor      string
	Class           *Class
	ResolvedMethod  *Method
}

func newMethodRef(info model.ConstantInfo, class *Class, file *model.ClassFile) *MethodRef {
	return &MethodRef{
		ClassName:       string(file.ConstantPool[binary.BigEndian.Uint16(info.Info[0:2])].Info),
		NameAndTypeName: string(file.ConstantPool[binary.BigEndian.Uint16(info.Info[2:4])].Info),
		Class:           class,
	}
}

type FieldRef struct {
	ClassName       string
	NameAndTypeName string
	Class           *Class
	ResolvedField   *Field
}

func newFieldRef(info model.ConstantInfo, class *Class, file *model.ClassFile) *FieldRef {
	return &FieldRef{
		ClassName:       string(file.ConstantPool[binary.BigEndian.Uint16(info.Info[0:2])].Info),
		NameAndTypeName: string(file.ConstantPool[binary.BigEndian.Uint16(info.Info[2:4])].Info),
		Class:           class,
	}
}
