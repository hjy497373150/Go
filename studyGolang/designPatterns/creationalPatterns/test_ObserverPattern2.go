// 观察者模式2 更为复杂的demo 江湖百晓生
package main

import "fmt"


const (
	PGaibang string = "丐帮"
	PMingJiao string = "明教"
)
//--------抽象层----------
// 江湖事件
type Event struct {
	Noti HeroNotifier //被知晓的通知者
	One HeroListener // 打人者
	Another HeroListener // 被打者
	Msg string // 具体消息
}

type HeroListener interface {
	// 观察者英雄
	OnFriendBeFighted(event *Event) //同帮派成员被揍了怎么办？
	GetInfo() string //返回英雄信息
	GetName() string //返回姓名
	GetParty() string //返回帮派
}

type HeroNotifier interface {
	// 添加观察者
	AddHeroListener(listener HeroListener)
	// 删除观察者
	DeleteHeroListener(listener HeroListener)
	// 广播事件
	Notify(event *Event) 
}

//-------------实现层----------
// 具体英雄
type Hero struct {
	Name string
	Party string // 帮派
}

func (h *Hero) GetInfo() string {
	return fmt.Sprintf("[%s]%s", h.Party, h.Name) 
}

func (h *Hero) GetName() string {
	return h.Name
}

func (h *Hero) GetParty() string {
	return h.Party
}

func (h *Hero) OnFriendBeFighted(event *Event) {
	// 1.判断自己是不是当事人
	if h.Name == event.One.GetName() || h.Name == event.Another.GetName() {
		return // 忽略
	}
	// 2.判断是否本帮成员被揍
	if h.Party == event.Another.GetParty() {
		fmt.Println(h.GetInfo(), "得知消息，发起报仇反击！")
		h.Fight(event.One,event.Noti) //循环调用
		return
	}
	// 3.判断是否本帮成员揍了别人
	if h.Party == event.One.GetParty() {
		fmt.Println(h.GetInfo(), "得知消息，拍手叫好！")
		return
	}
}

func (h *Hero) Fight(another HeroListener, baixiao HeroNotifier) {
	msg := fmt.Sprintf("%s 把 %s 揍了...",h.GetInfo(),another.GetInfo())

	// 生成江湖事件
	event := new(Event)
	event.Noti = baixiao
	event.One = h
	event.Another = another
	event.Msg = msg

	// 百晓生广播事件
	baixiao.Notify(event)
}

//百晓生
type BaiXiao struct {
	herolist []HeroListener
}

func (bx *BaiXiao) AddHeroListener(listener HeroListener) {
	bx.herolist = append(bx.herolist, listener)
}

func (bx *BaiXiao) DeleteHeroListener(listener HeroListener) {
	// 需要判断list中是否有要删除的listener
	for index,l := range bx.herolist {
		// 找到要删除的元素位置
		if l == listener {
			// 将删除的点前后连接
			bx.herolist = append(bx.herolist[:index],bx.herolist[index+1:]...)
			break
		}
	}
}

func (bx *BaiXiao) Notify(event *Event) {
	fmt.Println("[世界消息] 百晓生广播消息: ",event.Msg)
	for _,l := range bx.herolist {
		//依次调用全部观察的具体动作
		l.OnFriendBeFighted(event)
	}
}


func main() {
	hero1 := Hero{
		"黄蓉",
		PGaibang,
	}

	hero2 := Hero{
		"洪七公",
		PGaibang,
	}

	hero3 := Hero{
		"乔峰",
		PGaibang,
	}

	hero4 := Hero{
		"张无忌",
		PMingJiao,
	}

	hero5 := Hero{
		"金毛狮王",
		PMingJiao,
	}

	hero6 := Hero{
		"杨逍",
		PMingJiao,
	}

	baixiao := new(BaiXiao)

	baixiao.AddHeroListener(&hero1)
	baixiao.AddHeroListener(&hero2)
	baixiao.AddHeroListener(&hero3)
	baixiao.AddHeroListener(&hero4)
	baixiao.AddHeroListener(&hero5)
	baixiao.AddHeroListener(&hero6)
	
	fmt.Println("武林一片平静...")
	hero1.Fight(&hero4,baixiao)

}