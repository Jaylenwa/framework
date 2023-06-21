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
	router.GET("/user/:id", h.getUserById) // 查询UserById
	router.GET("/user", h.getUserList)     // 查询User列表
	router.POST("/user", h.addUser)        // 创建User
	router.PUT("/user/:id", h.updateUser)  // 修改User
	router.DELETE("/user/:id", h.delUser)  // 删除User
}

// RegisterRouterPrivate 注册内部API
func (h *userHttpHandler) RegisterRouterPrivate(router *gin.RouterGroup) {
}

func (h *userHttpHandler) getUserById(c *gin.Context) {

	type reqPath struct {
		Id uint64 `validate:"required" uri:"id" json:"id"`
	}

	var req reqPath

	if err := c.ShouldBindUri(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.userService.GetUserById(c, req.Id)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h *userHttpHandler) getUserList(c *gin.Context) {

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

	total, res, err := h.userService.GetUserList(c, filter)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total":   total,
		"entries": res,
	})

}

func (h *userHttpHandler) addUser(c *gin.Context) {

	var req dto.AddUserReq

	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := validate.Validate(req)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	id, err := h.userService.AddUser(c, req)
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

	type reqPath struct {
		Id uint64 `validate:"required" uri:"id" json:"id"`
	}

	var reqUri reqPath

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

	type reqPath struct {
		Id uint64 `validate:"required" uri:"id" json:"id"`
	}

	var req reqPath

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
