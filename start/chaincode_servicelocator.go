package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type ServiceLocatorChaincode struct {
}

func main() {
	err := shim.Start(new(ServiceLocatorChaincode))
	if err != nil {
		fmt.Printf("Error starting Service Locator Chaincode: %s", err)
	}
}

func (t *ServiceLocatorChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}

func (t *ServiceLocatorChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} else if function == "test" {
		return nil, errors.New("probando..")
	}
	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *ServiceLocatorChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "dummy_query" {											//read a variable
		fmt.Println("hi there " + function)						//error
		return nil, nil;
	} else if function == "test_query" {
		fmt.Println("TEST QUERY");
		return nil, nil;
	}
	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query: " + function)
}
