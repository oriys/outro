package main

import (
	"outro/interpreter"
	"outro/rtda"
)

func main() {
	// read class file from java/Main.class
	// parse class file
	// create a new thread
	// create a new frame
	// push the frame to the thread
	// execute the frame
	loader := rtda.NewApplicationClassLoader()
	class, err := loader.LoadClass("java/classes/HelloWorld.class")
	checkErr(err)
	mainMethod, err := class.GetMainMethod()
	checkErr(err)
	thread := rtda.NewThread()
	thread.NewFrame(mainMethod)
	jvm := interpreter.JVM{}
	jvm.Thread = thread
	jvm.Execute()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
