package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtWidgets/qgraphicswidget.h
// dst-file: /src/widgets/qgraphicswidget.go
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
  // proto:  void QGraphicsWidget::setAutoFillBackground(bool enabled);
extern void _ZN15QGraphicsWidget21setAutoFillBackgroundEb(void* qthis, bool arg0);
  // proto:  void QGraphicsWidget::setWindowTitle(const QString & title);
extern void _ZN15QGraphicsWidget14setWindowTitleERK7QString(void* qthis, void* arg0);
  // proto:  void QGraphicsWidget::setLayout(QGraphicsLayout * layout);
extern void _ZN15QGraphicsWidget9setLayoutEP15QGraphicsLayout(void* qthis, void* arg0);
  // proto:  void QGraphicsWidget::setGeometry(const QRectF & rect);
extern void _ZN15QGraphicsWidget11setGeometryERK6QRectF(void* qthis, void* arg0);
  // proto:  QRectF QGraphicsWidget::rect();
extern void demth_ZNK15QGraphicsWidget4rectEv(void* qthis);
  // proto:  QSizeF QGraphicsWidget::size();
extern void _ZNK15QGraphicsWidget4sizeEv(void* qthis);
  // proto:  void QGraphicsWidget::releaseShortcut(int id);
extern void _ZN15QGraphicsWidget15releaseShortcutEi(void* qthis, int32_t arg0);
  // proto:  void QGraphicsWidget::setWindowFrameMargins(qreal left, qreal top, qreal right, qreal bottom);
extern void _ZN15QGraphicsWidget21setWindowFrameMarginsEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  int QGraphicsWidget::type();
extern void _ZNK15QGraphicsWidget4typeEv(void* qthis);
  // proto:  void QGraphicsWidget::unsetLayoutDirection();
extern void _ZN15QGraphicsWidget20unsetLayoutDirectionEv(void* qthis);
  // proto:  void QGraphicsWidget::QGraphicsWidget(const QGraphicsWidget & );
extern void* dector_ZN15QGraphicsWidgetC1ERKS_(void* arg0);
extern void _ZN15QGraphicsWidgetC1ERKS_(void* qthis, void* arg0);
  // proto:  QRectF QGraphicsWidget::windowFrameGeometry();
extern void _ZNK15QGraphicsWidget19windowFrameGeometryEv(void* qthis);
  // proto:  void QGraphicsWidget::resize(qreal w, qreal h);
extern void demth_ZN15QGraphicsWidget6resizeEdd(void* qthis, double arg0, double arg1);
  // proto:  QRectF QGraphicsWidget::windowFrameRect();
extern void _ZNK15QGraphicsWidget15windowFrameRectEv(void* qthis);
  // proto:  void QGraphicsWidget::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void _ZN15QGraphicsWidget5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QGraphicsWidget::adjustSize();
extern void _ZN15QGraphicsWidget10adjustSizeEv(void* qthis);
  // proto:  void QGraphicsWidget::paintWindowFrame(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
extern void _ZN15QGraphicsWidget16paintWindowFrameEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  QPalette QGraphicsWidget::palette();
extern void _ZNK15QGraphicsWidget7paletteEv(void* qthis);
  // proto:  void QGraphicsWidget::unsetWindowFrameMargins();
extern void _ZN15QGraphicsWidget23unsetWindowFrameMarginsEv(void* qthis);
  // proto:  void QGraphicsWidget::resize(const QSizeF & size);
extern void _ZN15QGraphicsWidget6resizeERK6QSizeF(void* qthis, void* arg0);
  // proto:  void QGraphicsWidget::setPalette(const QPalette & palette);
extern void _ZN15QGraphicsWidget10setPaletteERK8QPalette(void* qthis, void* arg0);
  // proto:  bool QGraphicsWidget::autoFillBackground();
extern void _ZNK15QGraphicsWidget18autoFillBackgroundEv(void* qthis);
  // proto:  QStyle * QGraphicsWidget::style();
extern void _ZNK15QGraphicsWidget5styleEv(void* qthis);
  // proto:  QPainterPath QGraphicsWidget::shape();
