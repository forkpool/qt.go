package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qgesturerecognizer.h
// dst-file: /src/widgets/qgesturerecognizer.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QGestureRecognizer)=8
type QGestureRecognizer struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QGestureRecognizer) FreeQGestureRecognizer(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QGestureRecognizer", "~QGestureRecognizer", args)
 }

}


func NewQGestureRecognizer(args ...interface{})() {
}


func (this *QGestureRecognizer) reset(args ...interface{}) () {
  // reset(class QGesture *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGesture{}) // "QGesture *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QGestureRecognizer5resetEP8QGesture
  default:
    qtrt.ErrorResolve("QGestureRecognizer", "reset", args)
 }

}


func (this *QGestureRecognizer) create(args ...interface{}) () {
  // create(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QGestureRecognizer6createEP7QObject
  default:
    qtrt.ErrorResolve("QGestureRecognizer", "create", args)
 }

}

// <= body block end

