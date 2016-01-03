package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtCore/qtextstream.h
// dst-file: /src/core/qtextstream.go
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
  // proto:  void QTextStreamManipulator::exec(QTextStream & s);
extern void demth_ZN22QTextStreamManipulator4execER11QTextStream(void* qthis, void* arg0);
  // proto:  void QTextStreamManipulator::QTextStreamManipulator(QTSMFI m, int a);
extern void* dector_ZN22QTextStreamManipulatorC1EM11QTextStreamFviEi(void* arg0, int32_t arg1);
extern void _ZN22QTextStreamManipulatorC1EM11QTextStreamFviEi(void* qthis, void* arg0, int32_t arg1);
  // proto:  void QTextStreamManipulator::QTextStreamManipulator(QTSMFC m, QChar c);
extern void* dector_ZN22QTextStreamManipulatorC1EM11QTextStreamFv5QCharES1_(void* arg0, void* arg1);
extern void _ZN22QTextStreamManipulatorC1EM11QTextStreamFv5QCharES1_(void* qthis, void* arg0, void* arg1);
  // proto:  QTextCodec * QTextStream::codec();
extern void _ZNK11QTextStream5codecEv(void* qthis);
  // proto:  void QTextStream::QTextStream(QIODevice * device);
extern void* dector_ZN11QTextStreamC1EP9QIODevice(void* arg0);
extern void _ZN11QTextStreamC1EP9QIODevice(void* qthis, void* arg0);
  // proto:  void QTextStream::setLocale(const QLocale & locale);
extern void _ZN11QTextStream9setLocaleERK7QLocale(void* qthis, void* arg0);
  // proto:  void QTextStream::QTextStream();
extern void* dector_ZN11QTextStreamC1Ev();
extern void _ZN11QTextStreamC1Ev(void* qthis);
  // proto:  bool QTextStream::atEnd();
extern void _ZNK11QTextStream5atEndEv(void* qthis);
  // proto:  bool QTextStream::readLineInto(QString * line, qint64 maxlen);
extern void _ZN11QTextStream12readLineIntoEP7QStringx(void* qthis, void* arg0, int64_t arg1);
  // proto:  void QTextStream::setRealNumberPrecision(int precision);
extern void _ZN11QTextStream22setRealNumberPrecisionEi(void* qthis, int32_t arg0);
  // proto:  void QTextStream::setDevice(QIODevice * device);
extern void _ZN11QTextStream9setDeviceEP9QIODevice(void* qthis, void* arg0);
  // proto:  void QTextStream::reset();
extern void _ZN11QTextStream5resetEv(void* qthis);
  // proto:  bool QTextStream::seek(qint64 pos);
extern void _ZN11QTextStream4seekEx(void* qthis, int64_t arg0);
  // proto:  QString * QTextStream::string();
extern void _ZNK11QTextStream6stringEv(void* qthis);
  // proto:  void QTextStream::setAutoDetectUnicode(bool enabled);
extern void _ZN11QTextStream20setAutoDetectUnicodeEb(void* qthis, bool arg0);
  // proto:  QChar QTextStream::padChar();
extern void _ZNK11QTextStream7padCharEv(void* qthis);
  // proto:  QIODevice * QTextStream::device();
extern void _ZNK11QTextStream6deviceEv(void* qthis);
  // proto:  void QTextStream::resetStatus();
extern void _ZN11QTextStream11resetStatusEv(void* qthis);
  // proto:  bool QTextStream::autoDetectUnicode();
extern void _ZNK11QTextStream17autoDetectUnicodeEv(void* qthis);
  // proto:  int QTextStream::fieldWidth();
extern void _ZNK11QTextStream10fieldWidthEv(void* qthis);
  // proto:  bool QTextStream::generateByteOrderMark();
extern void _ZNK11QTextStream21generateByteOrderMarkEv(void* qthis);
  // proto:  void QTextStream::setGenerateByteOrderMark(bool generate);
extern void _ZN11QTextStream24setGenerateByteOrderMarkEb(void* qthis, bool arg0);
  // proto:  void QTextStream::setCodec(QTextCodec * codec);
extern void _ZN11QTextStream8setCodecEP10QTextCodec(void* qthis, void* arg0);
  // proto:  void QTextStream::flush();
