package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

/* 客户端功能 */

// 声明命令行解析环境变量
var serverIP string
var serverPort int

// 设置命令行解析环境变量
func init() {
	flag.StringVar(&serverIP, "ip", "127.0.0.1", "设置服务器IP地址（默认127.0.0.1）")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口（默认8888）")
}

func main() {
	// 命令行解析，会解析在当前进程中设置的环境变量
	flag.Parse()

	client := NewClient(serverIP, serverPort)
	if client == nil {
		fmt.Println(">>>>> 连接服务器失败...")
		return
	}

	fmt.Println(">>>>> 连接服务器成功...")

	// 单独开启一个goroutine去处理服务端回执的消息
	go client.DoResponse()

	// 启动客户端的业务
	client.run()
}

// Client 客户端
type Client struct {
	ServerIP   string
	ServerPort int
	// 客户端名称
	Name string
	Conn net.Conn
	// 当前客户端模式
	Mode int
}

// NewClient 创建客户端
func NewClient(serverIP string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIP:   serverIP,
		ServerPort: serverPort,
		Mode:       999,
	}

	// 连接服务器
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIP, serverPort))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}

	client.Conn = conn
	return client
}

// 菜单显示
func (client *Client) menu() bool {
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	// 用户输入
	var mode int
	//goland:noinspection GoUnhandledErrorResult
	fmt.Scanln(&mode)

	if mode >= 0 && mode <= 3 {
		client.Mode = mode
		return true
	} else {
		fmt.Println(">>>>>请输入合法范围内数值<<<<<")
		return false
	}
}

// 客户端业务
func (client *Client) run() {
	for client.Mode != 0 {
		// 若没有正确输入，则一直循环执行
		for client.menu() != true {

		}

		// 根据不同的模式处理不同的业务
		switch client.Mode {
		// 公聊模式
		case 1:
			client.publicChat()
			break
		// 私聊模式
		case 2:
			client.privateChat()
			break
		// 更新用户名
		case 3:
			client.updateName()
			break
		// 退出
		case 0:
			fmt.Println(0)
			break
		}
	}
}

// DoResponse 接收服务端消息，直接显示到标准输出
func (client *Client) DoResponse() {
	// io.Copy()，会不断阻塞连接（永久阻塞监听），若有数据则会写到Stdout
	_, err := io.Copy(os.Stdout, client.Conn)
	if err != nil {
		return
	}
}

// 更新用户名
func (client *Client) updateName() bool {
	// 提示用户输入更新名
	fmt.Println(">>>>>请输入用户名：")
	//goland:noinspection GoUnhandledErrorResult
	fmt.Scanln(&client.Name)

	// 通过约定协议向服务端发送请求
	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.Conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}

	return true
}

// 公聊模式
func (client *Client) publicChat() {
	// 提示用户输入消息
	fmt.Println(">>>>>请输入聊天内容，exit表示退出")
	var chatMsg string
	//goland:noinspection GoUnhandledErrorResult
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		// 消息不为空则发送给服务器
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.Conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn write err:", err)
				break
			}
		}

		chatMsg = ""
		fmt.Println(">>>>>请输入聊天内容，exit表示退出")
		//goland:noinspection GoUnhandledErrorResult
		fmt.Scanln(&chatMsg)
	}
}

// 查询在线用户
func (client *Client) selectUser() {
	sendMsg := "who\n"
	_, err := client.Conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn write err:", err)
		return
	}
}

// 私聊模式
func (client *Client) privateChat() {
	// 聊天对象用户名
	var remoteName string
	// 聊天内容
	var chatMsg string

	// 查询当前在线用户
	client.selectUser()
	// 提示选择一个用户进入私聊
	fmt.Println(">>>>>请输入聊天对象[用户名]，exit退出：")
	//goland:noinspection GoUnhandledErrorResult
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>>>请输入消息内容，exit退出：")
		//goland:noinspection GoUnhandledErrorResult
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			// 消息不为空则发送给服务器
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
				_, err := client.Conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn write err:", err)
					break
				}
			}

			// 不断循环直到exit
			chatMsg = ""
			fmt.Println(">>>>>请输入消息内容，exit退出")
			//goland:noinspection GoUnhandledErrorResult
			fmt.Scanln(&chatMsg)
		}

		// 重新选择私聊对象用户
		client.selectUser()
		fmt.Println(">>>>>请输入聊天对象[用户名]，exit退出：")
		//goland:noinspection GoUnhandledErrorResult
		fmt.Scanln(&remoteName)
	}
}
