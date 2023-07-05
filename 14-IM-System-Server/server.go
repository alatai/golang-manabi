package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
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
	go server.ListenMessage()

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
	user := NewUser(conn, server)
	// 用户上线
	user.Online()

	// 监听用户是否活跃的channel
	isLive := make(chan bool)

	// 接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			// read：读取的长度
			read, err := conn.Read(buf)
			if read == 0 {
				// 用户下线
				user.Offline()
				return
			}

			// EOF：文件末尾
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}

			// 提取用户的消息（去除'\n'）
			msg := string(buf[:read-1])
			// 用于针对msg进行处理
			user.DoMessage(msg)
			// 用户的任意消息，代表当前用户时活跃状态
			isLive <- true
		}
	}()

	// 当前handler阻塞
	for {
		// select会阻塞当前channel
		select {
		// 当前用户是活跃的
		case <-isLive:
		// 不在任何处理，激活下一个case的语句，达到重置定时器的目的
		// time.After：golang中的定时器
		// time.After()执行后重置该定时器
		case <-time.After(time.Second * 120):
			// 已经超时，将当前的用户强制关闭
			user.SendMsg("您被下线\n")
			// 销毁用户资源
			close(user.Channel)
			// 关闭连接
			//goland:noinspection GoUnhandledErrorResult
			conn.Close()
			// 退出当前handler
			// runtime.exit()
			return
		}
	}
}

// BroadCase 广播消息
func (server *Server) BroadCase(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	server.Message <- sendMsg
}

// ListenMessage 监听Message广播消息channel的goroutine，一旦有消息就发送给全部的在线用户
func (server *Server) ListenMessage() {
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
