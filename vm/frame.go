package vm

import (
	"monkey/code"
	"monkey/object"
)

type Frame struct {
	fn          *object.CompiledFunction
	ip          int
	basePointer int // 为了存本地变量，需要基址信息
}

func NewFrame(fn *object.CompiledFunction, basePointer int) *Frame {
	return &Frame{fn: fn, ip: -1, basePointer: basePointer}
}

func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}
