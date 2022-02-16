package handler

import (
	"govue/helpers"
	"govue/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegistUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIresponse("Create Account Failed!", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helpers.APIresponse("Create Account Failed!", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
	}

	formatter := user.FormatUser(newUser, "tokentokentoken")
	response := helpers.APIresponse("Account has been registered", http.StatusOK, "Success", formatter)

	c.JSON(http.StatusOK, response)

}
