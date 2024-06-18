package handler

import (
	"DatabaseLab/internal/app/product/service"
	"DatabaseLab/middleware/response"
	"github.com/gin-gonic/gin"
)

func HandleCategory(c *gin.Context) {
	res := service.GetCategory(c)
	response.HttpResponse(c, res)
}

func HandleAddProduct(c *gin.Context) {
	res := service.AddProduct(c)
	response.HttpResponse(c, res)
}

func HandleGetProduct(c *gin.Context) {
	res := service.GetProduct(c)
	response.HttpResponse(c, res)
}

func HandleDeleteProduct(c *gin.Context) {
	res := service.DeleteProduct(c)
	response.HttpResponse(c, res)
}

func HandleUpdateProduct(c *gin.Context) {
	res := service.UpdateProduct(c)
	response.HttpResponse(c, res)
}

func HandleGetProductList(c *gin.Context) {
	res := service.GetProductList(c)
	response.HttpResponse(c, res)
}
