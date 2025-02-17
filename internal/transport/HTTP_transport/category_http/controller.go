package category_http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"leather-shop/internal/models"
	"leather-shop/internal/services"
	"leather-shop/internal/transport/HTTP_transport"
)

type categoryController struct {
	categoryService services.Category
}

func New(categoryService services.Category) HTTP_transport.Category {
	return &categoryController{
		categoryService: categoryService,
	}
}

func (cc *categoryController) GetCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}

	category, err := cc.categoryService.GetCategory(categoryId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func (cc *categoryController) GetCategories(ctx *gin.Context) {
	categories, err := cc.categoryService.GetCategories()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

func (cc *categoryController) CreateCategory(ctx *gin.Context) {
	var category models.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
	}

	if err := cc.categoryService.CreateCategory(&category); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
	}
	ctx.JSON(http.StatusOK, category)
}

func (cc *categoryController) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}

	if err := cc.categoryService.DeleteCategory(categoryId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Категория успешно удалёна"})
}

func (cc *categoryController) EditCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}

	var category models.Category

	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Ошибка": err.Error()})
		return
	}
	category.Id = categoryId

	if err := cc.categoryService.EditCategory(&category); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Ошибка": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, category)
}
