package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtCore/qreadwritelock.h
// dst-file: /src/core/qreadwritelock.go
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
  // proto:  QReadWriteLock * QWriteLocker::readWriteLock();
extern void demth_ZNK12QWriteLocker13readWriteLockEv(void* qthis);
  // proto:  void QWriteLocker::QWriteLocker(QReadWriteLock * readWriteLock);
extern void* dector_ZN12QWriteLockerC1EP14QReadWriteLock(void* arg0);
extern void demth_ZN12QWriteLockerC1EP14QReadWriteLock(void* qthis, void* arg0);
  // proto:  void QWriteLocker::unlock();
extern void demth_ZN12QWriteLocker6unlockEv(void* qthis);
  // proto:  void QWriteLocker::~QWriteLocker();
extern void demth_ZN12QWriteLockerD0Ev(void* qthis);
  // proto:  void QWriteLocker::relock();
extern void demth_ZN12QWriteLocker6relockEv(void* qthis);
  // proto:  void QWriteLocker::QWriteLocker(const QWriteLocker & );
extern void* dector_ZN12QWriteLockerC1ERKS_(void* arg0);
extern void _ZN12QWriteLockerC1ERKS_(void* qthis, void* arg0);
  // proto:  void QReadWriteLock::~QReadWriteLock();
extern void _ZN14QReadWriteLockD0Ev(void* qthis);
  // proto:  void QReadWriteLock::QReadWriteLock(const QReadWriteLock & );
extern void* dector_ZN14QReadWriteLockC1ERKS_(void* arg0);
extern void _ZN14QReadWriteLockC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QReadWriteLock::tryLockForRead();
extern void _ZN14QReadWriteLock14tryLockForReadEv(void* qthis);
  // proto:  void QReadWriteLock::lockForWrite();
extern void _ZN14QReadWriteLock12lockForWriteEv(void* qthis);
  // proto:  bool QReadWriteLock::tryLockForWrite();
extern void _ZN14QReadWriteLock15tryLockForWriteEv(void* qthis);
  // proto:  void QReadWriteLock::unlock();
extern void _ZN14QReadWriteLock6unlockEv(void* qthis);
  // proto:  bool QReadWriteLock::tryLockForRead(int timeout);
extern void _ZN14QReadWriteLock14tryLockForReadEi(void* qthis, int32_t arg0);
  // proto:  void QReadWriteLock::lockForRead();
extern void _ZN14QReadWriteLock11lockForReadEv(void* qthis);
  // proto:  bool QReadWriteLock::tryLockForWrite(int timeout);
extern void _ZN14QReadWriteLock15tryLockForWriteEi(void* qthis, int32_t arg0);
  // proto:  QReadWriteLock * QReadLocker::readWriteLock();
extern void demth_ZNK11QReadLocker13readWriteLockEv(void* qthis);
  // proto:  void QReadLocker::~QReadLocker();
extern void demth_ZN11QReadLockerD0Ev(void* qthis);
  // proto:  void QReadLocker::QReadLocker(QReadWriteLock * readWriteLock);
extern void* dector_ZN11QReadLockerC1EP14QReadWriteLock(void* arg0);
extern void demth_ZN11QReadLockerC1EP14QReadWriteLock(void* qthis, void* arg0);
  // proto:  void QReadLocker::QReadLocker(const QReadLocker & );
extern void* dector_ZN11QReadLockerC1ERKS_(void* arg0);
extern void _ZN11QReadLockerC1ERKS_(void* qthis, void* arg0);
  // proto:  void QReadLocker::relock();
extern void demth_ZN11QReadLocker6relockEv(void* qthis);
  // proto:  void QReadLocker::unlock();
extern void demth_ZN11QReadLocker6unlockEv(void* qthis);
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

// class sizeof(QWriteLocker)=4
type QWriteLocker struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QReadWriteLock)=8
type QReadWriteLock struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QReadLocker)=4
type QReadLocker struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  QReadWriteLock * QWriteLocker::readWriteLock();
func (this *QWriteLocker) readWriteLock(args ...interface{}) () {
  // readWriteLock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QWriteLocker13readWriteLockEv
    // invoke: QReadWriteLock * readWriteLock()
    C.demth_ZNK12QWriteLocker13readWriteLockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWriteLocker", "readWriteLock", args)
  }

}

  // proto:  void QWriteLocker::QWriteLocker(QReadWriteLock * readWriteLock);
func NewQWriteLocker(args ...interface{}) QWriteLocker {
  return QWriteLocker{}
}

  // proto:  void QWriteLocker::unlock();
