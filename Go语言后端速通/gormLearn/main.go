package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gormLearn/entity"
	// "gorm.io/driver/sqlite"
	"log"
	"os"
)

/*
gorm 用于与SQL进行结构体与表结构映射
*/

//var db *gorm.DB

//func init() {
//
//	// 配置数据源
//	dsn := "root:Asashishi107QS.@tcp(8.218.247.195:3306)/gorm_learn?charset=utf8&parseTime=True&loc=Local"
//
//	// gorm处理MySQL链接和操作
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic(err.Error())
//	}
//
//	// 设置数据库连接池
//	sqlDB, _ := db.DB()
//
//	// 最大连接数
//	sqlDB.SetMaxIdleConns(25)
//	// 最大空闲连接数
//	sqlDB.SetMaxOpenConns(10)
//
//	//SetMaxOpenConns(n int)：
//	//功能：设置最大打开的连接数，包括正在使用的连接和连接池中的连接。
//	//默认值：0，表示不限制。
//	//示例：db.SetMaxOpenConns(100) 设置最大打开连接数为100。
//	//SetMaxIdleConns(n int)：
//	//功能：设置连接池中的最大空闲连接数。
//	//默认值：0，表示不保持空闲连接。
//	//示例：db.SetMaxIdleConns(10) 设置最大空闲连接数为10。
//	//SetConnMaxLifetime(d time.Duration)：
//	//功能：设置连接的最大生命周期，超过这个时间的连接将被关闭。
//	//示例：db.SetConnMaxLifetime(time.Hour) 设置连接的最大生命周期为1小时。
//	//SetConnMaxIdleTime(d time.Duration)：
//	//功能：设置连接的最大空闲时间，超过这个时间的空闲连接将被关闭。
//	//示例：db.SetConnMaxIdleTime(30 * time.Minute) 设置连接的最大空闲时间为30分钟。
//	//SetConnMaxIdleTime(d time.Duration)：
//	//功能：设置连接的最大空闲时间，超过这个时间的空闲连接将被关闭。
//	//示例：db.SetConnMaxIdleTime(30 * time.Minute) 设置连接的最大空闲时间为30分钟。
//	//Ping()：
//	//功能：验证数据库连接是否有效。通常在初始化时调用以确保连接池可用。
//	//示例：err := db.Ping() 如果连接失败，将返回错误。
//	//Conn()：
//	//功能：从连接池中获取一个连接。
//	//示例：conn, err := db.Conn(context.Background()) 获取一个连接。
//	//Close()：
//	//功能：关闭数据库连接池，释放所有资源。
//	//示例：db.Close() 关闭连接池
//
//}

//func GetDB() *gorm.DB {
//	return db
//}

// crud
// insert
func addRecord(db *gorm.DB) {

	// 结构体对象和表记录做映射

	// 添加老师对象
	teacher := entity.Teacher{
		Name: "Asashishi",
		Tno:  1072903224,
		Pwd:  "Asashishi107QS.",
	}
	fmt.Println(teacher)
	// 按照实例化对象的指针生成SQL语句,gorm会进行格式化并对数据进行回写 插入记录
	db.Debug().Create(&teacher)
	fmt.Println(teacher)

	// 班级
	class01 := entity.Class{
		Name:      "软件一班",
		Num:       78,
		TeacherID: 1,
	}
	class02 := entity.Class{
		Name:      "软件二班",
		Num:       70,
		TeacherID: 1,
	}
	class03 := entity.Class{
		Name:      "软件三班",
		Num:       45,
		TeacherID: 1,
	}
	// 批量插入记录
	classes := []entity.Class{class01, class02, class03}
	db.Create(&classes)

	//课程
	course01 := entity.Course{
		Name:      "计算机网络",
		Credit:    3,
		Period:    16,
		TeacherID: 1,
	}
	course02 := entity.Course{
		Name:      "C语言程序设计",
		Credit:    3,
		Period:    16,
		TeacherID: 1,
	}
	course03 := entity.Course{
		Name:      "Go语言程序设计",
		Credit:    3,
		Period:    16,
		TeacherID: 1,
	}
	// 批量插入记录
	courses := []entity.Course{course01, course02, course03}
	db.Create(&courses)

	// 多对多添加记录
	// 先查后增加
	var students []entity.Student
	db.Where("name in ?", []string{"计算机网络", "C语言程序设计"}).Find(&courses)
	fmt.Println(courses)
	student01 := entity.Student{
		Name:    "AsaNeon",
		Sno:     20901,
		Pwd:     "AsaNeon107",
		Gender:  0,
		ClassID: 1,
		Courses: courses,
	}
	// 新增记录
	db.Create(&student01)
	fmt.Println(student01)
	db.Find(&students)
	fmt.Println(students)

	student02 := entity.Student{
		Name:    "AsaNeonNeon",
		Sno:     30304,
		Pwd:     "AsaNeon107",
		Gender:  1,
		ClassID: 2,
	}

	db.Create(&student02)

	// Association Append 方法为存在的字段添加新的关联记录
	db.Model(&student02).Association("Courses").Append(courses)

	// 查询后绑定切片并添加
	var student03 = entity.Student{}
	db.First(&student03, "name = ?", "AsaNeon")
	fmt.Println(student03)
	// 表中需要有字段,查询到对应字段并赋值给实例化结构体对象之后将关联字段插入表关系模型
	db.Model(&student03).Association("Courses").Append(courses)

	// 解除多对多绑定关系
	// db.Model(&student03).Association("Courses").Delete(courses)
	// db.Model(&student03).Association("Courses").Clear()

	// 查询关联关系
	var courses_res []entity.Course
	db.Model(&student03).Association("Courses").Find(&courses_res)
}

