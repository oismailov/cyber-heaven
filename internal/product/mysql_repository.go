package product

import (
	"gorm.io/gorm"
)

type MysqlRepository struct {
	DB *gorm.DB
}

func NewProductRepository(
	db *gorm.DB,
) Repository {
	return &MysqlRepository{
		DB: db,
	}
}

func (r *MysqlRepository) FindByID(id int) (*Product, error) {
	var product Product
	result := r.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}
