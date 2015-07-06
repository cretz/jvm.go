package instructions

import (
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
)

// Extend local variable index by additional bytes
type Wide struct {
	ModifiedInstruction Instruction
}

func (self *Wide) fetchOperands(decoder *InstructionDecoder) {
	opcode := decoder.readUint8()
	switch opcode {
	case 0x15:
		inst := &iload{}
		inst.Index = uint(decoder.readUint16())
		self.ModifiedInstruction = inst
	case 0x16:
		inst := &lload{}
		inst.Index = uint(decoder.readUint16())
		self.ModifiedInstruction = inst
	case 0x17:
		inst := &fload{}
		inst.Index = uint(decoder.readUint16())
		self.ModifiedInstruction = inst
	case 0x18:
		inst := &dload{}
		inst.Index = uint(decoder.readUint16())
		self.ModifiedInstruction = inst
	case 0x19:
		inst := &aload{}
		inst.Index = uint(decoder.readUint16())
		self.ModifiedInstruction = inst
	case 0x36:
		inst := &istore{}
		inst.Index = uint(decoder.readUint16())
		self.ModifiedInstruction = inst
	case 0x37:
		inst := &lstore{}
		inst.Index = uint(decoder.readUint16())
		self.ModifiedInstruction = inst
	case 0x38:
		inst := &fstore{}
		inst.Index = uint(decoder.readUint16())
		self.ModifiedInstruction = inst
	case 0x39:
		inst := &dstore{}
		inst.Index = uint(decoder.readUint16())
		self.ModifiedInstruction = inst
	case 0x3a:
		inst := &astore{}
		inst.Index = uint(decoder.readUint16())
		self.ModifiedInstruction = inst
	case 0xa9:
		inst := &ret{}
		inst.Index = uint(decoder.readUint16())
		self.ModifiedInstruction = inst
	case 0x84:
		inst := &Iinc{}
		inst.Index = uint(decoder.readUint16())
		inst.Const = int32(decoder.readInt16())
		self.ModifiedInstruction = inst
	}
}

func (self *Wide) Execute(frame *rtda.Frame) {
	self.ModifiedInstruction.Execute(frame)
}

func (self *Wide) ByteSize() int {
	return self.ByteSize() * 2 - 1
}
