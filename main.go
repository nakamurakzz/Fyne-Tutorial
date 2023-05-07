package main

import (
	"log"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	helloLabel := widget.NewLabel("Hello Fyne!")
	myWindow.SetContent(helloLabel)

	timeFormatted := time.Now().Format("2006-01-02 15:04:05")
	helloLabel.SetText(timeFormatted)

	// 画面更新する関数はShowAndRunの前に書く
	go func() {
		for range time.Tick(time.Second) {
			timeFormatted := time.Now().Format("2006-01-02 15:04:05")
			helloLabel.SetText(timeFormatted)
		}
	}()

	myWindow.ShowAndRun() // 起動

	tidyUp()
}

func tidyUp() {
	// ここで後始末をする
	log.Println("exit.")
}
