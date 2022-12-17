package model

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
