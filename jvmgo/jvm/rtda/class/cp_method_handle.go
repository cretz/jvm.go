package class

import (
	cf "github.com/cretz/jvm.go/jvmgo/classfile"
)

type ConstantMethodHandle struct {
	referenceKind  uint8
	referenceIndex uint16
}

func newConstantMethodHandle(mhInfo *cf.ConstantMethodHandleInfo) *ConstantMethodHandle {
	return &ConstantMethodHandle{
		referenceKind:  mhInfo.ReferenceKind(),
		referenceIndex: mhInfo.ReferenceIndex(),
	}
}
