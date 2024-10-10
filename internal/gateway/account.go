package gateway

import (
	"github.com.br/yankev12/fc-ms-wallet/internal/entity"
)

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
