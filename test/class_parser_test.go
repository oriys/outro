package test

import (
	cy "github.com/smartystreets/goconvey/convey"
	"io"
	"os"
	"outro/util"
	"testing"
)

func TestClassParse(t *testing.T) {
	file, err := os.Open("../java/Solution.class")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	class := util.NewClassFileParser(util.NewByteReader(bytes)).Parse()
	cy.ShouldEqual(class.Magic, 0xCAFEBABE)
	cy.ShouldEqual(class.MinorVersion, 0)
	cy.ShouldEqual(class.MajorVersion, 52)
	cy.ShouldEqual(class.ConstantPoolCount, 27)
	cy.ShouldEqual(class.AccessFlags, 32)
	cy.ShouldEqual(class.ThisClass, 2)
	cy.ShouldEqual(class.SuperClass, 3)
	cy.ShouldEqual(class.InterfacesCount, 0)
	cy.ShouldEqual(class.FieldsCount, 0)
	cy.ShouldEqual(class.MethodsCount, 2)
	cy.ShouldEqual(class.AttributesCount, 1)

}
