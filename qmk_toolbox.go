package main

import (
	"context"
	"fmt"
	"github.com/therecipe/qt/gui"
	"log"
	"net/http"
	"time"

	"github.com/google/go-github/github"
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

	createHexGroup(widget)
	createConfigGroup(widget)
	createConsoleGroup(widget)

	return mainWindow

}

func createHexGroup(widget *widgets.QWidget) {
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
	// mcu selection component
	var mcuList []string
	mcuList = []string{"atmega32u4", "at90usb1286", "atmega32u2", "atmega16u2", "atmega328p", "atmega32a"}
	mcuComboBoxWidget := widgets.NewQComboBox(nil)
	mcuComboBoxWidget.AddItems(mcuList)
	mcuComboBoxWidget.ConnectCurrentIndexChanged(func(index int) {
		fmt.Println(index)
	})
	// Assign sub component to layout
	hexLayout.AddWidget(hexFileInputWidget, 1, core.Qt__AlignLeft)
	hexLayout.AddWidget(hexButtonWidget, 1, core.Qt__AlignCenter)
	hexLayout.AddWidget(mcuComboBoxWidget, 1, core.Qt__AlignRight)
	widget.Layout().AddWidget(hexWrapper)
}

func createConfigGroup(widget *widgets.QWidget) {
	var ctx context.Context
	ctx = context.Background()

	apiClient := http.Client{
		Timeout: time.Second * 2,
	}

	gitClient := github.NewClient(&apiClient)

	keyboardList := GetKeyBoardList(apiClient)
	keyMapList, err := GetKeyMapList(ctx, gitClient, keyboardList[0])
	if err != nil {
		log.Fatal(err)
	}

	var selectedKeyboard string
	var selectedKeymap string

	// configLoaderGrouping component
	configWrapper := widgets.NewQGroupBox2("Keyboard from qmk.fm", widget)
	configLayout := widgets.NewQHBoxLayout2(configWrapper)

	// configLayout component
	keyboardSelectionWidget := widgets.NewQComboBox(nil)
	keymapSelectionWidget := widgets.NewQComboBox(nil)

	keyboardSelectionWidget.AddItems(keyboardList)
	keyboardSelectionWidget.ConnectCurrentTextChanged(func(keyboard string) {
		keyMapList, err = GetKeyMapList(ctx, gitClient, keyboard)
		if err != nil {
			log.Fatal(err)
		}
		keymapSelectionWidget.Clear()
		keymapSelectionWidget.AddItems(keyMapList)
		keymapSelectionWidget.Update()
	})

	keymapSelectionWidget.AddItems(keyMapList)
	keymapSelectionWidget.ConnectCurrentTextChanged(func(keymap string) {
		fmt.Println(keymap)
	})

	// configButton component
	configButtonWidget := widgets.NewQPushButton2("load", nil)
	configButtonWidget.SetText("Load")
	configButtonWidget.ConnectReleased(func() {
		selectedKeyboard = keyboardSelectionWidget.CurrentText()
		selectedKeymap = keymapSelectionWidget.CurrentText()
		log.Printf("%s/%s", selectedKeyboard, selectedKeymap)
		widget.
	})

	configLayout.AddWidget(keyboardSelectionWidget, 1, core.Qt__AlignLeft)
	configLayout.AddWidget(keymapSelectionWidget, 1, core.Qt__AlignCenter)
	configLayout.AddWidget(configButtonWidget, 1, core.Qt__AlignRight)

	widget.Layout().AddWidget(configWrapper)
}

func createConsoleGroup(widget *widgets.QWidget) {
	textFont := gui.NewQFont2("monospace", -1, -1, false)

	consoleWrapper := widgets.NewQGroupBox2("Console", widget)
	consoleLayout := widgets.NewQGridLayout(consoleWrapper)

	consoleWidget := widgets.NewQTextEdit(widget)
	consoleWidget.SetReadOnly(true)
	consoleWidget.SetFont(textFont)

	consoleLayout.AddWidget(consoleWidget)

	widget.Layout().AddWidget(consoleWrapper)
}
