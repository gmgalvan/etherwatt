package transport

import (
	"errors"
	"fmt"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Transacion struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int64  `json:"amount"`
}

func (h *Handler) EtherWattWallet(g *gin.Context) {
	add := g.Param("address")
	result, err := h.uc.GetTokenBalance(add)
	if err != nil {
		msgErr := fmt.Sprintf("Error during getting balance, err: %v", err)
		err := errors.New(msgErr)
		g.JSON(http.StatusBadRequest, err)
	}
	g.JSON(http.StatusOK, result)
}

func (h *Handler) EtherWallet(g *gin.Context) {
	add := g.Param("address")
	result, err := h.uc.GetEtherBalance(add)
	if err != nil {
		msgErr := fmt.Sprintf("Error during getting balance, err: %v", err)
		err := errors.New(msgErr)
		g.JSON(http.StatusBadRequest, err)
	}
	g.JSON(http.StatusOK, result)

}

func (h *Handler) Conection(g *gin.Context) {
	add := g.Param("address")
	err := h.uc.AskReward(add)
	if err != nil {
		msgErr := fmt.Sprintf("Error during asking reward, err: %v", err)
		err := errors.New(msgErr)
		g.JSON(http.StatusBadRequest, err)
	}

	result, err := h.uc.GetTokenBalance(add)
	if err != nil {
		msgErr := fmt.Sprintf("Error during getting balance, err: %v", err)
		err := errors.New(msgErr)
		g.JSON(http.StatusBadRequest, err)
	}

	g.JSON(http.StatusOK, "new balance: "+result.String())
}

func (h *Handler) Transfer(g *gin.Context) {
	var t Transacion
	if err := g.BindJSON(&t); err != nil {
		msgErr := fmt.Sprintf("Error during unmarshalling, err: %v", err)
		err := errors.New(msgErr)
		g.JSON(http.StatusBadRequest, err)
	}

	amount := big.NewInt(t.Amount)
	err := h.uc.Transfer(t.From, t.To, amount)
	if err != nil {
		fmt.Println("error", err)
		msgErr := fmt.Sprintf("Error during transfering, err: %v", err)
		err := errors.New(msgErr)
		g.JSON(http.StatusBadRequest, err)
	}
}
