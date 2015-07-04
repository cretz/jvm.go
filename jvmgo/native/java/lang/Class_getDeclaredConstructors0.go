package lang

import (
	. "github.com/cretz/jvm.go/jvmgo/any"
	"github.com/cretz/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/cretz/jvm.go/jvmgo/jvm/rtda/class"
)

/*
Constructor(Class<T> declaringClass,
            Class<?>[] parameterTypes,
            Class<?>[] checkedExceptions,
            int modifiers,
            int slot,
            String signature,
            byte[] annotations,
            byte[] parameterAnnotations)
}
*/
const _constructorConstructorDescriptor = "" +
	"(Ljava/lang/Class;" +
	"[Ljava/lang/Class;" +
	"[Ljava/lang/Class;" +
	"II" +
	"Ljava/lang/String;" +
	"[B[B)V"

func init() {
	_class(getDeclaredConstructors0, "getDeclaredConstructors0", "(Z)[Ljava/lang/reflect/Constructor;")
}

// private native Constructor<T>[] getDeclaredConstructors0(boolean publicOnly);
// (Z)[Ljava/lang/reflect/Constructor;
func getDeclaredConstructors0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	classObj := vars.GetThis()
	publicOnly := vars.GetBoolean(1)

	class := classObj.Extra().(*rtc.Class)
	constructors := class.GetConstructors(publicOnly)
	constructorCount := uint(len(constructors))

	constructorClass := rtc.BootLoader().LoadClass("java/lang/reflect/Constructor")
	constructorArr := constructorClass.NewArray(constructorCount)

	stack := frame.OperandStack()
	stack.PushRef(constructorArr)

	if constructorCount > 0 {
		thread := frame.Thread()
		constructorObjs := constructorArr.Refs()
		constructorInitMethod := constructorClass.GetConstructor(_constructorConstructorDescriptor)
		for i, constructor := range constructors {
			constructorObj := constructorClass.NewObjWithExtra(constructor)
			constructorObjs[i] = constructorObj

			// init constructorObj
			thread.InvokeMethodWithShim(constructorInitMethod, []Any{
				constructorObj, // this
				classObj,       // declaringClass
				getParameterTypeArr(constructor),    // parameterTypes
				getExceptionTypeArr(constructor),    // checkedExceptions
				int32(constructor.GetAccessFlags()), // modifiers
				int32(0), // todo slot
				getSignatureStr(constructor.Signature()),                    // signature
				getAnnotationByteArr(constructor.AnnotationData()),          // annotations
				getAnnotationByteArr(constructor.ParameterAnnotationData()), // parameterAnnotations
			})
		}
	}
}
