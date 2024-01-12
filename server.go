package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// 创建server接口
func NewSever(ip string, port int) *Server {
	server := Server{
		Ip:   ip,
		Port: port,
	}
	return &server
}

func (this *Server) handler(conn net.Conn) {
	//...当前连接的业务
	fmt.Println("连接建立成功！")
}

// 启动服务器接口
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listent err:", err)
		return
	}
	//close listen socket
	defer listener.Close()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accpet err:", err)
			continue
		}
		//do handler
		go this.handler((conn))
	}
}
