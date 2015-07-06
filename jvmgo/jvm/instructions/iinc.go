package instructions

import (
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
)

// Increment local variable by constant
type Iinc struct {
	Index uint
	Const int32
}

func (self *Iinc) fetchOperands(decoder *InstructionDecoder) {
	self.Index = uint(decoder.readUint8())
	self.Const = int32(decoder.readInt8())
}

func (self *Iinc) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.SetInt(self.Index, val)
}

func (self *Iinc) ByteSize() int {
	return 3
}
