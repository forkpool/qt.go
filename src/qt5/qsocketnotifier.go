package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qsocketnotifier.h
// dst-file: /src/core/qsocketnotifier.go
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
// class sizeof(QSocketNotifier)=1
type QSocketNotifier struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _activated QSocketNotifier_activated_signal;
}


func NewQSocketNotifier(args ...interface{})() {
}


func (this *QSocketNotifier) socket(args ...interface{}) () {
  // socket()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QSocketNotifier6socketEv
  default:
    qtrt.ErrorResolve("QSocketNotifier", "socket", args)
 }

}


func (this *QSocketNotifier) isEnabled(args ...interface{}) () {
  // isEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QSocketNotifier9isEnabledEv
  default:
    qtrt.ErrorResolve("QSocketNotifier", "isEnabled", args)
 }

}


func (this *QSocketNotifier) setEnabled(args ...interface{}) () {
  // setEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QSocketNotifier10setEnabledEb
  default:
    qtrt.ErrorResolve("QSocketNotifier", "setEnabled", args)
 }

}


func (this *QSocketNotifier) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QSocketNotifier10metaObjectEv
  default:
    qtrt.ErrorResolve("QSocketNotifier", "metaObject", args)
 }

}


func (this *QSocketNotifier) FreeQSocketNotifier(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QSocketNotifier", "~QSocketNotifier", args)
 }

}

// <= body block end

