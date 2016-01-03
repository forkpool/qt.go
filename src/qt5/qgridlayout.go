package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtWidgets/qgridlayout.h
// dst-file: /src/widgets/qgridlayout.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QGridLayout::setRowMinimumHeight(int row, int minSize);
extern void _ZN11QGridLayout19setRowMinimumHeightEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  QLayoutItem * QGridLayout::takeAt(int index);
extern void _ZN11QGridLayout6takeAtEi(void* qthis, int32_t arg0);
  // proto:  void QGridLayout::getItemPosition(int idx, int * row, int * column, int * rowSpan, int * columnSpan);
extern void _ZNK11QGridLayout15getItemPositionEiPiS0_S0_S0_(void* qthis, int32_t arg0, int32_t* arg1, int32_t* arg2, int32_t* arg3, int32_t* arg4);
  // proto:  int QGridLayout::minimumHeightForWidth(int );
extern void _ZNK11QGridLayout21minimumHeightForWidthEi(void* qthis, int32_t arg0);
  // proto:  int QGridLayout::rowMinimumHeight(int row);
extern void _ZNK11QGridLayout16rowMinimumHeightEi(void* qthis, int32_t arg0);
  // proto:  void QGridLayout::invalidate();
extern void _ZN11QGridLayout10invalidateEv(void* qthis);
  // proto:  int QGridLayout::count();
extern void _ZNK11QGridLayout5countEv(void* qthis);
  // proto:  void QGridLayout::setColumnStretch(int column, int stretch);
extern void _ZN11QGridLayout16setColumnStretchEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  int QGridLayout::spacing();
extern void _ZNK11QGridLayout7spacingEv(void* qthis);
  // proto:  int QGridLayout::rowStretch(int row);
extern void _ZNK11QGridLayout10rowStretchEi(void* qthis, int32_t arg0);
  // proto:  QSize QGridLayout::sizeHint();
extern void _ZNK11QGridLayout8sizeHintEv(void* qthis);
  // proto:  int QGridLayout::rowCount();
extern void _ZNK11QGridLayout8rowCountEv(void* qthis);
  // proto:  void QGridLayout::setGeometry(const QRect & );
extern void _ZN11QGridLayout11setGeometryERK5QRect(void* qthis, void* arg0);
  // proto:  void QGridLayout::setVerticalSpacing(int spacing);
extern void _ZN11QGridLayout18setVerticalSpacingEi(void* qthis, int32_t arg0);
  // proto:  void QGridLayout::setHorizontalSpacing(int spacing);
extern void _ZN11QGridLayout20setHorizontalSpacingEi(void* qthis, int32_t arg0);
  // proto:  int QGridLayout::columnStretch(int column);
extern void _ZNK11QGridLayout13columnStretchEi(void* qthis, int32_t arg0);
  // proto:  void QGridLayout::QGridLayout(const QGridLayout & );
extern void* dector_ZN11QGridLayoutC1ERKS_(void* arg0);
extern void _ZN11QGridLayoutC1ERKS_(void* qthis, void* arg0);
  // proto:  int QGridLayout::columnCount();
extern void _ZNK11QGridLayout11columnCountEv(void* qthis);
  // proto:  int QGridLayout::columnMinimumWidth(int column);
extern void _ZNK11QGridLayout18columnMinimumWidthEi(void* qthis, int32_t arg0);
  // proto:  QSize QGridLayout::minimumSize();
extern void _ZNK11QGridLayout11minimumSizeEv(void* qthis);
  // proto:  bool QGridLayout::hasHeightForWidth();
extern void _ZNK11QGridLayout17hasHeightForWidthEv(void* qthis);
  // proto:  QRect QGridLayout::cellRect(int row, int column);
extern void _ZNK11QGridLayout8cellRectEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  void QGridLayout::setRowStretch(int row, int stretch);
extern void _ZN11QGridLayout13setRowStretchEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  QLayoutItem * QGridLayout::itemAtPosition(int row, int column);
extern void _ZNK11QGridLayout14itemAtPositionEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  const QMetaObject * QGridLayout::metaObject();
extern void _ZNK11QGridLayout10metaObjectEv(void* qthis);
  // proto:  int QGridLayout::verticalSpacing();
