
package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type sampleSmartContract struct {

}

func main() {
	err := shim.Start(new(sampleSmartContract))
	if err != nil {
		fmt.Println("Error starting sampleSmartContract: %s", err)
	}
}

func (s *sampleSmartContract) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){

		if(len(args) < 1){
			return nil, errors.New("Please input the initial value to store")
		}

		err := stub.PutState("myBlockchain",[]byte(args[1]));
		if err != nil{
			return nil, err
		}

		return nil, nil
}

func (s *sampleSmartContract) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){

	if(len(args) < 1){
			return nil, errors.New("Please input the value to update")
	}

	err := stub.PutState("myBlockchain",[]byte(args[1]));
		if err != nil{
			return nil, err
		}

	return nil, nil

}

func (s *sampleSmartContract) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){

	valAsBytes, err := stub.GetState("myBlockchain")
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state }"
        return nil, errors.New(jsonResp)
	}

	return valAsBytes, nil

}
