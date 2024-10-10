package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Client    *Client
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(client *Client) (*Account, error) {
	account := &Account{
		ID:        uuid.New().String(),
		Client:    client,
		Balance:   0,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	err := account.Validate()

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (account *Account) Validate() error {
	if account.Client == nil {
		return errors.New("client is required")
	}

	return nil
}

func (account *Account) Credit(amount float64) {
	account.Balance += amount
	account.UpdatedAt = time.Now()
}

func (account *Account) Debit(amount float64) {
	account.Balance -= amount
	account.UpdatedAt = time.Now()
}
