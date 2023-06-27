// internal/app/user/user_handler.go

package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mark-server/internal/domain"
	"mark-server/internal/helpers"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// CreateUser
// @Tags    users
// @Summary 创建用户
// @Accept   json
// @Produce  json
// @Param data body domain.UpdateUser true "User Info"
// @Success 200 {object} domain.CreateUserResp
// @Router /user/create [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	user := &domain.UpdateUser{}
	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.WriteResponse(c, 1001, err.Error(), nil)
		return
	}
	respUser, err := h.service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	helpers.WriteResponse(c, 0, "Success", respUser)
}

// GetUsers
// @Tags    users
// @Summary 获取用户列表
// @Produce  json
// @Param page query string true "page index"
// @Param size query string true "page size"
// @Success 200 {object} domain.UserList
// @Router /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		size = 10
	}

	users, total, err := h.service.GetUsers(page, size)
	if err != nil {
		helpers.WriteResponse(c, 1001, "Failed to get users", nil)
		return
	}

	helpers.WriteResponse(c, 0, "Success", gin.H{"users": users, "total": total})
}

// UpdateUser
// @Tags    users
// @Summary 创建用户
// @Accept   json
// @Produce  json
// @Param data body domain.UpdateUser true "User Info"
// @Success 200 {object} domain.CreateUserResp
// @Router /user/{id} [PUT]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.WriteResponse(c, 1001, "Invalid user id", nil)
		return
	}

	user := &domain.UpdateUser{}
	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.WriteResponse(c, 1001, err.Error(), nil)
		return
	}
	err = h.service.UpdateUser(user, userID)
	if err != nil {
		helpers.WriteResponse(c, 1001, err.Error(), nil)
		return
	}

	helpers.WriteResponse(c, 0, "User updated successfully", nil)
}

// DeleteUser
// @Tags    users
// @Summary 删除用户
// @Accept   json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} json "{"code":0,"data":"Success","message":""}"
// @Router /user/{id} [DELETE]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.WriteResponse(c, 1001, "Invalid user id", nil)
		return
	}

	err = h.service.DeleteUser(userID)
	if err != nil {
		helpers.WriteResponse(c, 1001, "Failed to delete user", nil)
		return
	}
	helpers.WriteResponse(c, 0, "", "Success")
}
