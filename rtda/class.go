package rtda

type ConstantPool struct {
	constants []interface{}
}

func (c *ConstantPool) GetConstant(index uint16) interface{} {
	return c.constants[index]
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
	ConstantPool      *ConstantPool
	Fields            []*Field
	Methods           []*Method
	Loader            *ApplicationClassLoader
	SuperClass        *Class
	Interfaces        []*Class
	InstanceSlotCount uint
	StaticSlotCount   uint
	StaticVars        []*interface{}
}