extern void _ZNK15QGraphicsWidget5shapeEv(void* qthis);
  // proto:  void QGraphicsWidget::setShortcutEnabled(int id, bool enabled);
extern void _ZN15QGraphicsWidget18setShortcutEnabledEib(void* qthis, int32_t arg0, bool arg1);
  // proto:  void QGraphicsWidget::removeAction(QAction * action);
extern void _ZN15QGraphicsWidget12removeActionEP7QAction(void* qthis, void* arg0);
  // proto:  void QGraphicsWidget::insertAction(QAction * before, QAction * action);
extern void _ZN15QGraphicsWidget12insertActionEP7QActionS1_(void* qthis, void* arg0, void* arg1);
  // proto:  bool QGraphicsWidget::close();
extern void _ZN15QGraphicsWidget5closeEv(void* qthis);
  // proto:  const QMetaObject * QGraphicsWidget::metaObject();
extern void _ZNK15QGraphicsWidget10metaObjectEv(void* qthis);
  // proto:  QRectF QGraphicsWidget::boundingRect();
extern void _ZNK15QGraphicsWidget12boundingRectEv(void* qthis);
  // proto:  void QGraphicsWidget::setContentsMargins(qreal left, qreal top, qreal right, qreal bottom);
extern void _ZN15QGraphicsWidget18setContentsMarginsEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  void QGraphicsWidget::setFont(const QFont & font);
extern void _ZN15QGraphicsWidget7setFontERK5QFont(void* qthis, void* arg0);
  // proto:  QString QGraphicsWidget::windowTitle();
extern void _ZNK15QGraphicsWidget11windowTitleEv(void* qthis);
  // proto:  QGraphicsLayout * QGraphicsWidget::layout();
extern void _ZNK15QGraphicsWidget6layoutEv(void* qthis);
  // proto:  void QGraphicsWidget::~QGraphicsWidget();
extern void _ZN15QGraphicsWidgetD0Ev(void* qthis);
  // proto:  QGraphicsWidget * QGraphicsWidget::focusWidget();
extern void _ZNK15QGraphicsWidget11focusWidgetEv(void* qthis);
  // proto:  void QGraphicsWidget::addAction(QAction * action);
extern void _ZN15QGraphicsWidget9addActionEP7QAction(void* qthis, void* arg0);
  // proto:  QFont QGraphicsWidget::font();
extern void _ZNK15QGraphicsWidget4fontEv(void* qthis);
  // proto:  QList<QAction *> QGraphicsWidget::actions();
extern void _ZNK15QGraphicsWidget7actionsEv(void* qthis);
  // proto:  void QGraphicsWidget::setShortcutAutoRepeat(int id, bool enabled);
extern void _ZN15QGraphicsWidget21setShortcutAutoRepeatEib(void* qthis, int32_t arg0, bool arg1);
  // proto: static void QGraphicsWidget::setTabOrder(QGraphicsWidget * first, QGraphicsWidget * second);
extern void _ZN15QGraphicsWidget11setTabOrderEPS_S0_(void* arg0, void* arg1);
  // proto:  void QGraphicsWidget::getWindowFrameMargins(qreal * left, qreal * top, qreal * right, qreal * bottom);
extern void _ZNK15QGraphicsWidget21getWindowFrameMarginsEPdS0_S0_S0_(void* qthis, double* arg0, double* arg1, double* arg2, double* arg3);
  // proto:  void QGraphicsWidget::setStyle(QStyle * style);
extern void _ZN15QGraphicsWidget8setStyleEP6QStyle(void* qthis, void* arg0);
  // proto:  void QGraphicsWidget::getContentsMargins(qreal * left, qreal * top, qreal * right, qreal * bottom);
extern void _ZNK15QGraphicsWidget18getContentsMarginsEPdS0_S0_S0_(void* qthis, double* arg0, double* arg1, double* arg2, double* arg3);
  // proto:  bool QGraphicsWidget::isActiveWindow();
extern void _ZNK15QGraphicsWidget14isActiveWindowEv(void* qthis);
  // proto:  void QGraphicsWidget::setGeometry(qreal x, qreal y, qreal w, qreal h);
