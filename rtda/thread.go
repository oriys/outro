package rtda

type Thread struct {
	PC           int
	stack        []*Frame
	currentClass *Class
	ThreadGroup  []*Thread
}

func NewThread() *Thread {
	thread := Thread{}
	thread.stack = make([]*Frame, 0)
	return &thread
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

func (t *Thread) NewFrame(method *Method) {
	frame := Frame{}
	frame.Method = method
	frame.Thread = t
	frame.localVariables = make([]interface{}, method.MaxLocals)
	frame.operandStack = make([]interface{}, method.MaxStack)
	t.PushFrame(&frame)
}
