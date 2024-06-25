package handlers

import (
	"crowdfunding-backend/domains/users"
	"crowdfunding-backend/utils/helpers"
	"fmt"
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

func (h *userHandler) LoginUser(c *gin.Context) {
	var input users.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Login failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loggedInUser, err := h.userService.Login(input)

	if err != nil {
		errors := helpers.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Login failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := users.FormatUsers(loggedInUser, "tokentokentoken")

	response := helpers.APIResponse("Login succeeded", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input users.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errors := helpers.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_available": isEmailAvailable}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helpers.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helpers.APIResponse("Failed to upload avatar images", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//harusnya dapat dari jwt
	userID := 1

	//path := "images/" + file.Filename
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helpers.APIResponse("Failed to upload avatar images", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helpers.APIResponse("Failed to upload avatar images", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helpers.APIResponse("Avatar successfuly uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}
