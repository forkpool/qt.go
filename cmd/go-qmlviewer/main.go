package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	funk "github.com/thoas/go-funk"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtqml"
	"github.com/kitech/qt.go/qtquickwidgets"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
)

func main() {
	file := "./main.qml"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	bcc, err := ioutil.ReadFile(file)
	qtrt.ErrPrint(err, file)
	if err != nil {
		os.Exit(1)
	}
	iswin := funk.Contains(strings.Split(string(bcc), "\n"), "Window {")

	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)

	if iswin {
		log.Println("app mod qml:")
		qmle := qtqml.NewQQmlApplicationEngine(nil)
		qmle.AddImportPath("./")
		qmle.Load1(file)
	} else {
		log.Println("rect mod qml:")
		qw := qtquickwidgets.NewQQuickWidget(nil)
		qw.Engine().AddImportPath("./")
		qw.SetSource(qtcore.NewQUrl1(file, 0))
		qw.Show()
	}

	app.Exec()
}
