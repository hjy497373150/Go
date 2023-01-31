// 观察者模式

package main

import (
	"fmt"

)

// ---------抽象层----------
// 抽象的观察者
type Listener interface {
	OnTeacherComing() //老师来了怎么办
}

// 抽象的通知者
type Notifier interface {
	AddListener(listener Listener) //添加要通知的对象（观察者）
	DeleteListener(listener Listener) //删除要通知的对象
	Notify() //通知
}

//----------实现层-----------
// 观察者学生
type Zhang3 struct {
	badthing string
}

func (z3 *Zhang3) OnTeacherComing() {
	fmt.Println("张三停止 ",z3.badthing)
}

func (z3 *Zhang3) DoBadThing() {
	fmt.Println("张三正在 ",z3.badthing)
}

type Li4 struct {
	badthing string
}

func (l4 *Li4) OnTeacherComing() {
	fmt.Println("李四停止 ",l4.badthing)
}

func (l4 *Li4) DoBadThing() {
	fmt.Println("李四正在 ",l4.badthing)
}

type Wang5 struct {
	badthing string
}

func (w5 *Wang5) OnTeacherComing() {
	fmt.Println("王五停止 ",w5.badthing)
}

func (w5 *Wang5) DoBadThing() {
	fmt.Println("王五正在 ",w5.badthing)
}

// 通知者班长
type ClassMonitor struct {
	listenerlist []Listener // 通知者需要知道所有的观察者对象
}

func (cm *ClassMonitor) AddListener(listener Listener) {
	cm.listenerlist = append(cm.listenerlist, listener)
}

func (cm *ClassMonitor) DeleteListener(listener Listener) {
	// 需要判断list中是否有要删除的listener
	for index,l := range cm.listenerlist {
		// 找到要删除的元素位置
		if l == listener {
			// 将删除的点前后连接
			cm.listenerlist = append(cm.listenerlist[:index],cm.listenerlist[index+1:]...)
			break
		}
	}
}

func (cm *ClassMonitor) Notify() {
	for _,l := range cm.listenerlist {
		l.OnTeacherComing()
	}
}

//---------业务逻辑层----------
func main() {
	zhang3 := new(Zhang3)
	zhang3.badthing = "抄作业"

	li4 := Li4 {
		badthing: "玩手机",
	}

	wang5 := Wang5 {
		badthing: "看小说",
	}

	classMonitor := new(ClassMonitor)
	classMonitor.AddListener(zhang3)
	classMonitor.AddListener(&li4)
	classMonitor.AddListener(&wang5)

	fmt.Println("上课铃响了，但老师还没来，大家都在...")
	zhang3.DoBadThing()
	li4.DoBadThing()
	wang5.DoBadThing()

	fmt.Println("这时候老师还没来，班长给大家使了个眼色...")
	classMonitor.Notify()
}