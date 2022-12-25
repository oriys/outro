package rtda

import "outro/model"

type Thread struct {
	PC           int
	stack        []*Frame
	currentClass *model.Class
}

func (t *Thread) Execute() {
}

func (t *Thread) PushFrame(frame *Frame) {
	t.stack = append(t.stack, frame)
}

func (t *Thread) PopFrame() *Frame {
	frame := t.stack[len(t.stack)-1]
	t.stack = t.stack[:len(t.stack)-1]
	return frame
}

func (t *Thread) CurrentFrame() *Frame {
	return t.stack[len(t.stack)-1]
}

func (t *Thread) TopFrame() *Frame {
	return t.stack[len(t.stack)-1]
}

func NewThread(frame *Frame, class *model.Class) *Thread {
	return &Thread{
		currentFrame: frame,
		currentClass: class,
	}
}
