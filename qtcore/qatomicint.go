package qtcore

// /usr/include/qt/QtCore/qatomic.h
// #include <qatomic.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin

type QAtomicInt struct {
	*qtrt.CObject
}

func (this *QAtomicInt) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAtomicInt) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAtomicIntFromPointer(cthis unsafe.Pointer) *QAtomicInt {
	return &QAtomicInt{&qtrt.CObject{cthis}}
}
func (*QAtomicInt) NewFromPointer(cthis unsafe.Pointer) *QAtomicInt {
	return NewQAtomicIntFromPointer(cthis)
}

// /usr/include/qt/QtCore/qatomic.h:162
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QAtomicInt(int)
func NewQAtomicInt(value int) *QAtomicInt {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QAtomicIntC2Ei", ffiqt.FFI_TYPE_POINTER, value)
	gopp.ErrPrint(err, rv)
	gothis := NewQAtomicIntFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAtomicInt)
	return gothis
}

func DeleteQAtomicInt(this *QAtomicInt) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QAtomicIntD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end