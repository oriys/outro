package interpreter

import (
	"errors"
	"math"
	"outro/rtda"
)

type Instruct uint8

const (
	NOP             Instruct = 0x00
	ACONST_NULL     Instruct = 0x01
	ICONST_M1       Instruct = 0x02
	ICONST_0        Instruct = 0x03
	ICONST_1        Instruct = 0x04
	ICONST_2        Instruct = 0x05
	ICONST_3        Instruct = 0x06
	ICONST_4        Instruct = 0x07
	ICONST_5        Instruct = 0x08
	LCONST_0        Instruct = 0x09
	LCONST_1        Instruct = 0x0a
	FCONST_0        Instruct = 0x0b
	FCONST_1        Instruct = 0x0c
	FCONST_2        Instruct = 0x0d
	DCONST_0        Instruct = 0x0e
	DCONST_1        Instruct = 0x0f
	BIPUSH          Instruct = 0x10
	SIPUSH          Instruct = 0x11
	LDC             Instruct = 0x12
	LDC_W           Instruct = 0x13
	LDC2_W          Instruct = 0x14
	ILOAD           Instruct = 0x15
	LLOAD           Instruct = 0x16
	FLOAD           Instruct = 0x17
	DLOAD           Instruct = 0x18
	ALOAD           Instruct = 0x19
	ILOAD_0         Instruct = 0x1a
	ILOAD_1         Instruct = 0x1b
	ILOAD_2         Instruct = 0x1c
	ILOAD_3         Instruct = 0x1d
	LLOAD_0         Instruct = 0x1e
	LLOAD_1         Instruct = 0x1f
	LLOAD_2         Instruct = 0x20
	LLOAD_3         Instruct = 0x21
	FLOAD_0         Instruct = 0x22
	FLOAD_1         Instruct = 0x23
	FLOAD_2         Instruct = 0x24
	FLOAD_3         Instruct = 0x25
	DLOAD_0         Instruct = 0x26
	DLOAD_1         Instruct = 0x27
	DLOAD_2         Instruct = 0x28
	DLOAD_3         Instruct = 0x29
	ALOAD_0         Instruct = 0x2a
	ALOAD_1         Instruct = 0x2b
	ALOAD_2         Instruct = 0x2c
	ALOAD_3         Instruct = 0x2d
	IALOAD          Instruct = 0x2e
	LALOAD          Instruct = 0x2f
	FALOAD          Instruct = 0x30
	DALOAD          Instruct = 0x31
	AALOAD          Instruct = 0x32
	BALOAD          Instruct = 0x33
	CALOAD          Instruct = 0x34
	SALOAD          Instruct = 0x35
	ISTORE          Instruct = 0x36
	LSTORE          Instruct = 0x37
	FSTORE          Instruct = 0x38
	DSTORE          Instruct = 0x39
	ASTORE          Instruct = 0x3a
	ISTORE_0        Instruct = 0x3b
	ISTORE_1        Instruct = 0x3c
	ISTORE_2        Instruct = 0x3d
	ISTORE_3        Instruct = 0x3e
	LSTORE_0        Instruct = 0x3f
	LSTORE_1        Instruct = 0x40
	LSTORE_2        Instruct = 0x41
	LSTORE_3        Instruct = 0x42
	FSTORE_0        Instruct = 0x43
	FSTORE_1        Instruct = 0x44
	FSTORE_2        Instruct = 0x45
	FSTORE_3        Instruct = 0x46
	DSTORE_0        Instruct = 0x47
	DSTORE_1        Instruct = 0x48
	DSTORE_2        Instruct = 0x49
	DSTORE_3        Instruct = 0x4a
	ASTORE_0        Instruct = 0x4b
	ASTORE_1        Instruct = 0x4c
	ASTORE_2        Instruct = 0x4d
	ASTORE_3        Instruct = 0x4e
	IASTORE         Instruct = 0x4f
	LASTORE         Instruct = 0x50
	FASTORE         Instruct = 0x51
	DASTORE         Instruct = 0x52
	AASTORE         Instruct = 0x53
	BASTORE         Instruct = 0x54
	CASTORE         Instruct = 0x55
	SASTORE         Instruct = 0x56
	POP             Instruct = 0x57
	POP2            Instruct = 0x58
	DUP             Instruct = 0x59
	DUP_X1          Instruct = 0x5a
	DUP_X2          Instruct = 0x5b
	DUP2            Instruct = 0x5c
	DUP2_X1         Instruct = 0x5d
	DUP2_X2         Instruct = 0x5e
	SWAP            Instruct = 0x5f
	IADD            Instruct = 0x60
	LADD            Instruct = 0x61
	FADD            Instruct = 0x62
	DADD            Instruct = 0x63
	ISUB            Instruct = 0x64
	LSUB            Instruct = 0x65
	FSUB            Instruct = 0x66
	DSUB            Instruct = 0x67
	IMUL            Instruct = 0x68
	LMUL            Instruct = 0x69
	FMUL            Instruct = 0x6a
	DMUL            Instruct = 0x6b
	IDIV            Instruct = 0x6c
	LDIV            Instruct = 0x6d
	FDIV            Instruct = 0x6e
	DDIV            Instruct = 0x6f
	IREM            Instruct = 0x70
	LREM            Instruct = 0x71
	FREM            Instruct = 0x72
	DREM            Instruct = 0x73
	INEG            Instruct = 0x74
	LNEG            Instruct = 0x75
	FNEG            Instruct = 0x76
	DNEG            Instruct = 0x77
	ISHL            Instruct = 0x78
	LSHL            Instruct = 0x79
	ISHR            Instruct = 0x7a
	LSHR            Instruct = 0x7b
	IUSHR           Instruct = 0x7c
	LUSHR           Instruct = 0x7d
	IAND            Instruct = 0x7e
	LAND            Instruct = 0x7f
	IOR             Instruct = 0x80
	LOR             Instruct = 0x81
	IXOR            Instruct = 0x82
	LXOR            Instruct = 0x83
	IINC            Instruct = 0x84
	I2L             Instruct = 0x85
	I2F             Instruct = 0x86
	I2D             Instruct = 0x87
	L2I             Instruct = 0x88
	L2F             Instruct = 0x89
	L2D             Instruct = 0x8a
	F2I             Instruct = 0x8b
	F2L             Instruct = 0x8c
	F2D             Instruct = 0x8d
	D2I             Instruct = 0x8e
	D2L             Instruct = 0x8f
	D2F             Instruct = 0x90
	I2B             Instruct = 0x91
	I2C             Instruct = 0x92
	I2S             Instruct = 0x93
	LCMP            Instruct = 0x94
	FCMPL           Instruct = 0x95
	FCMPG           Instruct = 0x96
	DCMPL           Instruct = 0x97
	DCMPG           Instruct = 0x98
	IFEQ            Instruct = 0x99
	IFNE            Instruct = 0x9a
	IFLT            Instruct = 0x9b
	IFGE            Instruct = 0x9c
	IFGT            Instruct = 0x9d
	IFLE            Instruct = 0x9e
	IF_ICMPEQ       Instruct = 0x9f
	IF_ICMPNE       Instruct = 0xa0
	IF_ICMPLT       Instruct = 0xa1
	IF_ICMPGE       Instruct = 0xa2
	IF_ICMPGT       Instruct = 0xa3
	IF_ICMPLE       Instruct = 0xa4
	IF_ACMPEQ       Instruct = 0xa5
	IF_ACMPNE       Instruct = 0xa6
	GOTO            Instruct = 0xa7
	JSR             Instruct = 0xa8
	RET             Instruct = 0xa9
	TABLESWITCH     Instruct = 0xaa
	LOOKUPSWITCH    Instruct = 0xab
	IRETURN         Instruct = 0xac
	LRETURN         Instruct = 0xad
	FRETURN         Instruct = 0xae
	DRETURN         Instruct = 0xaf
	ARETURN         Instruct = 0xb0
	RETURN          Instruct = 0xb1
	GETSTATIC       Instruct = 0xb2
	PUTSTATIC       Instruct = 0xb3
	GETFIELD        Instruct = 0xb4
	PUTFIELD        Instruct = 0xb5
	INVOKEVIRTUAL   Instruct = 0xb6
	INVOKESPECIAL   Instruct = 0xb7
	INVOKESTATIC    Instruct = 0xb8
	INVOKEINTERFACE Instruct = 0xb9
	INVOKEDYNAMIC   Instruct = 0xba
	NEW             Instruct = 0xbb
	NEWARRAY        Instruct = 0xbc
	ANEWARRAY       Instruct = 0xbd
	ARRAYLENGTH     Instruct = 0xbe
	ATHROW          Instruct = 0xbf
	CHECKCAST       Instruct = 0xc0
	INSTANCEOF      Instruct = 0xc1
	MONITORENTER    Instruct = 0xc2
	MONITOREXIT     Instruct = 0xc3
	WIDE            Instruct = 0xc4
	MULTIANEWARRAY  Instruct = 0xc5
	IFNULL          Instruct = 0xc6
	IFNONNULL       Instruct = 0xc7
	GOTO_W          Instruct = 0xc8
	JSR_W           Instruct = 0xc9
	BREAKPOINT      Instruct = 0xca
	IMPDEP1         Instruct = 0xfe
	IMPDEP2         Instruct = 0xff
)

