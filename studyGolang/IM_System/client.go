package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIp string
	ServerPort int
	Name string
	Conn net.Conn
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIp: serverIp,
		ServerPort: serverPort,
	}

	// 与服务器连接
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

func main() {

	// 命令行解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>>连接服务器失败....")
		return
	}
	fmt.Println(">>>>>>>>连接服务器成功!")

	// 启动客户端任务
	select {}
}