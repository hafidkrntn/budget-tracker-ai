package category

import (
	"backend-go/internal/model/form"
	"backend-go/internal/model/migrate"
	"backend-go/internal/model/response"
	"backend-go/internal/repository/category"
	"backend-go/pkg/pagination"
	"fmt"
)

func CreateCategory(req form.Category) (*migrate.Category, error) {
	category, err := category.InsertCategory(req)

	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	return category, nil
}

func GetAllPagination(params form.CategoryParams) (pagination.PaginatedResponse[response.Category], error) {
	result, err := category.GetAllPagination(params)
	if err != nil {
		return pagination.PaginatedResponse[response.Category]{}, fmt.Errorf("failed to get paginated categories: %w", err)
	}
	return result, nil
}

func GetAllCategory() (*[]response.Category, error) {
	results, err := category.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}

	return results, nil
}

func GetCategoryById(id string) (*response.Category, error) {
	results, err := category.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get category by id: %w", err)
	}

	return results, nil
}

func UpdateCategory(req form.Category) (*migrate.Category, error) {
	results, err := category.UpdateCategory(req)
	if err != nil {
		return nil, fmt.Errorf("failed to update category: %w", err)
	}

	return results, nil
}

func DeletedCategory(id string) (bool, error) {
	category, err := category.DeleteCategory(id)
	if err != nil || !category {
		return false, fmt.Errorf("failed to delete category: %w", err)
	}

	return category, nil
}
