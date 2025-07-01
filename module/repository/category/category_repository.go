package category

import (
	"backend-go/config"
	"backend-go/module/model/form"
	"backend-go/module/model/migrate"
	"backend-go/module/model/response"
	"backend-go/pkg/pagination"
	"fmt"
	"time"
)

func InsertCategory(req form.Category) (*migrate.Category, error) {
	body := migrate.Category{
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now(),
	}

	if err := config.DbConn.Create(&body).Error; err != nil {
		return nil, err
	}

	return &body, nil
}

func GetAllPagination(params form.CategoryParams) (pagination.PaginatedResponse[response.Category], error) {
	p := pagination.Resolve(params.Page, params.Limit)

	query := config.DbConn.Table("categories")

	var records int64
	if err := query.Count(&records).Error; err != nil {
		return pagination.PaginatedResponse[response.Category]{}, err
	}

	var results []response.Category
	if err := query.Offset(p.Offset).Limit(p.Limit).Scan(&results).Error; err != nil {
		return pagination.PaginatedResponse[response.Category]{}, err
	}

	return pagination.NewPaginatedResponse(results, int(records), p.Page, p.Limit), nil
}

func GetAll() (*[]response.Category, error) {
	var results []response.Category
	if err := config.DbConn.Table("categories").Scan(&results).Error; err != nil {
		return nil, err
	}

	return &results, nil
}

func GetById(id string) (*response.Category, error) {
	var results response.Category
	if err := config.DbConn.Table("categories").Where("id = ?", id).Find(&results).Error; err != nil {
		return nil, err
	}

	return &results, nil
}

func UpdateCategory(req form.Category) (*migrate.Category, error) {
	category := &migrate.Category{
		Name:        req.Name,
		Description: req.Description,
		UpdatedAt:   time.Now(),
	}

	tx := config.DbConn.
		Table("categories").
		Where("id = ?", req.ID).
		Updates(category)

	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("category with id %v not found", req.ID)
	}

	category.ID = req.ID
	return category, nil
}

func DeleteCategory(id string) (bool, error) {
	tx := config.DbConn.
		Table("categories").
		Where("id = ?", id).
		Delete(nil)

	if tx.Error != nil {
		return false, tx.Error
	}
	if tx.RowsAffected == 0 {
		return false, fmt.Errorf("category with id %s not found", id)
	}

	return true, nil
}
