package createaccount

import (
	"testing"

	"github.com.br/yankev12/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "j@j")
	clientMock := &ClientGatewayMock{}
	clientMock.On("Get", client.ID).Return(client, nil)

	accountMonk := &AccountGatewayMock{}

	accountMonk.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountMonk, clientMock)
	inputDto := CreateAccountInputDTO{
		ClientID: client.ID,
	}
	output, err := uc.Execute(inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, output.ID)
	clientMock.AssertExpectations(t)
	accountMonk.AssertExpectations(t)
	clientMock.AssertNumberOfCalls(t, "Get", 1)
	accountMonk.AssertNumberOfCalls(t, "Save", 1)
}
