package transaction

import (
	"backend-go/internal/model/form"
	"backend-go/internal/model/migrate"
	"backend-go/internal/model/response"
	"backend-go/pkg/pagination"
	"backend-go/pkg/utilities"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func InsertTransaction(db *gorm.DB, data *migrate.Transaction) (*migrate.Transaction, error) {
	if err := db.Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func GetMonthlyTransactionsByUser(db *gorm.DB, userID string, start, end time.Time) ([]migrate.Transaction, error) {
	var transactions []migrate.Transaction
	if err := db.Where("user_id = ? AND date BETWEEN ? AND ?", userID, start, end).
		Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func GetTransactionPagination(db *gorm.DB, params form.Params) (pagination.PaginatedResponse[response.TransactionPagination], error) {
	p := pagination.Resolve(params.Page, params.Limit)

	query := db.Table("transactions AS a").
		Select(`a.id, a.type, a.amount, a.date, a.note, b.name AS user_name, b.email AS user_email, c.name AS category_name, c.description AS category_description`).
		Joins("INNER JOIN users AS b ON a.user_id = b.id").
		Joins("INNER JOIN categories AS c ON a.category_id = c.id")

	if !utilities.IsNilOrEmpty(&params.Search) {
		searchTerm := "%" + params.Search + "%"
		query = query.Where(`
			a.type ILIKE ? OR 
			a.note ILIKE ? OR 
			a.date::TEXT ILIKE ? OR 
			b.name ILIKE ? OR 
			b.email ILIKE ? OR 
			c.name ILIKE ? OR 
			c.description ILIKE ?`,
			searchTerm, searchTerm, searchTerm, searchTerm, searchTerm, searchTerm, searchTerm)
	}

	var records int64
	if err := query.Count(&records).Error; err != nil {
		return pagination.PaginatedResponse[response.TransactionPagination]{}, err
	}

	var data []response.TransactionPagination
	if err := query.Offset(p.Offset).Limit(p.Limit).Scan(&data).Error; err != nil {
		return pagination.PaginatedResponse[response.TransactionPagination]{}, err
	}

	return pagination.NewPaginatedResponse(data, int(records), p.Page, p.Limit), nil
}

func GetTransactionById(db *gorm.DB, id string) (*response.TransactionPagination, error) {
	var data response.TransactionPagination
	if err := db.Table("transactions AS a").
		Select(`a.id, a.type, a.amount, a.date, a.note, b.name AS user_name, b.email AS user_email, c.name AS category_name, c.description AS category_description`).
		Joins("INNER JOIN users AS b ON a.user_id = b.id").
		Joins("INNER JOIN categories AS c ON a.category_id = c.id").
		Where("a.id = ?", id).
		First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func UpdateTransaction(db *gorm.DB, id string, data *migrate.Transaction) (*migrate.Transaction, error) {
	tx := db.Table("transactions").Where("id = ?", id).Updates(data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("transaction with id %s not found", id)
	}
	return data, nil
}

func DeleteTransaction(db *gorm.DB, id string) (bool, error) {
	tx := db.Table("transactions").Where("id = ?", id).Delete(nil)
	if tx.Error != nil {
		return false, tx.Error
	}
	if tx.RowsAffected == 0 {
		return false, fmt.Errorf("transactions with id %s not found", id)
	}
	return true, nil
}
