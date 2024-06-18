package dto

import "DatabaseLab/internal/app/product/model"

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func BuildCategoryResponse(category []*model.Category) []*CategoryResponse {
	var categoryResponse []*CategoryResponse
	for _, v := range category {
		categoryResponse = append(categoryResponse, &CategoryResponse{
			ID:   v.ID,
			Name: v.CategoryName,
		})
	}
	return categoryResponse
}

type ProductRequest struct {
	ProductName   string `json:"product_name"`
	CategoryId    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	OnSale        bool   `json:"on_sale"`
	Num           int    `json:"num"`
	BossID        uint   `json:"boss_id"`
}

type ProductResponse struct {
	ID            uint   `json:"id"`
	ProductName   string `json:"product_name"`
	CategoryId    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	OnSale        bool   `json:"on_sale"`
	Num           int    `json:"num"`
	BossID        uint   `json:"boss_id"`
}

func BuildProductResponse(product *model.Product) *ProductResponse {
	return &ProductResponse{
		ID:            product.ID,
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
}

func BuildProductListResponse(products []*model.Product) []*ProductResponse {
	var productResponse []*ProductResponse
	for _, v := range products {
		productResponse = append(productResponse, &ProductResponse{
			ID:            v.ID,
			ProductName:   v.ProductName,
			CategoryId:    v.CategoryId,
			Title:         v.Title,
			Info:          v.Info,
			ImgPath:       v.ImgPath,
			Price:         v.Price,
			DiscountPrice: v.DiscountPrice,
			OnSale:        v.OnSale,
			Num:           v.Num,
			BossID:        v.BossID,
		})
	}
	return productResponse
}

type ProductListRequest struct {
	ProductName string `json:"product_name"`
	CategoryId  uint   `json:"category_id"`
	BossId      uint   `json:"boss_id"`
}
