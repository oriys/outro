package model

type Thread struct {
	PC           int
	currentFrame *Frame
	currentClass *Class
}
