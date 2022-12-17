package model

import (
	"fmt"
)

var ConstantTagNameMap = map[uint8]string{
	7:  "CONSTANT_Class",
	9:  "CONSTANT_FieldRef",
	10: "CONSTANT_MethodRef",
	11: "CONSTANT_InterfaceMethodRef",
	8:  "CONSTANT_String",
	3:  "CONSTANT_Integer",
	4:  "CONSTANT_Float",
	5:  "CONSTANT_Long",
	6:  "CONSTANT_Double",
	12: "CONSTANT_NameAndType",
	1:  "CONSTANT_Utf8",
	15: "CONSTANT_MethodHandle",
	16: "CONSTANT_MethodType",
	18: "CONSTANT_InvokeDynamic",
}
var flagMap = map[uint16]string{
	0x0001: "ACC_PUBLIC",
	0x0010: "ACC_FINAL",
	0x0020: "ACC_SUPER",
	0x0200: "ACC_INTERFACE",
	0x0400: "ACC_ABSTRACT",
	0x1000: "ACC_SYNTHETIC",
	0x2000: "ACC_ANNOTATION",
	0x4000: "ACC_ENUM",
}

func (info MethodInfo) Print() {
	fmt.Println("Access Flags: ", info.AccessFlags)
	fmt.Println("Name Index: ", info.NameIndex)
	fmt.Println("Descriptor Index: ", info.DescriptorIndex)
	fmt.Println("Attributes Count: ", info.AttributesCount)
	for _, attributeInfo := range info.Attributes {
		attributeInfo.Print()
	}
}

func (c *Class) Print() {
	c.printMagic()
	c.printMinorVersion()
	c.printMajorVersion()
	c.printConstantPoolCount()
	c.printConstantPool()
	c.printAccessFlags()
	c.printThisClass()
	c.printSuperClass()
	c.printInterfacesCount()
	c.printInterfaces()
	c.printFieldsCount()
	c.printFields()
	c.printMethodsCount()
	c.printMethods()
	c.printAttributesCount()
	c.printAttributes()
}

func (c *Class) printMagic() {
	fmt.Printf("Magic: %x\n", c.Magic)
}

func (c *Class) printMinorVersion() {
	fmt.Println("Minor Version: ", c.MinorVersion)
}

func (c *Class) printMajorVersion() {
	fmt.Println("Major Version: ", c.MajorVersion)
}

func (c *Class) printConstantPoolCount() {
	fmt.Println("Constant Pool Count: ", c.ConstantPoolCount)
}

func (c *Class) printConstantPool() {
	for _, constantInfo := range c.ConstantPool {
		constantInfo.Print()
	}
}

func (c *Class) printAccessFlags() {
	fmt.Println("Access Flags: ", c.AccessFlags)
	fmt.Print("Access Flags: ")
	for flag, flagName := range flagMap {
		if c.AccessFlags&flag == flag {
			fmt.Print(flagName, " ")
		}
	}
}

func (c *Class) printThisClass() {
	fmt.Println("This Class: ", c.ThisClass)
}

func (c *Class) printSuperClass() {
	fmt.Println("Super Class: ", c.SuperClass)
}

func (c *Class) printInterfacesCount() {
	fmt.Print("Interfaces Count: ", c.InterfacesCount)
}

func (c *Class) printInterfaces() {
	for _, interfaceIndex := range c.Interfaces {
		fmt.Println("Interface: ", interfaceIndex)
	}
}

func (c *Class) printFieldsCount() {
	fmt.Println("Fields Count: ", c.FieldsCount)
}

func (c *Class) printFields() {
	for _, fieldInfo := range c.Fields {
		fieldInfo.Print()
	}
}

func (c *Class) printMethodsCount() {
	fmt.Println("Methods Count: ", c.MethodsCount)
}

func (c *Class) printMethods() {
	for _, methodInfo := range c.Methods {
		methodInfo.Print()
	}
}

func (c *Class) printAttributesCount() {
	fmt.Println("Attributes Count: ", c.AttributesCount)
}

func (c *Class) printAttributes() {
	for _, attributeInfo := range c.Attributes {
		attributeInfo.Print()
	}
}

func (info *ConstantInfo) Print() {
	fmt.Print("Tag: ", ConstantTagNameMap[info.Tag])
	switch info.Tag {
	case 1:
		fmt.Println(" ", string(info.Info))
	case 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 15, 16, 18:
		fmt.Println(" ", info.Info)
	}
}

func (info *FieldInfo) Print() {
	fmt.Println("Access Flags: ", info.AccessFlags)
	fmt.Println("Name Index: ", info.NameIndex)
	fmt.Println("Descriptor Index: ", info.DescriptorIndex)
	fmt.Println("Attributes Count: ", info.AttributesCount)
	for _, attributeInfo := range info.Attributes {
		attributeInfo.Print()
	}

}

func (i *AttributeInfo) Print() {
	fmt.Println("Attribute Name Index: ", i.AttributeNameIndex)
	fmt.Println("Attribute Length: ", i.AttributeLength)
	fmt.Println("Info: ", i.Info)
}
