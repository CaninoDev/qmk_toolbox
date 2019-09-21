package main

import (
	"encoding/json"
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"io/ioutil"
	"log"
	"net/http"
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
	var selectedKeyboard string
	var selectedKeymap string

	// configLoaderGrouping component
	configWrapper := widgets.NewQGroupBox2("Keyboard from qmk.fm", widget)
	configLayout := widgets.NewQHBoxLayout2(configWrapper)

	// configLayout component
	keyboardList := populateKeyboardList()
	keyboardSelectionWidget := widgets.NewQComboBox(nil)
	keyboardSelectionWidget.AddItems(keyboardList)
	keyboardSelectionWidget.ConnectCurrentTextChanged(func(keyboard string) {
		fmt.Println(keyboard)
		selectedKeyboard = keyboard
	})

	//
	var keymapList []string
	keymapList = []string{"keymap1", "keymap2"}
	keymapSelectionWidget := widgets.NewQComboBox(nil)
	keymapSelectionWidget.AddItems(keymapList)
	keymapSelectionWidget.ConnectCurrentTextChanged(func(keymap string) {
		fmt.Println(keymap)
		selectedKeymap = keymap
	})

	// configButton component
	configButtonWidget := widgets.NewQPushButton2("load", nil)
	configButtonWidget.SetText("Load")
	configButtonWidget.ConnectReleased(func() {
		fmt.Println(fmt.Sprintf("%s%s", selectedKeyboard, selectedKeymap))
	})

	configLayout.AddWidget(keyboardSelectionWidget, 1, core.Qt__AlignLeft)
	configLayout.AddWidget(keymapSelectionWidget, 1, core.Qt__AlignCenter)
	configLayout.AddWidget(configButtonWidget, 1, core.Qt__AlignRight)

	widget.Layout().AddWidget(configWrapper)
}

func populateKeyboardList() (keyboardList []string) {
	url := "http://compile.qmk.fm/v1/keyboards"


	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	_ = json.Unmarshal([]byte(body), &keyboardList)

	return keyboardList
}