package HTTP_transport

import (
	"github.com/gin-gonic/gin"
)

type Category interface {
	GetCategory(ctx *gin.Context)
	GetCategories(ctx *gin.Context)
	CreateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
	EditCategory(ctx *gin.Context)
}
