package test

import (
	"github.com/smartystreets/goconvey/convey"
	"io"
	"os"
	"outro/util"
	"testing"
)

func TestMethodInvoke(t *testing.T) {
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
	convey.ShouldEqual(class.Magic, 0xCAFEBABE)
	convey.ShouldEqual(class.MinorVersion, 0)
	convey.ShouldEqual(class.MajorVersion, 52)
	convey.ShouldEqual(class.ConstantPoolCount, 62)
	convey.ShouldEqual(class.AccessFlags, 33)
	convey.ShouldEqual(class.ThisClass, 10)
	convey.ShouldEqual(class.SuperClass, 11)
	convey.ShouldEqual(class.InterfacesCount, 0)
	convey.ShouldEqual(class.FieldsCount, 0)
	convey.ShouldEqual(class.MethodsCount, 5)
	convey.ShouldEqual(class.AttributesCount, 1)
	class.Print()
}
