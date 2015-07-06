package instructions

import (
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
)

type Instruction interface {
	fetchOperands(decoder *InstructionDecoder)
	Execute(frame *rtda.Frame)
	ByteSize() int
}

type NoOperandsInstruction struct {
	// empty
}

func (self *NoOperandsInstruction) fetchOperands(decoder *InstructionDecoder) {
	// nothing to do
}

func (self *NoOperandsInstruction) ByteSize() int {
	return 1
}

type BranchInstruction struct {
	Offset int // todo target
}

func (self *BranchInstruction) fetchOperands(decoder *InstructionDecoder) {
	self.Offset = int(decoder.readInt16())
}

func (self *BranchInstruction) ByteSize() int {
	return 3
}

type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) fetchOperands(decoder *InstructionDecoder) {
	self.Index = uint(decoder.readUint8())
}

func (self *Index8Instruction) ByteSize() int {
	return 2
}

type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) fetchOperands(decoder *InstructionDecoder) {
	self.Index = uint(decoder.readUint16())
}

func (self *Index16Instruction) ByteSize() int {
	return 3
}

func branch(frame *rtda.Frame, offset int) {
	nextPC := frame.Thread().PC() + offset
	frame.SetNextPC(nextPC)
}
