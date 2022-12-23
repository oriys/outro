package rtda

type ClassLoader interface {
	LoadClass(className string) *Class
	ReadClass(className string) ([]byte, error)
}

type ApplicationClassLoader struct {
	classMap             map[string]*Class
	extensionClassLoader *ExtensionClassLoader
}

type ExtensionClassLoader struct {
	classMap             map[string]*Class
	bootstrapClassLoader *BootstrapClassLoader
}

type BootstrapClassLoader struct {
	classMap map[string]*Class
}