var InstructDisplayNameMap = map[Instruct]string{

	0x00: "nop",
	0x01: "aconst_null",
	0x02: "iconst_m1",
	0x03: "iconst_0",
	0x04: "iconst_1",
	0x05: "iconst_2",
	0x06: "iconst_3",
	0x07: "iconst_4",
	0x08: "iconst_5",
	0x09: "lconst_0",
	0x0a: "lconst_1",
	0x0b: "fconst_0",
	0x0c: "fconst_1",
	0x0d: "fconst_2",
	0x0e: "dconst_0",
	0x0f: "dconst_1",
	0x10: "bipush",
	0x11: "sipush",
	0x12: "ldc",
	0x13: "ldc_w",
	0x14: "ldc2_w",
	0x15: "iload",
	0x16: "lload",
	0x17: "fload",
	0x18: "dload",
	0x19: "aload",
	0x1a: "iload_0",
	0x1b: "iload_1",
	0x1c: "iload_2",
	0x1d: "iload_3",
	0x1e: "lload_0",
	0x1f: "lload_1",
	0x20: "lload_2",
	0x21: "lload_3",
	0x22: "fload_0",
	0x23: "fload_1",
	0x24: "fload_2",
	0x25: "fload_3",
	0x26: "dload_0",
	0x27: "dload_1",
	0x28: "dload_2",
	0x29: "dload_3",
	0x2a: "aload_0",
	0x2b: "aload_1",
	0x2c: "aload_2",
	0x2d: "aload_3",
	0x2e: "iaload",
	0x2f: "laload",
	0x30: "faload",
	0x31: "daload",
	0x32: "aaload",
	0x33: "baload",
	0x34: "caload",
	0x35: "saload",
	0x36: "istore",
	0x37: "lstore",
	0x38: "fstore",
	0x39: "dstore",
	0x3a: "astore",
	0x3b: "istore_0",
	0x3c: "istore_1",
	0x3d: "istore_2",
	0x3e: "istore_3",
	0x3f: "lstore_0",
	0x40: "lstore_1",
	0x41: "lstore_2",
	0x42: "lstore_3",
	0x43: "fstore_0",
	0x44: "fstore_1",
	0x45: "fstore_2",
	0x46: "fstore_3",
	0x47: "dstore_0",
	0x48: "dstore_1",
	0x49: "dstore_2",
	0x4a: "dstore_3",
	0x4b: "astore_0",
	0x4c: "astore_1",
	0x4d: "astore_2",
	0x4e: "astore_3",
	0x4f: "iastore",
	0x50: "lastore",
	0x51: "fastore",
	0x52: "dastore",
	0x53: "aastore",
	0x54: "bastore",
	0x55: "castore",
	0x56: "sastore",
	0x57: "pop",
	0x58: "pop2",
	0x59: "dup",
	0x5a: "dup_x1",
	0x5b: "dup_x2",
	0x5c: "dup2",
	0x5d: "dup2_x1",
	0x5e: "dup2_x2",
	0x5f: "swap",
	0x60: "iadd",
	0x61: "ladd",
	0x62: "fadd",
	0x63: "dadd",
	0x64: "isub",
	0x65: "lsub",
	0x66: "fsub",
	0x67: "dsub",
	0x68: "imul",
	0x69: "lmul",
	0x6a: "fmul",
	0x6b: "dmul",
	0x6c: "idiv",
	0x6d: "ldiv",
	0x6e: "fdiv",
	0x6f: "ddiv",
	0x70: "irem",
	0x71: "lrem",
	0x72: "frem",
	0x73: "drem",
	0x74: "ineg",
	0x75: "lneg",
	0x76: "fneg",
	0x77: "dneg",
	0x78: "ishl",
	0x79: "lshl",
	0x7a: "ishr",
	0x7b: "lshr",
	0x7c: "iushr",
	0x7d: "lushr",
	0x7e: "iand",
	0x7f: "land",
	0x80: "ior",
	0x81: "lor",
	0x82: "ixor",
	0x83: "lxor",
	0x84: "iinc",
	0x85: "i2l",
	0x86: "i2f",
	0x87: "i2d",
	0x88: "l2i",
	0x89: "l2f",
	0x8a: "l2d",
	0x8b: "f2i",
	0x8c: "f2l",
	0x8d: "f2d",
	0x8e: "d2i",
	0x8f: "d2l",
	0x90: "d2f",
	0x91: "i2b",
	0x92: "i2c",
	0x93: "i2s",
	0x94: "lcmp",
	0x95: "fcmpl",
	0x96: "fcmpg",
	0x97: "dcmpl",
	0x98: "dcmpg",
	0x99: "ifeq",
	0x9a: "ifne",
	0x9b: "iflt",
	0x9c: "ifge",
	0x9d: "ifgt",
	0x9e: "ifle",
	0x9f: "if_icmpeq",
	0xa0: "if_icmpne",
	0xa1: "if_icmplt",
	0xa2: "if_icmpge",
	0xa3: "if_icmpgt",
	0xa4: "if_icmple",
	0xa5: "if_acmpeq",
	0xa6: "if_acmpne",
	0xa7: "goto",
	0xa8: "jsr",
	0xa9: "ret",
	0xaa: "tableswitch",
	0xab: "lookupswitch",
	0xac: "ireturn",
	0xad: "lreturn",
	0xae: "freturn",
	0xaf: "dreturn",
	0xb0: "areturn",
	0xb1: "return",
	0xb2: "getstatic",
	0xb3: "putstatic",
	0xb4: "getfield",
	0xb5: "putfield",
	0xb6: "invokevirtual",
	0xb7: "invokespecial",
	0xb8: "invokestatic",
	0xb9: "invokeinterface",
	0xba: "invokedynamic",
	0xbb: "new",
	0xbc: "newarray",
	0xbd: "anewarray",
	0xbe: "arraylength",
	0xbf: "athrow",
	0xc0: "checkcast",
	0xc1: "instanceof",
	0xc2: "monitorenter",
	0xc3: "monitorexit",
	0xc4: "wide",
	0xc5: "multianewarray",
	0xc6: "ifnull",
	0xc7: "ifnonnull",
	0xc8: "goto_w",
	0xc9: "jsr_w",
	0xca: "breakpoint",
	0xfe: "impdep1",
	0xff: "impdep2",
}

