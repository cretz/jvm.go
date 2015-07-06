package instructions

import (
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
)

// Return void from method
type Return struct{ NoOperandsInstruction }

func (self *Return) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	thread.PopFrame()
}

// Return reference from method
type areturn struct{ NoOperandsInstruction }

func (self *areturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	ref := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(ref)
}

// Return double from method
type dreturn struct{ NoOperandsInstruction }

func (self *dreturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(val)
}

// Return float from method
type freturn struct{ NoOperandsInstruction }

func (self *freturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(val)
}

// Return int from method
type ireturn struct{ NoOperandsInstruction }

func (self *ireturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(val)
}

// Return double from method
type lreturn struct{ NoOperandsInstruction }

func (self *lreturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(val)
}