func (this *QWriteLocker) unlock(args ...interface{}) () {
  // unlock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QWriteLocker6unlockEv
    // invoke: void unlock()
    C.demth_ZN12QWriteLocker6unlockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWriteLocker", "unlock", args)
  }

}

  // proto:  void QWriteLocker::~QWriteLocker();
func (this *QWriteLocker) FreeQWriteLocker(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWriteLocker", "~QWriteLocker", args)
  }

}

  // proto:  void QWriteLocker::relock();
func (this *QWriteLocker) relock(args ...interface{}) () {
  // relock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QWriteLocker6relockEv
    // invoke: void relock()
    C.demth_ZN12QWriteLocker6relockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWriteLocker", "relock", args)
  }

}

  // proto:  void QReadWriteLock::~QReadWriteLock();
func (this *QReadWriteLock) FreeQReadWriteLock(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QReadWriteLock", "~QReadWriteLock", args)
  }

}

  // proto:  void QReadWriteLock::QReadWriteLock(const QReadWriteLock & );
func NewQReadWriteLock(args ...interface{}) QReadWriteLock {
  return QReadWriteLock{}
}

  // proto:  bool QReadWriteLock::tryLockForRead();
func (this *QReadWriteLock) tryLockForRead(args ...interface{}) () {
  // tryLockForRead()
  // tryLockForRead(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QReadWriteLock14tryLockForReadEv
    // invoke: bool tryLockForRead()
    C._ZN14QReadWriteLock14tryLockForReadEv(this.qclsinst)
  case 1:
    // invoke: _ZN14QReadWriteLock14tryLockForReadEi
    // invoke: bool tryLockForRead(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN14QReadWriteLock14tryLockForReadEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QReadWriteLock", "tryLockForRead", args)
  }

}

  // proto:  void QReadWriteLock::lockForWrite();
func (this *QReadWriteLock) lockForWrite(args ...interface{}) () {
  // lockForWrite()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QReadWriteLock12lockForWriteEv
    // invoke: void lockForWrite()
    C._ZN14QReadWriteLock12lockForWriteEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QReadWriteLock", "lockForWrite", args)
  }

}

  // proto:  bool QReadWriteLock::tryLockForWrite();
func (this *QReadWriteLock) tryLockForWrite(args ...interface{}) () {
  // tryLockForWrite()
  // tryLockForWrite(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QReadWriteLock15tryLockForWriteEv
    // invoke: bool tryLockForWrite()
    C._ZN14QReadWriteLock15tryLockForWriteEv(this.qclsinst)
  case 1:
    // invoke: _ZN14QReadWriteLock15tryLockForWriteEi
    // invoke: bool tryLockForWrite(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN14QReadWriteLock15tryLockForWriteEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QReadWriteLock", "tryLockForWrite", args)
  }

}

  // proto:  void QReadWriteLock::unlock();
func (this *QReadWriteLock) unlock(args ...interface{}) () {
  // unlock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QReadWriteLock6unlockEv
    // invoke: void unlock()
    C._ZN14QReadWriteLock6unlockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QReadWriteLock", "unlock", args)
  }

}

  // proto:  void QReadWriteLock::lockForRead();
func (this *QReadWriteLock) lockForRead(args ...interface{}) () {
  // lockForRead()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QReadWriteLock11lockForReadEv
    // invoke: void lockForRead()
    C._ZN14QReadWriteLock11lockForReadEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QReadWriteLock", "lockForRead", args)
  }

}

  // proto:  QReadWriteLock * QReadLocker::readWriteLock();
func (this *QReadLocker) readWriteLock(args ...interface{}) () {
  // readWriteLock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QReadLocker13readWriteLockEv
    // invoke: QReadWriteLock * readWriteLock()
    C.demth_ZNK11QReadLocker13readWriteLockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QReadLocker", "readWriteLock", args)
  }

}

  // proto:  void QReadLocker::~QReadLocker();
func (this *QReadLocker) FreeQReadLocker(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QReadLocker", "~QReadLocker", args)
  }

}

  // proto:  void QReadLocker::QReadLocker(QReadWriteLock * readWriteLock);
func NewQReadLocker(args ...interface{}) QReadLocker {
  return QReadLocker{}
}

  // proto:  void QReadLocker::relock();
func (this *QReadLocker) relock(args ...interface{}) () {
  // relock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QReadLocker6relockEv
    // invoke: void relock()
    C.demth_ZN11QReadLocker6relockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QReadLocker", "relock", args)
  }

}

  // proto:  void QReadLocker::unlock();
func (this *QReadLocker) unlock(args ...interface{}) () {
  // unlock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QReadLocker6unlockEv
    // invoke: void unlock()
    C.demth_ZN11QReadLocker6unlockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QReadLocker", "unlock", args)
  }

}

// <= body block end
