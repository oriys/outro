package main

import (
	"io"
	"os"
	"outro/util"
)

func main() {
	// read class file from java/Main.class
	// parse class file
	// create a new thread
	// create a new frame
	// push the frame to the thread
	// execute the frame

	file, err := os.Open("java/Main.class")
	defer file.Close()
	checkErr(err)
	bytes, err := io.ReadAll(file)
	reader := util.NewByteReader(bytes)
	parser := util.NewClassFileParser(reader)
	parser.Parse()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
