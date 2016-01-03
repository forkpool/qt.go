package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtCore/qdatastream.h
// dst-file: /src/core/qdatastream.go
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
  // proto:  QDataStream & QDataStream::readBytes(char *& , uint & len);
extern void _ZN11QDataStream9readBytesERPcRj(void* qthis, unsigned char* arg0, int32_t* arg1);
  // proto:  void QDataStream::unsetDevice();
extern void _ZN11QDataStream11unsetDeviceEv(void* qthis);
  // proto:  void QDataStream::QDataStream(const QDataStream & );
extern void* dector_ZN11QDataStreamC1ERKS_(void* arg0);
extern void _ZN11QDataStreamC1ERKS_(void* qthis, void* arg0);
  // proto:  void QDataStream::QDataStream(QIODevice * );
extern void* dector_ZN11QDataStreamC1EP9QIODevice(void* arg0);
extern void _ZN11QDataStreamC1EP9QIODevice(void* qthis, void* arg0);
  // proto:  void QDataStream::~QDataStream();
extern void _ZN11QDataStreamD0Ev(void* qthis);
  // proto:  int QDataStream::skipRawData(int len);
extern void _ZN11QDataStream11skipRawDataEi(void* qthis, int32_t arg0);
  // proto:  QDataStream & QDataStream::writeBytes(const char * , uint len);
extern void _ZN11QDataStream10writeBytesEPKcj(void* qthis, unsigned char* arg0, int32_t arg1);
  // proto:  void QDataStream::QDataStream();
extern void* dector_ZN11QDataStreamC1Ev();
extern void _ZN11QDataStreamC1Ev(void* qthis);
  // proto:  void QDataStream::resetStatus();
extern void _ZN11QDataStream11resetStatusEv(void* qthis);
  // proto:  void QDataStream::QDataStream(const QByteArray & );
extern void* dector_ZN11QDataStreamC1ERK10QByteArray(void* arg0);
extern void _ZN11QDataStreamC1ERK10QByteArray(void* qthis, void* arg0);
  // proto:  int QDataStream::version();
extern void demth_ZNK11QDataStream7versionEv(void* qthis);
  // proto:  bool QDataStream::atEnd();
extern void _ZNK11QDataStream5atEndEv(void* qthis);
  // proto:  void QDataStream::setVersion(int );
extern void demth_ZN11QDataStream10setVersionEi(void* qthis, int32_t arg0);
  // proto:  void QDataStream::setDevice(QIODevice * );
extern void _ZN11QDataStream9setDeviceEP9QIODevice(void* qthis, void* arg0);
  // proto:  int QDataStream::writeRawData(const char * , int len);
extern void _ZN11QDataStream12writeRawDataEPKci(void* qthis, unsigned char* arg0, int32_t arg1);
  // proto:  int QDataStream::readRawData(char * , int len);
extern void _ZN11QDataStream11readRawDataEPci(void* qthis, unsigned char* arg0, int32_t arg1);
  // proto:  QIODevice * QDataStream::device();
extern void demth_ZNK11QDataStream6deviceEv(void* qthis);
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

// class sizeof(QDataStream)=1
type QDataStream struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  QDataStream & QDataStream::readBytes(char *& , uint & len);
func (this *QDataStream) readBytes(args ...interface{}) () {
  // readBytes(char *&, uint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.StringTy(false) // "char *&"
  vtys[0][1] = qtrt.Int32Ty(false) // "uint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDataStream9readBytesERPcRj
    // invoke: QDataStream & readBytes(char *&, uint &)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    C._ZN11QDataStream9readBytesERPcRj(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDataStream", "readBytes", args)
  }

}

  // proto:  void QDataStream::unsetDevice();
func (this *QDataStream) unsetDevice(args ...interface{}) () {
  // unsetDevice()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDataStream11unsetDeviceEv
    // invoke: void unsetDevice()
    C._ZN11QDataStream11unsetDeviceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataStream", "unsetDevice", args)
  }

}

  // proto:  void QDataStream::QDataStream(const QDataStream & );
func NewQDataStream(args ...interface{}) QDataStream {
  return QDataStream{}
}

  // proto:  void QDataStream::~QDataStream();
func (this *QDataStream) FreeQDataStream(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDataStream", "~QDataStream", args)
  }

}

  // proto:  int QDataStream::skipRawData(int len);
func (this *QDataStream) skipRawData(args ...interface{}) () {
  // skipRawData(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDataStream11skipRawDataEi
    // invoke: int skipRawData(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN11QDataStream11skipRawDataEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataStream", "skipRawData", args)
  }

}

  // proto:  QDataStream & QDataStream::writeBytes(const char * , uint len);
func (this *QDataStream) writeBytes(args ...interface{}) () {
  // writeBytes(const char *, uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDataStream10writeBytesEPKcj
    // invoke: QDataStream & writeBytes(const char *, uint)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN11QDataStream10writeBytesEPKcj(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDataStream", "writeBytes", args)
  }

}

  // proto:  void QDataStream::resetStatus();
func (this *QDataStream) resetStatus(args ...interface{}) () {
  // resetStatus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDataStream11resetStatusEv
    // invoke: void resetStatus()
    C._ZN11QDataStream11resetStatusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataStream", "resetStatus", args)
  }

}

  // proto:  int QDataStream::version();
func (this *QDataStream) version(args ...interface{}) () {
  // version()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QDataStream7versionEv
    // invoke: int version()
    C.demth_ZNK11QDataStream7versionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataStream", "version", args)
  }

}

  // proto:  bool QDataStream::atEnd();
func (this *QDataStream) atEnd(args ...interface{}) () {
  // atEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QDataStream5atEndEv
    // invoke: bool atEnd()
    C._ZNK11QDataStream5atEndEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataStream", "atEnd", args)
  }

}

  // proto:  void QDataStream::setVersion(int );
func (this *QDataStream) setVersion(args ...interface{}) () {
  // setVersion(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDataStream10setVersionEi
    // invoke: void setVersion(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZN11QDataStream10setVersionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataStream", "setVersion", args)
  }

}

  // proto:  void QDataStream::setDevice(QIODevice * );
func (this *QDataStream) setDevice(args ...interface{}) () {
  // setDevice(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDataStream9setDeviceEP9QIODevice
    // invoke: void setDevice(class QIODevice *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QDataStream9setDeviceEP9QIODevice(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDataStream", "setDevice", args)
  }

}

  // proto:  int QDataStream::writeRawData(const char * , int len);
func (this *QDataStream) writeRawData(args ...interface{}) () {
  // writeRawData(const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDataStream12writeRawDataEPKci
    // invoke: int writeRawData(const char *, int)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN11QDataStream12writeRawDataEPKci(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDataStream", "writeRawData", args)
  }

}

  // proto:  int QDataStream::readRawData(char * , int len);
func (this *QDataStream) readRawData(args ...interface{}) () {
  // readRawData(char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDataStream11readRawDataEPci
    // invoke: int readRawData(char *, int)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN11QDataStream11readRawDataEPci(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QDataStream", "readRawData", args)
  }

}

  // proto:  QIODevice * QDataStream::device();
func (this *QDataStream) device(args ...interface{}) () {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QDataStream6deviceEv
    // invoke: QIODevice * device()
    C.demth_ZNK11QDataStream6deviceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDataStream", "device", args)
  }

}

// <= body block end