// select
func selectRcord(db *gorm.DB) {

	// 查询全部记录
	var status []entity.Class
	db.Find(&status)
	// 查询后 返回值为结构体的数组对象
	fmt.Println(status)

	// 基于String的Where语句
	db.Where("teacher_id = ?", 1).Find(&status)
	fmt.Println(status)
	var total int64
	// Count()接受指针类型的int64
	db.Model(&entity.Teacher{}).Where("Name = ?", "Asashishi").Count(&total)
	fmt.Println(total)

	// 基于struct/map的where语句
	db.Where(entity.Class{Name: "软件一班", TeacherID: 1}).Find(&status)
	fmt.Println(status)
	db.Where(map[string]interface{}{"Name": "Asashishi", "TeacherID": 1}).Find(&status)
	fmt.Println(status)

	// 查询单条记录
	class := entity.Class{}
	db.Take(&class)
	fmt.Println(class)
	db.First(&class)
	fmt.Println(class)
	db.Last(&class)
	fmt.Println(class)

	// 其他查询
	// 查询指定记录的指定字段
	var course []entity.Course
	db.Select("name,credit").Where("teacher_id = ?", 1).Find(&course)
	fmt.Println(course)
	// 查询指定记录的并忽略指定字段
	db.Omit("teacher_id").Where("teacher_id = ?", 1).Find(&course)
	fmt.Println(course)
	// 分页查询略过前一条取两条
	db.Order("teacher_id desc").Limit(2).Offset(1).Find(&course)
	fmt.Println(course)
	// 分组查询, 使用结构体承接分组
	type GroupedCourse struct {
		TeacherID int
		Count     int
	}
	var grouped []GroupedCourse
	db.Model(&entity.Class{}).Select("teacher_id,Count(*) as count").Group("teacher_id").Having("Count > ?", 1).Scan(&grouped)
	fmt.Println(grouped)
	// ...
}

// Preload 预加载联表查询
func Preload(db *gorm.DB) {

	// 多对一
	// 查询-预加载关联数据
	var class entity.Class
	// 会将关联表数据值值付给Class结构体的属性
	db.Preload("Teacher").Where("teacher_id = ?", 1).Find(&class)
	// 直接指向结构体内的结构体属性
	fmt.Println(class.Teacher)
	// eg
	var students []entity.Student
	db.Preload("Class").Where("class_id = ?", 1).Find(&students)
	fmt.Println(students[0].Class.Num)

	// 多对多预加载查询
	db.Preload("Class").Preload("Courses").Where("class_id = ?", 1).Find(&students)
	fmt.Println(students[0].Class, students[0].Courses)

	// 嵌套预加载
	db.Preload("Class").Preload("Class.Teacher").Where("class_id = ?", 1).Find(&students)
	fmt.Println(students[0].Class.Teacher.Name)

	// 预加载所有直接关联
	db.Preload(clause.Associations).Where("teacher_id = ?", 1).Find(&students)
	fmt.Println(students[0])

	// 反向查询
	db.Preload("Students").Where("teacher_id = ?", 1).Take(&class)
	fmt.Println(class.Students)
}

