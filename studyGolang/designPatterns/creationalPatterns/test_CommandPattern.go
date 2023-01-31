// 命令模式

package main

import "fmt"

// 医生，命令的接受者
type Doctor struct {

}

func (d *Doctor) treatEye() {
	fmt.Println("医生治疗眼睛")
}

func (d *Doctor) treatNose() {
	fmt.Println("医生治疗鼻子")
}

// 护士，命令的收集者，也是调用命令者
type Nurse struct {
	CmdList []Command //收集的命令集合
}

// 发送病单，发送命令的方法
func (n *Nurse)Notify() {
	if n.CmdList == nil {
		return // 病单为空，直接返回
	}
	for _,cmd := range n.CmdList {
		cmd.Treat()
	}
}

// 抽象的命令(病单)
type Command interface {
	Treat() //执行病单绑定的命令(这里会调用病单已经绑定的医生的诊断方法)
}

//治疗眼睛的病单
type CommandTreatEye struct {
	doctor *Doctor
}

// 实现了command接口
func (cmd *CommandTreatEye) Treat() {
	cmd.doctor.treatEye()
}

//治疗鼻子的病单
type CommandTreatNose struct {
	doctor *Doctor
}

func (cmd *CommandTreatNose) Treat() {
	cmd.doctor.treatNose()
}

// 业务层，病人
func main() {
	//依赖病单，通过填写病单，让医生看病
	doctor := new(Doctor)

	// 治疗眼睛的病单
	cmdEye := CommandTreatEye{doctor: doctor}

	// 治疗鼻子的病单
	cmdNose := CommandTreatNose{doctor: doctor}

	// 护士
	nurse := new(Nurse)

	nurse.CmdList = append(nurse.CmdList, &cmdEye, &cmdNose)

	nurse.Notify()

}