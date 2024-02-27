package vm

import (
	"go_compiler/code"
	"go_compiler/object"
)

type Frame struct {
	fn *object.CompiledFunction // Points to the compiled function referenced by the frame
	ip int                      // The instruction pointer in this frame
}

func NewFrame(fn *object.CompiledFunction) *Frame {
	return &Frame{fn: fn, ip: -1}
}

func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}

