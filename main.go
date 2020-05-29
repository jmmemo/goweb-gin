package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//路由组v1
	v1 := r.Group("/v1/topics")
	{
		v1.GET("", func(c *gin.Context) {
			if c.Query("username") == "" {
				c.String(200, "默认获取所有帖子列表")
			} else {
				c.String(200, "获取用户名=%s的帖子列表", c.Query("username"))
			}
		})

		v1.GET(":topic_id", func(c *gin.Context) {
			c.String(200, "获取用户id=%s的帖子", c.Param("topic_id"))
		})

	}

	r.Run()
}
