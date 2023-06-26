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
