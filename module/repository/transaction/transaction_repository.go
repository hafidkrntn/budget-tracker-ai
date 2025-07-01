package transaction

import (
	"backend-go/config"
	"backend-go/module/model/form"
	"backend-go/module/model/migrate"
	"backend-go/module/model/response"
	"backend-go/module/repository/category"
	"backend-go/module/repository/users"
	"backend-go/pkg/pagination"
	"backend-go/pkg/utilities"
	"fmt"
	"time"
)

func InsertTransaction(req form.TransactionForm) (*migrate.Transaction, error) {
	if err := verifyTransaction(req); err != nil {
		return nil, err
	}

	newTransaction := migrate.Transaction{
		Type:       req.Type,
		Amount:     req.Amount,
		Date:       req.Date,
		Note:       req.Note,
		UserID:     req.UserId,
		CategoryID: req.CategoryId,
		CreatedAt:  time.Now(),
	}

	if err := config.DbConn.Create(&newTransaction).Error; err != nil {
		return nil, err
	}

	return &newTransaction, nil
}

func verifyTransaction(req form.TransactionForm) error {
	categoryData, err := category.GetById(req.CategoryId.String())
	if err != nil {
		return fmt.Errorf("invalid category ID: %w", err)
	}
	if categoryData == nil {
		return fmt.Errorf("category not found")
	}

	userData, err := users.GetUserById(req.UserId.String())
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}
	if userData == nil {
		return fmt.Errorf("user not found")
	}

	if req.Amount <= 0 {
		return fmt.Errorf("amount must be greater than zero")
	}

	if req.Date.After(time.Now()) {
		return fmt.Errorf("transaction date cannot be in the future")
	}

	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, -1)

	var transactions []migrate.Transaction
	if err := config.DbConn.Where("user_id = ? AND date BETWEEN ? AND ?", req.UserId, startOfMonth, endOfMonth).
		Find(&transactions).Error; err != nil {
		return err
	}

	var totalIncome, totalExpense float64
	for _, tx := range transactions {
		switch tx.Type {
		case "income":
			totalIncome += tx.Amount
		case "expense":
			totalExpense += tx.Amount
		}
	}

	if req.Type == "expense" && totalIncome < totalExpense+req.Amount {
		return fmt.Errorf("this month's income is insufficient")
	}

	return nil
}

func GetTransactionPagination(params form.TransactionParams) (pagination.PaginatedResponse[response.TransactionPagination], error) {
	p := pagination.Resolve(params.Page, params.Limit)

	query := config.DbConn.Table("transactions AS a").
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

func GetTransactionById(id string) (*response.TransactionPagination, error) {
	var data response.TransactionPagination
	if err := config.DbConn.Table("transactions AS a").
		Select(`a.id, a.type, a.amount, a.date, a.note, b.name AS user_name, b.email AS user_email, c.name AS category_name, c.description AS category_description`).
		Joins("INNER JOIN users AS b ON a.user_id = b.id").
		Joins("INNER JOIN categories AS c ON a.category_id = c.id").
		First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func UpdateTransaction(req form.TransactionForm) (*migrate.Transaction, error) {
	if err := verifyTransaction(req); err != nil {
		return nil, err
	}

	transaction := &migrate.Transaction{
		Type:       req.Type,
		Amount:     req.Amount,
		Date:       req.Date,
		Note:       req.Note,
		UserID:     req.UserId,
		CategoryID: req.CategoryId,
		UpdatedAt:  time.Now(),
	}

	tx := config.DbConn.Table("transactions").Where("id = ?", req.ID).Updates(transaction)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("transaction with id %v not found", req.ID)
	}

	transaction.ID = req.ID
	return transaction, nil
}

func DeleteTransaction(id string) (bool, error) {
	tx := config.DbConn.Table("transactions").Where("id = ?", id).Delete(nil)
	if tx.Error != nil {
		return false, tx.Error
	}
	if tx.RowsAffected == 0 {
		return false, fmt.Errorf("transactions with id %s not found", id)
	}

	return true, nil
}
