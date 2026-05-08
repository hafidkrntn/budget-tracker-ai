package category

import (
	"backend-go/config"
	"backend-go/internal/model/form"
	"backend-go/internal/model/migrate"
	"backend-go/internal/model/response"
	"backend-go/internal/repository/category"
	"backend-go/pkg/pagination"
	"fmt"
	"time"
)

func CreateCategory(req form.Category) (*migrate.Category, error) {
	db := config.GetDB()

	data := &migrate.Category{
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now(),
	}

	result, err := category.InsertCategory(db, data)
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	return result, nil
}

func GetAllPagination(params form.Params) (pagination.PaginatedResponse[response.Category], error) {
	db := config.GetDB()

	result, err := category.GetAllPagination(db, params)
	if err != nil {
		return pagination.PaginatedResponse[response.Category]{}, fmt.Errorf("failed to get paginated categories: %w", err)
	}

	return result, nil
}

func GetAllCategory() (*[]response.Category, error) {
	db := config.GetDB()

	results, err := category.GetAll(db)
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}

	return results, nil
}

func GetCategoryById(id string) (*response.Category, error) {
	db := config.GetDB()

	results, err := category.GetById(db, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get category by id: %w", err)
	}

	return results, nil
}

func UpdateCategory(req form.Category) (*migrate.Category, error) {
	db := config.GetDB()

	data := &migrate.Category{
		Name:        req.Name,
		Description: req.Description,
		UpdatedAt:   time.Now(),
	}

	result, err := category.UpdateCategory(db, req.ID.String(), data)
	if err != nil {
		return nil, fmt.Errorf("failed to update category: %w", err)
	}

	result.ID = req.ID
	return result, nil
}

func DeletedCategory(id string) (bool, error) {
	db := config.GetDB()

	ok, err := category.DeleteCategory(db, id)
	if err != nil || !ok {
		return false, fmt.Errorf("failed to delete category: %w", err)
	}

	return ok, nil
}
