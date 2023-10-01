package screen

import (
	"automaton/models"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"strings"
)

type AutomatonScreen struct {
	Window fyne.Window
}

func NewAutomatonScreen(w fyne.Window) *AutomatonScreen {
	return &AutomatonScreen{
		Window: w,
	}
}

func (auth *AutomatonScreen) Start() {
	input := widget.NewEntry()

	button := widget.NewButton("Validar", func() {

		result := models.Validation(input.Text)

		stringResult := strings.Join(result.States, " -> ")

		if result.Status {
			message := "La entrada es válida"
			resultAlert := dialog.NewInformation("Resultado", message, auth.Window)
			statesAlert := dialog.NewInformation("Estados", stringResult, auth.Window)

			statesAlert.Show()
			resultAlert.Show()
		} else {
			message := "La entrada es invalida"
			resultAlert := dialog.NewInformation("Resultado", message, auth.Window)
			statesAlert := dialog.NewInformation("Estados", stringResult, auth.Window)
			statesAlert.Show()
			resultAlert.Show()
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

	auth.Window.SetContent(content)
}
