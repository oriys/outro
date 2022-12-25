package rtda

import (
	"io"
	"os"
	"outro/parser"
)

type ClassLoader interface {
	LoadClass(className string) (*Class, error)
}

type ApplicationClassLoader struct {
	classMap map[string]*Class
}

func NewApplicationClassLoader() *ApplicationClassLoader {
	return &ApplicationClassLoader{classMap: make(map[string]*Class)}
}

func (a *ApplicationClassLoader) LoadClass(className string) (*Class, error) {
	class := a.classMap[className]
	if class != nil {
		return class, nil
	}
	newClass := ParseClassFile(className)
	a.classMap[className] = newClass
	return newClass, nil
}

func ParseClassFile(name string) *Class {
	file, err := os.Open("java/classes/MethodInvoke.class")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	reader := parser.NewByteReader(bytes)
	parser := parser.NewClassFileParser(reader)
	class := parser.Parse()
	return NewClass(class)

}
