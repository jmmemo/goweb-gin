package src

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTopicDetail(c *gin.Context) {
	c.String(200, "获取用户id=%s的帖子", c.Param("topic_id"))
}
func NewTopic(c *gin.Context) {
	c.String(200, "新增帖子")
}

func DelTopic(c *gin.Context) {
	c.String(200, "删除帖子")
}
func GetTopicList() {

}

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
