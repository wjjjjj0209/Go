package main

import (
	"encoding/json"
	"go_code/multiChat/client/model"
	"go_code/multiChat/client/process"
	"go_code/multiChat/client/utils"
	"log"
	"net"

	. "github.com/lxn/walk/declarative"

	"github.com/lxn/walk"
)

var outTE *walk.TextEdit

func showMenu(usrProcess *process.UsrProcess) {
	go serverProcessMes(usrProcess.Conn)
	var mw *walk.MainWindow
	var message *walk.LineEdit
	defer usrProcess.Conn.Close()

	if _, err := (MainWindow{
		AssignTo: &mw,
		Title:    "MultiChat",
		Size:     Size{Width: 1000, Height: 600},
		Layout:   VBox{},
		Children: []Widget{
			TextEdit{
				AssignTo: &outTE,
				ReadOnly: true,
			},
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					LineEdit{
						AssignTo: &message,
					},
					PushButton{
						Text: "发送",
						OnClicked: func() {
							if message.Text() != "" {
								err := usrProcess.SendMes(message.Text())
								if err != nil {
									outTE.SetText(outTE.Text() + err.Error())
								}
							}
						},
					},
				},
			},
			PushButton{
				Text: "显示在线用户",
				OnClicked: func() {
					err := usrProcess.GetOnlineUsrs()
					if err != nil {
						outTE.SetText(outTE.Text() + err.Error())
					}
				},
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}

	// exit := false
	// for !exit {
	// 	fmt.Println("--------------------1.显示在线用户------------------")
	// 	fmt.Println("--------------------2.发送消息----------------------")
	// 	fmt.Println("--------------------3.退出系统----------------------")
	// 	var key int
	// 	fmt.Scanf("%d\n", &key)
	// 	switch key {
	// 	case 1:
	// 		err := usrProcess.GetOnlineUsrs()
	// 		if err != nil {
	// 			fmt.Println("获取失败!")
	// 		}
	// 	case 2:
	// 		var words string
	// 		fmt.Scanf("%s\n", &words)
	// 		err := usrProcess.SendMes(words)
	// 		if err != nil {
	// 			fmt.Println("发送失败!")
	// 		}
	// 	case 3:
	// 		exit = true
	// 	default:
	// 		fmt.Println("输入有误！")
	// 	}
	// }
}

func serverProcessMes(conn net.Conn) {
	for outTE == nil {
	}
	tf := &utils.Transfer{
		Conn: conn,
	}
	var notifyMes model.NotifyMes
	for {
		mes, err := tf.ReadPkg()
		if err != nil {
			outTE.SetText(outTE.Text() + err.Error())
			continue
		}
		err = json.Unmarshal([]byte(mes.Data), &notifyMes)
		if err != nil {
			outTE.SetText(outTE.Text() + err.Error())
			continue
		} else {
			outTE.SetText(outTE.Text() + notifyMes.Data)
		}
	}
}
