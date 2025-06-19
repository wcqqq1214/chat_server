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

// 修改用户信息
func ModifyUserInfo(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)
	log.Logger.Debug("user", log.Any("user", user))

	if err := service.UserService.ModifyUserInfo(&user); err != nil {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessMsg(nil))
}

// 获取用户的详细信息
func GetUserDetails(c *gin.Context) {
	uuid := c.Param("uuid")

	c.JSON(http.StatusOK, response.SuccessMsg(service.UserService.GetUserDetails(uuid)))
}

// 通过名称查找群组或者用户
func GetUserOrGroupByName(c *gin.Context) {
	name := c.Query("name")

	c.JSON(http.StatusOK, response.SuccessMsg(service.UserService.GetUserOrGroupByName(name)))
}

// 获取用户列表
func GetUserList(c *gin.Context) {
	uuid := c.Query("uuid")

	c.JSON(http.StatusOK, response.SuccessMsg(service.UserService.GetUserList(uuid)))
}
