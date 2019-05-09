package user

import (
	"apiserver-study/handler"
	"apiserver-study/model"
	"apiserver-study/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Update(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))

	//绑定用户数据
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u.CreatedAt = time.Now()

	u.Id = uint64(userId)

	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.Errvalidation, nil)
		return
	}

	//加密密码
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Update(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, nil)

}
