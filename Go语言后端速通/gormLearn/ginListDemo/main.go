package main

import (
	"errors"
	"ginListDemo/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
)

var db *gorm.DB

func initDB() (err error) {

	// 定义log对象并输出至控制台
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			//// 超时阈值
			//SlowThreshold: time.Second,
			LogLevel: logger.Info,
			// 使用彩色日志
			Colorful: true,
		})

	dsn := "root:123456@tcp(127.0.0.1:3306)/gin_list?charset=utf8&parseTime=True&loc=Local"
	// 初始化全局的db
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Println(err)
	}
	return err
}

// CRUD

// InsertHandler 增 增加待办事项
func InsertHandler(ctr *gin.Context) {
	var todo entity.Todo
	// 获取请求参数
	err0 := ctr.ShouldBind(&todo)
	if err0 != nil {
		log.Println("Error", err0)
		ctr.JSON(200, gin.H{
			"code": 500,
			"msg":  "unsupported args",
		})
		return
	}
	// 业务逻辑
	err1 := db.Create(&todo).Error
	if err1 != nil {
		log.Println("Error", err1)
		ctr.JSON(200, gin.H{
			"code": 500,
			"msg":  "database error",
		})
		return
	}
	// 返回响应
	ctr.JSON(200, gin.H{
		"code": 200,
		"msg":  "succeed",
		"data": todo,
	})
}

// DeleteHandler 删 删除待办事项
func DeleteHandler(ctr *gin.Context) {

	// 请求参数
	todoID, err0 := strconv.Atoi(ctr.Param("id"))
	if err0 != nil {
		log.Println("Error", err0)
		ctr.JSON(200, gin.H{
			"code": 500,
			"msg":  "unsupported args",
		})
		return
	}

	// 业务逻辑
	err1 := db.First(&entity.Todo{}, todoID).Error
	if err1 != nil {
		log.Println("Error", err1)
		if errors.Is(err1, gorm.ErrRecordNotFound) {
			ctr.JSON(200, gin.H{
				"code": 500,
				"msg":  "record not found",
			})
		} else {
			ctr.JSON(200, gin.H{
				"code": 500,
				"msg":  "database error",
			})
		}
		return
	}

	err2 := db.Delete(&entity.Todo{}, todoID).Error
	if err2 != nil {
		log.Println("Error", err2)
		ctr.JSON(200, gin.H{
			"code": 500,
			"msg":  "delete error",
		})
	}

	// 返回响应
	ctr.JSON(200, gin.H{
		"code": 200,
		"msg":  "succeed",
	})
}

// UpdateHandler 改 修改待办事项
func UpdateHandler(ctr *gin.Context) {
	var todo entity.Todo
	// 获取请求参数
	err0 := ctr.ShouldBind(&todo)
	if err0 != nil {
		log.Println("Error", err0)
		ctr.JSON(200, gin.H{
			"code": 500,
			"msg":  "unsupported args",
		})
		return
	}

	// 业务逻辑
	// 检查数据是否存在
	err1 := db.Where("title = ?", todo.Title).First(&todo).Error
	if err1 != nil {
		log.Println("Error", err1)
		if errors.Is(err1, gorm.ErrRecordNotFound) {
			ctr.JSON(200, gin.H{
				"code": 500,
				"msg":  "record not found",
			})
		} else {
			ctr.JSON(200, gin.H{
				"code": 500,
				"msg":  "database error",
			})
		}
		return
	}
	err2 := db.Model(&todo).Where("title = ?", todo.Title).Update("status", todo.Status).Error
	if err2 != nil {
		log.Println("Error", err2)
		ctr.JSON(200, gin.H{
			"code": 500,
			"msg":  "database error",
		})
		return
	}

	// 返回响应
	ctr.JSON(200, gin.H{
		"code": 200,
		"msg":  "succeed",
		"data": todo,
	})
}

// SelectHandler 查 查询待办事项
func SelectHandler(ctr *gin.Context) {
	// 业务逻辑
	var todoList []entity.Todo
	err0 := db.Find(&todoList).Error
	if err0 != nil {
		log.Println("Error", err0)
		ctr.JSON(200, gin.H{
			"code": 500,
			"msg":  "database error",
		})
		return
	}
	// 返回响应
	ctr.JSON(200, gin.H{
		"code": 200,
		"msg":  "succeed",
		"data": todoList,
	})
}

func main() {
	// 链接数据库
	err0 := initDB()
	if err0 != nil {
		log.Println("Error", err0)
	}

	// 数据库迁移并进行映射
	//err1 := db.AutoMigrate(&entity.Todo{})
	//if err1 != nil {
	//	log.Println("Error", err1)
	//}

	// gin框架彩色日志
	gin.ForceConsoleColor()
	// 启动Server端
	Server := gin.Default()

	// 路由
	todo := Server.Group("/api/todo")
	{
		// 增加
		todo.POST("/insert", InsertHandler)
		// 删除
		todo.DELETE("/delete/:id", DeleteHandler)
		// 更新
		todo.PUT("/update", UpdateHandler)
		// 查询
		todo.GET("/select", SelectHandler)
	}

	// 启动http服务端
	err2 := Server.Run("0.0.0.0:80")
	if err2 != nil {
		log.Println("Error", err2)
	}
}
