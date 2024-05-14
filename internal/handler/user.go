package handler

import (
	"fmt"
	"github.com/f1rdavsi/reporter/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUser(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.service.GetUserById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var request signUpData
	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "JSON: "+err.Error())
		return
	}

	user, err := h.service.UpdateUser(
		id,
		request.Firstname, request.Lastname, request.Email, request.Username, request.Password,
	)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    user,
	})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.DeleteUserById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *Handler) RestoreUser(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.RestoreUserById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *Handler) ChangePictureUser(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	image, err := c.FormFile("picture")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	filePath, err := utils.GenFilenameWithDir(image.Filename)
	if err != nil {
		newErrorResponse(c, http.StatusUnsupportedMediaType, err.Error())
		return
	}

	user, err := h.service.UpdatePictureUser(id, filePath)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.SaveUploadedFile(image, fmt.Sprintf("./file/layouts/%s", filePath)); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    user,
	})

}

func (h *Handler) UploadUserPicture(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	image, err := c.FormFile("picture")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	filePath, err := utils.GenFilenameWithDir(image.Filename)
	if err != nil {
		newErrorResponse(c, http.StatusUnsupportedMediaType, err.Error())
		return
	}

	user, err := h.service.UploadUserPicture(id, filePath)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.SaveUploadedFile(image, fmt.Sprintf("./files/layouts/%s", filePath)); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    user,
	})
}
