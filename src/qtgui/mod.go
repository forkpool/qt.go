package qtgui

/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -lQtInline
//  -lQt5Core -lQt5Gui -lQt5Widgets
*/
import "C"
import "qtcore"

func init() {
	if false {
		qtcore.KeepMe()
	}
}

func KeepMe() {}
