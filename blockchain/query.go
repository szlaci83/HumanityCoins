package blockchain

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/api/apitxn/chclient"
)

func (setup *FabricSetup) InvokeGetByID(ID string) (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "getUser")
	args = append(args, ID)

	response, err := setup.client.Query(chclient.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}

func (setup *FabricSetup) InvokeGetKeys() (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "getKeys")

	response, err := setup.client.Query(chclient.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}