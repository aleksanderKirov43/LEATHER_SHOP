package HTTP_transport

import (
	"github.com/gin-gonic/gin"
)

type Category interface {
	GetCategory(c *gin.Context)
	GetCategories(c *gin.Context)
	CreateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	EditCategory(c *gin.Context)
}
