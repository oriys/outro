package interpret

import "outro/rtda"

type JVM struct {
	thread *rtda.Thread
}

func (jvm *JVM) Execute() {

	jvm.thread.Execute()

}

func NewJVM(thread *rtda.Thread) *JVM {
	return &JVM{
		thread: thread,
	}
}
