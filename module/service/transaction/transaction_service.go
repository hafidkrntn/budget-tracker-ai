package transaction

import (
	"backend-go/module/model/form"
	"backend-go/module/model/migrate"
	"backend-go/module/model/response"
	"backend-go/module/repository/transaction"
	"backend-go/pkg/pagination"
	"fmt"
)

func CreateTransaction(req form.TransactionForm) (*migrate.Transaction, error) {
	transaction, err := transaction.InsertTransaction(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	return transaction, nil
}

func GetTransactionPagination(params form.TransactionParams) (pagination.PaginatedResponse[response.TransactionPagination], error) {
	data, err := transaction.GetTransactionPagination(params)
	if err != nil {
		return pagination.PaginatedResponse[response.TransactionPagination]{}, fmt.Errorf("failed to get data paginated transcation: %w", err)
	}

	return data, nil
}

func GetTransactionById(id string) (*response.TransactionPagination, error) {
	results, err := transaction.GetTransactionById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction by id: %w", err)
	}

	return results, nil
}

func UpdateTransaction(req form.TransactionForm) (*migrate.Transaction, error) {
	transaction, err := transaction.UpdateTransaction(req)
	if err != nil {
		return nil, fmt.Errorf("failed to updated transaction: %w", err)
	}

	return transaction, nil
}

func DeleteTransaction(id string) (bool, error) {
	transaction, err := transaction.DeleteTransaction(id)
	if err != nil || !transaction {
		return false, fmt.Errorf("failed to delete transaction: %w", err)
	}

	return transaction, nil
}
