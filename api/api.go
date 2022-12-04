package api

import (
	"github.com/gin-gonic/gin"

	_ "github.com/samandar2605/booking/api/docs" // for swagger
	v1 "github.com/samandar2605/booking/api/v1"
	"github.com/samandar2605/booking/config"
	"github.com/samandar2605/booking/storage"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type RouterOptions struct {
	Cfg      *config.Config
	Storage  storage.StorageI
	InMemory storage.InMemoryStorageI
}

// @title           Swagger for blog api
// @version         1.0
// @description     This is a blog service api.
// @host      		localhost:8000
// @BasePath  		/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Security ApiKeyAuth
func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Cfg:      opt.Cfg,
		Storage:  opt.Storage,
		InMemory: opt.InMemory,
	})
	router.Static("/media", "./media")
	apiV1 := router.Group("/v1")

	// // Category
	// apiV1.GET("/categories", handlerV1.GetAllCategories)
	// apiV1.GET("/categories/:id", handlerV1.GetCategory)
	// apiV1.POST("/categories", handlerV1.AuthMiddleware, handlerV1.CreateCategory)
	// apiV1.PUT("/categories/:id", handlerV1.AuthMiddleware, handlerV1.UpdateCategory)
	// apiV1.DELETE("/categories/:id", handlerV1.AuthMiddleware, handlerV1.DeleteCategory)

	// // Like
	// apiV1.POST("/likes", handlerV1.AuthMiddleware, handlerV1.CreateOrUpdateLike)
	// apiV1.GET("/likes/user-post", handlerV1.AuthMiddleware, handlerV1.GetLike)

	// User
	apiV1.POST("/users", handlerV1.CreateUser)
	apiV1.GET("/users", handlerV1.GetAllUsers)
	apiV1.GET("/users/:id", handlerV1.GetUser)
	apiV1.PUT("/users/:id", handlerV1.UpdateUser)
	apiV1.DELETE("/users/:id", handlerV1.DeleteUser)

	// // Comment
	// apiV1.GET("/comments", handlerV1.GetAllComment)
	// apiV1.GET("/comments/:id", handlerV1.GetComment)
	// apiV1.POST("/comments", handlerV1.AuthMiddleware, handlerV1.CreateComment)
	// apiV1.PUT("/comments/:id", handlerV1.AuthMiddleware, handlerV1.UpdateComment)
	// apiV1.DELETE("/comments/:id", handlerV1.AuthMiddleware, handlerV1.DeleteComment)

	// // Post
	// apiV1.GET("/posts", handlerV1.GetAllPost)
	// apiV1.GET("/posts/:id", handlerV1.GetPost)
	// apiV1.POST("/posts", handlerV1.AuthMiddleware, handlerV1.CreatePost)
	// apiV1.PUT("/posts/:id", handlerV1.AuthMiddleware, handlerV1.UpdatePost)
	// apiV1.DELETE("/posts/:id", handlerV1.AuthMiddleware, handlerV1.DeletePost)

	// file upload
	apiV1.POST("/file-upload", handlerV1.UploadFile)

	// hotel
	apiV1.GET("/hotels", handlerV1.GetAllHotels)
	apiV1.POST("/hotels", handlerV1.CreateHotel)
	apiV1.GET("/hotels/:id", handlerV1.GetHotel)
	apiV1.POST("/hotels/add-room", handlerV1.AddRoom)

	// Register
	apiV1.POST("/auth/register", handlerV1.Register)
	apiV1.POST("/auth/verify", handlerV1.Verify)
	apiV1.POST("/auth/login", handlerV1.Login)
	apiV1.POST("/auth/forgot-password", handlerV1.ForgotPassword)
	apiV1.POST("/auth/verify-forgot-password", handlerV1.VerifyForgotPassword)
	apiV1.POST("/auth/update-password", handlerV1.UpdatePassword)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