extern void _ZNK11QGridLayout15verticalSpacingEv(void* qthis);
  // proto:  void QGridLayout::QGridLayout(QWidget * parent);
extern void* dector_ZN11QGridLayoutC1EP7QWidget(void* arg0);
extern void _ZN11QGridLayoutC1EP7QWidget(void* qthis, void* arg0);
  // proto:  int QGridLayout::horizontalSpacing();
extern void _ZNK11QGridLayout17horizontalSpacingEv(void* qthis);
  // proto:  void QGridLayout::setColumnMinimumWidth(int column, int minSize);
extern void _ZN11QGridLayout21setColumnMinimumWidthEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  void QGridLayout::QGridLayout();
extern void* dector_ZN11QGridLayoutC1Ev();
extern void _ZN11QGridLayoutC1Ev(void* qthis);
  // proto:  int QGridLayout::heightForWidth(int );
extern void _ZNK11QGridLayout14heightForWidthEi(void* qthis, int32_t arg0);
  // proto:  void QGridLayout::~QGridLayout();
extern void _ZN11QGridLayoutD0Ev(void* qthis);
  // proto:  void QGridLayout::setSpacing(int spacing);
extern void _ZN11QGridLayout10setSpacingEi(void* qthis, int32_t arg0);
  // proto:  void QGridLayout::addWidget(QWidget * w);
extern void demth_ZN11QGridLayout9addWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  QLayoutItem * QGridLayout::itemAt(int index);
extern void _ZNK11QGridLayout6itemAtEi(void* qthis, int32_t arg0);
  // proto:  QSize QGridLayout::maximumSize();
extern void _ZNK11QGridLayout11maximumSizeEv(void* qthis);
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QGridLayout)=1
type QGridLayout struct {
  /*qbase*/ QLayout;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QGridLayout::setRowMinimumHeight(int row, int minSize);
func (this *QGridLayout) setRowMinimumHeight(args ...interface{}) () {
  // setRowMinimumHeight(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout19setRowMinimumHeightEii
    // invoke: void setRowMinimumHeight(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN11QGridLayout19setRowMinimumHeightEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGridLayout", "setRowMinimumHeight", args)
  }

}

  // proto:  QLayoutItem * QGridLayout::takeAt(int index);
func (this *QGridLayout) takeAt(args ...interface{}) () {
  // takeAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout6takeAtEi
    // invoke: QLayoutItem * takeAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN11QGridLayout6takeAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGridLayout", "takeAt", args)
  }

}

  // proto:  void QGridLayout::getItemPosition(int idx, int * row, int * column, int * rowSpan, int * columnSpan);
func (this *QGridLayout) getItemPosition(args ...interface{}) () {
  // getItemPosition(int, int *, int *, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"
  vtys[0][4] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout15getItemPositionEiPiS0_S0_S0_
    // invoke: void getItemPosition(int, int *, int *, int *, int *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    var arg4 = (*C.int32_t)(args[4].(*int32))
    if false {fmt.Println(arg4)}
    C._ZNK11QGridLayout15getItemPositionEiPiS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QGridLayout", "getItemPosition", args)
  }

}

  // proto:  int QGridLayout::minimumHeightForWidth(int );
func (this *QGridLayout) minimumHeightForWidth(args ...interface{}) () {
  // minimumHeightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout21minimumHeightForWidthEi
    // invoke: int minimumHeightForWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK11QGridLayout21minimumHeightForWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGridLayout", "minimumHeightForWidth", args)
  }

}

  // proto:  int QGridLayout::rowMinimumHeight(int row);
func (this *QGridLayout) rowMinimumHeight(args ...interface{}) () {
  // rowMinimumHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout16rowMinimumHeightEi
    // invoke: int rowMinimumHeight(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK11QGridLayout16rowMinimumHeightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGridLayout", "rowMinimumHeight", args)
  }

}

  // proto:  void QGridLayout::invalidate();
func (this *QGridLayout) invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout10invalidateEv
    // invoke: void invalidate()
    C._ZN11QGridLayout10invalidateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGridLayout", "invalidate", args)
  }

}

  // proto:  int QGridLayout::count();
func (this *QGridLayout) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout5countEv
    // invoke: int count()
    C._ZNK11QGridLayout5countEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGridLayout", "count", args)
  }

}

  // proto:  void QGridLayout::setColumnStretch(int column, int stretch);