// delete
func deleteRecord(db *gorm.DB) {
	var course entity.Course

	// 删除一条记录
	db.Where("name = ?", "计算机网络").Take(&course)
	fmt.Println(course)
	db.Delete(&course)

	// 按条件删除
	db.Where("Credit < ?", 3).Delete(&entity.Course{})

	// 全部删除
	db.Where("1 = 1").Delete(&entity.Course{})
}

// update
func updateRecord(db *gorm.DB) {

	// 先查后修改
	var course entity.Course
	db.Where("name = ?", "C语言程序设计").Take(&course)
	fmt.Println(course)
	course.Name = "R#程序设计"
	db.Save(&course)
	fmt.Println(course)

	// 更新所有字段
	db.Model(&entity.Course{}).Where("1 = 1").Update("Credit", 3)

	// 更新表达式 - 字段值自增1
	db.Model(&entity.Course{}).Where("1 = 1").Update("Credit", gorm.Expr("Credit + ?", 1))

	// 按条件更新字段
	db.Model(&entity.Course{}).Where("teacher_id", 1).Update("Credit", 5)

	// 通过结构体更新多个字段
	db.Model(&entity.Course{}).Where("teacher_id", 1).Updates(&entity.Course{Credit: 5, Period: 20})

	// 通过Map字典更新多个字段
	db.Model(&entity.Course{}).Where("teacher_id", 1).Updates(map[string]interface{}{"Credit": 5, "Period": 2})
}

func main() {

	// 定义数据源
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_learn?charset=utf8&parseTime=True&loc=Local"

	// 定义log对象并输出至控制台
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),

		// log.New() 函数：
		// log.New(writer io.Writer, prefix string, flag int) *log.Logger
		// 用于创建一个新的日志记录器（*log.Logger）。
		// 它接受三个参数：
		// 1. writer：日志输出的目标，可以是任何实现了 io.Writer 接口的对象。在这里是 os.Stdout，表示将日志输出到控制台（标准输出）。
		// 2. prefix：每行日志的前缀字符串。在此例中，使用了 "\r\n" 作为前缀，这会在每行日志之前添加换行符，但这并不是常见的用法。通常情况下，前缀可以是字符串标识符，比如 "[INFO] " 或 "[ERROR] " 来区分不同类型的日志信息。
		// 3. flag：日志的格式标志，控制日志记录器如何格式化日志信息。log.LstdFlags 是标准格式，包含日期和时间（log.Ldate 和 log.Ltime 的组合），即 "2009/01/23 01:23:23" 的格式。

		// log.LstdFlags 是一个常量，等价于 log.Ldate | log.Ltime
		// log.Ldate：在日志条目中添加日期部分，格式是 YYYY/MM/DD。
		// log.Ltime：在日志条目中添加时间部分，格式是 HH:MM:SS。

		logger.Config{

			//// 超时阈值
			//SlowThreshold: time.Second,

			LogLevel: logger.Info,
		})

	// gorm处理MySQL链接和操作

	db, err0 := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 日志配置
		Logger: newLogger,
	})

	//db, err0 := gorm.Open(sqlite.Open("gorm_learn.db"), &gorm.Config{
	//	// 日志配置
	//	Logger: newLogger,
	//})

	if err0 != nil {
		log.Println(err0.Error())
	}

	/*
		// 迁移表 如果表不存在则会自动创建 会按照结构体映射创建
		err1 := db.AutoMigrate(&entity.Teacher{})
		if err1 != nil {
			log.Println(err1.Error())
		}
		err2 := db.AutoMigrate(&entity.Class{})
		if err2 != nil {
			log.Println(err2.Error())
		}
		err3 := db.AutoMigrate(&entity.Course{})
		if err3 != nil {
			log.Println(err3.Error())
		}
		err4 := db.AutoMigrate(&entity.Student{})
		if err4 != nil {
			log.Println(err4.Error())
		}
	*/

	//addRecord(db)
	//selectRcord(db)
	Preload(db)
	//deleteRecord(db)
	//updateRecord(db)

}
