package handler

import (
	"fmt"
	"github.com/f1rdavsi/reporter/pkg/utils"
	"net/http"

	"github.com/f1rdavsi/reporter/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type transactionData struct {
	From    int     `json:"from" binding:"required"`
	To      int     `json:"to" binding:"required"`
	ToType  string  `json:"to_type" binding:"required"`
	Comment string  `json:"comment" binding:"required"`
	Amount  float64 `json:"amount" binding:"required"`
	Type    string  `json:"type" binding:"required"`
}

func (h *Handler) CreateTransaction(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var request transactionData

	if err := c.ShouldBindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ok := request.Type == "income" || request.Type == "expense"
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "transaction type out of 'income' and 'expense'")
		return
	}

	ok = request.ToType == "account" || request.ToType == "category"
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "transaction to_type out of 'account' and 'category'")
		return
	}

	account, err := h.service.GetAccount(request.From, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var transaction models.Transaction
	if request.ToType == "account" {
		_account, err := h.service.ExistsAccount(request.To)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("there is no account with id: %v", request.To))
			return
		}
		transaction, err = h.makeTransaction(account, _account, request.ToType, request.Comment, request.Amount, request.Type)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		_cat, err := h.service.GetCategory(request.To)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("there is no category with id: %v", request.To))
			return
		}
		transaction, err = h.makeTransaction(account, _cat, request.ToType, request.Comment, request.Amount, request.Type)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data":    transaction,
	})
}

func (h *Handler) GetTransactions(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	trs, err := h.service.GetTransactions(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    trs,
	})
}

func (h *Handler) GetTransaction(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id := cast.ToInt(c.Param("id"))

	tr, err := h.service.GetTransaction(id, userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    tr,
	})
}

func (h *Handler) UpdateTransaction(c *gin.Context) {
	var data map[string]string

	id := cast.ToInt(c.Param("id"))
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tr, err := h.service.UpdateTransaction(id, userId, data["comment"])
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    tr,
	})
}

func (h *Handler) makeTransaction(account models.Account, _to any, to_type, comment string, amount float64, trType string) (models.Transaction, error) {
	var tr models.Transaction

	tr.From = account.ID
	tr.Type = trType
	tr.Comment = comment
	tr.Amount = amount

	switch t := _to.(type) {
	case models.Account:
		tr.To = t.ID
		tr.ToType = to_type
		if trType == "expense" {
			account.Balance -= amount
			t.Balance += amount
		} else {
			account.Balance += amount
			t.Balance -= amount
		}

		_, err := h.service.UpdateAccount(account.ID, account.UserID, account.Name, account.Balance)
		if err != nil {
			return tr, err
		}

		_, err = h.service.UpdateAccount(t.ID, t.UserID, t.Name, t.Balance)
		if err != nil {
			return tr, err
		}
	case models.Category:
		tr.To = t.ID
		tr.ToType = to_type
		if trType == "expense" {
			account.Balance -= t.Price
		} else {
			account.Balance += t.Price
		}

		_, err := h.service.UpdateAccount(account.ID, account.UserID, account.Name, account.Balance)
		if err != nil {
			return tr, err
		}
	default:
		return tr, utils.ErrInvalidType
	}

	return h.service.CreateTransaction(tr)
}
