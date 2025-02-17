package category_http

import (
	"github.com/gin-gonic/gin"

	"leather-shop/internal/transport/HTTP_transport"
)

func NewRouterCategory(engine *gin.RouterGroup, cointroller HTTP_transport.Category) {
	categoryGroup := engine.Group("")

	categoryGroup.POST("", cointroller.CreateCategory)
	categoryGroup.GET("/:id", cointroller.GetCategory)
	categoryGroup.GET("", cointroller.GetCategories)
	categoryGroup.DELETE("/:id", cointroller.DeleteCategory)
	categoryGroup.PUT("/:id", cointroller.EditCategory)
}
