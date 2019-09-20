package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type Gui struct {
	fileInput *widgets.QLineEdit
	pushButton *widgets.QPushButton
}

func newWindow(gui *Gui) *widgets.QMainWindow {
	mainWindow := widgets.NewQMainWindow(nil, 0)
	mainWindow.SetMinimumSize2(600, 600)
	mainWindow.SetWindowTitle("QMK Toolbox")

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())

	mainWindow.SetCentralWidget(widget)

	// hexLoaderGrouping component
	hexWrapper := widgets.NewQGroupBox2("Load", widget)
	hexVBoxLayout := widgets.NewQVBoxLayout2(hexWrapper)

	// hexLoadInput component
	hexFileInputWidget := widgets.NewQLineEdit2("Load", nil)

	// hexButton component
	var fileName []string
	hexButtonWidget := widgets.NewQPushButton2("load", nil)
	hexButtonWidget.SetText("Load")
	hexButtonWidget.ConnectClicked(func(bool) {
		hexFileDialogWidget := widgets.NewQFileDialog(nil, core.Qt__Dialog)
		hexFileDialogWidget.SetFileMode(widgets.QFileDialog__ExistingFile)
		hexFileDialogWidget.GetOpenFileName(hexButtonWidget,"Select .hex to flash", "$HOME/", "Hex (*.hex);;;", ".hex", 0)
		hexFileDialogWidget.SetNameFilter("hex (*.hex)")
		fileName = hexFileDialogWidget.SelectedFiles()
	})

	// Assign subcomponent to layout
	hexVBoxLayout.AddWidget(hexFileInputWidget,0, 0)
	hexVBoxLayout.AddWidget(hexButtonWidget,0 ,0)

	widget.Layout().AddWidget(hexWrapper)

	return mainWindow

}
