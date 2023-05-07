package main

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}

	timeText := canvas.NewText("Hello", green)
	timeText2 := canvas.NewText("World", blue)
	timeText2.Move(fyne.NewPos(0, 20))
	content := fyne.NewContainerWithoutLayout(timeText, timeText2)

	// timeFormatted := time.Now().Format("2006-01-02 15:04:05")
	// helloLabel.SetText(timeFormatted)

	// 画面更新する関数はShowAndRunの前に書く
	// go func() {
	// 	for range time.Tick(time.Second) {
	// 		timeFormatted := time.Now().Format("2006-01-02 15:04:05")
	// 		helloLabel.SetText(timeFormatted)
	// 	}
	// }()

	myWindow.Resize(fyne.NewSize(300, 300)) // 画面サイズの定義
	myWindow.SetContent(content)
	// 2つめのウィンドウを作成する
	// w2 := myApp.NewWindow("Larger")
	// w2.SetContent(widget.NewLabel("More content"))
	// w2.Resize(fyne.NewSize(100, 100))
	// w2.Show()

	myWindow.Show() // 起動
	myApp.Run()     // 起動

	tidyUp()
}

func tidyUp() {
	// ここで後始末をする
	log.Println("exit.")
}
