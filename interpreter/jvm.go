package interpreter

import "outro/rtda"

type JVM struct {
	Thread *rtda.Thread
}

func (jvm *JVM) Execute() {
	frame := jvm.Thread.CurrentFrame()
	run(frame)
}

func run(frame *rtda.Frame) {
	opcodes := frame.Method.Code
	for {
		println(InstructDisplayNameMap[Instruct(opcodes[frame.Thread.PC])])
		opcode := Instruct(opcodes[frame.Thread.PC])
		pc, err := InstructFuncMap[opcode](frame)
		if err != nil {
			panic(err)
		}
		frame.Thread.PC = pc
	}
}
