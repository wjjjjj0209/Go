package main

import (
	"errors"
	"go_code/multiChat/client/process"

	"log"
	"strconv"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {

	var mw *walk.MainWindow
	var UsrID *walk.LineEdit
	var UsrPwd *walk.LineEdit

	if _, err := (MainWindow{
		AssignTo: &mw,
		Title:    "MultiChat Login",
		Size:     Size{Width: 300, Height: 200},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "用户ID:",
					},
					LineEdit{
						AssignTo: &UsrID,
					},
					Label{
						Text: "密码:",
					},
					LineEdit{
						AssignTo: &UsrPwd,
					},
				},
			},

			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					PushButton{
						Text: "登录",
						OnClicked: func() {
							id, err := strconv.Atoi(UsrID.Text())
							if err != nil {
								errorTip(errors.New("用户ID格式错误"))
								return
							}
							pwd := UsrPwd.Text()
							up := &process.UsrProcess{}
							err = up.Login(id, pwd)
							if err != nil {
								errorTip(err)
								return
							}
							mw.Close()
							showMenu(up)
						},
					},
					PushButton{
						Text: "注册",
						OnClicked: func() {
							if _, err := RegisterDialog(mw); err != nil {
								log.Fatal(err)
							}
						},
					},
				},
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}
