package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type MyServer struct {
	IP string
	Port int
	// 在线用户列表 key为Name value是User
	onlineMapUser map[string] *User
	mapLock sync.RWMutex

	// 消息广播的channel
	Message chan string
}

// 创建一个Server
func NewServer(ip string, port int) *MyServer {
	server := &MyServer{
		IP : ip,
		Port : port,
		onlineMapUser: make(map[string]*User),
		Message: make(chan string),
	}
	return server
}

// 为MyServer定义start接口 启动服务器
func (this *MyServer) Start() {
	// socket listen
	listener,err := net.Listen("tcp", fmt.Sprintf("%s:%d",this.IP,this.Port))
	if err != nil {
		fmt.Println("net.Listen err:",err)
		return
	}
	// close listener Server
	defer listener.Close()

	// 启动监听message广播消息

	go this.listenMessager()

	// accept 接受连接请求
	for {
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:",err)
			continue
		}
		// do handler
		go this.Handler(conn)
	}

}

func (this *MyServer)Handler(conn net.Conn) {
	// fmt.Println("连接建立成功...")

	// 创建用户
	user := NewUser(conn, this)

	// 用户上线
	user.Online()

	// 判断用户是否活跃
	isActive := make(chan bool)

	// 接收客户端发送的消息，核心是启动一个针对客户端conn的读gorountine
	go func() {
		buf := make([]byte,4096) // 4K大小，超过会有问题
		for {
			n, err := conn.Read(buf) //把信息读在buf中
			if n == 0 {
				// 用户下线
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("read conn err:", err)
				return
			}

			// 提取用户的信息，去除\n
			msg := string(buf[:n-1])

			// 用户针对msg进行处理
			user.DoMessage(msg)

			isActive <-true //只要有发消息就说明该用户活跃
		}
	}()

	// 定时器监控当前用户是否活跃
	for {
		select {
		case <-isActive:
			// 当前用户活跃，应该重置定时器
			// 不做任何事情，激活select重置time即可
		
		case <-time.After(time.Second * 300):
			// 如果超过300s没发消息就要把用户踢出去
			user.SendMessage("由于你长时间没发消息,系统已将你强制踢出!\n")

			// 销毁使用的资源
			close(user.C)

			// 关闭连接
			conn.Close()

			return
		}
	}
}

// 广播消息的方法
func (this *MyServer)BoardCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":"  + msg

	this.Message <- sendMsg // 把sendMsg发给消息广播的channel
}

// 监听message广播消息，一旦 this.Message中有消息就通知全部的user
func (this *MyServer)listenMessager() {
	for {
		msg := <-this.Message

		// 将msg发送给全部的user
		this.mapLock.Lock()
		for _,user := range this.onlineMapUser {
			user.C <- msg // 然后就会调用User的listenMessage方法
		}
		this.mapLock.Unlock()
	}
}