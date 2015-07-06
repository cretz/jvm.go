package instructions

import (
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
)

// Branch always
type Goto struct{ BranchInstruction }

func (self *Goto) Execute(frame *rtda.Frame) {
	branch(frame, self.Offset)
}

// Branch always (wide index)
type GotoW struct {
	offset int
}

func (self *GotoW) fetchOperands(decoder *InstructionDecoder) {
	self.offset = int(decoder.readInt32())
}
func (self *GotoW) Execute(frame *rtda.Frame) {
	branch(frame, self.offset)
}

func (self *GotoW) ByteSize() int {
	return 5
}
