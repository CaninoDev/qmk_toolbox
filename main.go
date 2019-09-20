package main

import (
	"github.com/therecipe/qt/core"
	"os"
	"github.com/therecipe/qt/widgets"
)

	var qApp *widgets.QApplication

func main() {

	// needs to be called once before you can start using the QWidgets
	qApp = widgets.NewQApplication(len(os.Args), os.Args)
	core.QCoreApplication_SetOrganizationName("QMK")
	core.QCoreApplication_SetApplicationName("QMK ToolBox")
	core.QCoreApplication_SetApplicationVersion("0.0.1")

	gui := &Gui{}

	mainWindow := newWindow(gui)

	// make the window visible
	mainWindow.Show()

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	qApp.Exec()
}
