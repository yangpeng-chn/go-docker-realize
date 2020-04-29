package model

import "errors"

type Transaction struct {
	ID          int    `json:"id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

var transactions []Transaction

func ListTransactions() ([]Transaction, error) {
	return transactions, nil
}

func ShowTransaction(id int) (Transaction, error) {
	for i := range transactions {
		if transactions[i].ID == id {
			return transactions[i], nil
		}
	}
	return Transaction{}, errors.New("transaction not found")
}

func CreateTransaction(transaction Transaction) error {
	for i := range transactions {
		if transactions[i].ID == transaction.ID {
			return errors.New("transaction ID exists")
		}
	}
	transactions = append(transactions, transaction)
	return nil
}

func DeleteTransactions() error {
	transactions = nil
	return nil
}
