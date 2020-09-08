package process

import (
	"errors"
	"go_code/multiChat/server/model"
	"go_code/multiChat/server/utils"

	"net"
)

//Processor 控制器
type Processor struct {
	Conn net.Conn
	Up   *UsrProcess
}

func (processor *Processor) serverProcessMes(mes *model.Message) (err error) {

	switch mes.Type {
	case model.LoginMesType:
		err = processor.Up.ServerProcessLogin(mes)
	case model.RegistMesType:
		err = processor.Up.ServerProcessRegist(mes)
	case model.StatusReplyType:
		err = processor.Up.NotifyStatus()
	case model.SendMesType:
		err = processor.Up.SendMes(mes)
	default:
		err = errors.New("不存在该消息类型")
	}
	return
}

//Control 控制
func (processor *Processor) Control() (err error) {
	tf := &utils.Transfer{
		Conn: processor.Conn,
	}
	var mes model.Message
	for {
		mes, err = tf.ReadPkg()
		if err != nil {
			processor.Up.DeleteConn()
			return
		}
		err = processor.serverProcessMes(&mes)
		if err != nil {
			processor.Up.DeleteConn()
			return
		}
	}
}
