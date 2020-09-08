package main

import (
	"errors"
	"go_code/multiChat/client/process"
	"strconv"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

//RegisterDialog 注册界面
func RegisterDialog(mw *walk.MainWindow) (int, error) {
	var dlg *walk.Dialog
	var registPB, cancelPB *walk.PushButton
	var UsrName *walk.LineEdit
	var UsrID *walk.LineEdit
	var UsrPwd *walk.LineEdit

	return Dialog{
		AssignTo:      &dlg,
		Title:         "新用户注册",
		DefaultButton: &registPB,
		CancelButton:  &cancelPB,
		Size:          Size{Width: 300, Height: 300},
		Layout:        VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "昵称:",
					},
					LineEdit{
						AssignTo: &UsrName,
					},
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
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &registPB,
						Text:     "OK",
						OnClicked: func() {

							id, err := strconv.Atoi(UsrID.Text())
							if err != nil {
								errorTip(errors.New("用户ID格式错误"))
								return
							}
							pwd := UsrPwd.Text()
							name := UsrName.Text()
							if strings.Trim(name, " ") == "" {
								errorTip(errors.New("昵称不能为空"))
								return
							}
							if strings.Trim(pwd, " ") == "" {
								errorTip(errors.New("密码不能为空"))
								return
							}
							up := &process.UsrProcess{}
							err = up.Regist(id, pwd, name)
							if err != nil {
								errorTip(err)
								return
							}
							dlg.Accept()
						},
					},
					PushButton{
						AssignTo:  &cancelPB,
						Text:      "Cancel",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}.Run(mw)
}
