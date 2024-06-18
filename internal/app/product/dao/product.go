package dao

import (
	"DatabaseLab/internal/app/product/dto"
	"DatabaseLab/internal/app/product/model"
	"gorm.io/gorm"
)

type product struct {
	*gorm.DB
}

func (u *product) Init(db *gorm.DB) (err error) {
	u.DB = db
	return db.AutoMigrate(&model.ProductImage{}, &model.Product{}, &model.ProductInfoImg{}, &model.ProductParamsImg{}, &model.Category{}, &model.Carousel{})
}

func (u *product) GetCategoryNames() ([]*model.Category, error) {
	var categories []*model.Category
	// 使用 Select 指定只查询 id 和 name 字段
	err := u.DB.Model(&model.Category{}).Select("id", "category_name").Find(&categories).Error
	return categories, err
}

func (u *product) AddProduct(product *dto.ProductRequest) error {
	return u.DB.Model(&model.Product{}).Create(product).Error
}

func (u *product) GetProduct(id string) (*model.Product, error) {
	var product model.Product
	err := u.DB.Model(&model.Product{}).Where("id = ?", id).Find(&product).Error
	return &product, err
}

func (u *product) DeleteProduct(id string) error {
	return u.DB.Model(&model.Product{}).Where("id = ?", id).Delete(&model.Product{}).Error
}

func (u *product) UpdateProduct(product *model.Product) error {
	return u.DB.Model(&model.Product{}).Where("id = ?", product.ID).Updates(product).Error
}

func (u *product) GetProductList(productRequest *dto.ProductListRequest) ([]*dto.ProductResponse, error) {
	var products []*model.Product
	query := u.DB.Model(&model.Product{})

	// 添加基于ProductName的模糊查询
	if productRequest.ProductName != "" {
		query = query.Where("product_name LIKE ?", "%"+productRequest.ProductName+"%")
	}
	// 添加基于CategoryId的精确查询
	if productRequest.CategoryId != 0 {
		query = query.Where("category_id = ?", productRequest.CategoryId)
	}
	// 添加基于BossId的精确查询
	if productRequest.BossId != 0 {
		query = query.Where("boss_id = ?", productRequest.BossId)
	}

	err := query.Find(&products).Error
	if err != nil {
		return nil, err
	}

	// 转换products到DTO
	responses := dto.BuildProductListResponse(products)
	return responses, nil
}
