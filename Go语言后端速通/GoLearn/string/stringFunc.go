package main

// 字符串处理函数

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "Golang"

	// 内置函数len 仅用于字符串,返回字符串字节长度
	fmt.Println(len(str))

	// 字符串遍历
	// 方式1 利用键值循环 for-range
	jaStr := "こにちは!"
	for i, value := range jaStr {
		// 打印索引值以及字符的编码值
		fmt.Println(i, value)
	}
	// 方式2 利用切片
	r := []rune(jaStr)
	for i := 0; i < len(r); i++ {
		// 使用类型转换进行字符内容输出
		fmt.Println(string(r[i]))
	}

	// 字符串转整型 来自strconv
	strAtoi, err := strconv.Atoi("1024")
	if err != nil {
		return
	}
	fmt.Println(strAtoi)

	// 整型转字符串 来自strconv
	intItoA := strconv.Itoa(1024)
	println(intItoA)

	// 计算一个子字符串在另一个字符串中出现的次数 来自strings 返回 int 类型
	count := strings.Count("JavaEE", "J")
	fmt.Println(count)

	// 不区分大小写进行字符串比较 来自strings 返回true 区分大小写直接用 ==
	fmt.Println(strings.EqualFold("hello", "Hello"))

	// 返回子字符串在字符串中第一次出现的索引值,没有则返回-1
	fmt.Println(strings.Index("JavaEE", "J"))
}
