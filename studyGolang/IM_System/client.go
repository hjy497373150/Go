package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp string
	ServerPort int
	Name string
	Conn net.Conn
	Flag int //当前客户端的模式
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIp: serverIp,
		ServerPort: serverPort,
		Flag: 999,
	}

	// 与服务器建立连接
	conn,err := net.Dial("tcp", fmt.Sprintf("%s:%d",serverIp,serverPort))
	// 连接失败直接返回nil
	if err != nil {
		fmt.Println("net.dial error:",err)
		return nil
	}

	client.Conn = conn

	return client

}

var serverIp string
var serverPort int
func init(){
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址（默认值是127.0.0.1）")
	flag.IntVar(&serverPort,"port", 8888, "设置服务器端口（默认为8888）")
}

// 客户端菜单封装
func (client *Client) menu() bool {
	var flag int // 模式
	fmt.Println("<-------功能菜单，请选择指定的功能编号-------->")
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.修改用户名")
	fmt.Println("0.退出")
	
	// 读取用户选择的模式
	fmt.Scanln(&flag)

	if flag >=0 && flag <= 3 {
		client.Flag = flag
		return true
	} else {
		fmt.Println("请输入正确的功能编号")
		return false
	}
}

// 开启客户端任务
func (client *Client) run() {
	for client.Flag != 0 {
		for client.menu() != true {

		}

		// 根据不同的模式处理不同的业务
		switch client.Flag {
		case 1:
			// 公聊模式
			client.PublicChat()
			break
		case 2:
			// 私聊模式
			client.PrivateChat()
			break
		case 3:
			// 修改用户名
			client.UpdateName()
			break
		}
	}
	fmt.Println("退出客户端...")
}

// 更新用户名方法
func (client *Client) UpdateName() bool{
	fmt.Println("请输入您要修改的用户名:")

	fmt.Scanln(&client.Name)

	sendMsg := "rename:" + client.Name + "\n"
	_,err := client.Conn.Write([]byte(sendMsg)) 

	if err != nil {
		fmt.Println("client.conn error:",err)
		return false
	}
	return true
}

// 公聊模式
func (client *Client) PublicChat() {
	fmt.Println("请输入您要发送的信息, 输入exit退出")

	var chatMsg string
	fmt.Scanln(&chatMsg)

	// 发送给服务器
	for chatMsg != "exit" {
		// 消息不为空则发送
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.Conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("client.Conn write error:",err)
				break
			}
		}
		// 重置chatmsg
		chatMsg = ""
		fmt.Println("请输入您要发送的信息, exit退出")
		fmt.Scanln(&chatMsg)
	}
}

// 查询在线用户
func (client *Client) SelectOnlineUser() {
	sendMsg := "who\n"

	_, err := client.Conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("查询用户失败")
		return
	}
}

// 私聊模式
func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	// 查询在线用户
	client.SelectOnlineUser()

	fmt.Println("请选择你要私聊的对象，exit退出")
	fmt.Scanln(&remoteName)
	for remoteName != "exit" {
		fmt.Println("请输入您要发送的信息, exit退出")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			// 消息不为空则发送
			if len(chatMsg) != 0 {
				sendMsg := "to:" + remoteName + ":" + chatMsg + "\n"
				_, err := client.Conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("client.Conn write error:",err)
					break
				}
			}
			// 重置chatmsg
			chatMsg = ""
			fmt.Println("请输入您要发送的信息, exit退出")
			fmt.Scanln(&chatMsg)
		}
		remoteName = ""
		fmt.Println("请选择你要私聊的对象，exit退出")
		fmt.Scanln(&remoteName)
	}

}

// 处理服务器完成操作后的回调
func (client *Client) DoResponse() {
	// 一旦clinet.conn中有数据就直接copy到标准输出上，永久阻塞监听。效果与下面的for循环一样
	io.Copy(os.Stdout, client.Conn)

	// for {
	// 	buf := make([]byte, 4096)
	// 	client.Conn.Read(buf)
	// 	fmt.Println(string(buf))
	// }
	
}
func main() {

	// 命令行解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>>连接服务器失败....")
		return
	}
	fmt.Println(">>>>>>>>连接服务器成功!")
	
	//单独开一个goroutine处理服务器的回调
	go client.DoResponse()

	// 启动客户端任务
	client.run()
}