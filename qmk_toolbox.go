package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/go-github/github"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type App struct {
	app *widgets.QApplication
	window *widgets.QMainWindow

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

	apiClient *http.Client
	githubClient github.Client

}

func NewApp(qt *widgets.QApplication) *App {
	app := App{app: qt}
	return &app
}

func (a *App) Run() {
	a.apiClient = &http.Client{
		Timeout: time.Second * 2,
	}

	a.githubClient = *github.NewClient(a.apiClient)

	a.window = widgets.NewQMainWindow(nil, 0)
	a.window.SetWindowTitle("QMK Toolbox")

	a.hexFilePath = widgets.NewQLineEdit(nil)
	a.hexFilePath.SetReadOnly(true)

	a.hexLoadButton = widgets.NewQPushButton2("Load...", nil)
	a.hexLoadButton.ConnectClicked(a.onHexLoadButtonClicked)

	a.mcuSelector = widgets.NewQComboBox(nil)

	a.keyboardSelector = widgets.NewQComboBox(nil)
	a.keyboardSelector.ConnectCurrentTextChanged(a.populateKeyMapSelector)
	a.populateKeyboardSelector()

	a.keymapSelector = widgets.NewQComboBox(nil)

	a.keymapLoadButton = widgets.NewQPushButton2("Load...", nil)
	a.keymapLoadButton.ConnectClicked(a.onKeyMapLoadButtonClicked)

	a.console = widgets.NewQTextEdit2("console", nil)
	a.console.SetReadOnly(true)
	a.console.SetReadOnly(true)
	textFont := gui.NewQFont2("monospace", -1, -1, false)
	a.console.SetFont(textFont)
	colorPalette := gui.NewQPalette()
	colorPalette.SetColor(gui.QPalette__All, gui.QPalette__Base, gui.NewQColor6("black"))
	colorPalette.SetColor(gui.QPalette__All, gui.QPalette__Text, gui.NewQColor6("white"))
	a.console.SetPalette(colorPalette)

	a.flashButton = widgets.NewQPushButton2("Flash", nil)
	a.flashButton.ConnectClicked(a.onFlashButtonClicked)

	a.resetButton = widgets.NewQPushButton2("Reset", nil)
	a.resetButton.ConnectClicked(a.onResetButtonClicked)

	hexLayout := widgets.NewQHBoxLayout()
	hexLayout.AddWidget(a.hexFilePath, 1, 0)
	hexLayout.AddWidget(a.hexLoadButton, 1,0)
	hexLayout.AddWidget(a.mcuSelector, 1,0)

	configLayout := widgets.NewQHBoxLayout()
	configLayout.AddWidget(a.keyboardSelector, 1, 0)
	configLayout.AddWidget(a.keymapSelector, 1, 0)
	configLayout.AddWidget(a.keymapLoadButton, 1,0)

	consoleLayout := widgets.NewQHBoxLayout()
	consoleLayout.AddWidget(a.console, 1,0)

	masterLayout := widgets.NewQVBoxLayout()
	masterLayout.AddLayout(hexLayout, 1)
	masterLayout.AddLayout(configLayout, 1)
	masterLayout.AddLayout(consoleLayout, 1)

	centralWidget := widgets.NewQWidget(a.window, 0)
	centralWidget.SetLayout(masterLayout)
	a.window.SetCentralWidget(centralWidget)

	a.window.Show()
	a.app.Exec()
}

func (a *App) onHexLoadButtonClicked(checked bool) {
	hexFileDialog := widgets.NewQFileDialog(nil, core.Qt__Dialog)
	hexFileDialog.SetFileMode(widgets.QFileDialog__ExistingFile)
	hexFileDialog.SetNameFilter("Hex (*.hex)")
	hexFileDialog.ConnectFileSelected(func(file string) {
		fmt.Println(file)
		a.hexFilePath.SetText(file)
	})
	hexFileDialog.ShowDefault()
}

func (a *App) populateKeyboardSelector() {
	keyboardList := GetKeyBoardList(a.apiClient)
	a.keyboardSelector.AddItems(keyboardList)
}

func (a *App) populateKeyMapSelector(keyboard string) {
	keymapList := GetKeyMapList(&a.githubClient, keyboard)
	a.keymapSelector.Clear()
	a.keymapSelector.AddItems(keymapList)
}


func (a *App) onKeyMapLoadButtonClicked(checked bool) {
	log.Print("button clicked")
}
func (a *App) onFlashButtonClicked(checked bool) {
	log.Print("button clicked")
}
func (a *App) onResetButtonClicked(checked bool) {
	log.Print("button clicked")
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