extern void demth_ZN15QGraphicsWidget11setGeometryEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
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

// class sizeof(QGraphicsWidget)=1
type QGraphicsWidget struct {
  /*qbase*/ QGraphicsObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _layoutChanged QGraphicsWidget_layoutChanged_signal;
//  _geometryChanged QGraphicsWidget_geometryChanged_signal;
}

  // proto:  void QGraphicsWidget::setAutoFillBackground(bool enabled);
func (this *QGraphicsWidget) setAutoFillBackground(args ...interface{}) () {
  // setAutoFillBackground(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget21setAutoFillBackgroundEb
    // invoke: void setAutoFillBackground(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN15QGraphicsWidget21setAutoFillBackgroundEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setAutoFillBackground", args)
  }

}

  // proto:  void QGraphicsWidget::setWindowTitle(const QString & title);
func (this *QGraphicsWidget) setWindowTitle(args ...interface{}) () {
  // setWindowTitle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget14setWindowTitleERK7QString
    // invoke: void setWindowTitle(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QGraphicsWidget14setWindowTitleERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setWindowTitle", args)
  }

}

  // proto:  void QGraphicsWidget::setLayout(QGraphicsLayout * layout);
func (this *QGraphicsWidget) setLayout(args ...interface{}) () {
  // setLayout(class QGraphicsLayout *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsLayout{}) // "QGraphicsLayout *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget9setLayoutEP15QGraphicsLayout
    // invoke: void setLayout(class QGraphicsLayout *)
    var arg0 = args[0].(QGraphicsLayout).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QGraphicsWidget9setLayoutEP15QGraphicsLayout(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setLayout", args)
  }

}

  // proto:  void QGraphicsWidget::setGeometry(const QRectF & rect);
func (this *QGraphicsWidget) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRectF &)
  // setGeometry(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget11setGeometryERK6QRectF
    // invoke: void setGeometry(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QGraphicsWidget11setGeometryERK6QRectF(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN15QGraphicsWidget11setGeometryEdddd
    // invoke: void setGeometry(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C.demth_ZN15QGraphicsWidget11setGeometryEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setGeometry", args)
  }

}

  // proto:  QRectF QGraphicsWidget::rect();
func (this *QGraphicsWidget) rect(args ...interface{}) () {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget4rectEv
    // invoke: QRectF rect()
    C.demth_ZNK15QGraphicsWidget4rectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "rect", args)
  }

}

  // proto:  QSizeF QGraphicsWidget::size();
func (this *QGraphicsWidget) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget4sizeEv
    // invoke: QSizeF size()
    C._ZNK15QGraphicsWidget4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "size", args)
  }

}

  // proto:  void QGraphicsWidget::releaseShortcut(int id);
func (this *QGraphicsWidget) releaseShortcut(args ...interface{}) () {
  // releaseShortcut(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget15releaseShortcutEi
    // invoke: void releaseShortcut(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN15QGraphicsWidget15releaseShortcutEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "releaseShortcut", args)
  }

}

  // proto:  void QGraphicsWidget::setWindowFrameMargins(qreal left, qreal top, qreal right, qreal bottom);
func (this *QGraphicsWidget) setWindowFrameMargins(args ...interface{}) () {
  // setWindowFrameMargins(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget21setWindowFrameMarginsEdddd
    // invoke: void setWindowFrameMargins(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C._ZN15QGraphicsWidget21setWindowFrameMarginsEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setWindowFrameMargins", args)
  }

}

  // proto:  int QGraphicsWidget::type();
func (this *QGraphicsWidget) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "type", args)
  }

}

  // proto:  void QGraphicsWidget::unsetLayoutDirection();
func (this *QGraphicsWidget) unsetLayoutDirection(args ...interface{}) () {
  // unsetLayoutDirection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget20unsetLayoutDirectionEv
    // invoke: void unsetLayoutDirection()
    C._ZN15QGraphicsWidget20unsetLayoutDirectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "unsetLayoutDirection", args)
  }

}

  // proto:  void QGraphicsWidget::QGraphicsWidget(const QGraphicsWidget & );
