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

  f := "add_consumer"
  invokeArgs := util.ToChaincodeArgs(f, stub.GetTxID(), "LCD Display device")
  stub.InvokeChaincode(args[0], invokeArgs)

  temp := "No disponible"
  bytes, err := json.Marshal(temp)
	if err != nil { return nil, errors.New("Error creating temp") }
  stub.PutState("consumer_temperature", bytes)

  return nil, nil
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  fmt.Println("invoke is running " + function)
  if function == "new_temp" {
    return t.new_temp(stub, args)
  }

  return nil, errors.New("Received unknown function invocation: " + function)
}

func (t *SimpleChaincode) new_temp(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
  fmt.Println("new_temp called")
  if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1 (new temp)")
	}

  temp := args[0]
  bytes, err := json.Marshal(temp)
	if err != nil { return nil, errors.New("Error updating temp") }
  stub.PutState("consumer_temperature", bytes)

  return nil, nil
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  fmt.Println("query is running " + function)

  if function == "get_temp" {
    return t.get_temperature(stub)
  }

  return nil, errors.New("Received unknown function query: " + function)
}

func (t *SimpleChaincode) get_temperature(stub shim.ChaincodeStubInterface) ([]byte, error) {
  fmt.Println("get_temperature called")
	bytes, err := stub.GetState("consumer_temperature")
	fmt.Println(bytes)

	if err != nil { return nil, errors.New("Error getting temperature record") }

	return bytes, nil
}
