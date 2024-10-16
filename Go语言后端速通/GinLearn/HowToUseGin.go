package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

// gin demo
func main() {

	// gin框架彩色日志
	gin.ForceConsoleColor()
	Server := gin.Default()

	/*RESTful API 基于不同操作分割不同的接口
	* GET 用于获取资源
	* POST 用于新建资源
	* PUT 用于更新资源
	* DELETE 用于删除资源
	* Any GO语言提供的方法, 接受所有请求, 在函数中做不同处理
	* 如：
	*		Server.Any("/" ,func(c *gin.Context){
	*			if c.Request.Method == http.MethodGet ...
	*		})
	 */

	/*后端模板渲染数据
	* Gin框架使用LoadHTMLGlob()或者LoadHTMLFiles()进行模板渲染
	* 如:
	*		Server.LoadHTMLGlob("templates/*")
	*		Server.POST("/upload", func(ctl *gin.Context) {
	*			ctl.HTML(http.StatusOK, "upload.html",nil)
	*		})
	 */

	Server.GET("/hello", func(c *gin.Context) {
		/*返回Json
		* 方法一:
		* 		Server.GET("/", func(c *gin.Context){
		*			var Msg struct{
		*				Message string `json:message`
		*			}
		*			Mag.Message = "Asashishi"
		*			Server.JSON(200,Msg)
		*		})
		* 方法二:
		 */
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	/*获取Json格式数据
	* Server。POST("/" func(c *gin.Context){
	* 	data, _ := c.GetRawData()
	*   // 定义map或结构体映射数据
	*   var dataMap map[string]interface{}
	*	// 反序列化
	* 	_ = json.Unmarshal(data, &dataMap)
	*	c.Json(200, dataMap)
	* })
	 */

	/*返回XML
	* Server.GET("/", func(c *gin.Context){
	* 	c.XML(200, gin.H{"message": "XML"})
	* })
	 */

	/*返回YMAL
	* Server.GET("/", func(c *gin.Context){
	* 	c.YAML(200, gin.H{"message": "ok","status": 200})
	* })
	 */

	/*返回pb(protobuf)
	* Server.GET("/", func(c *gin.Context){
	* 	reps := []int64{int64(1), int64(2)}
	* 	label := "test"
	* 	data := &protoexample.Test{
	*		Label: &label,
	*		Reps: reps,
	*	}
	*   c.ProtoBuf(200, data)
	* })
	 */

	/*获取querystring(http查询字符串)参数
	* Server.GET("/", func(c *gin.Context){
	* 	username := c.DefaultQuery("username","value")
	*   address := c.Query("address")
	* 	c.JSON(200, gin.H{"username": username,"address": address})
	* }
	 */

	/*获取path(如/value1/value2)参数
	* Server.GET("/:username/:address", func(c *gin.Context){
	* 	username := c.Param("username")
	* 	address := c.Param("address")
	*	c.JSON(200, gin.H{"username": username,"address": address})
	* }
	 */

	/*获取form表单
	* Server.POST("/", func(c *gin.Context){
	* 	username := c.PostForm("username")
	* 	address := c.PostForm("address")
	*  	c.JSON(200, gin.H{"username": username,"address": address})
	* })
	 */

	/*解析文件
	* c.FormFile("file") 解析单个文件
	* c.MultipartForm() 解析多个文件
	 */

	/*重定向
	* 方法一:
	* 	Server.GET("/", func(c *gin.Context) {
	*		c.Redirect(301,"http://...")
	*	})
	* 方法二:
	* 	Server.GET("/", func(c *gin.Context){
	*   	c.Request.URL.Path = "/"(本地路由)
	* 		Server.HandleContext(c)
	* 	})
	 */

	/*路由组
	* userGroup := Server.Group("/user")
	* {
	* 	userGroup.GET("/index",func(c *gin.Context){...})
	*   userGroup.POST("/...",func(c *gin.Context){...})
	*   // 路由组支持嵌套
	* 	subGroup := userGroup("/sub")
	*   {
	*		// 访问 user/sub/...
	*		subGroup.PUT("/...",func(c *gin.Context){...})
	*	}
	* }
	* */

	/*中间件
	* 定义中间件 类似于Flask的before_request()
	* func StatCost() gin.HandlerFunc{
	*	return func(c *gin.Context){
	*		start := time.Now()
	*  		// 执行下一个处理函数
	*       c.Next()
	*       // c.Abort() 终止处理, 不执行以下部分
	*		cost := time.Since(start)
	* 		log.Println(cost)
	* 	}
	* Server.POST("/", StatCost(), func(c *gin.Context){...}
	*
	* 为全局路由注册中间件, 即所有路由默认使用该中间件
	* Server.Use(StatCost)
	*
	* 为路由组注册中间件
	* userGroup := Server.Group("/user",StatCost()) 或者 userGroup := userGroup := Server.Group("/user"){userGroup.Use(StatCost())}
	* 注意 gin中间件中启用goroutine应当使用拷贝的上下文即c.Copy()而不是c *gin.Context
	 */

	/*gin.New()
	* 不启用Logger()和Recovery()的情况下创建一个http路由服务器对象
	* 等于去掉以上两个中间件的gin.Default()
	 */

	// 参数绑定
	type User struct {
		UserName string `form:"username" json:"username" binding:"required"`
		PassWord string `form:"password" json:"password" binding:"required"`
	}
	Server.POST("/post", func(c *gin.Context) {
		var user User
		// ShouldBind()会根据Header中的Content-Type自行选择绑定器来映射数据, 如果数据名称和类型对不上, 则可根据错误自定义返回信息
		// 如果是GET请求,会使用Form/query绑定引擎
		// 如果是POST请求, 首先检查content-type是否为JSON或者XML, 都不是则使用form-data
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(200, gin.H{
				"status":  403,
				"message": "args error",
			})
			return
		}
		c.JSON(200, gin.H{
			"status":   200,
			"username": user.UserName,
			"password": user.PassWord,
		})
	})

	err := Server.Run("0.0.0.0:80")
	if err != nil {
		log.Panicln(err)
	}

	/*启用多个服务
	* var (
	* 	EG errgroup.Group
	* )
	* func Router0() http.Handler{
	*	router := gin.Default
	* 	router.GET("/", func(c *gin.Context){...})
	* 	return router
	* }
	* func Router1() http.Handler{
	*	router := gin.Default
	* 	router.GET("/", func(c *gin.Context){...})
	* 	return router
	* }
	* Server0 := &http.Server{
	* 	Addr: ":8080",
	*	Handler: Router0(),
	*	ReadTimeout: 5 * time.Second,
	*	WriteTimeout: 10 * time.Second
	* }
	* Server1 := &http.Server{
	* 	Addr: ":8081",
	*	Handler: Router1(),
	*	ReadTimeout: 5 * time.Second,
	*	WriteTimeout: 10 * time.Second
	* }
	* EG.Go(func() err{
	* 	return server0.ListenAndServe（）
	* })
	* EG.Go(func() err{
	* 	return server1.ListenAndServe（）
	* })
	* err := EG.Wait()
	* if err != nil{
	*	log.Fatal(err)
	* }
	 */
}
