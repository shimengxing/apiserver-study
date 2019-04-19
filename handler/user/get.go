package user

import (
	"apiserver-study/handler"
	"apiserver-study/model"
	"apiserver-study/pkg/errno"
	"github.com/gin-gonic/gin"
)

//根据用户名查询用户信息
func Get(c *gin.Context) {
	username := c.Param("username")
	user, err := model.GetUser(username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
	}
	handler.SendResponse(c, nil, user)
}
