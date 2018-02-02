package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin

type QAccessibleValueInterface struct {
	*qtrt.CObject
}

func (this *QAccessibleValueInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAccessibleValueInterface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAccessibleValueInterfaceFromPointer(cthis unsafe.Pointer) *QAccessibleValueInterface {
	return &QAccessibleValueInterface{&qtrt.CObject{cthis}}
}
func (*QAccessibleValueInterface) NewFromPointer(cthis unsafe.Pointer) *QAccessibleValueInterface {
	return NewQAccessibleValueInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessible.h:566
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleValueInterface()
func DeleteQAccessibleValueInterface(this *QAccessibleValueInterface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QAccessibleValueInterfaceD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessible.h:568
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QVariant currentValue()
func (this *QAccessibleValueInterface) CurrentValue() *qtcore.QVariant /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleValueInterface12currentValueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtGui/qaccessible.h:569
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setCurrentValue(const QVariant &)
func (this *QAccessibleValueInterface) SetCurrentValue(value *qtcore.QVariant) {
	var convArg0 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QAccessibleValueInterface15setCurrentValueERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:570
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QVariant maximumValue()
func (this *QAccessibleValueInterface) MaximumValue() *qtcore.QVariant /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleValueInterface12maximumValueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtGui/qaccessible.h:571
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QVariant minimumValue()
func (this *QAccessibleValueInterface) MinimumValue() *qtcore.QVariant /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleValueInterface12minimumValueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtGui/qaccessible.h:572
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QVariant minimumStepSize()
func (this *QAccessibleValueInterface) MinimumStepSize() *qtcore.QVariant /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleValueInterface15minimumStepSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

//  body block end