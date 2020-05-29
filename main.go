package main

import (
	"github.com/gin-gonic/gin"
	"goweb-gin/src"
)

func main() {
	r := gin.Default()
	//路由组v1
	v1 := r.Group("/v1/topics")
	{
		v1.GET("", src.GetTopicList)

		v1.GET(":topic_id", src.GetTopicDetail) //查询某个id的帖子

		v1.Use(src.MustLogin()) //中间件：要求登录
		{
			v1.POST("", src.NewTopic)   //新增帖子
			v1.DELETE("/:topic_id", src.DelTopic) //删除某个id的帖子

		}
	}

	r.Run()
}
