package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	// 当前用户绑定的channel
	Channel chan string
	// 当前用户与客户端通信的连接
	Conn net.Conn
	// 当前用户属于哪个Server
	Server *Server
}

// NewUser 创建一个用户
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:    userAddr,
		Addr:    userAddr,
		Channel: make(chan string),
		Conn:    conn,
		Server:  server,
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

// Online 用户上线功能
func (user *User) Online() {
	// 用户上线，将用户加入到onlineMap中
	user.Server.mapLock.Lock()
	user.Server.OnlineMap[user.Name] = user
	user.Server.mapLock.Unlock()

	// 广播当前用户上线消息
	user.Server.BroadCase(user, "已上线")
}

// Offline 用户下线功能
func (user *User) Offline() {
	// 用户下线，将用户加入到onlineMap中
	user.Server.mapLock.Lock()
	delete(user.Server.OnlineMap, user.Name)
	user.Server.mapLock.Unlock()

	// 广播当前用户下线消息
	user.Server.BroadCase(user, "已下线")
}

// DoMessage 用户处理消息功能
func (user *User) DoMessage(msg string) {
	// 通过“who”指令查询在线用户
	if msg == "who" {
		user.Server.mapLock.Lock()
		for _, onlineUser := range user.Server.OnlineMap {
			onlineMsg := "[" + onlineUser.Addr + "]" + onlineUser.Name + " : 在线...\n"
			// 将该消息发送给当前用户
			user.SendMsg(onlineMsg)
		}
		user.Server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" { // 通过“rename”指令修改用户名
		// 消息格式：rename|newName
		// Split：通过指定字符截取字符，截取的内容放到数组中
		newName := strings.Split(msg, "|")[1] // [rename, newName]
		// 判断newName是否已经存在
		_, ok := user.Server.OnlineMap[newName]
		if ok {
			user.SendMsg("当前用户名称已经被使用\n")
		} else {
			user.Server.mapLock.Lock()
			// 删除旧名称的用户（服务器登录用户）
			delete(user.Server.OnlineMap, user.Name)
			// 创建新名称用户（服务器登录用户）
			user.Server.OnlineMap[newName] = user
			user.Server.mapLock.Unlock()
			// 更新用户名
			user.Name = newName
			user.SendMsg("您已经更新用户名：" + user.Name + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" { // 通过“to”指令完成私聊
		// 消息格式：to|$username$|msg”
		// 获取对方用户名
		remoteName := strings.Split(msg, "|")[1] // [to, remoteName, msg]
		if remoteName == "" {
			user.SendMsg("消息格式不正确，请使用\"to|name|msg\"格式\n")
			return
		}

		// 根据用户名得到User对象
		remoteUser, ok := user.Server.OnlineMap[remoteName]
		// 当前用户不存在
		if !ok {
			user.SendMsg("该用户名不存在\n")
			return
		}

		// 获取消息内容，发送给指定用户
		content := strings.Split(msg, "|")[2] // [to, remoteName, msg]
		if content == "" {
			user.SendMsg("无消息内容，请重新发送\n")
			return
		}
		remoteUser.SendMsg(user.Name + "对您说：" + content + "\n")
	} else {
		user.Server.BroadCase(user, msg)
	}
}

// SendMsg 给指定用户发送消息
func (user *User) SendMsg(msg string) {
	//goland:noinspection GoUnhandledErrorResult
	user.Conn.Write([]byte(msg))
}
