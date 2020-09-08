package process

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_code/multiChat/server/model"
	"go_code/multiChat/server/utils"

	"net"
)

//UsrProcess 用户进程管理
type UsrProcess struct {
	Conn  net.Conn
	UsrID int
}

//DeleteConn 断开连接
func (usrProcess *UsrProcess) DeleteConn() {
	usrM.deleteOnlineUsr(usrProcess.UsrID)
}

//SendMes 群发消息
func (usrProcess *UsrProcess) SendMes(mes *model.Message) (err error) {
	var sendMes model.SendMes
	err = json.Unmarshal([]byte(mes.Data), &sendMes)
	if err != nil {
		return
	}
	conn := model.MyUsrDao.GetRedisConn()
	defer conn.Close()
	usr, err := model.MyUsrDao.GetUserByID(conn, usrProcess.UsrID)
	if err != nil {
		return
	}
	words := usr.UsrName + ": " + sendMes.Data
	err = usrProcess.NotifyOtherUsr(words)

	err = usrProcess.NotifySelf(words)
	return
}

//NotifyStatus 通知在线情况
func (usrProcess *UsrProcess) NotifyStatus() (err error) {
	words := "当前在线用户:\n"
	conn := model.MyUsrDao.GetRedisConn()
	defer conn.Close()
	for id := range usrM.onlineUsrs {
		usr, err := model.MyUsrDao.GetUserByID(conn, id)
		if err != nil {
			continue
		}
		words += usr.UsrName + "\n"
	}
	err = usrProcess.NotifySelf(words)
	return
}

//NotifySelf 通知自己
func (usrProcess *UsrProcess) NotifySelf(words string) (err error) {
	var mes model.Message
	mes.Type = model.NotifyMesType

	var notifyMes model.NotifyMes
	notifyMes.Data = words

	data, err := json.Marshal(notifyMes)
	if err != nil {
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		return
	}

	err = usrProcess.Notify(data)

	return
}

//NotifyOtherUsr 通知其他人
func (usrProcess *UsrProcess) NotifyOtherUsr(words string) (err error) {
	var mes model.Message
	mes.Type = model.NotifyMesType

	var notifyMes model.NotifyMes
	notifyMes.Data = words

	data, err := json.Marshal(notifyMes)
	if err != nil {
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		return
	}

	for id, up := range usrM.onlineUsrs {
		if id == usrProcess.UsrID {
			continue
		}
		err = up.Notify(data)
	}
	return
}

//Notify 通知消息
func (usrProcess *UsrProcess) Notify(data []byte) (err error) {
	tf := &utils.Transfer{
		Conn: usrProcess.Conn,
	}
	err = tf.WritePkg(data)
	return
}

//ServerProcessLogin 登录管理
func (usrProcess *UsrProcess) ServerProcessLogin(mes *model.Message) (err error) {
	//处理mes
	var loginMes model.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		return
	}
	var resMes model.Message
	resMes.Type = model.LoginResMesType
	var loginResMes model.LoginResMes

	_, err = usrM.getConnByUsrID(loginMes.UsrID)
	if err == nil {
		err = errors.New("用户已在线！")
		loginResMes.Code = 404
		loginResMes.Error = err.Error()
	} else {
		usr, err := model.MyUsrDao.Login(loginMes.UsrID, loginMes.UsrPwd)
		fmt.Println(usr)
		if err != nil {
			loginResMes.Code = 404
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 200
			usrProcess.UsrID = loginMes.UsrID
			usrM.addOnlineUsr(usrProcess)
			words := fmt.Sprintf("%s上线了!", usr.UsrName)
			usrProcess.NotifyOtherUsr(words)
		}
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		return
	}
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		return
	}
	tf := &utils.Transfer{
		Conn: usrProcess.Conn,
	}
	err = tf.WritePkg(data)
	return
}

//ServerProcessRegist 注册
func (usrProcess *UsrProcess) ServerProcessRegist(mes *model.Message) (err error) {
	var registMes model.RgistMes
	err = json.Unmarshal([]byte(mes.Data), &registMes)
	if err != nil {
		return
	}
	var resMes model.Message
	resMes.Type = model.RegistResType
	var registResMes model.RegistResMes

	err = model.MyUsrDao.Regist(registMes.Usr)
	if err != nil {
		registResMes.Code = 404
		registResMes.Error = err.Error()
	} else {
		registResMes.Code = 200
	}

	data, err := json.Marshal(registResMes)
	if err != nil {
		return
	}
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		return
	}
	tf := &utils.Transfer{
		Conn: usrProcess.Conn,
	}
	err = tf.WritePkg(data)
	return
}
