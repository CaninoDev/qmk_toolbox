package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func main() {

	// needs to be called once before you can start using the QWidgets
	qApp := widgets.NewQApplication(len(os.Args), os.Args)
	core.QCoreApplication_SetOrganizationName("QMK")
	core.QCoreApplication_SetApplicationName("QMK ToolBox")
	core.QCoreApplication_SetApplicationVersion("0.0.1")
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	mainWindow := NewWindow()

	// make the window visible
	mainWindow.Show()

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	qApp.Exec()
}
