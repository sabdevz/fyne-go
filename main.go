package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func updateTime(c *widget.Label) {
	timeFormatted := time.Now().Format("3:4:05 pm")
	c.SetText(timeFormatted)
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	clockWidget := widget.NewLabel("Hello Fyne!")
	updateTime(clockWidget)

	w.SetContent(container.NewVBox(
		hello,
		clockWidget,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clockWidget)
		}
	}()

	w.ShowAndRun()
}
