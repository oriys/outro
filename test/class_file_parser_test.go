package test

import (
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"os"
	"outro/util"
	"testing"
)

func TestClassParse(t *testing.T) {
	file, err := os.Open("../java/classes/HelloWorld.class")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	class := util.NewClassFileParser(util.NewByteReader(bytes)).Parse()
	Convey("Test Class Parse", t, func() {
		So(class.Magic, ShouldEqual, 0xCAFEBABE)
		So(class.MinorVersion, ShouldEqual, 0)
		So(class.MajorVersion, ShouldEqual, 61)
		So(class.ConstantPoolCount, ShouldEqual, 34)
		So(class.AccessFlags, ShouldEqual, 33)
		So(class.ThisClass, ShouldEqual, 21)
		So(class.SuperClass, ShouldEqual, 2)
		So(class.InterfacesCount, ShouldEqual, 0)
		So(class.FieldsCount, ShouldEqual, 0)
		So(class.MethodsCount, ShouldEqual, 2)
		So(class.AttributesCount, ShouldEqual, 1)
	})
}
