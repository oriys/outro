package rtda

type JString struct {
	chars string
	class *Class
}

func NewJString(classLoader *ApplicationClassLoader, str string) *JString {
	j := &JString{
		chars: str,
	}
	class, err := classLoader.LoadClass("java/lang/String")
	if err != nil {
		panic(err)
	}
	j.class = class
	return j
}
