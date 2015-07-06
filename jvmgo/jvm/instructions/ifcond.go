package instructions

import (
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
)

// Branch if int comparison with zero succeeds
type ifeq struct{ BranchInstruction }

func (self *ifeq) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		branch(frame, self.Offset)
	}
}

type ifne struct{ BranchInstruction }

func (self *ifne) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		branch(frame, self.Offset)
	}
}

type iflt struct{ BranchInstruction }

func (self *iflt) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		branch(frame, self.Offset)
	}
}

type ifle struct{ BranchInstruction }

func (self *ifle) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		branch(frame, self.Offset)
	}
}

type ifgt struct{ BranchInstruction }

func (self *ifgt) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		branch(frame, self.Offset)
	}
}

type ifge struct{ BranchInstruction }

func (self *ifge) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		branch(frame, self.Offset)
	}
}
