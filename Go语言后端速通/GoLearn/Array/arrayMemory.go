package main

import "fmt"

// 数组在定义时会开辟连续的数组长度的内存空间
var arr [15]int

func main() {

	fmt.Println(len(arr))
	// arr的初始值为0
	fmt.Println(arr)
	// arr地址指向数组起始地址
	fmt.Printf("%p\n", &arr)

	// 查看内存
	fmt.Printf("%p\n", &arr[1])
	fmt.Printf("%p\n", &arr[2])
	fmt.Printf("%p\n", &arr[3])
	fmt.Printf("%p\n", &arr[4])
	fmt.Printf("%p\n", &arr[5])

	// 按照指针访问并修改值,速度快
	// arr[0] = 1
	// arr[1] = 2
	// arr[2] = 3
	// arr[3] = 4
	// arr[4] = 5
	// arr[5] = 6

	for i := 0; i < len(arr); i++ {
		fmt.Print("Please input data：")
		// _ 代表不使用时的占位变量
		_, ok := fmt.Scanln(&arr[i])
		if ok != nil {
			return
		}
	}
	for j := range arr {
		fmt.Println(arr[j])
	}

}
