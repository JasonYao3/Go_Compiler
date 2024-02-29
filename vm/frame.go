package vm

import (
	"go_compiler/code"
	"go_compiler/object"
)

type Frame struct {
	fn          *object.CompiledFunction // Points to the compiled function referenced by the frame
	ip          int                      // The instruction pointer in this frame
	basePointer int                      // Points to the bottom of the stack of the current call frame. Frame pointer
}

func NewFrame(fn *object.CompiledFunction, basePointer int) *Frame {
	f := &Frame{
		fn:          fn,
		ip:          -1,
		basePointer: basePointer,
	}

	return f
}

func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}
