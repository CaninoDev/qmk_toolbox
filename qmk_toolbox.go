package main

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"log"
)

type GUI struct {
	hexGroup         *widgets.QGroupBox
	hexFilePath      *widgets.QLineEdit
	hexLoadButton    *widgets.QPushButton
	mcuSelector      *widgets.QComboBox
	configGroup      *widgets.QGroupBox
	keyboardSelector *widgets.QComboBox
	keymapSelector   *widgets.QComboBox
	keymapLoadButton *widgets.QPushButton
	dfuCheckBox		 *widgets.QCheckBox
	stm32CheckBox	 *widgets.QCheckBox
	halfkayCheckBox  *widgets.QCheckBox
	caterinaCheckBox *widgets.QCheckBox
	flashButton      *widgets.QPushButton
	resetButton      *widgets.QPushButton
	console          *widgets.QTextEdit
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

	g.console = widgets.NewQTextEdit2("", nil)
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

	g.dfuCheckBox = widgets.NewQCheckBox2("DFU", nil)
	g.stm32CheckBox = widgets.NewQCheckBox2("STM32", nil)
	g.halfkayCheckBox = widgets.NewQCheckBox2("Halfkay", nil)
	g.caterinaCheckBox = widgets.NewQCheckBox2("Caterina", nil)

	flasherGridLayout := widgets.NewQGridLayout2()
	flasherGridLayout.AddWidget3(g.dfuCheckBox, 0, 0,1, 1, core.Qt__AlignCenter)
	flasherGridLayout.AddWidget3(g.stm32CheckBox, 1, 0,1,1, core.Qt__AlignCenter)
	flasherGridLayout.AddWidget3(g.halfkayCheckBox, 0, 1, 1,1 ,core.Qt__AlignCenter)
	flasherGridLayout.AddWidget3(g.caterinaCheckBox, 1,1, 1, 1, core.Qt__AlignCenter)

	flasherGroupBox := widgets.NewQGroupBox(nil)
	flasherGroupBox.SetFlat(false)
	flasherGroupBox.SetTitle("MCUs")
	flasherGroupBox.SetAlignment(1)
	flasherGroupBox.SetLayout(flasherGridLayout)

	hexLayout := widgets.NewQHBoxLayout()
	hexLayout.AddWidget(g.hexFilePath, 1, 0)
	hexLayout.AddWidget(g.hexLoadButton, 1, 0)
	hexLayout.AddWidget(g.mcuSelector, 1, 0)

	hexGroupBox := widgets.NewQGroupBox(nil)
	hexGroupBox.SetTitle("Load File...")
	hexGroupBox.SetAlignment(1)
	hexGroupBox.SetLayout(hexLayout)

	configLayout := widgets.NewQHBoxLayout()
	configLayout.AddWidget(g.keyboardSelector, 1, 0)
	configLayout.AddWidget(g.keymapSelector, 1, 0)
	configLayout.AddWidget(g.keymapLoadButton, 1, 0)

	configGroupBox := widgets.NewQGroupBox(nil)
	configGroupBox.SetTitle("Keyboard from qmk.fm")
	configGroupBox.SetAlignment(1)
	configGroupBox.SetLayout(configLayout)

	consoleLayout := widgets.NewQHBoxLayout()
	consoleLayout.AddWidget(g.console, 1, 0)

	consoleGroupBox := widgets.NewQGroupBox(nil)
	consoleGroupBox.SetTitle("Console")
	consoleGroupBox.SetAlignment(1)
	consoleGroupBox.SetLayout(consoleLayout)

	masterLayout := widgets.NewQVBoxLayout()
	masterLayout.AddWidget(hexGroupBox, 1, core.Qt__AlignCenter)
	masterLayout.AddWidget(configGroupBox, 1, core.Qt__AlignCenter)
	masterLayout.AddWidget(consoleGroupBox, 1, core.Qt__AlignCenter)
	masterLayout.AddWidget(hexGroupBox, 1, core.Qt__AlignCenter)
	masterLayout.AddWidget(flasherGroupBox, 1, core.Qt__AlignCenter)
	
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
	cmd := core.NewQProcess(nil)
	cmd.ConnectReadyRead(func() {
		var stdOutput *core.QByteArray
		stdOutput = cmd.ReadAllStandardOutput()
		g.console.InsertPlainText(stdOutput.Data())
	})
	cmd.ConnectErrorOccurred(func(error core.QProcess__ProcessError) {
		log.Print(error)
	})
	cmd.Start("/home/caninodev/.local/bin/dotRepeater.sh", nil, core.QIODevice__ReadOnly)
}

func (g *GUI) onResetButtonClicked(checked bool) {
	log.Print("button clicked")
}