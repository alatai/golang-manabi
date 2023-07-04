package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	IP   string
	Port int
	// 在线用户的列表
	OnlineMap map[string]*User
	// map是全局的，所以需要加锁
	mapLock sync.RWMutex
	// 消息广播channel
	Message chan string
}

// NewServer 创建一个server接口
func NewServer(ip string, port int) *Server {
	// 指针类型
	server := &Server{
		IP:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

// Start 启动服务器接口
func (server *Server) Start() {
	// listen socket
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.IP, server.Port))
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}

	// 启动监听Message的goroutine
	go server.listenMessage()

	for {
		// accept
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen err: ", err)
			continue
		}

		// do handler
		// 利用goroutine异步处理
		go server.handler(conn)
	}

	// close listen socket
	//goland:noinspection ALL
	defer listen.Close()
}

// handler具体实现
func (server *Server) handler(conn net.Conn) {
	user := NewUser(conn)

	// 用户上线，将用户加入到onlineMap中
	server.mapLock.Lock()
	server.OnlineMap[user.Name] = user
	server.mapLock.Unlock()

	// 广播当前用户上线消息
	server.broadCase(user, "已上线")

	// 当前handler阻塞
	select {}
}

// 广播消息
func (server *Server) broadCase(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	server.Message <- sendMsg
}

// 监听Message广播消息channel的goroutine，一旦有消息就发送给全部的在线用户
func (server *Server) listenMessage() {
	for {
		msg := <-server.Message
		// 将msg发送给全部的在线用户
		server.mapLock.Lock()
		for _, client := range server.OnlineMap {
			client.Channel <- msg
		}
		server.mapLock.Unlock()
	}
}
