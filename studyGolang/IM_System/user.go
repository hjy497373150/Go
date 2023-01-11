package main

import "net"

type User struct {
	Name string
	Addr string
	C chan string
	Conn net.Conn // 用户信息的读写

	Server *MyServer
}

// 创建一个用户的接口
func NewUser(conn net.Conn, server *MyServer) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C : make(chan string),
		Conn: conn,
		Server: server,
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

// 封装上线功能
func (this *User)Online() {
	// 用户上线 将用户加到onlineMapUser中
	this.Server.mapLock.Lock()
	this.Server.onlineMapUser[this.Name] = this
	this.Server.mapLock.Unlock()

	// 广播告知所有在线用户 有人上线了
	this.Server.BoardCast(this,"已上线")
} 

// 下线
func (this *User)Offline() {
	this.Server.mapLock.Lock()
	delete(this.Server.onlineMapUser, this.Name) //将用户从map中删除
	this.Server.mapLock.Unlock()

	this.Server.BoardCast(this,"下线")
}

// 用户处理消息的业务
func (this *User)DoMessage(msg string) {
	this.Server.BoardCast(this, msg)
}
