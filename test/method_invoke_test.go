package test

import (
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"os"
	"outro/model"
	"outro/parser"
	"testing"
)

func TestMethodInvoke(t *testing.T) {
	file, err := os.Open("../java/classes/MethodInvoke.class")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	class := parser.NewClassFileParser(parser.NewByteReader(bytes)).Parse()
	assertClassParse(t, class)
}

func assertClassParse(t *testing.T, class *model.Class) {
	Convey("Test Method Invoke Parse", t, func() {
		So(class.Magic, ShouldEqual, 0xCAFEBABE)
		So(class.MinorVersion, ShouldEqual, 0)
		So(class.MajorVersion, ShouldEqual, 61)
		So(class.ConstantPoolCount, ShouldEqual, 61)
		So(class.AccessFlags, ShouldEqual, 33)
		So(class.ThisClass, ShouldEqual, 8)
		So(class.SuperClass, ShouldEqual, 2)
		So(class.InterfacesCount, ShouldEqual, 0)
		So(class.FieldsCount, ShouldEqual, 0)
		So(class.MethodsCount, ShouldEqual, 5)
		So(class.AttributesCount, ShouldEqual, 1)
	})
}
