package handler

import (
	"fmt"
	"github.com/f1rdavsi/reporter/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type accountData struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func (h *Handler) CreateAccount(c *gin.Context) {
	var request accountData

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.CreateAccount(userId, request.Name, request.Balance)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"id":      id,
	})
}

func (h *Handler) GetAccount(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	account, err := h.service.GetAccount(id, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	hide := c.DefaultQuery("hide", "false")
	if hide == "true" {
		account.Balance = 0
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"account": account,
	})
}

func (h *Handler) GetAllAccounts(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	accounts, err := h.service.GetAllAccounts(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	hide := c.DefaultQuery("hide", "false")
	if hide == "true" {
		for i := range accounts {
			accounts[i].Balance = 0
		}
	}

	if len(accounts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "records not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    accounts,
	})
}

func (h *Handler) UpdateAccount(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))

	var request accountData
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	account, err := h.service.UpdateAccount(id, userId, request.Name, request.Balance)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    account,
	})
}

func (h *Handler) DeleteAccount(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	err = h.service.DeleteAccount(id, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *Handler) RestoreAccount(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	err = h.service.RestoreAccount(id, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *Handler) ChangePictureAccount(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))
	userId, err := getUserId(c)
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

	account, err := h.service.ChangePictureAccount(id, userId, filePath)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err = c.SaveUploadedFile(image, fmt.Sprintf("./files/layouts/%s", filePath)); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    account,
	})

}

func (h *Handler) UploadAccountPicture(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))
	userId, err := getUserId(c)
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

	account, err := h.service.UploadAccountPicture(id, userId, filePath)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err = c.SaveUploadedFile(image, fmt.Sprintf("./files/layouts/%s", filePath)); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    account,
	})
}
