package database_utils // Package database_utils 最好与包文件同名

import (
	"fmt"
)

// GetConn 首字母大写定义的函数才能被其他包访问(变量也是) 注意 一个 pakege 文件夹下的文件中不能有同名函数
func GetConn() {
	fmt.Println("Connection Established")
}
