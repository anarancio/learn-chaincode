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

	weatherOracles := []string{}
	weatherOracles = append(weatherOracles, "test")
	fmt.Println("oracles:")
	fmt.Println(weatherOracles)

	bytes, err := json.Marshal(weatherOracles)
	if err != nil { return nil, errors.New("Error creating weather oracles devices") }
	fmt.Println(bytes)

	fmt.Println("saving devices into the state")
	fmt.Println(bytes)
	stub.PutState("weather_oracles_devices", bytes)

	fmt.Println("TX ID:")
	fmt.Println(stub.GetTxID())

	return nil, nil
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} else if function == "add_weather_oracle" {
		return t.add_weather_oracle(stub, args)
	}

  fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation: " + function)
}

func (t *SimpleChaincode) add_weather_oracle(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("add_weather_oracle called")
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2 (UUID, name)")
	}

	bytes, err := stub.GetState("weather_oracles_devices")
	if err != nil { return nil, errors.New("Error getting Devices_List record") }
	var devices []string
	err = json.Unmarshal(bytes, &devices)

	devices = append(devices, args[0])
	bytes, err = json.Marshal(devices)
	if err != nil { return nil, errors.New("Error creating weather oracles devices") }
	stub.PutState("weather_oracles_devices", bytes)

	return nil, nil
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  fmt.Println("query is running " + function)
  if function == "get_oracle_weather_devices" {
   return t.get_oracle_weather_devices(stub)
  }
  return nil, errors.New("Received unknown function query: " + function)
}

func (t *SimpleChaincode) get_oracle_weather_devices(stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("get_oracle_weather_devices called")
	bytes, err := stub.GetState("weather_oracles_devices")
	fmt.Println(bytes)

	if err != nil { return nil, errors.New("Error getting Devices_List record") }

	return bytes, nil
}
