package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"strings"
)

type ToolBox struct {
	widgets.QMainWindow

	hexLoader *widgets.QGroupBox

}

func initToolBox() *ToolBox {
	var this = new(ToolBox)

	this.SetWindowTitle(core.QCoreApplication_ApplicationName())

	// create a regular widget
	// give it a QVBoxLayout
	// and make it the central widget of the window
	// widget := widgets.NewQWidget(nil, 0)
	// widget.SetLayout(widgets.NewQVBoxLayout())
	// this.SetCentralWidget(widget)

	this.hexLoader = createHexGroup()

	return this
}

func createHexGroup() *widgets.QGroupBox {
	var fileName []string
	var widget = new(widgets.QWidget)
	var groupBox = widgets.NewQGroupBox2("Local File", nil)
	var gridLayout = widgets.NewQGridLayout2()

	var fileInput = widgets.NewQLabel2(".hex file", widget, core.Qt__Window)

	var button = widgets.NewQPushButton2("Load", widget)
	button.ConnectClicked(func(bool) {
		// create a file dialog
		// restrict it to *.hex
		// and get file name
		hexFileSelect := widgets.NewQFileDialog(nil, core.Qt__Dialog)
		hexFileSelect.SetFileMode(widgets.QFileDialog__ExistingFile)
		hexFileSelect.GetOpenFileName(button,"Select hex to flash", "", "Hex (*.hex);;;", ".hex", 0)
		hexFileSelect.SetNameFilter("Hex (*.hex)")
		fileName = hexFileSelect.SelectedFiles()
	})
	button.ConnectReleased(func () {
		fileInput.SetText(strings.Join(fileName, "/"))
	})

	gridLayout.AddWidget(fileInput)
	gridLayout.AddWidget(button)
	groupBox.SetLayout(gridLayout)

	return groupBox

}