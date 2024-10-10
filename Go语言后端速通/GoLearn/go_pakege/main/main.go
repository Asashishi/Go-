package main // pakege 是进行包的声明 包的声明最好与当前文件夹同名 mian包为程序入口 与C类似

import (
	"fmt"
	"go_pakege/database_utils"   // 导入包 如有多个包 建议一次性导入
	tcp "go_pakege/socket_utils" // 可以给包取别名 但取别名之后不能用原名进行函数调用
)

func main() {

	fmt.Println("Enter process")

	database_utils.GetConn() // 使用包中的函数必须使到包名定位

	tcp.Access() // 别名调用
}
