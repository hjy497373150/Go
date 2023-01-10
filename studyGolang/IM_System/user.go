package main

import "net"

type User struct {
	Name string
	Addr string
	C chan string
	Conn net.Conn // 用户信息的读写
}

// 创建一个用户的接口
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C : make(chan string),
		Conn: conn,
	}
	// 启动监听当前user channel消息的goroutine

	go user.listenMessage()
	
	return user
}

// 监听当前用户的channel ，一旦有消息就发送给对端客户端
func (this *User) listenMessage() {
	for {
		msg := <-this.C

		this.Conn.Write([]byte(msg + "\n"))
	}
}