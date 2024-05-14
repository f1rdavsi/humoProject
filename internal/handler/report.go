package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type reportData struct {
	From     int    `json:"from_id"`
	To       int    `json:"to_id"`
	ToType   string `json:"to_type"`
	Limit    int    `json:"limit"`
	Page     int    `json:"page,omitempty"`
	Type     string `json:"type"`
	DateFrom string `json:"date_from,omitempty"`
	DateTo   string `json:"date_to,omitempty"`
}

func (h *Handler) GetReport(c *gin.Context) {
	var request reportData

	if err := c.ShouldBindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	page := cast.ToInt(c.Param("page"))
	if page != 0 {
		if request.Page == 0 {
			request.Page = page
		}
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	accountFrom, err := h.service.GetAccount(request.From, userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("not found account with id %v", request.From))
		return
	}

	accountTo, err := h.service.ExistsAccount(request.To)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("not found account with id %v", request.To))
		return
	}

	if request.Type != "" && request.Type != "expense" && request.Type != "income" {
		newErrorResponse(c, http.StatusBadRequest, "incorrect transaction type")
		return
	}

	if request.ToType != "" && request.ToType != "account" && request.ToType != "category" {
		newErrorResponse(c, http.StatusBadRequest, "incorrect `to_type` type")
		return
	}

	from := time.Time{}
	to := time.Time{}
	if request.DateFrom != "" {
		from, err = time.Parse("02-01-2006", request.DateFrom)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	if request.DateTo != "" {
		to, err = time.Parse("02-01-2006", request.DateTo)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	reports, err := h.service.GetReport(
		request.From, request.To, request.ToType,
		request.Limit, request.Page, request.Type,
		from, to,
	)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(reports) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "there no transaction yet",
		})
		return
	}

	ext := c.Query("ext")
	if ext != "" {
		userFrom, err := h.service.GetUserById(userId)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		userTo, err := h.service.GetUserById(request.To)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		switch ext {
		case "xlsx":
			excel, err := h.service.GetExcelReport(userFrom, userTo, accountFrom, accountTo, reports)
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
			buffer := new(bytes.Buffer)
			err = excel.Write(buffer)
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}

			c.Header("Content-Disposition", "attachment; filename=example.xlsx")
			c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buffer.Bytes())
			return
		case "csv":
			_csv, err := h.service.GetCSVReport(userFrom, userTo, accountFrom, accountTo, reports)
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}

			c.Header("Content-Disposition", "attachment; filename=example.csv")
			c.Header("Content-Type", "text/csv")
			c.File(_csv.Name())
			return
		default:
			newErrorResponse(c, http.StatusBadRequest, "invalid extension for report")
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    reports,
	})
}
