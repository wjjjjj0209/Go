package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"go_code/multiChat/server/model"

	"net"
)

//Transfer 接收与发送消息
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

//ReadPkg 接收
func (transfer *Transfer) ReadPkg() (mes model.Message, err error) {

	buf := transfer.Buf[:]

	//读取长度
	n, err := transfer.Conn.Read(buf[:4])
	if err != nil {
		return
	} else if n != 4 {
		err = errors.New("读取长度不匹配！")
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	//读取内容
	n, err = transfer.Conn.Read(buf[:pkgLen])
	if err != nil {
		return
	} else if n != int(pkgLen) {
		err = errors.New("读取长度不匹配！")
		return
	}

	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		return
	}
	return

}

//WritePkg 发送
func (transfer *Transfer) WritePkg(data []byte) (err error) {
	//发送长度与消息
	var pkgLen uint32 = uint32(len(data))
	buf := transfer.Buf[0:4]
	binary.BigEndian.PutUint32(buf, pkgLen)
	n, err := transfer.Conn.Write(buf)
	if err != nil {
		return
	} else if n != 4 {
		err = errors.New("发送长度不匹配！")
		return
	}

	n, err = transfer.Conn.Write(data)
	if err != nil {
		return
	} else if n != int(pkgLen) {
		err = errors.New("发送长度不匹配！")
		return
	}
	return
}
