package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account, _ := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountWithNilClient(t *testing.T) {
	account, error := NewAccount(nil)
	assert.Nil(t, account)
	assert.Error(t, error, "client is required")
}
