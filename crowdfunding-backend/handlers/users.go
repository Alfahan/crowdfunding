package handlers

import (
	"crowdfunding-backend/domains/users"
	"crowdfunding-backend/utils/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService users.Service
}

func NewUserHandler(userService users.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input users.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		errors := helpers.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Register account failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//token

	formatter := users.FormatUsers(newUser, "tokentokentoken")

	response := helpers.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
