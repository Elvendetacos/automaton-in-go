package main

import (
	"fyne.io/fyne/v2/canvas"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Automata")
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.SetFixedSize(true)

	input := widget.NewEntry()

	button := widget.NewButton("Validar", func() {

		result := automaton(input.Text)
		stringResult := strings.Join(result.states, " -> ")

		if result.status {
			message := "La entrada es válida"
			resultAlert := dialog.NewInformation("Resultado", message, myWindow)
			statesAlert := dialog.NewInformation("Estados", stringResult, myWindow)
			statesAlert.Show()
			resultAlert.Show()
		} else {
			message := "La entrada es invalida"
			alert := dialog.NewInformation("Resultado", message, myWindow)
			statesAlert := dialog.NewInformation("Estados", stringResult, myWindow)
			statesAlert.Show()
			alert.Show()
		}
	})

	img, _ := fyne.LoadResourceFromPath("assets/AutomatonShort.png")
	automate := canvas.NewImageFromResource(img)
	automate.Resize(fyne.NewSize(790, 470))
	containerImg := container.NewWithoutLayout(automate)

	content := container.NewVBox(
		widget.NewLabel("Ingresa tu matricula:"),
		input,
		button,
		containerImg,
	)

	myWindow.SetContent(content)

	myWindow.ShowAndRun()
}
