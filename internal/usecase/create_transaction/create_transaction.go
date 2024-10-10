package createtransaction

import (
	"github.com.br/yankev12/fc-ms-wallet/internal/entity"
	"github.com.br/yankev12/fc-ms-wallet/internal/gateway"
)

type CreateTransactionInputDTO struct {
	AccountIDFrom string
	AccountIDTo   string
	Amount        float64
}

type CreateTransactionOutput struct {
	ID string
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway     gateway.AccountGateway
}

func NewCreateTransactionUseCase(transactionGateway gateway.TransactionGateway, accountGateway gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: transactionGateway,
		AccountGateway:     accountGateway,
	}
}

func (uc *CreateTransactionUseCase) Execute(input CreateTransactionInputDTO) (*CreateTransactionOutput, error) {
	accountFrom, err := uc.AccountGateway.FindByID(input.AccountIDFrom)
	if err != nil {
		return nil, err
	}

	accountTo, err := uc.AccountGateway.FindByID(input.AccountIDTo)

	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)

	if err != nil {
		return nil, err
	}
	err = uc.TransactionGateway.Create(transaction)

	if err != nil {
		return nil, err
	}
	return &CreateTransactionOutput{ID: transaction.ID}, nil
}
