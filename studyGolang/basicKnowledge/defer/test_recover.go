// Go语言提供了专用于拦截运行时panic的内建函数recover。它可以是当前的程序从运行时panic的状态中恢复并重新获得流程控制权。
// 注意：recover只有在defer调用的函数中有效。
package main
import "fmt"


func Demo(i int) {
	//定义10个元素的数组
	var arr [10]int
	//错误拦截要在产生错误前设置
	defer func() {
		//设置recover拦截错误信息
		err := recover()
		//产生panic异常  打印错误信息
		if err != nil {
			fmt.Println(err)
		}
	}()
	//根据函数参数为数组元素赋值
	//如果i的值超过数组下标 会报错误：数组下标越界
	arr[i] = 10

}

func main() {
	Demo(10)
	//产生错误后 程序继续
	fmt.Println("程序继续执行...")
}