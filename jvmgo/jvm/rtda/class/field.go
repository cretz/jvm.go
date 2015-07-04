package class

import (
	. "github.com/cretz/jvm.go/jvmgo/any"
	cf "github.com/cretz/jvm.go/jvmgo/classfile"
)

type Field struct {
	ClassMember
	IsLongOrDouble  bool
	constValueIndex uint16
	slotId          uint
	_type           *Class
}

func newField(class *Class, fieldInfo *cf.MemberInfo) *Field {
	field := &Field{}
	field.class = class
	field.accessFlags = fieldInfo.AccessFlags()
	field.name = fieldInfo.Name()
	field.descriptor = fieldInfo.Descriptor()
	field.signature = fieldInfo.Signature()
	field.IsLongOrDouble = (field.descriptor == "J" || field.descriptor == "D")
	if kValAttr := fieldInfo.ConstantValueAttribute(); kValAttr != nil {
		field.constValueIndex = kValAttr.ConstantValueIndex()
	}
	return field
}

func (self *Field) ConstValueIndex() uint16 {
	return self.constValueIndex
}
func (self *Field) SlotId() uint {
	return self.slotId
}

func (self *Field) GetValue(ref *Obj) Any {
	fields := ref.fields.([]Any)
	return fields[self.slotId]
}
func (self *Field) PutValue(ref *Obj, val Any) {
	fields := ref.fields.([]Any)
	fields[self.slotId] = val
}

func (self *Field) GetStaticValue() Any {
	return self.class.staticFieldSlots[self.slotId]
}
func (self *Field) PutStaticValue(val Any) {
	self.class.staticFieldSlots[self.slotId] = val
}

func (self *Field) defaultValue() Any {
	switch self.descriptor[0] {
	case 'Z': // boolean
		return int32(0)
	case 'B': // byte
		return int32(0)
	case 'S': // short
		return int32(0)
	case 'C': // char
		return int32(0)
	case 'I': // int
		return int32(0)
	case 'J': // long
		return int64(0)
	case 'F': // float
		return float32(0)
	case 'D': // double
		return float64(0)
	case 'L': // Object
		return nil
	case '[': // Array
		return nil
	default:
		panic("BAD field descriptor: " + self.descriptor)
	}
}

// reflection
func (self *Field) Type() *Class {
	if self._type == nil {
		self._type = self.resolveType()
	}
	return self._type
}
func (self *Field) resolveType() *Class {
	descriptor := self.descriptor
	if len(descriptor) == 1 {
		switch descriptor[0] {
		case 'B':
			return bootLoader.GetPrimitiveClass("byte")
		case 'C':
			return bootLoader.GetPrimitiveClass("char")
		case 'D':
			return bootLoader.GetPrimitiveClass("double")
		case 'F':
			return bootLoader.GetPrimitiveClass("float")
		case 'I':
			return bootLoader.GetPrimitiveClass("int")
		case 'J':
			return bootLoader.GetPrimitiveClass("long")
		case 'S':
			return bootLoader.GetPrimitiveClass("short")
		case 'V':
			return bootLoader.GetPrimitiveClass("void")
		case 'Z':
			return bootLoader.GetPrimitiveClass("boolean")
		default:
			panic("BAD descriptor: " + descriptor)
		}
	}
	if descriptor[0] == 'L' {
		return bootLoader.LoadClass(descriptor[1 : len(descriptor)-1])
	}
	return bootLoader.LoadClass(descriptor)
}
