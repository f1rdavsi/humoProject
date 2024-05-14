package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type signInData struct {
	Email    string `json:"email" biding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type signUpData struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (h *Handler) SignUp(c *gin.Context) {
	var request signUpData

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.CreateUser(
		request.Firstname, request.Lastname, request.Username, request.Email, request.Password,
	)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"id":      id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var request signInData

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	token, err := h.service.GenerateToken(request.Email, request.Username, request.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