func NewQGraphicsWidget(args ...interface{}) QGraphicsWidget {
  return QGraphicsWidget{}
}

  // proto:  QRectF QGraphicsWidget::windowFrameGeometry();
func (this *QGraphicsWidget) windowFrameGeometry(args ...interface{}) () {
  // windowFrameGeometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget19windowFrameGeometryEv
    // invoke: QRectF windowFrameGeometry()
    C._ZNK15QGraphicsWidget19windowFrameGeometryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "windowFrameGeometry", args)
  }

}

  // proto:  void QGraphicsWidget::resize(qreal w, qreal h);
func (this *QGraphicsWidget) resize(args ...interface{}) () {
  // resize(qreal, qreal)
  // resize(const class QSizeF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget6resizeEdd
    // invoke: void resize(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.demth_ZN15QGraphicsWidget6resizeEdd(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN15QGraphicsWidget6resizeERK6QSizeF
    // invoke: void resize(const class QSizeF &)
    var arg0 = args[0].(QSizeF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QGraphicsWidget6resizeERK6QSizeF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "resize", args)
  }

}

  // proto:  QRectF QGraphicsWidget::windowFrameRect();
func (this *QGraphicsWidget) windowFrameRect(args ...interface{}) () {
  // windowFrameRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget15windowFrameRectEv
    // invoke: QRectF windowFrameRect()
    C._ZNK15QGraphicsWidget15windowFrameRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "windowFrameRect", args)
  }

}

  // proto:  void QGraphicsWidget::paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
func (this *QGraphicsWidget) paint(args ...interface{}) () {
  // paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
    // invoke: void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(QPainter).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStyleOptionGraphicsItem).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN15QGraphicsWidget5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "paint", args)
  }

}

  // proto:  void QGraphicsWidget::adjustSize();
func (this *QGraphicsWidget) adjustSize(args ...interface{}) () {
  // adjustSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget10adjustSizeEv
    // invoke: void adjustSize()
    C._ZN15QGraphicsWidget10adjustSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "adjustSize", args)
  }

}

  // proto:  void QGraphicsWidget::paintWindowFrame(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget);
func (this *QGraphicsWidget) paintWindowFrame(args ...interface{}) () {
  // paintWindowFrame(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget16paintWindowFrameEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget
    // invoke: void paintWindowFrame(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
    var arg0 = args[0].(QPainter).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStyleOptionGraphicsItem).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN15QGraphicsWidget16paintWindowFrameEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "paintWindowFrame", args)
  }

}

  // proto:  QPalette QGraphicsWidget::palette();
func (this *QGraphicsWidget) palette(args ...interface{}) () {
  // palette()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget7paletteEv
    // invoke: QPalette palette()
    C._ZNK15QGraphicsWidget7paletteEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "palette", args)
  }

}

  // proto:  void QGraphicsWidget::unsetWindowFrameMargins();
func (this *QGraphicsWidget) unsetWindowFrameMargins(args ...interface{}) () {
  // unsetWindowFrameMargins()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget23unsetWindowFrameMarginsEv
    // invoke: void unsetWindowFrameMargins()
    C._ZN15QGraphicsWidget23unsetWindowFrameMarginsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "unsetWindowFrameMargins", args)
  }

}

  // proto:  void QGraphicsWidget::setPalette(const QPalette & palette);
func (this *QGraphicsWidget) setPalette(args ...interface{}) () {
  // setPalette(const class QPalette &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPalette{}) // "const QPalette &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget10setPaletteERK8QPalette
    // invoke: void setPalette(const class QPalette &)
    var arg0 = args[0].(QPalette).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QGraphicsWidget10setPaletteERK8QPalette(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setPalette", args)
  }

}

  // proto:  bool QGraphicsWidget::autoFillBackground();
func (this *QGraphicsWidget) autoFillBackground(args ...interface{}) () {
  // autoFillBackground()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget18autoFillBackgroundEv
    // invoke: bool autoFillBackground()
    C._ZNK15QGraphicsWidget18autoFillBackgroundEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "autoFillBackground", args)
  }

}

  // proto:  QStyle * QGraphicsWidget::style();
