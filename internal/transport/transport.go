package transport

import "math/big"

type Usecases interface {
	GetEtherBalance(addr string) (*big.Float, error)
	GetTokenBalance(address string) (*big.Int, error)
	Transfer(addressFrom, addressTo string, amount *big.Int) error
	AskReward(address string) error
}

type Handler struct {
	uc Usecases
}

func NewTrasport(uc Usecases) *Handler {
	return &Handler{
		uc: uc,
	}
}
