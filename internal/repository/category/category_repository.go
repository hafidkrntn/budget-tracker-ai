package category

import (
	"backend-go/internal/model/form"
	"backend-go/internal/model/migrate"
	"backend-go/internal/model/response"
	"backend-go/pkg/pagination"
	"fmt"

	"gorm.io/gorm"
)

func InsertCategory(db *gorm.DB, data *migrate.Category) (*migrate.Category, error) {
	if err := db.Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func GetAllPagination(db *gorm.DB, params form.Params) (pagination.PaginatedResponse[response.Category], error) {
	p := pagination.Resolve(params.Page, params.Limit)

	sql := db.Table("categories")

	if params.Search != "" {
		searchTerm := "%" + params.Search + "%"
		sql = sql.Where("(name ILIKE ? OR description ILIKE ?)", searchTerm, searchTerm)
	}

	var records int64
	if err := sql.Count(&records).Error; err != nil {
		return pagination.PaginatedResponse[response.Category]{}, err
	}

	var results []response.Category
	if err := sql.Offset(p.Offset).Limit(p.Limit).Scan(&results).Error; err != nil {
		return pagination.PaginatedResponse[response.Category]{}, err
	}

	return pagination.NewPaginatedResponse(results, int(records), p.Page, p.Limit), nil
}

func GetAll(db *gorm.DB) (*[]response.Category, error) {
	var results []response.Category
	if err := db.Table("categories").Scan(&results).Error; err != nil {
		return nil, err
	}
	return &results, nil
}

func GetById(db *gorm.DB, id string) (*response.Category, error) {
	var results response.Category
	if err := db.Table("categories").Where("id = ?", id).Find(&results).Error; err != nil {
		return nil, err
	}
	return &results, nil
}

func UpdateCategory(db *gorm.DB, id string, data *migrate.Category) (*migrate.Category, error) {
	tx := db.Table("categories").Where("id = ?", id).Updates(data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("category with id %s not found", id)
	}
	return data, nil
}

func DeleteCategory(db *gorm.DB, id string) (bool, error) {
	tx := db.Table("categories").Where("id = ?", id).Delete(nil)
	if tx.Error != nil {
		return false, tx.Error
	}
	if tx.RowsAffected == 0 {
		return false, fmt.Errorf("category with id %s not found", id)
	}
	return true, nil
}
