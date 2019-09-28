package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func main() {

	// needs to be called once before you can start using the QWidgets
	qt := widgets.NewQApplication(len(os.Args), os.Args)
	core.QCoreApplication_SetOrganizationName("QMK")
	core.QCoreApplication_SetApplicationName("QMK ToolBox")
	core.QCoreApplication_SetApplicationVersion("0.0.1")
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := NewApp(qt)

	app.Run()
	app.window.Show()
}
