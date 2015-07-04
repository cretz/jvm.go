package instructions

import (
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
)

// Swap the top two operand stack values
type swap struct{ NoOperandsInstruction }

func (self *swap) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	stack.PushSlot(val1)
	stack.PushSlot(val2)
}
