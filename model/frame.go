package model

type Slot struct {
	Num int32
	Ref interface{}
}

type Frame struct {
	LocalVariables []Slot
	OperandStack   *OperandStack
	ConstantPool   *[]ConstantInfo
}

/*
		https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-2.html#jvms-2.6
	  	Each frame has its
			own array of local variables (§2.6.1),
			its own operand stack (§2.6.2),
			and a reference to the run-time constant pool (§2.5.5) of the class of the current method.
*/
func NewFrame(maxLocals, maxStack uint16, class *Class) *Frame {
	return &Frame{
		LocalVariables: make([]Slot, maxLocals),
		OperandStack:   &OperandStack{0, make([]Slot, maxStack)},
		ConstantPool:   &class.ConstantPool,
	}
}

type OperandStack struct {
	size  uint
	slots []Slot
}
