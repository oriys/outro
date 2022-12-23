package rtda

type ConstantPool struct {
	constants []interface{}
}

func (c *ConstantPool) GetConstant(index uint16) interface{} {
	return c.constants[index]
}

type Method struct {
	accessFlag uint16
	name       string
	descriptor string
	maxStack   uint16
	maxLocals  uint16
	code       []byte
	class      *Class
}

type Field struct {
	accessFlag uint16
	name       string
	descriptor string
	class      *Class
}

type Class struct {
	accessFlag        uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        *Slots
}
