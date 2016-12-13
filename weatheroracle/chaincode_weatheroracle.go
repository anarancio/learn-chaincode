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
  invokeArgs := util.ToChaincodeArgs(f, stub.GetTxID(), "1", "Temperatura en Montevideo en Farenheit")
  stub.InvokeChaincode(args[0], invokeArgs)

  listeners := []string{}
  bytes, err := json.Marshal(listeners)
	if err != nil { return nil, errors.New("Error creating listeners state") }

	stub.PutState("oracle_listeners", bytes)

  temp := "No disponible"
  bytes, err = json.Marshal(temp)
	if err != nil { return nil, errors.New("Error creating temp") }
  stub.PutState("temperature", bytes)

  return nil, nil
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  fmt.Println("invoke is running " + function)
  if function == "add_listener" {
    return t.add_listener(stub, args)
  } else if function == "new_temp" {
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
  stub.PutState("temperature", bytes)

  var listeners []string
  bytes, err = stub.GetState("oracle_listeners")
  err = json.Unmarshal(bytes, &listeners)
  for i := 0; i<len(listeners); i++ {
    f := "new_temp"
    invokeArgs := util.ToChaincodeArgs(f, temp)
    stub.InvokeChaincode(listeners[i], invokeArgs)
  }

  return nil, nil
}

func (t *SimpleChaincode) add_listener(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
  fmt.Println("add_listener called")
  if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1 (listener uuid)")
	}

  bytes, err := stub.GetState("oracle_listeners")
	if err != nil { return nil, errors.New("Error getting oracle_listeners state") }
	var listeners []string
	err = json.Unmarshal(bytes, &listeners)

  // create the record for the produce: <UUID>,<type>,<description
	listeners = append(listeners, args[0])
	bytes, err = json.Marshal(listeners)
	if err != nil { return nil, errors.New("Error creating oracle listeners") }
	stub.PutState("oracle_listeners", bytes)

	return nil, nil
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  fmt.Println("query is running " + function)
  if function == "get_temperature" {
    return t.get_temperature(stub)
  } else if function == "get_listeners" {
    return t.get_listeners(stub)
  }
  return nil, errors.New("Received unknown function query: " + function)
}

func (t *SimpleChaincode) get_listeners(stub shim.ChaincodeStubInterface) ([]byte, error) {
  fmt.Println("get_listeners called")
	bytes, err := stub.GetState("oracle_listeners")
	fmt.Println(bytes)

	if err != nil { return nil, errors.New("Error getting temperature record") }

	return bytes, nil
}

func (t *SimpleChaincode) get_temperature(stub shim.ChaincodeStubInterface) ([]byte, error) {
  fmt.Println("get_temperature called")
	bytes, err := stub.GetState("temperature")
	fmt.Println(bytes)

	if err != nil { return nil, errors.New("Error getting temperature record") }

	return bytes, nil
}
