package base

const (
	// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4
	ConstantClass              uint8 = 7
	ConstantFieldRef           uint8 = 9
	ConstantMethodRef          uint8 = 10
	ConstantInterfaceMethodRef uint8 = 11
	ConstantString             uint8 = 8
	ConstantInteger            uint8 = 3
	ConstantFloat              uint8 = 4
	ConstantLong               uint8 = 5
	ConstantDouble             uint8 = 6
	ConstantNameAndType        uint8 = 12
	ConstantUtf8               uint8 = 1
	ConstantMethodHandle       uint8 = 15
	ConstantMethodType         uint8 = 16
	ConstantInvokeDynamic      uint8 = 18
)
