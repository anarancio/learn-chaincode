package main

import (
	"errors"
	"fmt"
  "encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SimpleChaincode struct {
}

type Device_List struct {
  devices []string `json:"devices"`
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Service Locator Chaincode: %s", err)
	}
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

  var deviceList Device_List
  bytes, err := json.Marshal(deviceList)

  if err != nil { return nil, errors.New("Error creating the device list!") }

  stub.PutState("devices", bytes)

  return nil, nil
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} else if function == "test" {
		return stub.GetState("devices")
	}
	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)
  fmt.Println("TESTTTT")

	// Handle different functions
	if function == "dummy_query1" {											//read a variable
		fmt.Println("hi there " + function)						//error
		//return nil, nil;
    t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query: " + function)
}

func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
  valAsbytes, err := stub.GetState("devices")
  if err != nil {
    return nil, errors.New("Failed to get state for devices")
  }
  return valAsbytes,nil
}