func (this *QGraphicsWidget) style(args ...interface{}) () {
  // style()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget5styleEv
    // invoke: QStyle * style()
    C._ZNK15QGraphicsWidget5styleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "style", args)
  }

}

  // proto:  QPainterPath QGraphicsWidget::shape();
func (this *QGraphicsWidget) shape(args ...interface{}) () {
  // shape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget5shapeEv
    // invoke: QPainterPath shape()
    C._ZNK15QGraphicsWidget5shapeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "shape", args)
  }

}

  // proto:  void QGraphicsWidget::setShortcutEnabled(int id, bool enabled);
func (this *QGraphicsWidget) setShortcutEnabled(args ...interface{}) () {
  // setShortcutEnabled(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget18setShortcutEnabledEib
    // invoke: void setShortcutEnabled(int, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C._ZN15QGraphicsWidget18setShortcutEnabledEib(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setShortcutEnabled", args)
  }

}

  // proto:  void QGraphicsWidget::removeAction(QAction * action);
func (this *QGraphicsWidget) removeAction(args ...interface{}) () {
  // removeAction(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget12removeActionEP7QAction
    // invoke: void removeAction(class QAction *)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QGraphicsWidget12removeActionEP7QAction(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "removeAction", args)
  }

}

  // proto:  void QGraphicsWidget::insertAction(QAction * before, QAction * action);
func (this *QGraphicsWidget) insertAction(args ...interface{}) () {
  // insertAction(class QAction *, class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"
  vtys[0][1] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget12insertActionEP7QActionS1_
    // invoke: void insertAction(class QAction *, class QAction *)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QAction).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN15QGraphicsWidget12insertActionEP7QActionS1_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "insertAction", args)
  }

}

  // proto:  bool QGraphicsWidget::close();
func (this *QGraphicsWidget) close(args ...interface{}) () {
  // close()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget5closeEv
    // invoke: bool close()
    C._ZN15QGraphicsWidget5closeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "close", args)
  }

}

  // proto:  const QMetaObject * QGraphicsWidget::metaObject();
func (this *QGraphicsWidget) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK15QGraphicsWidget10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "metaObject", args)
  }

}

  // proto:  QRectF QGraphicsWidget::boundingRect();
func (this *QGraphicsWidget) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget12boundingRectEv
    // invoke: QRectF boundingRect()
    C._ZNK15QGraphicsWidget12boundingRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "boundingRect", args)
  }

}

  // proto:  void QGraphicsWidget::setContentsMargins(qreal left, qreal top, qreal right, qreal bottom);
func (this *QGraphicsWidget) setContentsMargins(args ...interface{}) () {
  // setContentsMargins(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget18setContentsMarginsEdddd
    // invoke: void setContentsMargins(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C._ZN15QGraphicsWidget18setContentsMarginsEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setContentsMargins", args)
  }

}

  // proto:  void QGraphicsWidget::setFont(const QFont & font);
func (this *QGraphicsWidget) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget7setFontERK5QFont
    // invoke: void setFont(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QGraphicsWidget7setFontERK5QFont(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setFont", args)
  }

}

  // proto:  QString QGraphicsWidget::windowTitle();
func (this *QGraphicsWidget) windowTitle(args ...interface{}) () {
  // windowTitle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget11windowTitleEv
    // invoke: QString windowTitle()
    C._ZNK15QGraphicsWidget11windowTitleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "windowTitle", args)
  }

}

  // proto:  QGraphicsLayout * QGraphicsWidget::layout();
func (this *QGraphicsWidget) layout(args ...interface{}) () {
  // layout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget6layoutEv
    // invoke: QGraphicsLayout * layout()
    C._ZNK15QGraphicsWidget6layoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "layout", args)
  }

}

  // proto:  void QGraphicsWidget::~QGraphicsWidget();
func (this *QGraphicsWidget) FreeQGraphicsWidget(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "~QGraphicsWidget", args)
  }

}

  // proto:  QGraphicsWidget * QGraphicsWidget::focusWidget();
