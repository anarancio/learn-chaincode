package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SimpleChaincode struct {
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1 (devicelocator uuid)")
	}
  fmt.Println("init using device locator: " + args[0])

  return nil, nil
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  fmt.Println("invoke is running " + function)

  return nil, errors.New("Received unknown function invocation: " + function)
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  fmt.Println("query is running " + function)
  
  return nil, errors.New("Received unknown function query: " + function)
}
