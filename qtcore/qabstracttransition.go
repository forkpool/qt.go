package qtcore

// /usr/include/qt/QtCore/qabstracttransition.h
// #include <qabstracttransition.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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
}

//  ext block end

//  body block begin
// bool eventTest(class QEvent *)
func (this *QAbstractTransition) InheritEventTest(f func(event *QEvent /*777 QEvent **/) bool) {
	ffiqt.SetAllInheritCallback(this, "eventTest", f)
}

// void onTransition(class QEvent *)
func (this *QAbstractTransition) InheritOnTransition(f func(event *QEvent /*777 QEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "onTransition", f)
}

// bool event(class QEvent *)
func (this *QAbstractTransition) InheritEvent(f func(e *QEvent /*777 QEvent **/) bool) {
	ffiqt.SetAllInheritCallback(this, "event", f)
}

type QAbstractTransition struct {
	*QObject
}

func (this *QAbstractTransition) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAbstractTransition) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQAbstractTransitionFromPointer(cthis unsafe.Pointer) *QAbstractTransition {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QAbstractTransition{bcthis0}
}
func (*QAbstractTransition) NewFromPointer(cthis unsafe.Pointer) *QAbstractTransition {
	return NewQAbstractTransitionFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstracttransition.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QAbstractTransition) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractTransition10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qabstracttransition.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractTransition(QState *)
func NewQAbstractTransition(sourceState *QState /*777 QState **/) *QAbstractTransition {
	var convArg0 = sourceState.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTransitionC1EP6QState", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractTransitionFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qabstracttransition.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractTransition()
func DeleteQAbstractTransition(this *QAbstractTransition) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTransitionD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qabstracttransition.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QState * sourceState()
func (this *QAbstractTransition) SourceState() *QState /*777 QState **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractTransition11sourceStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qabstracttransition.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractState * targetState()
func (this *QAbstractTransition) TargetState() *QAbstractState /*777 QAbstractState **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractTransition11targetStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractStateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qabstracttransition.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTargetState(QAbstractState *)
func (this *QAbstractTransition) SetTargetState(target *QAbstractState /*777 QAbstractState **/) {
	var convArg0 = target.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTransition14setTargetStateEP14QAbstractState", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracttransition.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractTransition::TransitionType transitionType()
func (this *QAbstractTransition) TransitionType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractTransition14transitionTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qabstracttransition.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransitionType(enum QAbstractTransition::TransitionType)
func (this *QAbstractTransition) SetTransitionType(type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTransition17setTransitionTypeENS_14TransitionTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracttransition.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] QStateMachine * machine()
func (this *QAbstractTransition) Machine() *QStateMachine /*777 QStateMachine **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractTransition7machineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStateMachineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qabstracttransition.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addAnimation(QAbstractAnimation *)
func (this *QAbstractTransition) AddAnimation(animation *QAbstractAnimation /*777 QAbstractAnimation **/) {
	var convArg0 = animation.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTransition12addAnimationEP18QAbstractAnimation", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracttransition.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeAnimation(QAbstractAnimation *)
func (this *QAbstractTransition) RemoveAnimation(animation *QAbstractAnimation /*777 QAbstractAnimation **/) {
	var convArg0 = animation.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTransition15removeAnimationEP18QAbstractAnimation", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracttransition.h:101
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [1] bool eventTest(QEvent *)
func (this *QAbstractTransition) EventTest(event *QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTransition9eventTestEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstracttransition.h:103
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [-2] void onTransition(QEvent *)
func (this *QAbstractTransition) OnTransition(event *QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTransition12onTransitionEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracttransition.h:105
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QAbstractTransition) Event(e *QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTransition5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QAbstractTransition__TransitionType = int

const QAbstractTransition__ExternalTransition QAbstractTransition__TransitionType = 0
const QAbstractTransition__InternalTransition QAbstractTransition__TransitionType = 1

//  body block end