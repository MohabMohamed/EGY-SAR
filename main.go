package main

import (
	"egy-sar/egysar"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func BindApis(app *wails.App) {
	app.Bind(egysar.NewLeeway)
}

func main() {
	js := mewn.String("./frontend/public/build/bundle.js")
	css := mewn.String("./frontend/public/build/bundle.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1024,
		Height:    768,
		Title:     "EGY-SAR",
		JS:        js,
		CSS:       css,
		Colour:    "#131313",
		Resizable: true,
	})

	BindApis(app)

	app.Run()

}
