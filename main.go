package main

import (
	"io"
	"os"
	"outro/parser"
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
	reader := parser.NewByteReader(bytes)
	parser := parser.NewClassFileParser(reader)
	class := parser.Parse()
	mainMethod, err := class.getMethod("main", "([Ljava/lang/String;)V")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