extern void _ZN11QTextStream5flushEv(void* qthis);
  // proto:  void QTextStream::setIntegerBase(int base);
extern void _ZN11QTextStream14setIntegerBaseEi(void* qthis, int32_t arg0);
  // proto:  void QTextStream::~QTextStream();
extern void _ZN11QTextStreamD0Ev(void* qthis);
  // proto:  QLocale QTextStream::locale();
extern void _ZNK11QTextStream6localeEv(void* qthis);
  // proto:  QString QTextStream::read(qint64 maxlen);
extern void _ZN11QTextStream4readEx(void* qthis, int64_t arg0);
  // proto:  void QTextStream::setPadChar(QChar ch);
extern void _ZN11QTextStream10setPadCharE5QChar(void* qthis, void* arg0);
  // proto:  int QTextStream::realNumberPrecision();
extern void _ZNK11QTextStream19realNumberPrecisionEv(void* qthis);
  // proto:  qint64 QTextStream::pos();
extern void _ZNK11QTextStream3posEv(void* qthis);
  // proto:  QString QTextStream::readAll();
extern void _ZN11QTextStream7readAllEv(void* qthis);
  // proto:  void QTextStream::skipWhiteSpace();
extern void _ZN11QTextStream14skipWhiteSpaceEv(void* qthis);
  // proto:  void QTextStream::setFieldWidth(int width);
extern void _ZN11QTextStream13setFieldWidthEi(void* qthis, int32_t arg0);
  // proto:  void QTextStream::setCodec(const char * codecName);
extern void _ZN11QTextStream8setCodecEPKc(void* qthis, unsigned char* arg0);
  // proto:  int QTextStream::integerBase();
extern void _ZNK11QTextStream11integerBaseEv(void* qthis);
  // proto:  QString QTextStream::readLine(qint64 maxlen);
extern void _ZN11QTextStream8readLineEx(void* qthis, int64_t arg0);
  // proto:  void QTextStream::QTextStream(const QTextStream & );
extern void* dector_ZN11QTextStreamC1ERKS_(void* arg0);
extern void _ZN11QTextStreamC1ERKS_(void* qthis, void* arg0);
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

// class sizeof(QTextStreamManipulator)=40
type QTextStreamManipulator struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextStream)=1
type QTextStream struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QTextStreamManipulator::exec(QTextStream & s);
func (this *QTextStreamManipulator) exec(args ...interface{}) () {
  // exec(class QTextStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextStream{}) // "QTextStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QTextStreamManipulator4execER11QTextStream
    // invoke: void exec(class QTextStream &)
    var arg0 = args[0].(QTextStream).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN22QTextStreamManipulator4execER11QTextStream(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStreamManipulator", "exec", args)
  }

}

  // proto:  void QTextStreamManipulator::QTextStreamManipulator(QTSMFI m, int a);
func NewQTextStreamManipulator(args ...interface{}) QTextStreamManipulator {
  return QTextStreamManipulator{}
}

  // proto:  QTextCodec * QTextStream::codec();
func (this *QTextStream) codec(args ...interface{}) () {
  // codec()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream5codecEv
    // invoke: QTextCodec * codec()
    C._ZNK11QTextStream5codecEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "codec", args)
  }

}

  // proto:  void QTextStream::QTextStream(QIODevice * device);
func NewQTextStream(args ...interface{}) QTextStream {
  return QTextStream{}
}

  // proto:  void QTextStream::setLocale(const QLocale & locale);
func (this *QTextStream) setLocale(args ...interface{}) () {
  // setLocale(const class QLocale &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLocale{}) // "const QLocale &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream9setLocaleERK7QLocale
    // invoke: void setLocale(const class QLocale &)
    var arg0 = args[0].(QLocale).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QTextStream9setLocaleERK7QLocale(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "setLocale", args)
  }

}

  // proto:  bool QTextStream::atEnd();
func (this *QTextStream) atEnd(args ...interface{}) () {
  // atEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream5atEndEv
    // invoke: bool atEnd()
    C._ZNK11QTextStream5atEndEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "atEnd", args)
  }

}

  // proto:  bool QTextStream::readLineInto(QString * line, qint64 maxlen);
