package main

import "fmt"

func main() {
	var score int
	fmt.Print("Please input your score: ")
	fmt.Scanln(&score)

	switch score / 10 {
	// 一个case可以处理多个条件
	case 10, 9, 8:
		fmt.Println("Nice")
	case 7, 6:
		fmt.Println("Good")
	case 5:
		fmt.Println("Emm...")
		//在case中增加fallthrough将会在执行完成后执行下一个case
		fallthrough
	case 4, 3, 2, 1:
		fmt.Println("Fuck")
	// default 可以放在分支的任意位置
	default:
		fmt.Println("Sorry, I don't know how to do that.")
	}

	// case表达式中可以声明,定义变量,但是是不推荐的做法

}
