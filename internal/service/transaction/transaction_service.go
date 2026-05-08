package transaction

import (
	"backend-go/config"
	"backend-go/internal/model/form"
	"backend-go/internal/model/migrate"
	"backend-go/internal/model/response"
	"backend-go/internal/repository/category"
	"backend-go/internal/repository/transaction"
	"backend-go/internal/repository/users"
	"backend-go/pkg/pagination"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func verifyTransaction(req form.TransactionForm) error {
	db := config.GetDB()

	categoryData, err := category.GetById(db, req.CategoryId.String())
	if err != nil {
		return fmt.Errorf("invalid category ID: %w", err)
	}
	if categoryData == nil || categoryData.ID == uuid.Nil {
		return fmt.Errorf("category not found")
	}

	userData, err := users.GetUserById(db, req.UserId.String())
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

	transactions, err := transaction.GetMonthlyTransactionsByUser(db, req.UserId.String(), startOfMonth, endOfMonth)
	if err != nil {
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

func CreateTransaction(req form.TransactionForm) (*migrate.Transaction, error) {
	db := config.GetDB()

	if err := verifyTransaction(req); err != nil {
		return nil, err
	}

	data := &migrate.Transaction{
		Type:       req.Type,
		Amount:     req.Amount,
		Date:       req.Date,
		Note:       req.Note,
		UserID:     req.UserId,
		CategoryID: req.CategoryId,
		CreatedAt:  time.Now(),
	}

	result, err := transaction.InsertTransaction(db, data)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	return result, nil
}

func GetTransactionPagination(params form.Params) (pagination.PaginatedResponse[response.TransactionPagination], error) {
	db := config.GetDB()

	data, err := transaction.GetTransactionPagination(db, params)
	if err != nil {
		return pagination.PaginatedResponse[response.TransactionPagination]{}, fmt.Errorf("failed to get paginated transaction: %w", err)
	}

	return data, nil
}

func GetTransactionById(id string) (*response.TransactionPagination, error) {
	db := config.GetDB()

	results, err := transaction.GetTransactionById(db, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction by id: %w", err)
	}

	return results, nil
}

func UpdateTransaction(req form.TransactionForm) (*migrate.Transaction, error) {
	db := config.GetDB()

	if err := verifyTransaction(req); err != nil {
		return nil, err
	}

	data := &migrate.Transaction{
		Type:       req.Type,
		Amount:     req.Amount,
		Date:       req.Date,
		Note:       req.Note,
		UserID:     req.UserId,
		CategoryID: req.CategoryId,
		UpdatedAt:  time.Now(),
	}

	result, err := transaction.UpdateTransaction(db, req.ID.String(), data)
	if err != nil {
		return nil, fmt.Errorf("failed to update transaction: %w", err)
	}

	result.ID = req.ID
	return result, nil
}

func DeleteTransaction(id string) (bool, error) {
	db := config.GetDB()

	ok, err := transaction.DeleteTransaction(db, id)
	if err != nil || !ok {
		return false, fmt.Errorf("failed to delete transaction: %w", err)
	}

	return ok, nil
}
