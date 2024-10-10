package socket_utils

import "fmt"

func Access() { //首字母大写定义的函数才能被其他包访问(变量也是) 注意 一个 pakege 文件夹下的文件中不能有同名函数
	fmt.Println("Tcp_Accessed")
}
