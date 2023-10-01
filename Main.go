package main

import (
	screen "automaton/scenes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var scene *screen.AutomatonScreen

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Automaton")
	myWindow.Resize(fyne.NewSize(800, 500))
	myWindow.SetFixedSize(true)
	img, _ := fyne.LoadResourceFromPath("assets/Icon.png")
	myWindow.SetIcon(img)

	scene = screen.NewAutomatonScreen(myWindow)
	scene.Start()

	myWindow.Show()
	myApp.Run()
}
