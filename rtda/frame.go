package rtda

import (
	"outro/model"
)

type Frame struct {
	LocalVariables []interface{}
	OperandStack   *OperandStack
	ConstantPool   *[]model.ConstantInfo
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
	return f.LocalVariables[u].(int32)
}

func (f *Frame) LocalVariableLong(u uint16) int64 {
	return f.LocalVariables[u].(int64)
}

func (f *Frame) LocalVariableFloat(u uint16) float32 {
	return f.LocalVariables[u].(float32)
}

func (f *Frame) LocalVariableDouble(u uint16) float64 {
	return f.LocalVariables[u].(float64)
}

func (f *Frame) LocalVariableRef(u uint16) *interface{} {
	return f.LocalVariables[u].(*interface{})
}

func (f *Frame) SetLocalVariableRef(u uint16, ref interface{}) {
	f.LocalVariables[u] = ref
}

func (f *Frame) SetLocalVariableInt(u uint16, val int32) {
	f.LocalVariables[u] = val
}

func (f *Frame) SetLocalVariableLong(u uint16, val int64) {
	f.LocalVariables[u] = val
}

func (f *Frame) SetLocalVariableFloat(u uint16, val float32) {
	f.LocalVariables[u] = val
}

func (f *Frame) SetLocalVariableDouble(u uint16, val float64) {
	f.LocalVariables[u] = val
}

/*
		https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-2.html#jvms-2.6
	  	Each frame has its
			own array of local variables (§2.6.1),
			its own operand stack (§2.6.2),
			and a reference to the run-time constant pool (§2.5.5) of the class of the current method.
*/
func NewFrame(maxLocals, maxStack uint16, class *model.Class) *Frame {
	return &Frame{
		LocalVariables: make([]interface{}, maxLocals),
		OperandStack:   &OperandStack{0, make([]interface{}, maxStack)},
		ConstantPool:   &class.ConstantPool,
	}
}

type OperandStack struct {
	size  uint
	slots []interface{}
}

func (s *OperandStack) Push(val interface{}) {
	s.slots[s.size] = val
	s.size++
}

func (s *OperandStack) Pop() interface{} {
	s.size--
	return s.slots[s.size]
}

func (s *OperandStack) PushInt(val int32) {
	s.Push(val)
}

func (s *OperandStack) PopInt() int32 {
	return s.Pop().(int32)
}

func (s *OperandStack) PushFloat(val float32) {
	s.Push(val)
}

func (s *OperandStack) PopFloat() float32 {
	return s.Pop().(float32)
}

func (s *OperandStack) PushLong(val int64) {
	s.Push(val)
}

func (s *OperandStack) PopLong() int64 {
	return s.Pop().(int64)
}

func (s *OperandStack) PushDouble(val float64) {
	s.Push(val)
}

func (s *OperandStack) PopDouble() float64 {
	return s.Pop().(float64)
}

func (s *OperandStack) PopIntArr() []int32 {
	return s.Pop().([]int32)
}

func (s *OperandStack) PopLongArr() []int64 {
	return s.Pop().([]int64)
}

func (s *OperandStack) PopFloatArr() []float32 {
	return s.Pop().([]float32)
}

func (s *OperandStack) PopDoubleArr() []float64 {
	return s.Pop().([]float64)
}

// pop aa
func (s *OperandStack) PopRefArr() []interface{} {
	return s.Pop().([]interface{})
}

func (s *OperandStack) PushIntArr(arr []int32) {
	s.Push(arr)
}

func (s *OperandStack) PushLongArr(arr []int64) {
	s.Push(arr)
}

func (s *OperandStack) PushFloatArr(arr []float32) {
	s.Push(arr)
}

func (s *OperandStack) PushDoubleArr(arr []float64) {
	s.Push(arr)
}

func (s *OperandStack) PopByteArr() []int8 {
	return s.Pop().([]int8)

}

func (s *OperandStack) PopCharArr() []rune {
	return s.Pop().([]rune)

}

func (s *OperandStack) PopShortArr() []int16 {
	return s.Pop().([]int16)
}
