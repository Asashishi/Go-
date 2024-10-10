package main

import "fmt"

func main() {

	// 定义并初始化Map 类似于字典
	var Map map[string]string
	// 创建方式 1 指定长度
	Map = make(map[string]string, 10)

	Map["Ja"] = "Asashishi"
	fmt.Println(Map["Ja"])

	// 创建方式2
	TM := map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
		"4": "4",
		"5": "5",
		"6": "6",
		"7": "7",
	}
	fmt.Println(TM)

	/*
		创建方式3 不指定长度
		Map := map[any][any]
	*/

	// 增 改 map[key] = value
	// 删 delete(key,value)
	// 查 map[key]
	// 清空数组 遍历清空或者直接new一个同名数组
}
