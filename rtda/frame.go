package rtda

import (
	"math"
	"outro/model"
)

type Slot struct {
	Num int32
	Ref interface{}
}

type Frame struct {
	LocalVariables []Slot
	OperandStack   *OperandStack
	ConstantPool   *[]model.ConstantInfo
	Method         *model.MethodInfo
	Thread         *Thread
}

func (f *Frame) Execute() {
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
		LocalVariables: make([]Slot, maxLocals),
		OperandStack:   &OperandStack{0, make([]Slot, maxStack)},
		ConstantPool:   &class.ConstantPool,
	}
}

type OperandStack struct {
	size  uint
	slots []Slot
}

func (s *OperandStack) Push(slot Slot) {
	s.slots[s.size] = slot
	s.size++
}

func (s *OperandStack) PushRef(t interface{}) {
	s.Push(Slot{Ref: t})
}

func (s *OperandStack) PopRef() interface{} {
	return s.Pop().Ref
}

func (s *OperandStack) Pop() Slot {
	s.size--
	return s.slots[s.size]
}

func (s *OperandStack) PushInt(val int32) {
	s.Push(Slot{Num: val})
}

func (s *OperandStack) PopInt() int32 {
	return s.Pop().Num
}

func (s *OperandStack) PopLong() int64 {
	high := s.Pop().Num
	low := s.Pop().Num
	return int64(high)<<32 | int64(low)
}

func (s *OperandStack) PushLong(val int64) {
	s.Push(Slot{Num: int32(val >> 32)})
	s.Push(Slot{Num: int32(val)})
}

func (s *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	s.Push(Slot{Num: int32(bits)})
	s.Push(Slot{Num: int32(bits >> 32)})
}

func (s *OperandStack) PopDouble() float64 {
	high := s.Pop().Num
	low := s.Pop().Num
	bits := uint64(high)<<32 | uint64(low)
	return math.Float64frombits(bits)
}

func (s *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	s.Push(Slot{Num: int32(bits)})
}

func (s *OperandStack) PopFloat() float32 {
	bits := uint32(s.Pop().Num)
	return math.Float32frombits(bits)
}
