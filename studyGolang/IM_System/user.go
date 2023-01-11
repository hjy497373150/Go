package main

import (
	"net"
	"strings"
)

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
	// 查询当前在线用户
	if msg == "who" {
		this.Server.mapLock.Lock()
		for _,user := range this.Server.onlineMapUser {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":在线...\n"
			this.SendMessage(onlineMsg)
		}
		this.Server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename:" {
		// 更改用户名 消息格式规定为rename:
		newName := msg[7:]

		// 查询onlineMap中有无newName
		_, ok := this.Server.onlineMapUser[newName]
		if ok {
			this.SendMessage("当前用户名已被占用!")
		} else {
			this.Server.mapLock.Lock()
			// 删除原来的kv 并将newName加到用户列表中
			delete(this.Server.onlineMapUser,this.Name)
			this.Server.onlineMapUser[newName] = this
			this.Server.mapLock.Unlock()

			this.Name = newName
			this.SendMessage("您已更新用户名为" + newName + "\n")
		}
	} else if len(msg) > 4 && msg[0:3] == "to:" {
		// 私聊功能 消息格式为"to:张三:你好啊"
		// 1.获取用户名
		remoteName := strings.Split(msg, ":")[1] //使用split方法将msg切成 to 张三 你好啊 的字符串数组格式

		if remoteName == "" {
			this.SendMessage("消息格式不正确，请使用\"to:张三:你好啊\"格式. \n")
			return
		}
		// 2.根据用户名，获取对方的user
		remoteUser,ok := this.Server.onlineMapUser[remoteName]
		if !ok {
			this.SendMessage("该用户不存在\n")
			return
		}
		// 3.获取消息内容并发送给remoteUser
		msgContent := strings.Split(msg,":")[2]
		if msgContent == "" {
			this.SendMessage("消息内容为空，请重发\n")
			return
		}
		remoteUser.SendMessage(this.Name + "对您说:" + msgContent + "\n")

	} else {
		this.Server.BoardCast(this, msg)
	}
	
}

func (this *User)SendMessage(msg string) {
	// this.C <- msg 
	// 两种写法都可以
	this.Conn.Write([]byte(msg)) 
}
