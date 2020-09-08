package process

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_code/multiChat/client/model"
	"go_code/multiChat/client/utils"

	"net"
)

//UsrProcess 用户管理
type UsrProcess struct {
	Conn  net.Conn
	UsrID int
}

//Login 用户登录
func (usrProcess *UsrProcess) Login(usrID int, usrPwd string) (err error) {
	//连接
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		return
	}
	fmt.Printf("%T", conn)
	//定义消息
	var mes model.Message
	mes.Type = model.LoginMesType

	var loginMes model.LoginMes
	loginMes.UsrID = usrID
	loginMes.UsrPwd = usrPwd

	//序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}

	//发送与读取
	err = tf.WritePkg(data)
	if err != nil {
		return
	}

	mes, err = tf.ReadPkg()
	if err != nil {
		return
	}

	var loginResMes model.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		return
	}
	if loginResMes.Code == 200 {

		usrProcess.Conn = conn
		usrProcess.UsrID = usrID
	} else {
		err = errors.New(loginResMes.Error)
		defer conn.Close()
	}

	return
}

//SendMes 发送消息
func (usrProcess *UsrProcess) SendMes(words string) (err error) {
	var mes model.Message
	mes.Type = model.SendMesType
	var sendMes = model.SendMes{
		UsrID: usrProcess.UsrID,
		Data:  words,
	}
	data, err := json.Marshal(sendMes)
	if err != nil {
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		return
	}
	tf := &utils.Transfer{
		Conn: usrProcess.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		return
	}

	return
}

//GetOnlineUsrs 获取在线信息
func (usrProcess *UsrProcess) GetOnlineUsrs() (err error) {
	var mes model.Message
	mes.Type = model.StatusReplyType

	var statusReplyMes = model.StatusReplyMes{}
	data, err := json.Marshal(statusReplyMes)
	if err != nil {
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		return
	}
	tf := &utils.Transfer{
		Conn: usrProcess.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		return
	}

	return
}

//Regist 用户注册
func (usrProcess *UsrProcess) Regist(usrID int,
	usrPwd string, usrName string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	defer conn.Close()
	//定义消息
	var mes model.Message
	mes.Type = model.RegistMesType

	var registMes model.RgistMes
	registMes.Usr.UsrID = usrID
	registMes.Usr.UsrPwd = usrPwd
	registMes.Usr.UsrName = usrName

	data, err := json.Marshal(registMes)
	if err != nil {
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		return
	}

	mes, err = tf.ReadPkg()
	if err != nil {
		return
	}

	var registResMes model.RegistResMes
	err = json.Unmarshal([]byte(mes.Data), &registResMes)
	if err != nil {
		return
	}

	if registResMes.Code != 200 {
		err = errors.New(registResMes.Error)
	}

	return
}
