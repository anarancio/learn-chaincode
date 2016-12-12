package main

import (
	"errors"
	"fmt"
  "encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  fmt.Println("init " + args[0])
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	producers := []string{}
	//weatherOracles = append(weatherOracles, "test")

	bytes, err := json.Marshal(producers)
	if err != nil { return nil, errors.New("Error creating producers devices") }

	stub.PutState("producers_devices", bytes)

	return nil, nil
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} else if function == "add_producer" {
		return t.add_producer(stub, args)
	}

  fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation: " + function)
}

func (t *SimpleChaincode) add_producer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("add_producer called")
	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3 (UUID, type, description)")
	}

	bytes, err := stub.GetState("producers_devices")
	if err != nil { return nil, errors.New("Error getting producers_devices state") }
	var devices []string
	err = json.Unmarshal(bytes, &devices)

  // create the record for the produce: <UUID>,<type>,<description
	devices = append(devices, args[0] + "," + args[1] + "," + args[2])
	bytes, err = json.Marshal(devices)
	if err != nil { return nil, errors.New("Error creating weather oracles devices") }
	stub.PutState("producers_devices", bytes)

	return nil, nil
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  fmt.Println("query is running " + function)
  if function == "get_producers" {
   return t.get_producers_devices(stub)
  }
  return nil, errors.New("Received unknown function query: " + function)
}

func (t *SimpleChaincode) get_producers_devices(stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("get_producers_devices called")
	bytes, err := stub.GetState("producers_devices")
	fmt.Println(bytes)

	if err != nil { return nil, errors.New("Error getting producers_devices record") }

	return bytes, nil
}
