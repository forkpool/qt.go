package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qtemporarydir.h
// dst-file: /src/core/qtemporarydir.go
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
// class sizeof(QTemporaryDir)=1
type QTemporaryDir struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QTemporaryDir) remove(args ...interface{}) () {
  // remove()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QTemporaryDir6removeEv
  default:
    qtrt.ErrorResolve("QTemporaryDir", "remove", args)
 }

}


func (this *QTemporaryDir) autoRemove(args ...interface{}) () {
  // autoRemove()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QTemporaryDir10autoRemoveEv
  default:
    qtrt.ErrorResolve("QTemporaryDir", "autoRemove", args)
 }

}


func (this *QTemporaryDir) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QTemporaryDir7isValidEv
  default:
    qtrt.ErrorResolve("QTemporaryDir", "isValid", args)
 }

}


func (this *QTemporaryDir) setAutoRemove(args ...interface{}) () {
  // setAutoRemove(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN13QTemporaryDir13setAutoRemoveEb
  default:
    qtrt.ErrorResolve("QTemporaryDir", "setAutoRemove", args)
 }

}


func (this *QTemporaryDir) FreeQTemporaryDir(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTemporaryDir", "~QTemporaryDir", args)
 }

}


func NewQTemporaryDir(args ...interface{})() {
}


func (this *QTemporaryDir) path(args ...interface{}) () {
  // path()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK13QTemporaryDir4pathEv
  default:
    qtrt.ErrorResolve("QTemporaryDir", "path", args)
 }

}

// <= body block end

