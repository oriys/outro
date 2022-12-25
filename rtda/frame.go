package rtda

import (
	"outro/model"
)

type Frame struct {
	localVariables []interface{}
	operandStack   []interface{}
	constantPool   *[]model.ConstantInfo
	Method         *Method
	Thread         *Thread
	class          *Class
}

func (f *Frame) Execute() {
}

func (f *Frame) NextByte() int32 {
	return int32(f.Method.Code[f.Thread.PC+1])
}

func (f *Frame) ReadOffset() int32 {
	return int32(f.Method.Code[f.Thread.PC+1])<<8 | int32(f.Method.Code[f.Thread.PC+2])
}

func (f *Frame) NextShort() int32 {
	return int32(f.Method.Code[f.Thread.PC+1])<<8 + int32(f.Method.Code[f.Thread.PC+2])
}

func (f *Frame) LocalVariableInt(u uint16) int32 {
	return f.localVariables[u].(int32)
}

func (f *Frame) LocalVariableLong(u uint16) int64 {
	return f.localVariables[u].(int64)
}

func (f *Frame) LocalVariableFloat(u uint16) float32 {
	return f.localVariables[u].(float32)
}

func (f *Frame) LocalVariableDouble(u uint16) float64 {
	return f.localVariables[u].(float64)
}

func (f *Frame) LocalVariableRef(u uint16) *interface{} {
	return f.localVariables[u].(*interface{})
}

func (f *Frame) SetLocalVariableRef(u uint16, ref interface{}) {
	f.localVariables[u] = ref
}

func (f *Frame) SetLocalVariableInt(u uint16, val int32) {
	f.localVariables[u] = val
}

func (f *Frame) SetLocalVariableLong(u uint16, val int64) {
	f.localVariables[u] = val
}

func (f *Frame) SetLocalVariableFloat(u uint16, val float32) {
	f.localVariables[u] = val
}

func (f *Frame) SetLocalVariableDouble(u uint16, val float64) {
	f.localVariables[u] = val
}

func (f *Frame) Push(t interface{}) {
	f.operandStack = append(f.operandStack, t)
}

func (f *Frame) Pop() interface{} {
	t := f.operandStack[len(f.operandStack)-1]
	f.operandStack = f.operandStack[:len(f.operandStack)-1]
	return t
}

func (f *Frame) PopInt() int32 {
	return f.Pop().(int32)
}

func (f *Frame) PopLong() int64 {
	return f.Pop().(int64)
}

func (f *Frame) PopFloat() float32 {
	return f.Pop().(float32)
}

func (f *Frame) PopDouble() float64 {
	return f.Pop().(float64)
}

func (f *Frame) PushInt(i int32) {
	f.Push(i)
}

func (f *Frame) PushLong(l int64) {
	f.Push(l)
}

func (f *Frame) PushFloat(f32 float32) {
	f.Push(f32)
}

func (f *Frame) PushDouble(f64 float64) {
	f.Push(f64)
}

func (f *Frame) PushRef(ref interface{}) {
	f.operandStack = append(f.operandStack, ref)
}

func (f *Frame) PopIntArr() []int32 {
	return f.Pop().([]int32)
}

func (f *Frame) PopLongArr() []int64 {
	return f.Pop().([]int64)
}

func (f *Frame) PopFloatArr() []float32 {
	return f.Pop().([]float32)
}

func (f *Frame) PopDoubleArr() []float64 {
	return f.Pop().([]float64)

}

func (f *Frame) PopRefArr() []interface{} {
	return f.Pop().([]interface{})
}

func (f *Frame) PopByteArr() []int8 {
	return f.Pop().([]int8)
}

func (f *Frame) PopCharArr() []uint16 {
	return f.Pop().([]uint16)
}

func (f *Frame) PopShortArr() []int16 {
	return f.Pop().([]int16)
}

/*
		https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-2.html#jvms-2.6
	  	Each frame has its
			own array of local variables (§2.6.1),
			its own operand stack (§2.6.2),
			and a reference to the run-time constant pool (§2.5.5) of the class of the current method.
*/
func NewFrame(maxLocals, maxStack uint16, class *model.ClassFile) *Frame {
	return &Frame{
		localVariables: make([]interface{}, maxLocals),
		operandStack:   make([]interface{}, maxStack),
		constantPool:   &class.ConstantPool,
	}
}
