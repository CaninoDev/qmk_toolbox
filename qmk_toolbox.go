package main

import (
	"bytes"
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"log"
	"os/exec"
)

type GUI struct {
	hexGroup *widgets.QGroupBox
	hexFilePath *widgets.QLineEdit
	hexLoadButton *widgets.QPushButton
	mcuSelector *widgets.QComboBox

	configGroup *widgets.QGroupBox
	keyboardSelector *widgets.QComboBox
	keymapSelector *widgets.QComboBox
	keymapLoadButton *widgets.QPushButton

	flashButton *widgets.QPushButton
	resetButton *widgets.QPushButton

	console *widgets.QTextEdit
}

func NewGUIWidget() (guiWidget *widgets.QWidget) {
	g := &GUI{}
	g.hexFilePath = widgets.NewQLineEdit(nil)
	g.hexFilePath.SetReadOnly(true)

	g.hexLoadButton = widgets.NewQPushButton2("Load...", nil)
	g.hexLoadButton.ConnectClicked(g.onHexLoadButtonClicked)

	g.mcuSelector = widgets.NewQComboBox(nil)

	g.keyboardSelector = widgets.NewQComboBox(nil)
	g.keyboardSelector.ConnectCurrentTextChanged(g.populateKeyMapSelector)
	g.populateKeyboardSelector()

	g.keymapSelector = widgets.NewQComboBox(nil)

	g.keymapLoadButton = widgets.NewQPushButton2("Load...", nil)
	g.keymapLoadButton.ConnectClicked(g.onFlashButtonClicked)

	g.console = widgets.NewQTextEdit2("console", nil)
	g.console.SetReadOnly(true)
	g.console.SetReadOnly(true)
	textFont := gui.NewQFont2("monospace", -1, -1, false)
	g.console.SetFont(textFont)
	colorPalette := gui.NewQPalette()
	colorPalette.SetColor(gui.QPalette__All, gui.QPalette__Base, gui.NewQColor6("black"))
	colorPalette.SetColor(gui.QPalette__All, gui.QPalette__Text, gui.NewQColor6("white"))
	g.console.SetPalette(colorPalette)

	g.flashButton = widgets.NewQPushButton2("Flash", nil)
	g.flashButton.ConnectClicked(g.onFlashButtonClicked)

	g.resetButton = widgets.NewQPushButton2("Reset", nil)
	g.resetButton.ConnectClicked(g.onResetButtonClicked)

	hexLayout := widgets.NewQHBoxLayout()
	hexLayout.AddWidget(g.hexFilePath, 1, 0)
	hexLayout.AddWidget(g.hexLoadButton, 1,0)
	hexLayout.AddWidget(g.mcuSelector, 1,0)

	configLayout := widgets.NewQHBoxLayout()
	configLayout.AddWidget(g.keyboardSelector, 1, 0)
	configLayout.AddWidget(g.keymapSelector, 1, 0)
	configLayout.AddWidget(g.keymapLoadButton, 1,0)

	consoleLayout := widgets.NewQHBoxLayout()
	consoleLayout.AddWidget(g.console, 1,0)

	masterLayout := widgets.NewQVBoxLayout()
	masterLayout.AddLayout(hexLayout, 1)
	masterLayout.AddLayout(configLayout, 1)
	masterLayout.AddLayout(consoleLayout, 1)

	guiWidget = widgets.NewQWidget(MainWindow, 0)
	guiWidget.SetLayout(masterLayout)

	return guiWidget
}

func (g *GUI) onHexLoadButtonClicked(checked bool) {
	hexFileDialog := widgets.NewQFileDialog(nil, core.Qt__Dialog)
	hexFileDialog.SetFileMode(widgets.QFileDialog__ExistingFile)
	hexFileDialog.SetNameFilter("Hex (*.hex)")
	hexFileDialog.ConnectFileSelected(func(file string) {
		fmt.Println(file)
		g.hexFilePath.SetText(file)
	})
	hexFileDialog.ShowDefault()
}

func (g *GUI) populateKeyboardSelector() {
	keyboardList := GetKeyBoardList()
	g.keyboardSelector.AddItems(keyboardList)
}

func (g *GUI) populateKeyMapSelector(keyboard string) {
	keymapList := GetKeyMapList(keyboard)
	g.keymapSelector.Clear()
	g.keymapSelector.AddItems(keymapList)
}


func (g *GUI) onKeyMapLoadButtonClicked(checked bool) {
	log.Print("button clicked")
}
func (g *GUI) onFlashButtonClicked(checked bool) {
	output := run("/usr/bin/bat")
	go func() {
		g.console.Append(output)
	}()
}

func (g *GUI) onResetButtonClicked(checked bool) {
	log.Print("button clicked")
}

