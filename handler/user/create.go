package user

import (
	"apiserver-study/handler"
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

func CreateUser(c *gin.Context) {
	var cr CreateRequest
	if err := c.Bind(&cr); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	admin2 := c.Param("username")
	log.Infof("url username: %s", admin2)

	desc := c.Query("desc")
	log.Infof("url key param desc: %s", desc)

	contentType := c.GetHeader("Content-type")
	log.Infof("Header Content-Type: %s", contentType)

	log.Debugf("username: [%s], password: [%s]", cr.Username, cr.Password)
	if cr.Username == "" {
		handler.SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("不能找到用户名")), nil)
		return
	}

	if cr.Password == "" {
		handler.SendResponse(c, errno.ErrPasswordNil, nil)
		return
	}

	rsp := CreateResponse{
		Username: cr.Username,
	}

	handler.SendResponse(c, nil, rsp)
}
