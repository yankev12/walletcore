package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("Client 1", "client1@email.com")
	account1, _ := NewAccount(client1)
	client2, _ := NewClient("Client 2", "client2@email.com")
	account2, _ := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 100)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 1100.0, account2.Balance)
	assert.Equal(t, 900.0, account1.Balance)
}

func TestCreateTransactionWithInsuficientBalance(t *testing.T) {
	client1, _ := NewClient("Client 1", "client1@email.com")
	account1, _ := NewAccount(client1)
	client2, _ := NewClient("Client 2", "client2@email.com")
	account2, _ := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 1500)
	assert.Error(t, err, "insufficient funds")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.0, account2.Balance)
	assert.Equal(t, 1000.0, account1.Balance)
}
