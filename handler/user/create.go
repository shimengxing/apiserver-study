package user

import (
	"apiserver-study/pkg/errno"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"net/http"
)

func Create(c *gin.Context) {
	var r struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var err error
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		return
	}

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)
	if r.Username == "" {
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found"+
			" in db: xx.xx.xx.xx")).Add("这是一段增加的错误信息")
		log.Errorf(err, "Get an error")
	}

	if errno.IsErrUserNotFound(err) {
		log.Debug("错误类型是 ErrUserNotFound")
	}

	if r.Password == "" {
		err = fmt.Errorf("密码为空")
	}

	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
