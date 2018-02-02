package qtwidgets

// /usr/include/qt/QtWidgets/qlayoutitem.h
// #include <qlayoutitem.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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

type QWidgetItem struct {
	*QLayoutItem
}

func (this *QWidgetItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QLayoutItem.GetCthis()
	}
}
func (this *QWidgetItem) SetCthis(cthis unsafe.Pointer) {
	this.QLayoutItem = NewQLayoutItemFromPointer(cthis)
}
func NewQWidgetItemFromPointer(cthis unsafe.Pointer) *QWidgetItem {
	bcthis0 := NewQLayoutItemFromPointer(cthis)
	return &QWidgetItem{bcthis0}
}
func (*QWidgetItem) NewFromPointer(cthis unsafe.Pointer) *QWidgetItem {
	return NewQWidgetItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:126
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QWidgetItem(QWidget *)
func NewQWidgetItem(w *QWidget /*777 QWidget **/) *QWidgetItem {
	var convArg0 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWidgetItemC2EP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:127
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWidgetItem()
func DeleteQWidgetItem(this *QWidgetItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWidgetItemD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:129
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QWidgetItem) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:130
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSize()
func (this *QWidgetItem) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem11minimumSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:131
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize maximumSize()
func (this *QWidgetItem) MaximumSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem11maximumSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:132
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::Orientations expandingDirections()
func (this *QWidgetItem) ExpandingDirections() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem19expandingDirectionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:133
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QWidgetItem) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:134
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRect &)
func (this *QWidgetItem) SetGeometry(arg0 *qtcore.QRect) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWidgetItem11setGeometryERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:135
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect geometry()
func (this *QWidgetItem) Geometry() *qtcore.QRect /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem8geometryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:136
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QWidget * widget()
func (this *QWidgetItem) Widget() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWidgetItem6widgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:138
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasHeightForWidth()
func (this *QWidgetItem) HasHeightForWidth() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem17hasHeightForWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:139
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int heightForWidth(int)
func (this *QWidgetItem) HeightForWidth(arg0 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem14heightForWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:140
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QSizePolicy::ControlTypes controlTypes()
func (this *QWidgetItem) ControlTypes() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWidgetItem12controlTypesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

//  body block end