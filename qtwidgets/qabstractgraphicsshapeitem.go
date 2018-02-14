package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 207
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

type QAbstractGraphicsShapeItem struct {
	*QGraphicsItem
}
type QAbstractGraphicsShapeItem_ITF interface {
	QGraphicsItem_ITF
	QAbstractGraphicsShapeItem_PTR() *QAbstractGraphicsShapeItem
}

func (ptr *QAbstractGraphicsShapeItem) QAbstractGraphicsShapeItem_PTR() *QAbstractGraphicsShapeItem {
	return ptr
}

func (this *QAbstractGraphicsShapeItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsItem.GetCthis()
	}
}
func (this *QAbstractGraphicsShapeItem) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsItem = NewQGraphicsItemFromPointer(cthis)
}
func NewQAbstractGraphicsShapeItemFromPointer(cthis unsafe.Pointer) *QAbstractGraphicsShapeItem {
	bcthis0 := NewQGraphicsItemFromPointer(cthis)
	return &QAbstractGraphicsShapeItem{bcthis0}
}
func (*QAbstractGraphicsShapeItem) NewFromPointer(cthis unsafe.Pointer) *QAbstractGraphicsShapeItem {
	return NewQAbstractGraphicsShapeItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:603
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractGraphicsShapeItem(QGraphicsItem *)
func NewQAbstractGraphicsShapeItem(parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QAbstractGraphicsShapeItem {
	var convArg0 = parent.QGraphicsItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAbstractGraphicsShapeItemC2EP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractGraphicsShapeItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAbstractGraphicsShapeItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:604
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractGraphicsShapeItem()
func DeleteQAbstractGraphicsShapeItem(this *QAbstractGraphicsShapeItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAbstractGraphicsShapeItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:606
// index:0
// Public Visibility=Default Availability=Available
// [8] QPen pen()
func (this *QAbstractGraphicsShapeItem) Pen() *qtgui.QPen /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QAbstractGraphicsShapeItem3penEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPenFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPen)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:607
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPen(const QPen &)
func (this *QAbstractGraphicsShapeItem) SetPen(pen qtgui.QPen_ITF) {
	var convArg0 = pen.QPen_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAbstractGraphicsShapeItem6setPenERK4QPen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:609
// index:0
// Public Visibility=Default Availability=Available
// [8] QBrush brush()
func (this *QAbstractGraphicsShapeItem) Brush() *qtgui.QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QAbstractGraphicsShapeItem5brushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:610
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBrush(const QBrush &)
func (this *QAbstractGraphicsShapeItem) SetBrush(brush qtgui.QBrush_ITF) {
	var convArg0 = brush.QBrush_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAbstractGraphicsShapeItem8setBrushERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:612
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isObscuredBy(const QGraphicsItem *)
func (this *QAbstractGraphicsShapeItem) IsObscuredBy(item QGraphicsItem_ITF /*777 const QGraphicsItem **/) bool {
	var convArg0 = item.QGraphicsItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QAbstractGraphicsShapeItem12isObscuredByEPK13QGraphicsItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:613
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath opaqueArea()
func (this *QAbstractGraphicsShapeItem) OpaqueArea() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QAbstractGraphicsShapeItem10opaqueAreaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

//  body block end

//  keep block begin

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
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
