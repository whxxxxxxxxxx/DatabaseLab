package service

import (
	"DatabaseLab/internal/app/product/dao"
	"DatabaseLab/internal/app/product/dto"
	"DatabaseLab/internal/app/product/model"
	"DatabaseLab/middleware/response"
	"DatabaseLab/pkg/ctl"
	"DatabaseLab/pkg/errorx"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetCategory(c *gin.Context) response.JsonResponse {
	code := errorx.Success
	categoryDao := dao.Product
	category, err := categoryDao.GetCategoryNames()
	if err != nil {
		code = errorx.Error
		return response.ReturnResponse(code, errorx.GetMsg(code), nil, nil)
	}
	return response.ReturnResponse(code, errorx.GetMsg(code), dto.BuildCategoryResponse(category), nil)
}

func AddProduct(c *gin.Context) response.JsonResponse {
	u, err := ctl.GetUserInfo(c.Request.Context())
	if err != nil {
		return response.ReturnResponse(errorx.Error, errorx.GetMsg(errorx.Error), nil, err)
	}
	code := errorx.Success
	productDao := dao.Product
	var product dto.ProductRequest
	product.BossID = u.Id
	if err := c.ShouldBind(&product); err != nil {
		code = errorx.Error
		return response.ReturnResponse(code, errorx.GetMsg(code), nil, nil)
	}
	productModel := model.Product{
		ProductName:   product.ProductName,
		CategoryId:    product.CategoryId,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		OnSale:        product.OnSale,
		Num:           product.Num,
		BossID:        product.BossID,
	}
	err = productDao.AddProduct(&productModel)
	if err != nil {
		code = errorx.Error
	}
	return response.ReturnResponse(code, errorx.GetMsg(code), nil, nil)
}

func GetProduct(c *gin.Context) response.JsonResponse {
	code := errorx.Success
	productDao := dao.Product
	id := c.Param("id")
	product, err := productDao.GetProduct(id)
	if err != nil {
		code = errorx.Error
		return response.ReturnResponse(code, errorx.GetMsg(code), nil, nil)
	}
	return response.ReturnResponse(code, errorx.GetMsg(code), dto.BuildProductResponse(product), nil)
}

func DeleteProduct(c *gin.Context) response.JsonResponse {
	code := errorx.Success
	productDao := dao.Product
	id := c.Param("id")
	err := productDao.DeleteProduct(id)
	if err != nil {
		code = errorx.Error
	}
	return response.ReturnResponse(code, errorx.GetMsg(code), nil, nil)
}

func UpdateProduct(c *gin.Context) response.JsonResponse {
	code := errorx.Success
	productDao := dao.Product
	id := c.Param("id")
	var product model.Product
	if err := c.ShouldBind(&product); err != nil {
		code = errorx.Error
		return response.ReturnResponse(code, errorx.GetMsg(code), nil, nil)
	}
	//将id从string转化为uint
	uid, _ := strconv.Atoi(id)
	product.ID = uint(uid)
	err := productDao.UpdateProduct(&product)
	if err != nil {
		code = errorx.Error
	}
	return response.ReturnResponse(code, errorx.GetMsg(code), nil, nil)
}

// 模糊查询获取商品列表
func GetProductList(c *gin.Context) response.JsonResponse {
	code := errorx.Success
	productDao := dao.Product
	var product dto.ProductListRequest
	if err := c.ShouldBind(&product); err != nil {
		code = errorx.Error
		return response.ReturnResponse(code, errorx.GetMsg(code), nil, nil)
	}
	products, err := productDao.GetProductList(&product)
	if err != nil {
		code = errorx.Error
	}
	return response.ReturnResponse(code, errorx.GetMsg(code), products, nil)
}
