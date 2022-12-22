package rtda

import "outro/model"

type Thread struct {
	PC           int
	currentFrame *Frame
	currentClass *model.Class
}

func (t *Thread) Execute() {
}

func NewThread(frame *Frame, class *model.Class) *Thread {
	return &Thread{
		currentFrame: frame,
		currentClass: class,
	}
}
