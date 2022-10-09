package rest

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/tarkovskynik/Golang-ninja-project/docs"
)

type Handler struct {
	usersService      Users
	filesService      FilesService
	maxUploadFileSize int64
	fileTypes         map[string]interface{}
}

func NewHandler(users Users, files FilesService, maxUploadFileSize int64, fileTypes map[string]interface{}) *Handler {
	return &Handler{
		usersService:      users,
		filesService:      files,
		maxUploadFileSize: maxUploadFileSize,
		fileTypes:         fileTypes,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	usersApi := router.Group("/auth")
	{
		usersApi.POST("/sign-up", h.signUp)
		usersApi.POST("/sign-in", h.signIn)
		usersApi.GET("/refresh", h.refresh)
	}

	filesApi := router.Group("/s3")
	filesApi.Use(h.authMiddleware())
	{
		filesApi.POST("/upload", h.fileUploadS3)
		filesApi.GET("/files", h.getFilesS3)
	}

	return router
}
