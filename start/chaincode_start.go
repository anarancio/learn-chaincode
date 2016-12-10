/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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

type Devices_List struct {
	weatherOracles 	[]string `json:"weatherOracles"`
	displayProviders []string `json:"displayProviders"`
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("init " + args[0])
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	weatherOracles := []string{}
	weatherOracles = append(weatherOracles, "test")
	fmt.Println("oracles:")
	fmt.Println(weatherOracles)

	bytes1, err := json.Marshal(weatherOracles)
	fmt.Println(bytes1)

	var devicesList Devices_List
	devicesList.weatherOracles = append(devicesList.weatherOracles, "test")
	bytes, err := json.Marshal(devicesList)

	if err != nil { return nil, errors.New("Error creating Devices_List record") }

	fmt.Println("saving devices into the state")
	fmt.Println(bytes)
	stub.PutState("devices", bytes)

	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	}
	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "dummy_query1" {											//read a variable
		fmt.Println("hi there " + function)						//error
		return nil, nil;
	} else if function == "get_devices" {
		return t.get_devices(stub)
	}
	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query: " + function)
}

func (t *SimpleChaincode) add_oracle(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("add_oracle called")
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2 (oracle type, oracle chaincode id)")
	}

	bytes, err := stub.GetState("devices")
	if err != nil { return nil, errors.New("Error getting Devices_List record") }

	var devicesList Devices_List
	err = json.Unmarshal(bytes, &devicesList)

	return nil, nil
}

func (t *SimpleChaincode) get_devices(stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("get_devices called")
	bytes, err := stub.GetState("devices")
	fmt.Println(bytes)

	if err != nil { return nil, errors.New("Error getting Devices_List record") }

	return bytes, nil
}
