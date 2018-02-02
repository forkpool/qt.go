package qtwidgets

// /usr/include/qt/QtWidgets/qapplication.h
// #include <qapplication.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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
// bool event(class QEvent *)
func (this *QApplication) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	ffiqt.SetAllInheritCallback(this, "event", f)
}

type QApplication struct {
	*qtgui.QGuiApplication
}

func (this *QApplication) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGuiApplication.GetCthis()
	}
}
func (this *QApplication) SetCthis(cthis unsafe.Pointer) {
	this.QGuiApplication = qtgui.NewQGuiApplicationFromPointer(cthis)
}
func NewQApplicationFromPointer(cthis unsafe.Pointer) *QApplication {
	bcthis0 := qtgui.NewQGuiApplicationFromPointer(cthis)
	return &QApplication{bcthis0}
}
func (*QApplication) NewFromPointer(cthis unsafe.Pointer) *QApplication {
	return NewQApplicationFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qapplication.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QApplication) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QApplication10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qapplication.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QApplication(int &, char **, int)
func NewQApplication(argc int, argv []string, arg2 int) *QApplication {
	var convArg1 = qtrt.StringSliceToCCharPP(argv)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplicationC2ERiPPci", ffiqt.FFI_TYPE_POINTER, &argc, convArg1, arg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQApplicationFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qapplication.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QApplication()
func DeleteQApplication(this *QApplication) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplicationD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qapplication.h:98
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStyle * style()
func (this *QApplication) Style() *QStyle /*777 QStyle **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication5styleEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStyleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QApplication_Style() *QStyle /*777 QStyle **/ {
	var nilthis *QApplication
	rv := nilthis.Style()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:99
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setStyle(QStyle *)
func (this *QApplication) SetStyle(arg0 *QStyle /*777 QStyle **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication8setStyleEP6QStyle", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetStyle(arg0 *QStyle /*777 QStyle **/) {
	var nilthis *QApplication
	nilthis.SetStyle(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:100
// index:1
// Public static Visibility=Default Availability=Available
// [8] QStyle * setStyle(const QString &)
func (this *QApplication) SetStyle_1(arg0 *qtcore.QString) *QStyle /*777 QStyle **/ {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication8setStyleERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStyleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QApplication_SetStyle_1(arg0 *qtcore.QString) *QStyle /*777 QStyle **/ {
	var nilthis *QApplication
	rv := nilthis.SetStyle_1(arg0)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:103
// index:0
// Public static Visibility=Default Availability=Available
// [4] int colorSpec()
func (this *QApplication) ColorSpec() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication9colorSpecEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QApplication_ColorSpec() int {
	var nilthis *QApplication
	rv := nilthis.ColorSpec()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:104
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setColorSpec(int)
func (this *QApplication) SetColorSpec(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication12setColorSpecEi", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetColorSpec(arg0 int) {
	var nilthis *QApplication
	nilthis.SetColorSpec(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:111
// index:0
// Public static Visibility=Default Availability=Available
// [16] QPalette palette(const QWidget *)
func (this *QApplication) Palette(arg0 *QWidget /*777 const QWidget **/) *qtgui.QPalette /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication7paletteEPK7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtgui.NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPalette)
	return rv2
}
func QApplication_Palette(arg0 *QWidget /*777 const QWidget **/) *qtgui.QPalette /*123*/ {
	var nilthis *QApplication
	rv := nilthis.Palette(arg0)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:112
// index:1
// Public static Visibility=Default Availability=Available
// [16] QPalette palette(const char *)
func (this *QApplication) Palette_1(className string) *qtgui.QPalette /*123*/ {
	var convArg0 = qtrt.CString(className)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication7paletteEPKc", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtgui.NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPalette)
	return rv2
}
func QApplication_Palette_1(className string) *qtgui.QPalette /*123*/ {
	var nilthis *QApplication
	rv := nilthis.Palette_1(className)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:113
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setPalette(const QPalette &, const char *)
func (this *QApplication) SetPalette(arg0 *qtgui.QPalette, className string) {
	var convArg0 = arg0.GetCthis()
	var convArg1 = qtrt.CString(className)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication10setPaletteERK8QPalettePKc", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetPalette(arg0 *qtgui.QPalette, className string) {
	var nilthis *QApplication
	nilthis.SetPalette(arg0, className)
}

// /usr/include/qt/QtWidgets/qapplication.h:114
// index:0
// Public static Visibility=Default Availability=Available
// [16] QFont font()
func (this *QApplication) Font() *qtgui.QFont /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication4fontEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}
func QApplication_Font() *qtgui.QFont /*123*/ {
	var nilthis *QApplication
	rv := nilthis.Font()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:115
// index:1
// Public static Visibility=Default Availability=Available
// [16] QFont font(const QWidget *)
func (this *QApplication) Font_1(arg0 *QWidget /*777 const QWidget **/) *qtgui.QFont /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication4fontEPK7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}
func QApplication_Font_1(arg0 *QWidget /*777 const QWidget **/) *qtgui.QFont /*123*/ {
	var nilthis *QApplication
	rv := nilthis.Font_1(arg0)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:116
// index:2
// Public static Visibility=Default Availability=Available
// [16] QFont font(const char *)
func (this *QApplication) Font_2(className string) *qtgui.QFont /*123*/ {
	var convArg0 = qtrt.CString(className)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication4fontEPKc", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}
func QApplication_Font_2(className string) *qtgui.QFont /*123*/ {
	var nilthis *QApplication
	rv := nilthis.Font_2(className)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:117
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setFont(const QFont &, const char *)
func (this *QApplication) SetFont(arg0 *qtgui.QFont, className string) {
	var convArg0 = arg0.GetCthis()
	var convArg1 = qtrt.CString(className)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication7setFontERK5QFontPKc", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetFont(arg0 *qtgui.QFont, className string) {
	var nilthis *QApplication
	nilthis.SetFont(arg0, className)
}

// /usr/include/qt/QtWidgets/qapplication.h:118
// index:0
// Public static Visibility=Default Availability=Available
// [8] QFontMetrics fontMetrics()
func (this *QApplication) FontMetrics() *qtgui.QFontMetrics /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication11fontMetricsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtgui.NewQFontMetricsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFontMetrics)
	return rv2
}
func QApplication_FontMetrics() *qtgui.QFontMetrics /*123*/ {
	var nilthis *QApplication
	rv := nilthis.FontMetrics()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:121
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setWindowIcon(const QIcon &)
func (this *QApplication) SetWindowIcon(icon *qtgui.QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication13setWindowIconERK5QIcon", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetWindowIcon(icon *qtgui.QIcon) {
	var nilthis *QApplication
	nilthis.SetWindowIcon(icon)
}

// /usr/include/qt/QtWidgets/qapplication.h:122
// index:0
// Public static Visibility=Default Availability=Available
// [8] QIcon windowIcon()
func (this *QApplication) WindowIcon() *qtgui.QIcon /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication10windowIconEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}
func QApplication_WindowIcon() *qtgui.QIcon /*123*/ {
	var nilthis *QApplication
	rv := nilthis.WindowIcon()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:128
// index:0
// Public static Visibility=Default Availability=Available
// [8] QDesktopWidget * desktop()
func (this *QApplication) Desktop() *QDesktopWidget /*777 QDesktopWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication7desktopEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDesktopWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QApplication_Desktop() *QDesktopWidget /*777 QDesktopWidget **/ {
	var nilthis *QApplication
	rv := nilthis.Desktop()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:130
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWidget * activePopupWidget()
func (this *QApplication) ActivePopupWidget() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication17activePopupWidgetEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QApplication_ActivePopupWidget() *QWidget /*777 QWidget **/ {
	var nilthis *QApplication
	rv := nilthis.ActivePopupWidget()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:131
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWidget * activeModalWidget()
func (this *QApplication) ActiveModalWidget() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication17activeModalWidgetEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QApplication_ActiveModalWidget() *QWidget /*777 QWidget **/ {
	var nilthis *QApplication
	rv := nilthis.ActiveModalWidget()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:132
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWidget * focusWidget()
func (this *QApplication) FocusWidget() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication11focusWidgetEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QApplication_FocusWidget() *QWidget /*777 QWidget **/ {
	var nilthis *QApplication
	rv := nilthis.FocusWidget()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:134
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWidget * activeWindow()
func (this *QApplication) ActiveWindow() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication12activeWindowEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QApplication_ActiveWindow() *QWidget /*777 QWidget **/ {
	var nilthis *QApplication
	rv := nilthis.ActiveWindow()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:135
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setActiveWindow(QWidget *)
func (this *QApplication) SetActiveWindow(act *QWidget /*777 QWidget **/) {
	var convArg0 = act.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication15setActiveWindowEP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetActiveWindow(act *QWidget /*777 QWidget **/) {
	var nilthis *QApplication
	nilthis.SetActiveWindow(act)
}

// /usr/include/qt/QtWidgets/qapplication.h:137
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWidget * widgetAt(const QPoint &)
func (this *QApplication) WidgetAt(p *qtcore.QPoint) *QWidget /*777 QWidget **/ {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication8widgetAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QApplication_WidgetAt(p *qtcore.QPoint) *QWidget /*777 QWidget **/ {
	var nilthis *QApplication
	rv := nilthis.WidgetAt(p)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:138
// index:1
// Public static inline Visibility=Default Availability=Available
// [8] QWidget * widgetAt(int, int)
func (this *QApplication) WidgetAt_1(x int, y int) *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication8widgetAtEii", ffiqt.FFI_TYPE_POINTER, x, y)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QApplication_WidgetAt_1(x int, y int) *QWidget /*777 QWidget **/ {
	var nilthis *QApplication
	rv := nilthis.WidgetAt_1(x, y)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:139
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWidget * topLevelAt(const QPoint &)
func (this *QApplication) TopLevelAt(p *qtcore.QPoint) *QWidget /*777 QWidget **/ {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication10topLevelAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QApplication_TopLevelAt(p *qtcore.QPoint) *QWidget /*777 QWidget **/ {
	var nilthis *QApplication
	rv := nilthis.TopLevelAt(p)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:140
// index:1
// Public static inline Visibility=Default Availability=Available
// [8] QWidget * topLevelAt(int, int)
func (this *QApplication) TopLevelAt_1(x int, y int) *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication10topLevelAtEii", ffiqt.FFI_TYPE_POINTER, x, y)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QApplication_TopLevelAt_1(x int, y int) *QWidget /*777 QWidget **/ {
	var nilthis *QApplication
	rv := nilthis.TopLevelAt_1(x, y)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:145
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void beep()
func (this *QApplication) Beep() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication4beepEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QApplication_Beep() {
	var nilthis *QApplication
	nilthis.Beep()
}

// /usr/include/qt/QtWidgets/qapplication.h:146
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void alert(QWidget *, int)
func (this *QApplication) Alert(widget *QWidget /*777 QWidget **/, duration int) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication5alertEP7QWidgeti", ffiqt.FFI_TYPE_POINTER, convArg0, duration)
	gopp.ErrPrint(err, rv)
}
func QApplication_Alert(widget *QWidget /*777 QWidget **/, duration int) {
	var nilthis *QApplication
	nilthis.Alert(widget, duration)
}

// /usr/include/qt/QtWidgets/qapplication.h:148
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setCursorFlashTime(int)
func (this *QApplication) SetCursorFlashTime(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication18setCursorFlashTimeEi", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetCursorFlashTime(arg0 int) {
	var nilthis *QApplication
	nilthis.SetCursorFlashTime(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:149
// index:0
// Public static Visibility=Default Availability=Available
// [4] int cursorFlashTime()
func (this *QApplication) CursorFlashTime() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication15cursorFlashTimeEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QApplication_CursorFlashTime() int {
	var nilthis *QApplication
	rv := nilthis.CursorFlashTime()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:151
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setDoubleClickInterval(int)
func (this *QApplication) SetDoubleClickInterval(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication22setDoubleClickIntervalEi", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetDoubleClickInterval(arg0 int) {
	var nilthis *QApplication
	nilthis.SetDoubleClickInterval(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:152
// index:0
// Public static Visibility=Default Availability=Available
// [4] int doubleClickInterval()
func (this *QApplication) DoubleClickInterval() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication19doubleClickIntervalEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QApplication_DoubleClickInterval() int {
	var nilthis *QApplication
	rv := nilthis.DoubleClickInterval()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:154
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setKeyboardInputInterval(int)
func (this *QApplication) SetKeyboardInputInterval(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication24setKeyboardInputIntervalEi", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetKeyboardInputInterval(arg0 int) {
	var nilthis *QApplication
	nilthis.SetKeyboardInputInterval(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:155
// index:0
// Public static Visibility=Default Availability=Available
// [4] int keyboardInputInterval()
func (this *QApplication) KeyboardInputInterval() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication21keyboardInputIntervalEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QApplication_KeyboardInputInterval() int {
	var nilthis *QApplication
	rv := nilthis.KeyboardInputInterval()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:158
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setWheelScrollLines(int)
func (this *QApplication) SetWheelScrollLines(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication19setWheelScrollLinesEi", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetWheelScrollLines(arg0 int) {
	var nilthis *QApplication
	nilthis.SetWheelScrollLines(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:159
// index:0
// Public static Visibility=Default Availability=Available
// [4] int wheelScrollLines()
func (this *QApplication) WheelScrollLines() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication16wheelScrollLinesEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QApplication_WheelScrollLines() int {
	var nilthis *QApplication
	rv := nilthis.WheelScrollLines()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:161
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setGlobalStrut(const QSize &)
func (this *QApplication) SetGlobalStrut(arg0 *qtcore.QSize) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication14setGlobalStrutERK5QSize", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetGlobalStrut(arg0 *qtcore.QSize) {
	var nilthis *QApplication
	nilthis.SetGlobalStrut(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:162
// index:0
// Public static Visibility=Default Availability=Available
// [8] QSize globalStrut()
func (this *QApplication) GlobalStrut() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication11globalStrutEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}
func QApplication_GlobalStrut() *qtcore.QSize /*123*/ {
	var nilthis *QApplication
	rv := nilthis.GlobalStrut()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:164
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setStartDragTime(int)
func (this *QApplication) SetStartDragTime(ms int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication16setStartDragTimeEi", ffiqt.FFI_TYPE_POINTER, ms)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetStartDragTime(ms int) {
	var nilthis *QApplication
	nilthis.SetStartDragTime(ms)
}

// /usr/include/qt/QtWidgets/qapplication.h:165
// index:0
// Public static Visibility=Default Availability=Available
// [4] int startDragTime()
func (this *QApplication) StartDragTime() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication13startDragTimeEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QApplication_StartDragTime() int {
	var nilthis *QApplication
	rv := nilthis.StartDragTime()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:166
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setStartDragDistance(int)
func (this *QApplication) SetStartDragDistance(l int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication20setStartDragDistanceEi", ffiqt.FFI_TYPE_POINTER, l)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetStartDragDistance(l int) {
	var nilthis *QApplication
	nilthis.SetStartDragDistance(l)
}

// /usr/include/qt/QtWidgets/qapplication.h:167
// index:0
// Public static Visibility=Default Availability=Available
// [4] int startDragDistance()
func (this *QApplication) StartDragDistance() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication17startDragDistanceEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QApplication_StartDragDistance() int {
	var nilthis *QApplication
	rv := nilthis.StartDragDistance()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:169
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isEffectEnabled(Qt::UIEffect)
func (this *QApplication) IsEffectEnabled(arg0 int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication15isEffectEnabledEN2Qt8UIEffectE", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QApplication_IsEffectEnabled(arg0 int) bool {
	var nilthis *QApplication
	rv := nilthis.IsEffectEnabled(arg0)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:170
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setEffectEnabled(Qt::UIEffect, _Bool)
func (this *QApplication) SetEffectEnabled(arg0 int, enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication16setEffectEnabledEN2Qt8UIEffectEb", ffiqt.FFI_TYPE_POINTER, arg0, enable)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetEffectEnabled(arg0 int, enable bool) {
	var nilthis *QApplication
	nilthis.SetEffectEnabled(arg0, enable)
}

// /usr/include/qt/QtWidgets/qapplication.h:179
// index:0
// Public static Visibility=Default Availability=Available
// [4] int exec()
func (this *QApplication) Exec() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication4execEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QApplication_Exec() int {
	var nilthis *QApplication
	rv := nilthis.Exec()
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:180
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool notify(QObject *, QEvent *)
func (this *QApplication) Notify(arg0 *qtcore.QObject /*777 QObject **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication6notifyEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qapplication.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void focusChanged(QWidget *, QWidget *)
func (this *QApplication) FocusChanged(old *QWidget /*777 QWidget **/, now *QWidget /*777 QWidget **/) {
	var convArg0 = old.GetCthis()
	var convArg1 = now.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication12focusChangedEP7QWidgetS1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qapplication.h:193
// index:0
// Public Visibility=Default Availability=Available
// [8] QString styleSheet()
func (this *QApplication) StyleSheet() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QApplication10styleSheetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qapplication.h:196
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyleSheet(const QString &)
func (this *QApplication) SetStyleSheet(sheet *qtcore.QString) {
	var convArg0 = sheet.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication13setStyleSheetERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qapplication.h:198
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoSipEnabled(const _Bool)
func (this *QApplication) SetAutoSipEnabled(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication17setAutoSipEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qapplication.h:199
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoSipEnabled()
func (this *QApplication) AutoSipEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QApplication14autoSipEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qapplication.h:200
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void closeAllWindows()
func (this *QApplication) CloseAllWindows() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication15closeAllWindowsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QApplication_CloseAllWindows() {
	var nilthis *QApplication
	nilthis.CloseAllWindows()
}

// /usr/include/qt/QtWidgets/qapplication.h:201
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void aboutQt()
func (this *QApplication) AboutQt() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication7aboutQtEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QApplication_AboutQt() {
	var nilthis *QApplication
	nilthis.AboutQt()
}

// /usr/include/qt/QtWidgets/qapplication.h:204
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QApplication) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QApplication__ColorSpec = int

const QApplication__NormalColor QApplication__ColorSpec = 0
const QApplication__CustomColor QApplication__ColorSpec = 1
const QApplication__ManyColor QApplication__ColorSpec = 2

//  body block end