package model

//定义消息类型
const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegistMesType   = "RgistMes"
	RegistResType   = "RegistResMes"
	NotifyMesType   = "NotifyMes"
	StatusReplyType = "StatusReplyMes"
	SendMesType     = "SendMes"
)

//Message 消息
type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

//LoginMes 登陆消息
type LoginMes struct {
	UsrID  int    `json:"usrID"`
	UsrPwd string `json:"usrPwd"`
}

//RgistMes 注册消息
type RgistMes struct {
	Usr usr `json:"usr"`
}

//LoginResMes 登录结果消息
type LoginResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

//RegistResMes 注册结果消息
type RegistResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

//StatusReplyMes 获取在线用户
type StatusReplyMes struct {
}

//SendMes 发送消息
type SendMes struct {
	UsrID int    `json:"usrID"`
	Data  string `json:"data"`
}

//NotifyMes 通知消息
type NotifyMes struct {
	Data string `json:"data"`
}
