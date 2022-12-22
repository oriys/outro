package main

import (
	"io"
	"os"
	"outro/model"
	"outro/rtda"
	"outro/util"
)

func main() {
	// read class file from java/Main.class
	// parse class file
	// create a new thread
	// create a new frame
	// push the frame to the thread
	// execute the frame

	file, err := os.Open("java/classes/MethodInvoke.class")
	defer file.Close()
	checkErr(err)
	bytes, err := io.ReadAll(file)
	reader := util.NewByteReader(bytes)
	parser := util.NewClassFileParser(reader)
	class := parser.Parse()
	mainMethod, err := class.FindMain()
	checkErr(err)
	frame := rtda.NewFrame(mainMethod.Attributes[0].MaxLocals, mainMethod.Attributes[0].MaxStack, class)
	thread := rtda.NewThread(frame, class)
	model.NewJVM(thread).Execute()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
