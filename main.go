package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	myWindow.SetContent(widget.NewLabel("Hello Fyne!"))

	myWindow.ShowAndRun() // 起動

	tidyUp()
}

func tidyUp() {
	// ここで後始末をする
	log.Println("exit.")
}
