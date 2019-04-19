package user

import (
	"apiserver-study/handler"
	"apiserver-study/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		handler.SendResponse(c, err, nil)
	}
	handler.SendResponse(c, nil, nil)
}
