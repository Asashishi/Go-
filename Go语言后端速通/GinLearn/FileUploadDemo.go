package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// Go语言中间件 类似于Flask的before_request
func StatCost() gin.HandlerFunc {
	// 记录请求耗时
	return func(ctl *gin.Context) {
		start := time.Now()
		// 可以通过c.Set在请求上下文中设置值, 后续的处理函数能够取到该值
		//ctl.Set("isLogin", true)
		// Next()调用该请求的剩余处理程序
		ctl.Next()
		// c.Abort() 终止执行
		cost := time.Since(start)
		log.Println("cost:", cost)
	}
}

func main() {
	gin.ForceConsoleColor()
	Server := gin.Default()
	// 加载模板
	Server.LoadHTMLGlob("templates/*")
	// 调用中间件
	Server.GET("/upload", StatCost(), func(ctl *gin.Context) {
		ctl.HTML(200, "upload.html", nil)
	})

	Server.POST("/upload/fileHandler", func(ctl *gin.Context) {
		// 处理MultiplepartForm时的最大内存大小限制默认为32Mib 可以使用 Server.MaxMultipartMemory = 128 << 20(128右移20位) 来修改
		form, err0 := ctl.MultipartForm()
		if err0 != nil {
			log.Println(err0)
			ctl.JSON(200, gin.H{
				"status":  500,
				"message": "Server error",
			})
			return
		}
		files := form.File["file"]
		for _, file := range files {
			// 指定文件保存路径
			dst := fmt.Sprintf("files/%v", file.Filename)
			// 保存文件
			err0 := ctl.SaveUploadedFile(file, dst)
			if err0 != nil {
				log.Println(err0)
				ctl.JSON(200, gin.H{
					"status":  400,
					"message": "Server save error please try again",
				})
				return
			}
		}
		ctl.JSON(200, gin.H{
			"status":  200,
			"message": "Files saved successfully",
		})
	})

	// 上传单个文件
	//Server.POST("/upload/fileHandler", func(ctl *gin.Context) {
	//	file, err0 := ctl.FormFile("file")
	//	if err0 != nil {
	//		log.Println(err0)
	//		ctl.JSON(200, gin.H{
	//			"status":  500,
	//			"message": "Server error",
	//		})
	//		return
	//	}
	//	dst := fmt.Sprintf("files/%v", file.Filename)
	//	err1 := ctl.SaveUploadedFile(file, dst)
	//	if err1 != nil {
	//		log.Println(err1)
	//		ctl.JSON(200, gin.H{
	//			"status":  400,
	//			"message": "Server save error please try again",
	//		})
	//	}
	//	ctl.JSON(200, gin.H{
	//		"status":  200,
	//		"message": "Files saved successfully",
	//	})
	//})

	err := Server.Run("0.0.0.0:80")
	if err != nil {
		log.Fatal(err)
	}
}
