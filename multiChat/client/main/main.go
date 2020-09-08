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

	// var key int
	// var loop = false

	// for !loop {
	// 	fmt.Println("---------------------欢迎使用多人聊天系统----------------------")
	// 	fmt.Println("\t\t\t 1 登录")
	// 	fmt.Println("\t\t\t 2 注册")
	// 	fmt.Println("\t\t\t 3 退出")
	// 	fmt.Scanf("%d\n", &key)
	// 	switch key {
	// 	case 1:
	// 		//登录
	// 		var (
	// 			usrID  int
	// 			usrPwd string
	// 		)
	// 		fmt.Println("输入ID:")
	// 		fmt.Scanf("%d\n", &usrID)
	// 		fmt.Println("输入密码:")
	// 		fmt.Scanf("%s\n", &usrPwd)
	// 		up := &process.UsrProcess{}
	// 		err := up.Login(usrID, usrPwd)
	// 		if err != nil {
	// 			fmt.Println("登陆失败:", err)
	// 		}
	// 	case 2:
	// 		//注册
	// 		var (
	// 			usrID   int
	// 			usrPwd  string
	// 			usrName string
	// 		)
	// 		fmt.Println("输入ID:")
	// 		fmt.Scanf("%d\n", &usrID)
	// 		fmt.Println("输入密码:")
	// 		fmt.Scanf("%s\n", &usrPwd)
	// 		fmt.Println("输入用户名:")
	// 		fmt.Scanf("%s\n", &usrName)
	// 		up := &process.UsrProcess{}
	// 		err := up.Regist(usrID, usrPwd, usrName)
	// 		if err != nil {
	// 			fmt.Println("注册失败:", err)
	// 		} else {
	// 			fmt.Println("注册成功！")
	// 		}
	// 	case 3:
	// 		//退出
	// 		loop = true
	// 	default:
	// 		fmt.Println("输入不是有效命令，请重新输入")
	// 	}
	// }
}
