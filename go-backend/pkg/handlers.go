package pkg

import (
	"context"
	"log"
	"math/big"
	"cashback-rewards/pkg/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func PurchaseProduct(client *ethclient.Client, auth *bind.TransactOpts, contractAddress string, productValue *big.Int) {
	contract, err := contracts.NewCashbackRewardsToken(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	tx, err := contract.BuyProduct(auth, productValue)
	if err != nil {
		log.Fatalf("Failed to execute transaction: %v", err)
	}

	log.Printf("Transaction sent: %s", tx.Hash().Hex())
}