package blockchain

import (
	"cointracker-assignment/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
)

type BcComClient struct {
}

func NewBcComClient() Client {
	return &BcComClient{}
}

func (client *BcComClient) GetTransactions(walletID string, nextOffset *string) (*models.TransactionsResponse, error) {
	// TODO: this is only getting one page. Need to Run it in a while len(transactions) == 50 to get them all,
	// but didn't add it for now to avoid getting rate limited
	url := "https://blockchain.info/rawaddr/" + walletID + "?limit=50"
	if nextOffset != nil {
		url += "&offset=" + *nextOffset
	}
	fmt.Println("URL is", url)
	response, err := http.Get(url)

	if err != nil {
		return nil, errors.New("failed to make client call")
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("failed to make client call")
	}
	// TODO: make this type specific for this implementation to avoid sharing same type
	var transactionResponse models.TransactionsResponse
	json.Unmarshal(responseData, &transactionResponse)

	newOffset := math.Min(50, float64(len(transactionResponse.Transactions)))
	if nextOffset != nil {
		intNextOffset, err := strconv.Atoi(*nextOffset)
		if err != nil {
			return nil, errors.New("invalid offset")
		}

		newOffset += float64(intNextOffset)
	}
	transactionResponse.NextOffset = strconv.FormatInt(int64(newOffset), 10)
	return &transactionResponse, nil
}
