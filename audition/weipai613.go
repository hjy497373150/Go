// 给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。

// 输入: temperatures = [73,74,75,71,69,72,76,73]
// 输出: [1,1,4,2,1,1,0,0]

// 输入: temperatures = [30,40,50,60]
// 输出: [1,1,1,0]

// 输入: temperatures = [30,60,90]
// 输出: [1,1,0]

package main

import "fmt"

func GetNextTemperature1(nums []int) []int {
	res := make([]int, len(nums))
	

	for i := 0;i < len(nums);i++ {
		for j := i+1; j < len(nums);j++ {
			// 如果到最后一天还没有更高的温度
			if j == len(nums)-1 && nums[i] >= nums[j] {
				res[i] = 0
			}
			if nums[j] > nums[i] {
				res[i] = j-i
				break
			}
		}
	}
	return res
}

func GetNextTemperature2(nums []int) []int {
	res := make([]int, len(nums))
	// 存放当前最高温度对应的下标
	stk := make([]int, 0)
	stk = append(stk, 0)
	for i := 1;i < len(nums);i++ {
		// 如果当前访问的温度高于栈中最高温度
		for len(stk) > 0 && nums[i] > nums[stk[len(stk)-1]] {
			res[stk[len(stk)-1]] = i - stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)
	}
	return res
}

func main() {
	temperatures := []int{73,74,75,71,69,72,76,73}
	temperatures1 := []int{30,40,50,60}
	temperatures2 := []int{30,60,90}
	res := GetNextTemperature2(temperatures)
	res1, res2 := GetNextTemperature2(temperatures1), GetNextTemperature2(temperatures2)
	fmt.Println(res)
	fmt.Println(res1)
	fmt.Println(res2)

}