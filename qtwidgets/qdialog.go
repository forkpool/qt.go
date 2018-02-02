package qtwidgets

// /usr/include/qt/QtWidgets/qdialog.h
// #include <qdialog.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
// void keyPressEvent(class QKeyEvent *)
func (this *QDialog) InheritKeyPressEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void closeEvent(class QCloseEvent *)
func (this *QDialog) InheritCloseEvent(f func(arg0 *qtgui.QCloseEvent /*777 QCloseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "closeEvent", f)
}

// void showEvent(class QShowEvent *)
func (this *QDialog) InheritShowEvent(f func(arg0 *qtgui.QShowEvent /*777 QShowEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "showEvent", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QDialog) InheritResizeEvent(f func(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void contextMenuEvent(class QContextMenuEvent *)
func (this *QDialog) InheritContextMenuEvent(f func(arg0 *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// bool eventFilter(class QObject *, class QEvent *)
func (this *QDialog) InheritEventFilter(f func(arg0 *qtcore.QObject /*777 QObject **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool) {
	ffiqt.SetAllInheritCallback(this, "eventFilter", f)
}

// void adjustPosition(class QWidget *)
func (this *QDialog) InheritAdjustPosition(f func(arg0 *QWidget /*777 QWidget **/)) {
	ffiqt.SetAllInheritCallback(this, "adjustPosition", f)
}

type QDialog struct {
	*QWidget
}

func (this *QDialog) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QDialog) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQDialogFromPointer(cthis unsafe.Pointer) *QDialog {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QDialog{bcthis0}
}
func (*QDialog) NewFromPointer(cthis unsafe.Pointer) *QDialog {
	return NewQDialogFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qdialog.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QDialog) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QDialog10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qdialog.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDialog(QWidget *, Qt::WindowFlags)
func NewQDialog(parent *QWidget /*777 QWidget **/, f int) *QDialog {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialogC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_POINTER, convArg0, f)
	gopp.ErrPrint(err, rv)
	gothis := NewQDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qdialog.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDialog()
func DeleteQDialog(this *QDialog) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialogD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qdialog.h:68
// index:0
// Public Visibility=Default Availability=Available
// [4] int result()
func (this *QDialog) Result() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QDialog6resultEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdialog.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setVisible(_Bool)
func (this *QDialog) SetVisible(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog10setVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOrientation(Qt::Orientation)
func (this *QDialog) SetOrientation(orientation int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:73
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Orientation orientation()
func (this *QDialog) Orientation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QDialog11orientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExtension(QWidget *)
func (this *QDialog) SetExtension(extension *QWidget /*777 QWidget **/) {
	var convArg0 = extension.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog12setExtensionEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * extension()
func (this *QDialog) Extension() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QDialog9extensionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qdialog.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QDialog) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QDialog8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qdialog.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
func (this *QDialog) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QDialog15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qdialog.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSizeGripEnabled(_Bool)
func (this *QDialog) SetSizeGripEnabled(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog18setSizeGripEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSizeGripEnabled()
func (this *QDialog) IsSizeGripEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QDialog17isSizeGripEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdialog.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModal(_Bool)
func (this *QDialog) SetModal(modal bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog8setModalEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), modal)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setResult(int)
func (this *QDialog) SetResult(r int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog9setResultEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void finished(int)
func (this *QDialog) Finished(result int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog8finishedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), result)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void accepted()
func (this *QDialog) Accepted() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog8acceptedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rejected()
func (this *QDialog) Rejected() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog8rejectedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:93
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void open()
func (this *QDialog) Open() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog4openEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:94
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int exec()
func (this *QDialog) Exec() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog4execEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdialog.h:95
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void done(int)
func (this *QDialog) Done(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog4doneEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void accept()
func (this *QDialog) Accept() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog6acceptEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:97
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void reject()
func (this *QDialog) Reject() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog6rejectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showExtension(_Bool)
func (this *QDialog) ShowExtension(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog13showExtensionEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:104
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QDialog) KeyPressEvent(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:105
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void closeEvent(QCloseEvent *)
func (this *QDialog) CloseEvent(arg0 *qtgui.QCloseEvent /*777 QCloseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog10closeEventEP11QCloseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:106
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)
func (this *QDialog) ShowEvent(arg0 *qtgui.QShowEvent /*777 QShowEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:107
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QDialog) ResizeEvent(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:109
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)
func (this *QDialog) ContextMenuEvent(arg0 *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdialog.h:111
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)
func (this *QDialog) EventFilter(arg0 *qtcore.QObject /*777 QObject **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdialog.h:112
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void adjustPosition(QWidget *)
func (this *QDialog) AdjustPosition(arg0 *QWidget /*777 QWidget **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QDialog14adjustPositionEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QDialog__DialogCode = int

const QDialog__Rejected QDialog__DialogCode = 0
const QDialog__Accepted QDialog__DialogCode = 1

//  body block end