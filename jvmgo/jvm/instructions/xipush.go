package instructions

import (
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
)

// Push byte
type Bipush struct {
	Val int8
}

func (self *Bipush) fetchOperands(decoder *InstructionDecoder) {
	self.Val = decoder.readInt8()
}
func (self *Bipush) Execute(frame *rtda.Frame) {
	i := int32(self.Val)
	frame.OperandStack().PushInt(i)
}

func (self *Bipush) ByteSize() int {
	return 2
}

// Push short
type Sipush struct {
	val int16
}

func (self *Sipush) fetchOperands(decoder *InstructionDecoder) {
	self.val = decoder.readInt16()
}
func (self *Sipush) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

func (self *Sipush) ByteSize() int {
	return 3
}