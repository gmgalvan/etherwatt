package usecases

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	contract "github.com/etherwatt/gen"
)

// Usecases struct with all usecases of the app
type Usecases struct {
	client   *ethclient.Client
	contract *contract.Etherwatt
}

// NewUsecases returns a usecase struct
func NewUsecases(client *ethclient.Client, contract *contract.Etherwatt) *Usecases {
	return &Usecases{
		client:   client,
		contract: contract,
	}
}

func (uc *Usecases) GetEtherBalance(addr string) (*big.Float, error) {
	address := common.HexToAddress(addr)

	balance, err := uc.client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return nil, err
	}

	// wei / 10^18
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethValue, nil
}

func (uc *Usecases) GetTokenBalance(address string) (*big.Int, error) {
	add := common.HexToAddress(address)
	tokenBalance, err := uc.contract.BalanceOf(&bind.CallOpts{}, add)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return tokenBalance, nil
}

func (uc *Usecases) Transfer(addressFrom, addressTo string, amount *big.Int) error {
	addF := common.HexToAddress(addressFrom)
	addT := common.HexToAddress(addressTo)

	txOpts := &bind.TransactOpts{}

	_, err := uc.contract.Transfer(txOpts, addF, addT, amount)
	if err != nil {
		return err
	}

	return nil
}

func (uc *Usecases) AskReward(address string) error {
	add := common.HexToAddress(address)
	amount := big.NewInt(10)

	transactOpts := &bind.TransactOpts{
		GasLimit: 0,
		GasPrice: nil,
		Value:    nil,
		Nonce:    nil,
	}

	_, err := uc.contract.Reward(transactOpts, add, amount)
	if err != nil {
		return err
	}

	return nil
}
