package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Product struct {
	Name        string `json:"name"`
	ID          string `json:"id"`
	Temperature int    `json:"temperature"`
	Dimensions  int    `json:"dimensions"`
	Weight      int    `json:"weight"`
	SellerID    string `json:"sellerId"`
	HubID       string `json:"hubId"`
	ScanTime    string `json:"scanTime"`
}

type ProductRegistrationContract struct {
	contractapi.Contract
}

func (p *ProductRegistrationContract) RegisterProduct(ctx contractapi.TransactionContextInterface, name string, id string, temperature int, dimensions int, weight int, sellerID string, hubID string, scanTime string) error {
	product := Product{
		Name:        name,
		ID:          id,
		Temperature: temperature,
		Dimensions:  dimensions,
		Weight:      weight,
		SellerID:    sellerID,
		HubID:       hubID,
		ScanTime:    scanTime,
	}

	// Convert the product to JSON format
	productJSON, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("failed to marshal product JSON: %v", err)
	}

	// Save the product to the ledger
	err = ctx.GetStub().PutState(id, productJSON)
	if err != nil {
		return fmt.Errorf("failed to put product on the ledger: %v", err)
	}

	return nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(ProductRegistrationContract))
	if err != nil {
		fmt.Printf("Error creating product registration chaincode: %v", err)
		return
	}

	err = chaincode.Start()
	if err != nil {
		fmt.Printf("Error starting product registration chaincode: %v", err)
	}
}
