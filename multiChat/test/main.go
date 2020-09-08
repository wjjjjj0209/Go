package main

import (
	"log"

	"github.com/lxn/walk"

	. "github.com/lxn/walk/declarative"
)

func main() {

	var mw *walk.MainWindow
	var mes *walk.TextEdit

	if _, err := (MainWindow{
		AssignTo: &mw,
		Title:    "MultiChat Login",
		Size:     Size{Width: 300, Height: 200},
		Layout:   VBox{},
		Children: []Widget{
			TextEdit{
				AssignTo: &mes,
				ReadOnly: true,
			},
			PushButton{
				Text: "发送",
				OnClicked: func() {
					mes.SetText(mes.Text() + "sadsadasdasd\nsadasdsa")
				},
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}
