package v1

import (
	"chat-room/internal/model"
	"chat-room/internal/service"
	"chat-room/pkg/common/response"
	"chat-room/pkg/global/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册
func Register(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)

	err := service.UserService.Register(&user)
	if err != nil {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessMsg(user))
}

// 登录
func Login(c *gin.Context) {
	var user model.User
	// c.BindJSON(&user)
	c.ShouldBindJSON(&user)
	log.Logger.Debug("user", log.Any("user", user))

	if service.UserService.Login(&user) {
		c.JSON(http.StatusOK, response.SuccessMsg(user))
		return
	}

	c.JSON(http.StatusOK, response.FailMsg("Login failed"))
}