func (this *QGridLayout) setColumnStretch(args ...interface{}) () {
  // setColumnStretch(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout16setColumnStretchEii
    // invoke: void setColumnStretch(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN11QGridLayout16setColumnStretchEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGridLayout", "setColumnStretch", args)
  }

}

  // proto:  int QGridLayout::spacing();
func (this *QGridLayout) spacing(args ...interface{}) () {
  // spacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout7spacingEv
    // invoke: int spacing()
    C._ZNK11QGridLayout7spacingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGridLayout", "spacing", args)
  }

}

  // proto:  int QGridLayout::rowStretch(int row);
func (this *QGridLayout) rowStretch(args ...interface{}) () {
  // rowStretch(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout10rowStretchEi
    // invoke: int rowStretch(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK11QGridLayout10rowStretchEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGridLayout", "rowStretch", args)
  }

}

  // proto:  QSize QGridLayout::sizeHint();
func (this *QGridLayout) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout8sizeHintEv
    // invoke: QSize sizeHint()
    C._ZNK11QGridLayout8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGridLayout", "sizeHint", args)
  }

}

  // proto:  int QGridLayout::rowCount();
func (this *QGridLayout) rowCount(args ...interface{}) () {
  // rowCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout8rowCountEv
    // invoke: int rowCount()
    C._ZNK11QGridLayout8rowCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGridLayout", "rowCount", args)
  }

}

  // proto:  void QGridLayout::setGeometry(const QRect & );
func (this *QGridLayout) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout11setGeometryERK5QRect
    // invoke: void setGeometry(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QGridLayout11setGeometryERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGridLayout", "setGeometry", args)
  }

}

  // proto:  void QGridLayout::setVerticalSpacing(int spacing);
func (this *QGridLayout) setVerticalSpacing(args ...interface{}) () {
  // setVerticalSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout18setVerticalSpacingEi
    // invoke: void setVerticalSpacing(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN11QGridLayout18setVerticalSpacingEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGridLayout", "setVerticalSpacing", args)
  }

}

  // proto:  void QGridLayout::setHorizontalSpacing(int spacing);
func (this *QGridLayout) setHorizontalSpacing(args ...interface{}) () {
  // setHorizontalSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout20setHorizontalSpacingEi
    // invoke: void setHorizontalSpacing(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN11QGridLayout20setHorizontalSpacingEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGridLayout", "setHorizontalSpacing", args)
  }

}

  // proto:  int QGridLayout::columnStretch(int column);
func (this *QGridLayout) columnStretch(args ...interface{}) () {
  // columnStretch(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout13columnStretchEi
    // invoke: int columnStretch(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK11QGridLayout13columnStretchEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGridLayout", "columnStretch", args)
  }

}

  // proto:  void QGridLayout::QGridLayout(const QGridLayout & );
func NewQGridLayout(args ...interface{}) QGridLayout {
  return QGridLayout{}
}

  // proto:  int QGridLayout::columnCount();
func (this *QGridLayout) columnCount(args ...interface{}) () {
  // columnCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout11columnCountEv
    // invoke: int columnCount()
    C._ZNK11QGridLayout11columnCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGridLayout", "columnCount", args)
  }

}

  // proto:  int QGridLayout::columnMinimumWidth(int column);
func (this *QGridLayout) columnMinimumWidth(args ...interface{}) () {
  // columnMinimumWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout18columnMinimumWidthEi
    // invoke: int columnMinimumWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK11QGridLayout18columnMinimumWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGridLayout", "columnMinimumWidth", args)
  }

}

  // proto:  QSize QGridLayout::minimumSize();
func (this *QGridLayout) minimumSize(args ...interface{}) () {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout11minimumSizeEv
    // invoke: QSize minimumSize()
    C._ZNK11QGridLayout11minimumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGridLayout", "minimumSize", args)
  }

}

  // proto:  bool QGridLayout::hasHeightForWidth();
func (this *QGridLayout) hasHeightForWidth(args ...interface{}) () {
  // hasHeightForWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout17hasHeightForWidthEv
    // invoke: bool hasHeightForWidth()
    C._ZNK11QGridLayout17hasHeightForWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGridLayout", "hasHeightForWidth", args)
  }

}

  // proto:  QRect QGridLayout::cellRect(int row, int column);
