package src

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MustLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, b := c.GetQuery("token"); !b {
			c.String(http.StatusUnauthorized, "没有token参数,无法继续完成操作")
			c.Abort()
		} else {
			c.Next()
		}
	}
}
func GetTopicDetail(c *gin.Context) {
	//c.String(200, "获取用户id=%s的帖子", c.Param("topic_id"))
	c.JSON(200, Create(101, "帖子标题"))
}
func NewTopic(c *gin.Context) {
	c.String(200, "新增帖子")
}

func DelTopic(c *gin.Context) {
	c.String(200, "删除帖子")
}
func GetTopicList(c *gin.Context) {
	//if c.Query("username") == "" {
	//	c.String(200, "默认获取所有帖子列表")
	//} else {
	//	c.String(200, "获取用户名=%s的帖子列表", c.Query("username"))
	//}
	query := TopicQuery{}
	err := c.BindQuery(&query)
	if err != nil {
		c.String(400, "参数错误%v", err)
	} else {
		c.JSON(200, query)
	}
}