func (this *QTextStream) readLineInto(args ...interface{}) () {
  // readLineInto(class QString *, qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "QString *"
  vtys[0][1] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream12readLineIntoEP7QStringx
    // invoke: bool readLineInto(class QString *, qint64)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int64_t(args[1].(int64))
    if false {fmt.Println(arg1)}
    C._ZN11QTextStream12readLineIntoEP7QStringx(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextStream", "readLineInto", args)
  }

}

  // proto:  void QTextStream::setRealNumberPrecision(int precision);
func (this *QTextStream) setRealNumberPrecision(args ...interface{}) () {
  // setRealNumberPrecision(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream22setRealNumberPrecisionEi
    // invoke: void setRealNumberPrecision(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN11QTextStream22setRealNumberPrecisionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "setRealNumberPrecision", args)
  }

}

  // proto:  void QTextStream::setDevice(QIODevice * device);
func (this *QTextStream) setDevice(args ...interface{}) () {
  // setDevice(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream9setDeviceEP9QIODevice
    // invoke: void setDevice(class QIODevice *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QTextStream9setDeviceEP9QIODevice(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "setDevice", args)
  }

}

  // proto:  void QTextStream::reset();
func (this *QTextStream) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream5resetEv
    // invoke: void reset()
    C._ZN11QTextStream5resetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "reset", args)
  }

}

  // proto:  bool QTextStream::seek(qint64 pos);
func (this *QTextStream) seek(args ...interface{}) () {
  // seek(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream4seekEx
    // invoke: bool seek(qint64)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    C._ZN11QTextStream4seekEx(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "seek", args)
  }

}

  // proto:  QString * QTextStream::string();
func (this *QTextStream) string(args ...interface{}) () {
  // string()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream6stringEv
    // invoke: QString * string()
    C._ZNK11QTextStream6stringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "string", args)
  }

}

  // proto:  void QTextStream::setAutoDetectUnicode(bool enabled);
func (this *QTextStream) setAutoDetectUnicode(args ...interface{}) () {
  // setAutoDetectUnicode(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream20setAutoDetectUnicodeEb
    // invoke: void setAutoDetectUnicode(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN11QTextStream20setAutoDetectUnicodeEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "setAutoDetectUnicode", args)
  }

}

  // proto:  QChar QTextStream::padChar();
func (this *QTextStream) padChar(args ...interface{}) () {
  // padChar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream7padCharEv
    // invoke: QChar padChar()
    C._ZNK11QTextStream7padCharEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "padChar", args)
  }

}

  // proto:  QIODevice * QTextStream::device();
func (this *QTextStream) device(args ...interface{}) () {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream6deviceEv
    // invoke: QIODevice * device()
    C._ZNK11QTextStream6deviceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "device", args)
  }

}

  // proto:  void QTextStream::resetStatus();
func (this *QTextStream) resetStatus(args ...interface{}) () {
  // resetStatus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream11resetStatusEv
    // invoke: void resetStatus()
    C._ZN11QTextStream11resetStatusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "resetStatus", args)
  }

}

  // proto:  bool QTextStream::autoDetectUnicode();
func (this *QTextStream) autoDetectUnicode(args ...interface{}) () {
  // autoDetectUnicode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream17autoDetectUnicodeEv
    // invoke: bool autoDetectUnicode()
    C._ZNK11QTextStream17autoDetectUnicodeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "autoDetectUnicode", args)
  }

}

  // proto:  int QTextStream::fieldWidth();
func (this *QTextStream) fieldWidth(args ...interface{}) () {
  // fieldWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream10fieldWidthEv
    // invoke: int fieldWidth()
    C._ZNK11QTextStream10fieldWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "fieldWidth", args)
  }

}

  // proto:  bool QTextStream::generateByteOrderMark();
func (this *QTextStream) generateByteOrderMark(args ...interface{}) () {
  // generateByteOrderMark()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream21generateByteOrderMarkEv
    // invoke: bool generateByteOrderMark()
    C._ZNK11QTextStream21generateByteOrderMarkEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "generateByteOrderMark", args)
  }

}

  // proto:  void QTextStream::setGenerateByteOrderMark(bool generate);
func (this *QTextStream) setGenerateByteOrderMark(args ...interface{}) () {
  // setGenerateByteOrderMark(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream24setGenerateByteOrderMarkEb
    // invoke: void setGenerateByteOrderMark(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN11QTextStream24setGenerateByteOrderMarkEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "setGenerateByteOrderMark", args)
  }

}

  // proto:  void QTextStream::setCodec(QTextCodec * codec);
