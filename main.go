package main

import (
	"ProjectZero/models"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
	"time"
)

var data [][]string

func updateEntry(newSlice []string) {
	err := models.NewEntry(newSlice)
	if err != nil {
		log.Fatal(err)
	}
}
func makeUI() (*widget.Entry, *widget.Entry, *widget.Entry, *widget.Entry, *fyne.Container) {
	inCharge := widget.NewEntry()
	inCharge.SetText("Charge:      ")

	inTank := widget.NewEntry()
	inTank.SetText("Tank:      ")
	inArbeitsgang := widget.NewEntry()
	inArbeitsgang.SetText("Arbeitsgang:      ")
	inStatus := widget.NewEntry()
	inStatus.SetText("Status:      ")

	var newSlice []string

	inCharge.OnChanged = func(content string) {
		inCharge.SetText("Charge: " + content)
		newSlice = append(newSlice, inCharge.Text)
	}
	inTank.OnChanged = func(content string) {
		inTank.SetText("Tank: " + content)
		newSlice = append(newSlice, inTank.Text)
	}
	inArbeitsgang.OnChanged = func(content string) {
		inArbeitsgang.SetText("Arbeitsgang: " + content)
		newSlice = append(newSlice, inArbeitsgang.Text)
	}
	inStatus.OnChanged = func(content string) {
		inStatus.SetText("Status: " + content)
		newSlice = append(newSlice, inStatus.Text)
	}

	submit := container.NewHBox(&*widget.NewButton("Enter", func() {
		defer fmt.Println("Oleee!!!")
		updateEntry(newSlice)
	}))

	return inCharge, inTank, inArbeitsgang, inStatus, submit
}

func updateTable() error {
	d, err := models.GetTable()
	data = d
	return err
}

func main() {
	application := app.New()
	w := application.NewWindow("Es werd ou einisch 4i!")
	w.Resize(fyne.NewSize(500, 400))

	err := updateTable()
	if err != nil {
		log.Fatal(err)
	}
	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template_Col")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})
	table.Resize(fyne.Size{500, 500})

	hB := container.NewHBox(
		widget.NewLabel("    "),
		widget.NewLabel("Charge:          "),
		widget.NewLabel("Tank:              "),
		widget.NewLabel("Arbeitsgang   "),
		widget.NewLabel("Status:      "))

	vB := container.NewVBox(
		widget.NewLabel("1"),
		widget.NewLabel("2"),
		widget.NewLabel("3"),
		widget.NewLabel("4"),
		widget.NewLabel("5"),
		widget.NewLabel("6"),
		widget.NewLabel("7"),
		widget.NewLabel("8"),
		widget.NewLabel("9"),
		widget.NewLabel("10"))
	vB.Resize(fyne.Size{10, 500})

	btn := container.NewHBox(&*widget.NewButton("refresh", func() {
		defer fmt.Println("Yeeeehaaa!!!")
		err := updateTable()
		if err != nil {
			fmt.Println(err)
		}
		table.Refresh()
	}))

	inCharge, inTank, inArbeitsgang, inStatus, btnEnter := makeUI()
	var commonWith = fyne.Size{200, 200}
	inCharge.Resize(commonWith)
	inChargeBox := container.NewHBox(inCharge)
	inChargeBox.Resize(commonWith)
	inTank.Resize(commonWith)
	inArbeitsgang.Resize(commonWith)
	inStatus.Resize(commonWith)

	lOut := container.NewVBox(
		container.NewBorder(hB, layout.NewSpacer(), vB, layout.NewSpacer(), table),
		container.NewHBox(layout.NewSpacer(), btn, btnEnter, layout.NewSpacer()),
		container.NewHBox(container.NewWithoutLayout(inChargeBox, inTank, inArbeitsgang, inStatus)),
	)

	container.NewWithoutLayout()
	//sumBox := container.NewVBox(bHNum, btn)
	w.SetContent(lOut)

	go func() {
		t := time.NewTicker(time.Second * 10)

		for range t.C {
			err = updateTable()
			if err != nil {
				log.Fatal(err)
			}
		}
	}()
	w.ShowAndRun()
}