func run(command string) string {
	cmd := exec.Command(command, "/home/caninodev/Developments/Go/src/github.com/caninodev/qmk_toolbox/sample.txt")
	var b bytes.Buffer
	cmd.Stdout = &b

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	return b.String()
}
//
//func createHexGroup(widget *widgets.QWidget) {
//	// hexLoaderGrouping component
//	hexWrapper := widgets.NewQGroupBox2("Load", widget)
//	hexLayout := widgets.NewQHBoxLayout2(hexWrapper)
//	// hexLoadInput component
//	hexFileInputWidget := widgets.NewQLineEdit2("Load", nil)
//	hexFileInputWidget.SetReadOnly(true)
//	// hexButton component
//	//var fileName []string
//	hexButtonWidget := widgets.NewQPushButton2("load", nil)
//	hexButtonWidget.SetText("Load")
//	hexButtonWidget.ConnectReleased(func() {
//	})
//	// mcu selection component
//	var mcuList []string
//	mcuList = []string{"atmega32u4", "at90usb1286", "atmega32u2", "atmega16u2", "atmega328p", "atmega32a"}
//	mcuComboBoxWidget := widgets.NewQComboBox(nil)
//	mcuComboBoxWidget.AddItems(mcuList)
//	mcuComboBoxWidget.ConnectCurrentIndexChanged(func(index int) {
//		fmt.Println(index)
//	})
//	// Assign sub component to layout
//	hexLayout.AddWidget(hexFileInputWidget, 1, core.Qt__AlignLeft)
//	hexLayout.AddWidget(hexButtonWidget, 1, core.Qt__AlignCenter)
//	hexLayout.AddWidget(mcuComboBoxWidget, 1, core.Qt__AlignRight)
//	widget.Layout().AddWidget(hexWrapper)
//}
//
//func createConfigGroup(widget *widgets.QWidget) {
//	var ctx context.Context
//	ctx = context.Background()
//
//	apiClient := http.Client{
//		Timeout: time.Second * 2,
//	}
//
//	gitClient := github.NewClient(&apiClient)
//
//	keyboardList := GetKeyBoardList(apiClient)
//	keyMapList, err := GetKeyMapList(ctx, gitClient, keyboardList[0])
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var selectedKeyboard string
//	var selectedKeymap string
//
//	// configLoaderGrouping component
//	configWrapper := widgets.NewQGroupBox2("Keyboard from qmk.fm", widget)
//	configLayout := widgets.NewQHBoxLayout2(configWrapper)
//
//	// configLayout component
//	keyboardSelectionWidget := widgets.NewQComboBox(nil)
//	keymapSelectionWidget := widgets.NewQComboBox(nil)
//
//	keyboardSelectionWidget.AddItems(keyboardList)
//	keyboardSelectionWidget.ConnectCurrentTextChanged(func(keyboard string) {
//		keyMapList, err = GetKeyMapList(ctx, gitClient, keyboard)
//		if err != nil {
//			log.Fatal(err)
//		}
//		keymapSelectionWidget.Clear()
//		keymapSelectionWidget.AddItems(keyMapList)
//		keymapSelectionWidget.Update()
//	})
//
//	keymapSelectionWidget.AddItems(keyMapList)
//	keymapSelectionWidget.ConnectCurrentTextChanged(func(keymap string) {
//		fmt.Println(keymap)
//	})
//
//	// configButton component
//	configButtonWidget := widgets.NewQPushButton2("load", nil)
//	configButtonWidget.SetText("Load")
//	configButtonWidget.ConnectReleased(func() {
//		selectedKeyboard = keyboardSelectionWidget.CurrentText()
//		selectedKeymap = keymapSelectionWidget.CurrentText()
//		log.Printf("%s/%s", selectedKeyboard, selectedKeymap)
//		widget.
//	})
//
//	configLayout.AddWidget(keyboardSelectionWidget, 1, core.Qt__AlignLeft)
//	configLayout.AddWidget(keymapSelectionWidget, 1, core.Qt__AlignCenter)
//	configLayout.AddWidget(configButtonWidget, 1, core.Qt__AlignRight)
//
//	widget.Layout().AddWidget(configWrapper)
//}
//
//func createConsoleGroup(widget *widgets.QWidget) {
//	textFont := gui.NewQFont2("monospace", -1, -1, false)
//
//	consoleWrapper := widgets.NewQGroupBox2("Console", widget)
//	consoleLayout := widgets.NewQGridLayout(consoleWrapper)
//
//	consoleWidget := widgets.NewQTextEdit(widget)
//	consoleWidget.SetReadOnly(true)
//	consoleWidget.SetFont(textFont)
//
//	consoleLayout.AddWidget(consoleWidget)
//
//	widget.Layout().AddWidget(consoleWrapper)
//}
