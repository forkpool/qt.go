package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin

type QStyleOptionGroupBox struct {
	*QStyleOptionComplex
}

func (this *QStyleOptionGroupBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOptionComplex.GetCthis()
	}
}
func (this *QStyleOptionGroupBox) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOptionComplex = NewQStyleOptionComplexFromPointer(cthis)
}
func NewQStyleOptionGroupBoxFromPointer(cthis unsafe.Pointer) *QStyleOptionGroupBox {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionGroupBox{bcthis0}
}
func (*QStyleOptionGroupBox) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionGroupBox {
	return NewQStyleOptionGroupBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:639
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionGroupBox()
func NewQStyleOptionGroupBox() *QStyleOptionGroupBox {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleOptionGroupBoxC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionGroupBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionGroupBox)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:642
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionGroupBox(int)
func NewQStyleOptionGroupBox_1(version int) *QStyleOptionGroupBox {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleOptionGroupBoxC2Ei", ffiqt.FFI_TYPE_POINTER, version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionGroupBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionGroupBox)
	return gothis
}

func DeleteQStyleOptionGroupBox(this *QStyleOptionGroupBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleOptionGroupBoxD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QStyleOptionGroupBox__StyleOptionType = int

const QStyleOptionGroupBox__Type QStyleOptionGroupBox__StyleOptionType = 983046

type QStyleOptionGroupBox__StyleOptionVersion = int

const QStyleOptionGroupBox__Version QStyleOptionGroupBox__StyleOptionVersion = 1

//  body block end