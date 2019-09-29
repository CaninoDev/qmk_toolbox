package main

import (
	"os"
	"runtime"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

var (
	QTApplication *widgets.QApplication
	MainWindow *widgets.QMainWindow
)
func main() {

	// needs to be called once before you can start using the QWidgets
	QTApplication = widgets.NewQApplication(len(os.Args), os.Args)
	core.QCoreApplication_SetOrganizationName("QMK")
	core.QCoreApplication_SetOrganizationDomain("qmk.fm")
	core.QCoreApplication_SetApplicationName("QMK ToolBox")
	core.QCoreApplication_SetApplicationVersion("0.0.1")
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	MainWindow = widgets.NewQMainWindow(nil, 0)
	if runtime.GOOS == "darwin" {
		MainWindow.SetUnifiedTitleAndToolBarOnMac(true)
	}

	centralWidget := NewGUIWidget()
	MainWindow.SetCentralWidget(centralWidget)
	MainWindow.Show()
	QTApplication.Exec()

}
