package main

import "fmt"

//数组

func main() {
	// 创建数组
	var scores [5]int

	// 按引索插入数据
	scores[0] = 100
	scores[1] = 97
	scores[2] = 98
	scores[3] = 78
	scores[4] = 90

	var score int
	// 遍历数组
	for i := range scores {
		score += scores[i]
	}

	fmt.Println(score / len(scores))

}