var InstructFuncMap = map[Instruct]func(frame *rtda.Frame) (pc int, err error){
	NOP: func(frame *rtda.Frame) (int, error) {
		return 1 + frame.Thread.PC, nil
	},
	ACONST_NULL: func(frame *rtda.Frame) (int, error) {
		frame.OperandStack.Push(nil)
		return 1 + frame.Thread.PC, nil
	},
	ICONST_M1: func(frame *rtda.Frame) (int, error) {
		frame.OperandStack.PushInt(-1)
		return 1 + frame.Thread.PC, nil
	},
	ICONST_0: func(frame *rtda.Frame) (int, error) {
		frame.OperandStack.PushInt(0)
		return 1 + frame.Thread.PC, nil
	},
	ICONST_1: func(frame *rtda.Frame) (int, error) {
		frame.OperandStack.PushInt(1)
		return 1 + frame.Thread.PC, nil
	},
	ICONST_2: func(frame *rtda.Frame) (int, error) {
		frame.OperandStack.PushInt(2)
		return 1 + frame.Thread.PC, nil
	},
	ICONST_3: func(frame *rtda.Frame) (int, error) {
		frame.OperandStack.PushInt(3)
		return 1 + frame.Thread.PC, nil
	},
	ICONST_4: func(frame *rtda.Frame) (int, error) {
		frame.OperandStack.PushInt(4)
		return 1 + frame.Thread.PC, nil
	},
	ICONST_5: func(frame *rtda.Frame) (int, error) {
		frame.OperandStack.PushInt(5)
		return 1 + frame.Thread.PC, nil
	},
	LCONST_0: func(frame *rtda.Frame) (int, error) {
		frame.OperandStack.PushLong(0)
		return 1 + frame.Thread.PC, nil
	},
	LCONST_1: func(frame *rtda.Frame) (int, error) {
		frame.OperandStack.PushLong(1)
		return 1 + frame.Thread.PC, nil
	},
	FCONST_0: func(frame *rtda.Frame) (int, error) {
		frame.OperandStack.PushFloat(0)
		return 1 + frame.Thread.PC, nil
	},
	FCONST_1: func(frame *rtda.Frame) (int, error) {
		frame.OperandStack.PushFloat(1)
		return 1 + frame.Thread.PC, nil
	},
	FCONST_2: func(frame *rtda.Frame) (int, error) {
		frame.OperandStack.PushFloat(2)
		return 1 + frame.Thread.PC, nil
	},
	DCONST_0: func(frame *rtda.Frame) (int, error) {
		frame.OperandStack.PushDouble(0)
		return 1 + frame.Thread.PC, nil
	},
	DCONST_1: func(frame *rtda.Frame) (int, error) {
		frame.OperandStack.PushDouble(1)
		return 1 + frame.Thread.PC, nil
	},
	BIPUSH: func(frame *rtda.Frame) (int, error) {
		i := int32(frame.NextByte())
		frame.OperandStack.PushInt(i)
		return 2 + frame.Thread.PC, nil
	},
	SIPUSH: func(frame *rtda.Frame) (int, error) {
		i := int32(frame.NextShort())
		frame.OperandStack.PushInt(i)
		return 3 + frame.Thread.PC, nil
	},
	LDC: func(frame *rtda.Frame) (int, error) {
		cp := frame.Method.Class.ConstantPool
		index := uint(frame.NextByte())
		c := cp.GetConstant(uint16(index))
		switch c.(type) {
		case int32:
			frame.OperandStack.PushInt(c.(int32))
		case float32:
			frame.OperandStack.PushFloat(c.(float32))
		case string:
			frame.OperandStack.Push(rtda.NewJString(frame.Method.Class.Loader, c.(string)))
		default:
			panic("todo: ldc!")
		}
		return 2 + frame.Thread.PC, nil
	},
	LDC_W: func(frame *rtda.Frame) (int, error) {
		cp := frame.Method.Class.ConstantPool
		index := frame.NextShort()
		c := cp.GetConstant(uint16(index))
		switch c.(type) {
		case int32:
			frame.OperandStack.PushInt(c.(int32))
		case float32:
			frame.OperandStack.PushFloat(c.(float32))
		case string:
			frame.OperandStack.Push(rtda.NewJString(frame.Method.Class.Loader, c.(string)))
		default:
			panic("todo: ldc_w!")
		}
		return 3 + frame.Thread.PC, nil
	},
	LDC2_W: func(frame *rtda.Frame) (int, error) {
		cp := frame.Method.Class.ConstantPool
		index := frame.NextShort()
		c := cp.GetConstant(uint16(index))
		switch c.(type) {
		case int64:
			frame.OperandStack.PushLong(c.(int64))
		case float64:
			frame.OperandStack.PushDouble(c.(float64))
		default:
			panic("todo: ldc2_w!")
		}
		return 3 + frame.Thread.PC, nil
	},
	ILOAD: func(frame *rtda.Frame) (int, error) {
		index := uint16(frame.NextByte())
		val := frame.LocalVariableInt(index)
		frame.OperandStack.PushInt(val)
		return 2 + frame.Thread.PC, nil
	},
	LLOAD: func(frame *rtda.Frame) (int, error) {
		index := uint16(frame.NextByte())
		val := frame.LocalVariableLong(index)
		frame.OperandStack.PushLong(val)
		return 2 + frame.Thread.PC, nil
	},
	FLOAD: func(frame *rtda.Frame) (int, error) {
		index := uint16(frame.NextByte())
		val := frame.LocalVariableFloat(index)
		frame.OperandStack.PushFloat(val)
		return 2 + frame.Thread.PC, nil
	},
	DLOAD: func(frame *rtda.Frame) (int, error) {
		index := uint16(frame.NextByte())
		val := frame.LocalVariableDouble(index)
		frame.OperandStack.PushDouble(val)
		return 2 + frame.Thread.PC, nil
	},
	ALOAD: func(frame *rtda.Frame) (int, error) {
		index := uint16(frame.NextByte())
		val := frame.LocalVariableRef(index)
		frame.OperandStack.Push(val)
		return 2 + frame.Thread.PC, nil
	},
	ILOAD_0: func(frame *rtda.Frame) (int, error) {
		val := frame.LocalVariableInt(0)
		frame.OperandStack.PushInt(val)
		return 1 + frame.Thread.PC, nil
	},
	ILOAD_1: func(frame *rtda.Frame) (int, error) {
		val := frame.LocalVariableInt(1)
		frame.OperandStack.PushInt(val)
		return 1 + frame.Thread.PC, nil
	},
	ILOAD_2: func(frame *rtda.Frame) (int, error) {
		val := frame.LocalVariableInt(2)
		frame.OperandStack.PushInt(val)
		return 1 + frame.Thread.PC, nil
	},
	ILOAD_3: func(frame *rtda.Frame) (int, error) {
		val := frame.LocalVariableInt(3)
		frame.OperandStack.PushInt(val)
		return 1 + frame.Thread.PC, nil
	},
	LLOAD_0: func(frame *rtda.Frame) (int, error) {
		val := frame.LocalVariableLong(0)
		frame.OperandStack.PushLong(val)
		return 1 + frame.Thread.PC, nil
	},
	LLOAD_1: func(frame *rtda.Frame) (int, error) {
		val := frame.LocalVariableLong(1)
		frame.OperandStack.PushLong(val)
		return 1 + frame.Thread.PC, nil
	},
	LLOAD_2: func(frame *rtda.Frame) (int, error) {
		val := frame.LocalVariableLong(2)
		frame.OperandStack.PushLong(val)
		return 1 + frame.Thread.PC, nil
	},
	LLOAD_3: func(frame *rtda.Frame) (int, error) {
		val := frame.LocalVariableLong(3)
		frame.OperandStack.PushLong(val)
		return 1 + frame.Thread.PC, nil
	},
	FLOAD_0: func(frame *rtda.Frame) (int, error) {
		val := frame.LocalVariableFloat(0)
		frame.OperandStack.PushFloat(val)
		return 1 + frame.Thread.PC, nil
	},
	FLOAD_1: func(frame *rtda.Frame) (int, error) {
		val := frame.LocalVariableFloat(1)
		frame.OperandStack.PushFloat(val)
		return 1 + frame.Thread.PC, nil
	},
	FLOAD_2: func(frame *rtda.Frame) (int, error) {
		val := frame.LocalVariableFloat(2)
		frame.OperandStack.PushFloat(val)
		return 1 + frame.Thread.PC, nil
	},
	FLOAD_3: func(frame *rtda.Frame) (int, error) {
		val := frame.LocalVariableFloat(3)
		frame.OperandStack.PushFloat(val)
		return 1 + frame.Thread.PC, nil
	},
	DLOAD_0: func(frame *rtda.Frame) (int, error) {
		val := frame.LocalVariableDouble(0)
		frame.OperandStack.PushDouble(val)
		return 1 + frame.Thread.PC, nil
	},
	DLOAD_1: func(frame *rtda.Frame) (int, error) {
		val := frame.LocalVariableDouble(1)
		frame.OperandStack.PushDouble(val)
		return 1 + frame.Thread.PC, nil
	},
	DLOAD_2: func(frame *rtda.Frame) (int, error) {
		val := frame.LocalVariableDouble(2)
		frame.OperandStack.PushDouble(val)
		return 1 + frame.Thread.PC, nil
	},
	DLOAD_3: func(frame *rtda.Frame) (int, error) {
		val := frame.LocalVariableDouble(3)
		frame.OperandStack.PushDouble(val)
		return 1 + frame.Thread.PC, nil
	},
	ALOAD_0: func(frame *rtda.Frame) (int, error) {
		val := frame.LocalVariableRef(0)
		frame.OperandStack.Push(val)
		return 1 + frame.Thread.PC, nil
	},
	ALOAD_1: func(frame *rtda.Frame) (int, error) {
		val := frame.LocalVariableRef(1)
		frame.OperandStack.Push(val)
		return 1 + frame.Thread.PC, nil
	},
	ALOAD_2: func(frame *rtda.Frame) (int, error) {
		val := frame.LocalVariableRef(2)
		frame.OperandStack.Push(val)
		return 1 + frame.Thread.PC, nil
	},
	ALOAD_3: func(frame *rtda.Frame) (int, error) {
		val := frame.LocalVariableRef(3)
		frame.OperandStack.Push(val)
		return 1 + frame.Thread.PC, nil
	},
	IALOAD: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		index := stack.PopInt()
		arr := stack.PopIntArr()
		if arr == nil {
			return 0, errors.New("java.lang.NullPointerException")
		}
		checkIndex(len(arr), index)
		stack.PushInt(arr[index])
		return 1 + frame.Thread.PC, nil
	},
	LALOAD: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		index := stack.PopInt()
		arr := stack.PopLongArr()
		if arr == nil {
			return 0, errors.New("java.lang.NullPointerException")
		}
		checkIndex(len(arr), index)
		stack.PushLong(arr[index])
		return 1 + frame.Thread.PC, nil
	},
	FALOAD: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		index := stack.PopInt()
		arr := stack.PopFloatArr()
		if arr == nil {
			return 0, errors.New("java.lang.NullPointerException")
		}
		checkIndex(len(arr), index)
		stack.PushFloat(arr[index])
		return 1 + frame.Thread.PC, nil
	},
	DALOAD: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		index := stack.PopInt()
		arr := stack.PopDoubleArr()
		if arr == nil {
			return 0, errors.New("java.lang.NullPointerException")
		}
		checkIndex(len(arr), index)
		stack.PushDouble(arr[index])
		return 1 + frame.Thread.PC, nil
	},
	AALOAD: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		index := stack.PopInt()
		arr := stack.PopRefArr()
		if arr == nil {
			return 0, errors.New("java.lang.NullPointerException")
		}
		checkIndex(len(arr), index)
		stack.Push(arr[index])
		return 1 + frame.Thread.PC, nil
	},
	BALOAD: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		index := stack.PopInt()
		arr := stack.PopByteArr()
		if arr == nil {
			return 0, errors.New("java.lang.NullPointerException")
		}
		checkIndex(len(arr), index)
		stack.PushInt(int32(arr[index]))
		return 1 + frame.Thread.PC, nil
	},
	CALOAD: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		index := stack.PopInt()
		arr := stack.PopCharArr()
		if arr == nil {
			return 0, errors.New("java.lang.NullPointerException")
		}
		checkIndex(len(arr), index)
		stack.PushInt(int32(arr[index]))
		return 1 + frame.Thread.PC, nil
	},
	SALOAD: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		index := stack.PopInt()
		arr := stack.PopShortArr()
		if arr == nil {
			return 0, errors.New("java.lang.NullPointerException")
		}
		checkIndex(len(arr), index)
		stack.PushInt(int32(arr[index]))
		return 1 + frame.Thread.PC, nil
	},
	ISTORE: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopInt()
		index := uint16(frame.NextByte())
		frame.SetLocalVariableInt(index, val)
		return 2 + frame.Thread.PC, nil
	},
	LSTORE: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopLong()
		index := uint16(frame.NextByte())
		frame.SetLocalVariableLong(index, val)
		return 2 + frame.Thread.PC, nil
	},
	FSTORE: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopFloat()
		index := uint16(frame.NextByte())
		frame.SetLocalVariableFloat(index, val)
		return 2 + frame.Thread.PC, nil
	},
	DSTORE: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopDouble()
		index := uint16(frame.NextByte())
		frame.SetLocalVariableDouble(index, val)
		return 2 + frame.Thread.PC, nil
	},
	ASTORE: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.Pop()
		index := uint16(frame.NextByte())
		frame.SetLocalVariableRef(index, val)
		return 2 + frame.Thread.PC, nil
	},
	ISTORE_0: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopInt()
		frame.SetLocalVariableInt(0, val)
		return 1 + frame.Thread.PC, nil
	},
	ISTORE_1: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopInt()
		frame.SetLocalVariableInt(1, val)
		return 1 + frame.Thread.PC, nil
	},
	ISTORE_2: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopInt()
		frame.SetLocalVariableInt(2, val)
		return 1 + frame.Thread.PC, nil
	},
	ISTORE_3: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopInt()
		frame.SetLocalVariableInt(3, val)
		return 1 + frame.Thread.PC, nil
	},
	LSTORE_0: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopLong()
		frame.SetLocalVariableLong(0, val)
		return 1 + frame.Thread.PC, nil
	},
	LSTORE_1: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopLong()
		frame.SetLocalVariableLong(1, val)
		return 1 + frame.Thread.PC, nil
	},
	LSTORE_2: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopLong()
		frame.SetLocalVariableLong(2, val)
		return 1 + frame.Thread.PC, nil
	},
	LSTORE_3: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopLong()
		frame.SetLocalVariableLong(3, val)
		return 1 + frame.Thread.PC, nil
	},
	FSTORE_0: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopFloat()
		frame.SetLocalVariableFloat(0, val)
		return 1 + frame.Thread.PC, nil
	},
	FSTORE_1: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopFloat()
		frame.SetLocalVariableFloat(1, val)
		return 1 + frame.Thread.PC, nil
	},
	FSTORE_2: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopFloat()
		frame.SetLocalVariableFloat(2, val)
		return 1 + frame.Thread.PC, nil
	},
	FSTORE_3: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopFloat()
		frame.SetLocalVariableFloat(3, val)
		return 1 + frame.Thread.PC, nil
	},
	DSTORE_0: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopDouble()
		frame.SetLocalVariableDouble(0, val)
		return 1 + frame.Thread.PC, nil
	},
	DSTORE_1: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopDouble()
		frame.SetLocalVariableDouble(1, val)
		return 1 + frame.Thread.PC, nil
	},
	DSTORE_2: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopDouble()
		frame.SetLocalVariableDouble(2, val)
		return 1 + frame.Thread.PC, nil
	},
	DSTORE_3: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopDouble()
		frame.SetLocalVariableDouble(3, val)
		return 1 + frame.Thread.PC, nil
	},
	ASTORE_0: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.Pop()
		frame.SetLocalVariableRef(0, val)
		return 1 + frame.Thread.PC, nil
	},
	ASTORE_1: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.Pop()
		frame.SetLocalVariableRef(1, val)
		return 1 + frame.Thread.PC, nil
	},
	ASTORE_2: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.Pop()
		frame.SetLocalVariableRef(2, val)
		return 1 + frame.Thread.PC, nil
	},
	ASTORE_3: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.Pop()
		frame.SetLocalVariableRef(3, val)
		return 1 + frame.Thread.PC, nil
	},
	IASTORE: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopInt()
		index := frame.OperandStack.PopInt()
		ints := frame.OperandStack.PopIntArr()
		if ints == nil {
			return 0, errors.New("java.lang.NullPointerException")
		}
		if index < 0 || index >= int32(len(ints)) {
			return 0, errors.New("ArrayIndexOutOfBoundsException")
		}
		ints[index] = val
		return 1 + frame.Thread.PC, nil
	},
	LASTORE: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopLong()
		index := frame.OperandStack.PopInt()
		longs := frame.OperandStack.PopLongArr()
		if longs == nil {
			return 0, errors.New("java.lang.NullPointerException")
		}
		if index < 0 || index >= int32(len(longs)) {
			return 0, errors.New("ArrayIndexOutOfBoundsException")
		}
		longs[index] = val
		return 1 + frame.Thread.PC, nil
	},
	FASTORE: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopFloat()
		index := frame.OperandStack.PopInt()
		floats := frame.OperandStack.PopFloatArr()
		if floats == nil {
			return 0, errors.New("java.lang.NullPointerException")
		}
		if index < 0 || index >= int32(len(floats)) {
			return 0, errors.New("ArrayIndexOutOfBoundsException")
		}
		floats[index] = val
		return 1 + frame.Thread.PC, nil
	},
	DASTORE: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopDouble()
		index := frame.OperandStack.PopInt()
		doubles := frame.OperandStack.PopDoubleArr()
		if doubles == nil {
			return 0, errors.New("java.lang.NullPointerException")
		}
		if index < 0 || index >= int32(len(doubles)) {
			return 0, errors.New("ArrayIndexOutOfBoundsException")
		}
		doubles[index] = val
		return 1 + frame.Thread.PC, nil
	},
	AASTORE: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.Pop()
		index := frame.OperandStack.PopInt()
		refs := frame.OperandStack.PopRefArr()
		if refs == nil {
			return 0, errors.New("java.lang.NullPointerException")
		}
		if index < 0 || index >= int32(len(refs)) {
			return 0, errors.New("ArrayIndexOutOfBoundsException")
		}
		refs[index] = val
		return 1 + frame.Thread.PC, nil
	},
	BASTORE: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopInt()
		index := frame.OperandStack.PopInt()
		bytes := frame.OperandStack.PopByteArr()
		if bytes == nil {
			return 0, errors.New("java.lang.NullPointerException")
		}
		if index < 0 || index >= int32(len(bytes)) {
			return 0, errors.New("ArrayIndexOutOfBoundsException")
		}
		bytes[index] = int8(val)
		return 1 + frame.Thread.PC, nil
	},
	CASTORE: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopInt()
		index := frame.OperandStack.PopInt()
		chars := frame.OperandStack.PopCharArr()
		if chars == nil {
			return 0, errors.New("java.lang.NullPointerException")
		}
		if index < 0 || index >= int32(len(chars)) {
			return 0, errors.New("ArrayIndexOutOfBoundsException")
		}
		chars[index] = rune(val)
		return 1 + frame.Thread.PC, nil
	},
	SASTORE: func(frame *rtda.Frame) (int, error) {
		val := frame.OperandStack.PopInt()
		index := frame.OperandStack.PopInt()
		shorts := frame.OperandStack.PopShortArr()
		if shorts == nil {
			return 0, errors.New("java.lang.NullPointerException")
		}
		if index < 0 || index >= int32(len(shorts)) {
			return 0, errors.New("ArrayIndexOutOfBoundsException")
		}
		shorts[index] = int16(val)
		return 1 + frame.Thread.PC, nil
	},
	POP: func(frame *rtda.Frame) (int, error) {
		frame.OperandStack.Pop()
		return 1 + frame.Thread.PC, nil
	},
	POP2: func(frame *rtda.Frame) (int, error) {
		frame.OperandStack.Pop()
		frame.OperandStack.Pop()
		return 1 + frame.Thread.PC, nil
	},
	DUP: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.Pop()
		stack.Push(val)
		stack.Push(val)
		return 1 + frame.Thread.PC, nil
	},
	DUP_X1: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.Pop()
		val2 := stack.Pop()
		stack.Push(val1)
		stack.Push(val2)
		stack.Push(val1)
		return 1 + frame.Thread.PC, nil
	},
	DUP_X2: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.Pop()
		val2 := stack.Pop()
		val3 := stack.Pop()
		stack.Push(val1)
		stack.Push(val3)
		stack.Push(val2)
		stack.Push(val1)
		return 1 + frame.Thread.PC, nil
	},
	DUP2: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.Pop()
		val2 := stack.Pop()
		stack.Push(val2)
		stack.Push(val1)
		stack.Push(val2)
		stack.Push(val1)
		return 1 + frame.Thread.PC, nil
	},
	DUP2_X1: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.Pop()
		val2 := stack.Pop()
		val3 := stack.Pop()
		stack.Push(val2)
		stack.Push(val1)
		stack.Push(val3)
		stack.Push(val2)
		stack.Push(val1)
		return 1 + frame.Thread.PC, nil
	},
	DUP2_X2: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.Pop()
		val2 := stack.Pop()
		val3 := stack.Pop()
		val4 := stack.Pop()
		stack.Push(val2)
		stack.Push(val1)
		stack.Push(val4)
		stack.Push(val3)
		stack.Push(val2)
		stack.Push(val1)
		return 1 + frame.Thread.PC, nil
	},
	SWAP: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.Pop()
		val2 := stack.Pop()
		stack.Push(val1)
		stack.Push(val2)
		return 1 + frame.Thread.PC, nil
	},
	IADD: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopInt()
		val2 := stack.PopInt()
		stack.PushInt(val1 + val2)
		return 1 + frame.Thread.PC, nil
	},
	LADD: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopLong()
		val2 := stack.PopLong()
		stack.PushLong(val1 + val2)
		return 1 + frame.Thread.PC, nil
	},
	FADD: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopFloat()
		val2 := stack.PopFloat()
		stack.PushFloat(val1 + val2)
		return 1 + frame.Thread.PC, nil
	},
	DADD: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopDouble()
		val2 := stack.PopDouble()
		stack.PushDouble(val1 + val2)
		return 1 + frame.Thread.PC, nil
	},
	ISUB: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopInt()
		val2 := stack.PopInt()
		stack.PushInt(val2 - val1)
		return 1 + frame.Thread.PC, nil
	},
	LSUB: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopLong()
		val2 := stack.PopLong()
		stack.PushLong(val2 - val1)
		return 1 + frame.Thread.PC, nil
	},
	FSUB: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopFloat()
		val2 := stack.PopFloat()
		stack.PushFloat(val2 - val1)
		return 1 + frame.Thread.PC, nil
	},
	DSUB: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopDouble()
		val2 := stack.PopDouble()
		stack.PushDouble(val2 - val1)
		return 1 + frame.Thread.PC, nil
	},
	IMUL: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopInt()
		val2 := stack.PopInt()
		stack.PushInt(val1 * val2)
		return 1 + frame.Thread.PC, nil
	},
	LMUL: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopLong()
		val2 := stack.PopLong()
		stack.PushLong(val1 * val2)
		return 1 + frame.Thread.PC, nil
	},
	FMUL: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopFloat()
		val2 := stack.PopFloat()
		stack.PushFloat(val1 * val2)
		return 1 + frame.Thread.PC, nil
	},
	DMUL: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopDouble()
		val2 := stack.PopDouble()
		stack.PushDouble(val1 * val2)
		return 1 + frame.Thread.PC, nil
	},
	IDIV: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopInt()
		val2 := stack.PopInt()
		if val1 == 0 {
			return 0, errors.New("java.lang.ArithmeticException: / by zero")
		}
		stack.PushInt(val2 / val1)
		return 1 + frame.Thread.PC, nil
	},
	LDIV: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopLong()
		val2 := stack.PopLong()
		if val1 == 0 {
			return 0, errors.New("java.lang.ArithmeticException: / by zero")
		}
		stack.PushLong(val2 / val1)
		return 1 + frame.Thread.PC, nil
	},
	FDIV: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopFloat()
		val2 := stack.PopFloat()
		stack.PushFloat(val2 / val1)
		return 1 + frame.Thread.PC, nil
	},
	DDIV: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopDouble()
		val2 := stack.PopDouble()
		stack.PushDouble(val2 / val1)
		return 1 + frame.Thread.PC, nil
	},
	IREM: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopInt()
		val2 := stack.PopInt()
		if val1 == 0 {
			return 0, errors.New("java.lang.ArithmeticException: / by zero")
		}
		stack.PushInt(val2 % val1)
		return 1 + frame.Thread.PC, nil
	},
	LREM: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopLong()
		val2 := stack.PopLong()
		if val1 == 0 {
			return 0, errors.New("java.lang.ArithmeticException: / by zero")
		}
		stack.PushLong(val2 % val1)
		return 1 + frame.Thread.PC, nil
	},
	FREM: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopFloat()
		val2 := stack.PopFloat()
		stack.PushFloat(float32(math.Mod(float64(val2), float64(val1))))
		return 1 + frame.Thread.PC, nil
	},
	DREM: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopDouble()
		val2 := stack.PopDouble()
		stack.PushDouble(math.Mod(val2, val1))
		return 1 + frame.Thread.PC, nil
	},
	INEG: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopInt()
		stack.PushInt(-val)
		return 1 + frame.Thread.PC, nil
	},
	LNEG: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopLong()
		stack.PushLong(-val)
		return 1 + frame.Thread.PC, nil
	},
	FNEG: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopFloat()
		stack.PushFloat(-val)
		return 1 + frame.Thread.PC, nil
	},
	DNEG: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopDouble()
		stack.PushDouble(-val)
		return 1 + frame.Thread.PC, nil
	},
	ISHL: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		s := stack.PopInt()
		val := stack.PopInt()
		stack.PushInt(val << uint(s))
		return 1 + frame.Thread.PC, nil
	},
	LSHL: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		s := stack.PopInt()
		val := stack.PopLong()
		stack.PushLong(val << uint(s))
		return 1 + frame.Thread.PC, nil
	},
	ISHR: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		s := stack.PopInt()
		val := stack.PopInt()
		stack.PushInt(val >> uint(s))
		return 1 + frame.Thread.PC, nil
	},
	LSHR: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		s := stack.PopInt()
		val := stack.PopLong()
		stack.PushLong(val >> uint(s))
		return 1 + frame.Thread.PC, nil
	},
	IUSHR: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		s := stack.PopInt()
		val := stack.PopInt()
		stack.PushInt(int32(uint32(val) >> uint(s)))
		return 1 + frame.Thread.PC, nil
	},
	LUSHR: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		s := stack.PopInt()
		val := stack.PopLong()
		stack.PushLong(int64(uint64(val) >> uint(s)))
		return 1 + frame.Thread.PC, nil
	},
	IAND: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopInt()
		val2 := stack.PopInt()
		stack.PushInt(val1 & val2)
		return 1 + frame.Thread.PC, nil
	},
	LAND: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopLong()
		val2 := stack.PopLong()
		stack.PushLong(val1 & val2)
		return 1 + frame.Thread.PC, nil
	},
	IOR: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopInt()
		val2 := stack.PopInt()
		stack.PushInt(val1 | val2)
		return 1 + frame.Thread.PC, nil
	},
	LOR: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopLong()
		val2 := stack.PopLong()
		stack.PushLong(val1 | val2)
		return 1 + frame.Thread.PC, nil
	},
	IXOR: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopInt()
		val2 := stack.PopInt()
		stack.PushInt(val1 ^ val2)
		return 1 + frame.Thread.PC, nil
	},
	LXOR: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val1 := stack.PopLong()
		val2 := stack.PopLong()
		stack.PushLong(val1 ^ val2)
		return 1 + frame.Thread.PC, nil
	},
	I2L: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopInt()
		stack.PushLong(int64(val))
		return 1 + frame.Thread.PC, nil
	},
	I2F: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopInt()
		stack.PushFloat(float32(val))
		return 1 + frame.Thread.PC, nil
	},
	I2D: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopInt()
		stack.PushDouble(float64(val))
		return 1 + frame.Thread.PC, nil
	},
	L2I: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopLong()
		stack.PushInt(int32(val))
		return 1 + frame.Thread.PC, nil
	},
	L2F: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopLong()
		stack.PushFloat(float32(val))
		return 1 + frame.Thread.PC, nil
	},
	L2D: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopLong()
		stack.PushDouble(float64(val))
		return 1 + frame.Thread.PC, nil
	},
	F2I: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopFloat()
		stack.PushInt(int32(val))
		return 1 + frame.Thread.PC, nil
	},
	F2L: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopFloat()
		stack.PushLong(int64(val))
		return 1 + frame.Thread.PC, nil
	},
	F2D: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopFloat()
		stack.PushDouble(float64(val))
		return 1 + frame.Thread.PC, nil
	},
	D2I: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopDouble()
		stack.PushInt(int32(val))
		return 1 + frame.Thread.PC, nil
	},
	D2L: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopDouble()
		stack.PushLong(int64(val))
		return 1 + frame.Thread.PC, nil
	},
	D2F: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopDouble()
		stack.PushFloat(float32(val))
		return 1 + frame.Thread.PC, nil
	},
	I2B: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopInt()
		stack.PushInt(int32(int8(val)))
		return 1 + frame.Thread.PC, nil
	},
	I2C: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopInt()
		stack.PushInt(int32(uint16(val)))
		return 1 + frame.Thread.PC, nil
	},
	I2S: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopInt()
		stack.PushInt(int32(int16(val)))
		return 1 + frame.Thread.PC, nil
	},
	LCMP: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val2 := stack.PopLong()
		val1 := stack.PopLong()
		if val1 > val2 {
			stack.PushInt(1)
		} else if val1 == val2 {
			stack.PushInt(0)
		} else {
			stack.PushInt(-1)
		}
		return 1 + frame.Thread.PC, nil
	},
	FCMPL: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val2 := stack.PopFloat()
		val1 := stack.PopFloat()
		if val1 > val2 {
			stack.PushInt(1)
		} else if val1 == val2 {
			stack.PushInt(0)
		} else if val1 < val2 || math.IsNaN(float64(val1)) || math.IsNaN(float64(val2)) {
			stack.PushInt(-1)
		}
		return 1 + frame.Thread.PC, nil
	},
	FCMPG: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val2 := stack.PopFloat()
		val1 := stack.PopFloat()
		if val1 > val2 {
			stack.PushInt(1)
		} else if val1 == val2 {
			stack.PushInt(0)
		} else if val1 < val2 || math.IsNaN(float64(val1)) || math.IsNaN(float64(val2)) {
			stack.PushInt(-1)
		}
		return 1 + frame.Thread.PC, nil
	},
	DCMPL: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val2 := stack.PopDouble()
		val1 := stack.PopDouble()
		if val1 > val2 {
			stack.PushInt(1)
		} else if val1 == val2 {
			stack.PushInt(0)
		} else if val1 < val2 || math.IsNaN(val1) || math.IsNaN(val2) {
			stack.PushInt(-1)
		}
		return 1 + frame.Thread.PC, nil
	},
	DCMPG: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val2 := stack.PopDouble()
		val1 := stack.PopDouble()
		if val1 > val2 {
			stack.PushInt(1)
		} else if val1 == val2 {
			stack.PushInt(0)
		} else if val1 < val2 || math.IsNaN(val1) || math.IsNaN(val2) {
			stack.PushInt(-1)
		}
		return 1 + frame.Thread.PC, nil
	},
	IFEQ: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopInt()
		if val == 0 {
			offset := int(frame.NextByte())
			return frame.Thread.PC + int(offset), nil
		}
		return 1 + frame.Thread.PC, nil
	},
	IFNE: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopInt()
		if val != 0 {
			offset := int(frame.NextByte())
			return frame.Thread.PC + int(offset), nil
		}
		return 1 + frame.Thread.PC, nil
	},
	IFLT: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopInt()
		if val < 0 {
			offset := int(frame.ReadOffset())
			return frame.Thread.PC + int(offset), nil
		}
		return 1 + frame.Thread.PC, nil
	},
	IFGE: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopInt()
		if val >= 0 {
			offset := int(frame.ReadOffset())
			return frame.Thread.PC + int(offset), nil
		}
		return 1 + frame.Thread.PC, nil
	},
	IFGT: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopInt()
		if val > 0 {
			offset := int(frame.ReadOffset())
			return frame.Thread.PC + int(offset), nil
		}
		return 1 + frame.Thread.PC, nil
	},
	IFLE: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val := stack.PopInt()
		if val <= 0 {
			offset := int(frame.ReadOffset())
			return frame.Thread.PC + int(offset), nil
		}
		return 1 + frame.Thread.PC, nil
	},
	IF_ICMPEQ: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val2 := stack.PopInt()
		val1 := stack.PopInt()
		if val1 == val2 {
			offset := int(frame.ReadOffset())
			return frame.Thread.PC + int(offset), nil
		}
		return 1 + frame.Thread.PC, nil
	},
	IF_ICMPNE: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val2 := stack.PopInt()
		val1 := stack.PopInt()
		if val1 != val2 {
			offset := int(frame.ReadOffset())
			return frame.Thread.PC + int(offset), nil
		}
		return 1 + frame.Thread.PC, nil
	},
	IF_ICMPLT: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val2 := stack.PopInt()
		val1 := stack.PopInt()
		if val1 < val2 {
			offset := int(frame.ReadOffset())
			return frame.Thread.PC + int(offset), nil
		}
		return 1 + frame.Thread.PC, nil
	},
	IF_ICMPGE: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val2 := stack.PopInt()
		val1 := stack.PopInt()
		if val1 >= val2 {
			offset := int(frame.ReadOffset())
			return frame.Thread.PC + int(offset), nil
		}
		return 1 + frame.Thread.PC, nil
	},
	IF_ICMPGT: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val2 := stack.PopInt()
		val1 := stack.PopInt()
		if val1 > val2 {
			offset := int(frame.ReadOffset())
			return frame.Thread.PC + int(offset), nil
		}
		return 1 + frame.Thread.PC, nil
	},
	IF_ICMPLE: func(frame *rtda.Frame) (int, error) {
		stack := frame.OperandStack
		val2 := stack.PopInt()
		val1 := stack.PopInt()
		if val1 <= val2 {
			offset := int(frame.ReadOffset())
			return frame.Thread.PC + int(offset), nil
		}
		return 1 + frame.Thread.PC, nil
	},
	IF_ACMPEQ: func(frame *rtda.Frame) (int, error) {
		panic("todo: if_acmpeq")
	},
	IF_ACMPNE: func(frame *rtda.Frame) (int, error) {
		panic("todo: if_acmpne")
	},
	GOTO: func(frame *rtda.Frame) (int, error) {
		offset := int(frame.ReadOffset())
		return frame.Thread.PC + int(offset), nil
	},
	JSR: func(frame *rtda.Frame) (int, error) {
		panic("todo: jsr")
	},
	RET: func(frame *rtda.Frame) (int, error) {
		panic("todo: ret")
	},
	TABLESWITCH: func(frame *rtda.Frame) (int, error) {
		panic("todo: tableswitch")
	},
	LOOKUPSWITCH: func(frame *rtda.Frame) (int, error) {
		panic("todo: lookupswitch")
	},
	IRETURN: func(frame *rtda.Frame) (int, error) {
		thread := frame.Thread
		currentFrame := thread.PopFrame()
		invokerFrame := thread.TopFrame()
		val := currentFrame.OperandStack.PopInt()
		invokerFrame.OperandStack.PushInt(val)
		return 0, nil
	},
	LRETURN: func(frame *rtda.Frame) (int, error) {
		thread := frame.Thread
		currentFrame := thread.PopFrame()
		invokerFrame := thread.TopFrame()
		val := currentFrame.OperandStack.PopLong()
		invokerFrame.OperandStack.PushLong(val)
		return 0, nil
	},
	FRETURN: func(frame *rtda.Frame) (int, error) {
		thread := frame.Thread
		currentFrame := thread.PopFrame()
		invokerFrame := thread.TopFrame()
		val := currentFrame.OperandStack.PopFloat()
		invokerFrame.OperandStack.PushFloat(val)
		return 0, nil
	},
	DRETURN: func(frame *rtda.Frame) (int, error) {
		thread := frame.Thread
		currentFrame := thread.PopFrame()
		invokerFrame := thread.TopFrame()
		val := currentFrame.OperandStack.PopDouble()
		invokerFrame.OperandStack.PushDouble(val)
		return 0, nil
	},
	ARETURN: func(frame *rtda.Frame) (int, error) {
		thread := frame.Thread
		currentFrame := thread.PopFrame()
		invokerFrame := thread.TopFrame()
		val := currentFrame.OperandStack.Pop()
		invokerFrame.OperandStack.Push(val)
		return 0, nil
	},
	RETURN: func(frame *rtda.Frame) (int, error) {
		thread := frame.Thread
		thread.PopFrame()
		return 0, nil
	},
	GETSTATIC: func(frame *rtda.Frame) (int, error) {
		panic("todo: getstatic")
	},
	PUTSTATIC: func(frame *rtda.Frame) (int, error) {
		panic("todo: putstatic")
	},
	GETFIELD: func(frame *rtda.Frame) (int, error) {
		panic("todo: getfield")
	},
	PUTFIELD: func(frame *rtda.Frame) (int, error) {
		panic("todo: putfield")
	},
	INVOKEVIRTUAL: func(frame *rtda.Frame) (int, error) {
		panic("todo: invokevirtual")
	},
	INVOKESPECIAL: func(frame *rtda.Frame) (int, error) {
		panic("todo: invokespecial")
	},
	INVOKESTATIC: func(frame *rtda.Frame) (int, error) {
		panic("todo: invokestatic")
	},
	INVOKEINTERFACE: func(frame *rtda.Frame) (int, error) {
		panic("todo: invokeinterface")
	},
	INVOKEDYNAMIC: func(frame *rtda.Frame) (int, error) {
		panic("todo: invokedynamic")
	},
	NEW: func(frame *rtda.Frame) (int, error) {
		panic("todo: new")
	},
	NEWARRAY: func(frame *rtda.Frame) (int, error) {
		panic("todo: newarray")
	},
	ANEWARRAY: func(frame *rtda.Frame) (int, error) {
		panic("todo: anewarray")
	},
	ARRAYLENGTH: func(frame *rtda.Frame) (int, error) {
		panic("todo: arraylength")
	},
	ATHROW: func(frame *rtda.Frame) (int, error) {
		panic("todo: athrow")
	},
	CHECKCAST: func(frame *rtda.Frame) (int, error) {
		panic("todo: checkcast")
	},
	INSTANCEOF: func(frame *rtda.Frame) (int, error) {
		panic("todo: instanceof")
	},
	MONITORENTER: func(frame *rtda.Frame) (int, error) {
		panic("todo: monitorenter")
	},
	MONITOREXIT: func(frame *rtda.Frame) (int, error) {
		panic("todo: monitorexit")
	},
	WIDE: func(frame *rtda.Frame) (int, error) {
		panic("todo: wide")
	},
	MULTIANEWARRAY: func(frame *rtda.Frame) (int, error) {
		panic("todo: multianewarray")
	},
	IFNULL: func(frame *rtda.Frame) (int, error) {
		panic("todo: ifnull")
	},
	IFNONNULL: func(frame *rtda.Frame) (int, error) {
		panic("todo: ifnonnull")
	},
	GOTO_W: func(frame *rtda.Frame) (int, error) {
		panic("todo: goto_w")
	},
	JSR_W: func(frame *rtda.Frame) (int, error) {
		panic("todo: jsr_w")
	},
}

func checkIndex(i int, index int32) {
	if i < 0 || i >= int(index) {
		panic("ArrayIndexOutOfBoundsException")
	}

}
