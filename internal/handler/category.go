package handler

import (
	"fmt"
	"github.com/f1rdavsi/reporter/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type catData struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (h *Handler) CreateCategory(c *gin.Context) {
	var request catData

	if err := c.ShouldBindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	cat, err := h.service.CreateCategory(request.Name, request.Description, request.Price)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data":    cat,
	})
}

func (h *Handler) GetCategories(c *gin.Context) {
	cats, err := h.service.GetCategories()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    cats,
	})
}

func (h *Handler) GetCategory(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))

	cat, err := h.service.GetCategory(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    cat,
	})
}

func (h *Handler) UpdateCategory(c *gin.Context) {
	var request catData

	id := cast.ToInt(c.Param("id"))

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	cat, err := h.service.UpdateCategory(id, request.Name, request.Description, request.Price)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    cat,
	})
}

func (h *Handler) UploadPictureCategory(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))

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

	cat, err := h.service.UploadPictureCategory(id, filePath)
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
		"data":    cat,
	})
}

func (h *Handler) ChangePictureCategory(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))

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

	cat, err := h.service.ChangePictureCategory(id, filePath)
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
		"data":    cat,
	})
}

func (h *Handler) DeleteCategory(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))

	if err := h.service.DeleteCategory(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *Handler) RestoreCategory(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))

	if err := h.service.RestoreCategory(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
