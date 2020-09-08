package main

import (
	"fmt"
	"go_code/multiChat/server/model"
	"go_code/multiChat/server/process"

	"net"
	"time"
)

func init() {
	initPool("localhost:6379", 16, 0, 300*time.Second)
	model.InitUsrDao(pool)
}

func control(conn net.Conn) {
	defer conn.Close()
	processor := &process.Processor{
		Conn: conn,
		Up: &process.UsrProcess{
			Conn: conn,
		},
	}
	err := processor.Control()
	fmt.Println("连接中断:", err)
	return
}

func main() {

	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("出错！")
		return
	}

	defer listen.Close()
	fmt.Println("服务器在8889端口监听...")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("该连接有错误！")
			continue
		}
		go control(conn)
	}
}
