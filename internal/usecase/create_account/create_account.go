package createaccount

import (
	"github.com.br/yankev12/fc-ms-wallet/internal/entity"
	"github.com.br/yankev12/fc-ms-wallet/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientID string
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(a gateway.AccountGateway, c gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: a,
		ClientGateway:  c,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := uc.ClientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}
	account, _ := entity.NewAccount(client)
	err = uc.AccountGateway.Save(account)

	if err != nil {
		return nil, err
	}
	return &CreateAccountOutputDTO{
		ID: account.ID,
	}, nil
}