func (this *QGridLayout) cellRect(args ...interface{}) () {
  // cellRect(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout8cellRectEii
    // invoke: QRect cellRect(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK11QGridLayout8cellRectEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGridLayout", "cellRect", args)
  }

}

  // proto:  void QGridLayout::setRowStretch(int row, int stretch);
func (this *QGridLayout) setRowStretch(args ...interface{}) () {
  // setRowStretch(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout13setRowStretchEii
    // invoke: void setRowStretch(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN11QGridLayout13setRowStretchEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGridLayout", "setRowStretch", args)
  }

}

  // proto:  QLayoutItem * QGridLayout::itemAtPosition(int row, int column);
func (this *QGridLayout) itemAtPosition(args ...interface{}) () {
  // itemAtPosition(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout14itemAtPositionEii
    // invoke: QLayoutItem * itemAtPosition(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK11QGridLayout14itemAtPositionEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGridLayout", "itemAtPosition", args)
  }

}

  // proto:  const QMetaObject * QGridLayout::metaObject();
func (this *QGridLayout) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK11QGridLayout10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGridLayout", "metaObject", args)
  }

}

  // proto:  int QGridLayout::verticalSpacing();
func (this *QGridLayout) verticalSpacing(args ...interface{}) () {
  // verticalSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout15verticalSpacingEv
    // invoke: int verticalSpacing()
    C._ZNK11QGridLayout15verticalSpacingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGridLayout", "verticalSpacing", args)
  }

}

  // proto:  int QGridLayout::horizontalSpacing();
func (this *QGridLayout) horizontalSpacing(args ...interface{}) () {
  // horizontalSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout17horizontalSpacingEv
    // invoke: int horizontalSpacing()
    C._ZNK11QGridLayout17horizontalSpacingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGridLayout", "horizontalSpacing", args)
  }

}

  // proto:  void QGridLayout::setColumnMinimumWidth(int column, int minSize);
func (this *QGridLayout) setColumnMinimumWidth(args ...interface{}) () {
  // setColumnMinimumWidth(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout21setColumnMinimumWidthEii
    // invoke: void setColumnMinimumWidth(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN11QGridLayout21setColumnMinimumWidthEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGridLayout", "setColumnMinimumWidth", args)
  }

}

  // proto:  int QGridLayout::heightForWidth(int );
func (this *QGridLayout) heightForWidth(args ...interface{}) () {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout14heightForWidthEi
    // invoke: int heightForWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK11QGridLayout14heightForWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGridLayout", "heightForWidth", args)
  }

}

  // proto:  void QGridLayout::~QGridLayout();
func (this *QGridLayout) FreeQGridLayout(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGridLayout", "~QGridLayout", args)
  }

}

  // proto:  void QGridLayout::setSpacing(int spacing);
func (this *QGridLayout) setSpacing(args ...interface{}) () {
  // setSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout10setSpacingEi
    // invoke: void setSpacing(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN11QGridLayout10setSpacingEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGridLayout", "setSpacing", args)
  }

}

  // proto:  void QGridLayout::addWidget(QWidget * w);
func (this *QGridLayout) addWidget(args ...interface{}) () {
  // addWidget(class QWidget *, int, int, Qt::Alignment)
  // addWidget(class QWidget *, int, int, int, int, Qt::Alignment)
  // addWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int64Ty(false) // "Qt::Alignment"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[1][4] = qtrt.Int32Ty(false) // "int"
  vtys[1][5] = qtrt.Int64Ty(false) // "Qt::Alignment"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout9addWidgetEP7QWidget
    // invoke: void addWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN11QGridLayout9addWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGridLayout", "addWidget", args)
  }

}

  // proto:  QLayoutItem * QGridLayout::itemAt(int index);
func (this *QGridLayout) itemAt(args ...interface{}) () {
  // itemAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout6itemAtEi
    // invoke: QLayoutItem * itemAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK11QGridLayout6itemAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGridLayout", "itemAt", args)
  }

}

  // proto:  QSize QGridLayout::maximumSize();
func (this *QGridLayout) maximumSize(args ...interface{}) () {
  // maximumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout11maximumSizeEv
    // invoke: QSize maximumSize()
    C._ZNK11QGridLayout11maximumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGridLayout", "maximumSize", args)
  }

}

// <= body block end