func (this *QTextStream) setCodec(args ...interface{}) () {
  // setCodec(class QTextCodec *)
  // setCodec(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream8setCodecEP10QTextCodec
    // invoke: void setCodec(class QTextCodec *)
    var arg0 = args[0].(QTextCodec).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QTextStream8setCodecEP10QTextCodec(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN11QTextStream8setCodecEPKc
    // invoke: void setCodec(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C._ZN11QTextStream8setCodecEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "setCodec", args)
  }

}

  // proto:  void QTextStream::flush();
func (this *QTextStream) flush(args ...interface{}) () {
  // flush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream5flushEv
    // invoke: void flush()
    C._ZN11QTextStream5flushEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "flush", args)
  }

}

  // proto:  void QTextStream::setIntegerBase(int base);
func (this *QTextStream) setIntegerBase(args ...interface{}) () {
  // setIntegerBase(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream14setIntegerBaseEi
    // invoke: void setIntegerBase(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN11QTextStream14setIntegerBaseEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "setIntegerBase", args)
  }

}

  // proto:  void QTextStream::~QTextStream();
func (this *QTextStream) FreeQTextStream(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextStream", "~QTextStream", args)
  }

}

  // proto:  QLocale QTextStream::locale();
func (this *QTextStream) locale(args ...interface{}) () {
  // locale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream6localeEv
    // invoke: QLocale locale()
    C._ZNK11QTextStream6localeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "locale", args)
  }

}

  // proto:  QString QTextStream::read(qint64 maxlen);
func (this *QTextStream) read(args ...interface{}) () {
  // read(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream4readEx
    // invoke: QString read(qint64)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    C._ZN11QTextStream4readEx(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "read", args)
  }

}

  // proto:  void QTextStream::setPadChar(QChar ch);
func (this *QTextStream) setPadChar(args ...interface{}) () {
  // setPadChar(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream10setPadCharE5QChar
    // invoke: void setPadChar(class QChar)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QTextStream10setPadCharE5QChar(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "setPadChar", args)
  }

}

  // proto:  int QTextStream::realNumberPrecision();
func (this *QTextStream) realNumberPrecision(args ...interface{}) () {
  // realNumberPrecision()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream19realNumberPrecisionEv
    // invoke: int realNumberPrecision()
    C._ZNK11QTextStream19realNumberPrecisionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "realNumberPrecision", args)
  }

}

  // proto:  qint64 QTextStream::pos();
func (this *QTextStream) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream3posEv
    // invoke: qint64 pos()
    C._ZNK11QTextStream3posEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "pos", args)
  }

}

  // proto:  QString QTextStream::readAll();
func (this *QTextStream) readAll(args ...interface{}) () {
  // readAll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream7readAllEv
    // invoke: QString readAll()
    C._ZN11QTextStream7readAllEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "readAll", args)
  }

}

  // proto:  void QTextStream::skipWhiteSpace();
func (this *QTextStream) skipWhiteSpace(args ...interface{}) () {
  // skipWhiteSpace()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream14skipWhiteSpaceEv
    // invoke: void skipWhiteSpace()
    C._ZN11QTextStream14skipWhiteSpaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "skipWhiteSpace", args)
  }

}

  // proto:  void QTextStream::setFieldWidth(int width);
func (this *QTextStream) setFieldWidth(args ...interface{}) () {
  // setFieldWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream13setFieldWidthEi
    // invoke: void setFieldWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN11QTextStream13setFieldWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "setFieldWidth", args)
  }

}

  // proto:  int QTextStream::integerBase();
func (this *QTextStream) integerBase(args ...interface{}) () {
  // integerBase()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream11integerBaseEv
    // invoke: int integerBase()
    C._ZNK11QTextStream11integerBaseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "integerBase", args)
  }

}

  // proto:  QString QTextStream::readLine(qint64 maxlen);
func (this *QTextStream) readLine(args ...interface{}) () {
  // readLine(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream8readLineEx
    // invoke: QString readLine(qint64)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    C._ZN11QTextStream8readLineEx(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "readLine", args)
  }

}

// <= body block end
