package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func png(c *gin.Context) {
	fmt.Println("ok")
}
func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 处理静态文件(这样处理后data文件夹里面的文件就可以被加载到浏览器中了)
	// 例如：http://localhost:5004/pic/1.jpg
	r.Static("/data", "../data/")
	// 创建一个get请求
	r.GET("png_", png)
	fmt.Println("启动成功！")
	r.Run(":3000")
}
