package main

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func NewWindow() *widgets.QMainWindow {
	mainWindow := widgets.NewQMainWindow(nil, 0)
	mainWindow.SetMinimumSize2(200, 400)
	mainWindow.SetWindowTitle("QMK Toolbox")

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())

	mainWindow.SetCentralWidget(widget)

	// hexLoaderGrouping component
	hexWrapper := widgets.NewQGroupBox2("Load", widget)
	hexLayout := widgets.NewQHBoxLayout2(hexWrapper)

	// hexLoadInput component
	hexFileInputWidget := widgets.NewQLineEdit2("Load", nil)
	hexFileInputWidget.SetReadOnly(true)

	// hexButton component
	//var fileName []string
	hexButtonWidget := widgets.NewQPushButton2("load", nil)
	hexButtonWidget.SetText("Load")
	hexButtonWidget.ConnectReleased(func() {
		// hexFileDialogWidget
		hexFileDialogWidget := widgets.NewQFileDialog(nil, core.Qt__Dialog)
		hexFileDialogWidget.SetFileMode(widgets.QFileDialog__ExistingFile)
//		hexFileDialogWidget.GetOpenFileName(nil, "Select .hex to flash", "$HOME", "Hex (*.hex);;;", ".hex", 0)
		hexFileDialogWidget.SetNameFilter("Hex (*.hex)")
		hexFileDialogWidget.ConnectFileSelected(func(file string) {
			fmt.Println(file)
			hexFileInputWidget.SetText(file)
		})
		hexFileDialogWidget.ShowDefault()
	})

	// Assign sub component to layout
	hexLayout.AddWidget(hexFileInputWidget, 1, core.Qt__AlignCenter)
	hexLayout.AddWidget(hexButtonWidget, 1, core.Qt__AlignCenter)

	widget.Layout().AddWidget(hexWrapper)

	return mainWindow

}
