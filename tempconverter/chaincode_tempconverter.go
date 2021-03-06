package main

import (
	"errors"
	"fmt"
  "encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
  "github.com/hyperledger/fabric/core/util"
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

  f := "add_producer"
  invokeArgs := util.ToChaincodeArgs(f, stub.GetTxID(), "2", "Temperatura en Celsius")
  stub.InvokeChaincode(args[0], invokeArgs)

  listeners := []string{}
  bytes, err := json.Marshal(listeners)
	if err != nil { return nil, errors.New("Error creating listeners state") }

	stub.PutState("converter_listeners", bytes)

  return nil, nil
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  fmt.Println("invoke is running " + function)
  if function == "add_listener" {
    return nil, nil
  } else if function == "new_temp" {
    return nil, nil
  }
  return nil, errors.New("Received unknown function invocation: " + function)
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  fmt.Println("query is running " + function)
  return nil, errors.New("Received unknown function query: " + function)
}
