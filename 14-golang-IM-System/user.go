package main

import "net"

type User struct {
	Name string
	Addr string
	// 当前用户绑定的channel
	Channel chan string
	// 当前用户与客户端通信的连接
	Conn net.Conn
}

// NewUser 创建一个用户
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:    userAddr,
		Addr:    userAddr,
		Channel: make(chan string),
		Conn:    conn,
	}

	// 启动监听当前用户channel消息的goroutine
	go user.listenMessage()
	return user
}

// 监听当前User channel的方法，一旦有消息就直接发送给对端客户端
func (user *User) listenMessage() {
	for {
		msg := <-user.Channel
		//goland:noinspection GoUnhandledErrorResult
		user.Conn.Write([]byte(msg + "\n"))
	}
}
