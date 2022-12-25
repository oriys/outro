package rtda

type ClassLoader interface {
	LoadClass(className string) (*Class, error)
}

type ApplicationClassLoader struct {
	classMap             map[string]*Class
	extensionClassLoader *ExtensionClassLoader
}

func (a *ApplicationClassLoader) LoadClass(className string) (*Class, error) {
	class := a.classMap[className]
	if class != nil {
		return class, nil
	}
	class, err := a.extensionClassLoader.LoadClass(className)
	if err != nil {
		return nil, err
	}
	a.classMap[className] = class
	return class, nil
}

type ExtensionClassLoader struct {
	classMap             map[string]*Class
	bootstrapClassLoader *BootstrapClassLoader
}

func (e ExtensionClassLoader) LoadClass(className string) (*Class, error) {
	class := e.classMap[className]
	if class != nil {
		return class, nil
	}
	class, err := e.bootstrapClassLoader.LoadClass(className)
	if err != nil {
		return nil, err
	}

	e.classMap[className] = class
	return class, nil
}

type BootstrapClassLoader struct {
	classMap map[string]*Class
}

func (b BootstrapClassLoader) LoadClass(className string) (*Class, error) {
	class := b.classMap[className]
	if class != nil {
		return class, nil
	}
	return nil, nil
}
