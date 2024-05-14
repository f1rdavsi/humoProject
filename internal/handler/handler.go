package handler

import (
	"github.com/f1rdavsi/reporter/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(h.CORSMiddleware())
	router.Use(gin.Recovery())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ping - pong",
		})
	})

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", h.SignUp)
			auth.POST("/login", h.signIn)
		}

		user := api.Group("/user", h.UserIdentity)
		{
			user.GET("/", h.GetUser)
			user.PUT("/", h.UpdateUser)
			user.GET("/restore", h.RestoreUser)
			user.DELETE("/delete", h.DeleteUser)
			user.PATCH("/change", h.ChangePictureUser)
			user.PATCH("/upload", h.UploadUserPicture)
		}

		account := api.Group("/account", h.UserIdentity)
		{
			account.GET("/", h.GetAllAccounts)
			account.GET("/:id", h.GetAccount)
			account.POST("/", h.CreateAccount)
			account.PUT("/:id", h.UpdateAccount)
			account.GET("/:id/restore", h.RestoreAccount)
			account.DELETE("/:id", h.DeleteAccount)
			account.PATCH("/:id/change", h.ChangePictureAccount)
			account.PATCH("/:id/upload", h.UploadAccountPicture)
		}

		category := api.Group("/category")
		{
			category.POST("/", h.CreateCategory)
			category.GET("/", h.GetCategories)
			category.GET("/:id", h.GetCategory)
			category.PUT("/:id", h.UpdateCategory)
			category.PATCH("/:id/upload", h.UploadPictureCategory)
			category.PATCH("/:id/change", h.ChangePictureCategory)
			category.DELETE("/:id/delete", h.DeleteCategory)
			category.GET("/:id/restore", h.RestoreCategory)
		}

		transaction := api.Group("/transaction", h.UserIdentity)
		{
			transaction.GET("/", h.GetTransactions)
			transaction.GET("/:id", h.GetTransaction)
			transaction.POST("/", h.CreateTransaction)
			transaction.PATCH("/", h.UpdateTransaction)
		}

		report := api.Group("/report", h.UserIdentity)
		{
			report.POST("/?page", h.GetReport)
		}
	}

	return router
}
