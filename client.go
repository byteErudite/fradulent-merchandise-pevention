package main

import (
	"fmt"
	"log"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

const (
	chaincodeID = "your-chaincode-id"
)

func run() {
	// Create a new SDK instance
	sdk, err := fabsdk.New(config.FromFile("connection.yaml"))
	if err != nil {
		log.Fatalf("Failed to create new SDK: %v", err)
	}
	defer sdk.Close()

	// Create a new client context
	clientContext := sdk.ChannelContext("your-channel-name", fabsdk.WithUser("your-username"))

	// Create a new channel client
	channelClient, err := channel.New(clientContext)
	if err != nil {
		log.Fatalf("Failed to create new channel client: %v", err)
	}

	// Prepare the registration payload
	name := "Product 1"
	id := "P1"
	temperature := 25
	dimensions := 10
	weight := 100
	sellerID := "S1"
	hubID := "H1"
	scanTime := "2023-06-08T12:00:00Z"

	// Create the request to invoke the chaincode function
	request := channel.Request{
		ChaincodeID: chaincodeID,
		Fcn:         "RegisterProduct",
		Args: [][]byte{
			[]byte(name),
			[]byte(id),
			[]byte(fmt.Sprintf("%d", temperature)),
			[]byte(fmt.Sprintf("%d", dimensions)),
			[]byte(fmt.Sprintf("%d", weight)),
			[]byte(sellerID),
			[]byte(hubID),
			[]byte(scanTime),
		},
	}

	// Invoke the chaincode function
	response, err := channelClient.Execute(request)
	if err != nil {
		log.Fatalf("Failed to execute chaincode function: %v", err)
	}

	fmt.Printf("Transaction ID: %s\n", response.TransactionID)
	fmt.Printf("Response: %s\n", response.Payload)
}