func (this *QGraphicsWidget) focusWidget(args ...interface{}) () {
  // focusWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget11focusWidgetEv
    // invoke: QGraphicsWidget * focusWidget()
    C._ZNK15QGraphicsWidget11focusWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "focusWidget", args)
  }

}

  // proto:  void QGraphicsWidget::addAction(QAction * action);
func (this *QGraphicsWidget) addAction(args ...interface{}) () {
  // addAction(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget9addActionEP7QAction
    // invoke: void addAction(class QAction *)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QGraphicsWidget9addActionEP7QAction(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "addAction", args)
  }

}

  // proto:  QFont QGraphicsWidget::font();
func (this *QGraphicsWidget) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget4fontEv
    // invoke: QFont font()
    C._ZNK15QGraphicsWidget4fontEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "font", args)
  }

}

  // proto:  QList<QAction *> QGraphicsWidget::actions();
func (this *QGraphicsWidget) actions(args ...interface{}) () {
  // actions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget7actionsEv
    // invoke: QList<QAction *> actions()
    C._ZNK15QGraphicsWidget7actionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "actions", args)
  }

}

  // proto:  void QGraphicsWidget::setShortcutAutoRepeat(int id, bool enabled);
func (this *QGraphicsWidget) setShortcutAutoRepeat(args ...interface{}) () {
  // setShortcutAutoRepeat(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget21setShortcutAutoRepeatEib
    // invoke: void setShortcutAutoRepeat(int, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C._ZN15QGraphicsWidget21setShortcutAutoRepeatEib(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setShortcutAutoRepeat", args)
  }

}

  // proto: static void QGraphicsWidget::setTabOrder(QGraphicsWidget * first, QGraphicsWidget * second);
func (this *QGraphicsWidget) setTabOrder_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setTabOrder", args)
  }

}

  // proto:  void QGraphicsWidget::getWindowFrameMargins(qreal * left, qreal * top, qreal * right, qreal * bottom);
func (this *QGraphicsWidget) getWindowFrameMargins(args ...interface{}) () {
  // getWindowFrameMargins(qreal *, qreal *, qreal *, qreal *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][1] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][2] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][3] = qtrt.DoubleTy(true) // "qreal *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget21getWindowFrameMarginsEPdS0_S0_S0_
    // invoke: void getWindowFrameMargins(qreal *, qreal *, qreal *, qreal *)
    var arg0 = (*C.double)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.double)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.double)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.double)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C._ZNK15QGraphicsWidget21getWindowFrameMarginsEPdS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "getWindowFrameMargins", args)
  }

}

  // proto:  void QGraphicsWidget::setStyle(QStyle * style);
func (this *QGraphicsWidget) setStyle(args ...interface{}) () {
  // setStyle(class QStyle *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyle{}) // "QStyle *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsWidget8setStyleEP6QStyle
    // invoke: void setStyle(class QStyle *)
    var arg0 = args[0].(QStyle).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QGraphicsWidget8setStyleEP6QStyle(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "setStyle", args)
  }

}

  // proto:  void QGraphicsWidget::getContentsMargins(qreal * left, qreal * top, qreal * right, qreal * bottom);
func (this *QGraphicsWidget) getContentsMargins(args ...interface{}) () {
  // getContentsMargins(qreal *, qreal *, qreal *, qreal *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][1] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][2] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][3] = qtrt.DoubleTy(true) // "qreal *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget18getContentsMarginsEPdS0_S0_S0_
    // invoke: void getContentsMargins(qreal *, qreal *, qreal *, qreal *)
    var arg0 = (*C.double)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.double)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.double)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.double)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C._ZNK15QGraphicsWidget18getContentsMarginsEPdS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "getContentsMargins", args)
  }

}

  // proto:  bool QGraphicsWidget::isActiveWindow();
func (this *QGraphicsWidget) isActiveWindow(args ...interface{}) () {
  // isActiveWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsWidget14isActiveWindowEv
    // invoke: bool isActiveWindow()
    C._ZNK15QGraphicsWidget14isActiveWindowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsWidget", "isActiveWindow", args)
  }

}

// <= body block end
