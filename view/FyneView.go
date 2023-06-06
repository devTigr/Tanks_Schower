package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var GlobalLable string = "none"
var w fyne.Window
var bH fyne.Container

//var bV fyne.Container

func setupScreen() (*fyne.Window, *fyne.Container) {
	application := app.New()
	w = application.NewWindow("Es werd ou einisch 4i!")
	w.Resize(fyne.NewSize(600, 400))
	bH.Resize(fyne.Size{Width: 600 / 4, Height: 400 / 11})
	//bV.Resize(fyne.Size{Width: 10})
	return &w, &bH

}

// BootScreen initialisiert die Applikation (GUI) und erstellt ein Fenster
func RunScreen() {
	label1 := widget.NewLabel("1")
	label2 := widget.NewLabel("2")
	label3 := widget.NewLabel("3")
	label4 := widget.NewLabel("4")
	label5 := widget.NewLabel("5")
	label6 := widget.NewLabel("6")
	label7 := widget.NewLabel("7")
	label8 := widget.NewLabel("8")
	label9 := *widget.NewLabel("9")
	label10 := *widget.NewLabel("10")
	btn := *widget.NewButton("refresh", func() {
		label1.SetText(GlobalLable)
	})

	setupScreen()

	//updateLabel(label)
	w.SetContent(&bH)

	/*
		go func() {
			t := time.NewTicker(time.Second)
			for range t.C {
				updateLabel(label)
			}
		}()
	*/
	window.ShowAndRun()
}

func updateLabel(label1 *widget.Label) {
	label1.SetText(GlobalLable)
}
func UpdateInfo(s string) {
	GlobalLable = s
}
func PrintMessage(s string) {
	GlobalLable = s
}
func PrintError(e error) {
	GlobalLable = e.Error()
}
