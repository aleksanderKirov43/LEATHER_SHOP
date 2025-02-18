package HTTP_transport

import (
	"github.com/gin-gonic/gin"
)

type Products interface {
	GetProduct(c *gin.Context)
	GetProducts(c *gin.Context)
	CreateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	EditProduct(c *gin.Context)
}
