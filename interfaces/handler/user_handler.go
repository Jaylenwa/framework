package handler

import (
	"encoding/json"
	"fmt"
	"framework/domain/service"
	"framework/infra/utils/validate"
	"framework/interfaces"
	"framework/interfaces/handler/dto"
	"framework/port/driver"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type userHttpHandler struct {
	userService driver.UserService
}

var (
	httpUserOnce sync.Once
	httpUserHand interfaces.HttpRouterInterface
)

func NewHttpUserHandler() interfaces.HttpRouterInterface {
	httpUserOnce.Do(func() {
		httpUserHand = &userHttpHandler{
			userService: service.NewUserService(),
		}
	})
	return httpUserHand
}

// RegisterRouterPublic 注册外部API
func (h *userHttpHandler) RegisterRouterPublic(router *gin.RouterGroup) {
	router.GET("/user/:id", h.createUserById) // 查询UserById
	router.GET("/user", h.findUserList)       // 查询User列表
	router.POST("/user", h.createUser)        // 创建User
	router.PUT("/user/:id", h.updateUser)     // 修改User
	router.DELETE("/user/:id", h.delUser)     // 删除User
}

// RegisterRouterPrivate 注册内部API
func (h *userHttpHandler) RegisterRouterPrivate(router *gin.RouterGroup) {
}

func (h *userHttpHandler) createUserById(c *gin.Context) {

	var req dto.FindUserByIdReq

	if err := c.ShouldBindUri(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.userService.FindUserById(c, req.Id)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h *userHttpHandler) findUserList(c *gin.Context) {

	var reqForm dto.GetUserListReq

	if err := c.ShouldBindQuery(&reqForm); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := validate.Validate(reqForm)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var filter map[string]interface{}
	reqBytes, err := json.Marshal(&reqForm)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(reqBytes, &filter)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	total, res, err := h.userService.FindUserList(c, filter)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total":   total,
		"entries": res,
	})

}

func (h *userHttpHandler) createUser(c *gin.Context) {

	var req dto.CreateUserReq

	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := validate.Validate(req)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	id, err := h.userService.CreateUser(c, req)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Header("Location", fmt.Sprintf("%s/%d", c.FullPath(), id))
	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (h *userHttpHandler) updateUser(c *gin.Context) {

	var reqUri dto.UpdateUserByIdReq

	var req dto.UpdateUserReq

	if err := c.ShouldBindUri(&reqUri); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := validate.Validate(req)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = h.userService.UpdateUser(c, reqUri.Id, req)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *userHttpHandler) delUser(c *gin.Context) {

	var req dto.DelUsersReq

	if err := c.ShouldBindUri(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := h.userService.DelUser(c, req.Id)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Status(http.StatusNoContent)
}
