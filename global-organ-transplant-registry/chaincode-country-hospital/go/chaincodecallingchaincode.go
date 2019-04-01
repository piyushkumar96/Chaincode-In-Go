package main

import (
    "bytes"
    "fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

func toChaincodeArgs(args ...string) [][]byte {
	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}
	return bargs
}

func callOtherChaincodeInvoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var channelName string // Sum entity
	//var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting atleast 2")
	}

	chaincodeName := args[0] // Expecting name of the chaincode you would like to call, this name would be given during chaincode install time
    channelName = args[1]


	// Query chaincode_example02
	f := "addCount"
	queryArgs := toChaincodeArgs(f, "ee")
        //return shim.Success(queryArgs)
       // ab := string(queryArgs[0]) + string(queryArgs[1]);
       // return shim.Success([]byte(ab))


	response := stub.InvokeChaincode(chaincodeName, queryArgs, channelName)
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", response.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}
	
	return shim.Success([]byte(string(response.Payload)))
}

func callOtherChaincodeQuery(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var channelName string // Sum entity
	//var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting atleast 2")
	}

	chaincodeName := args[0] // Expecting name of the chaincode you would like to call, this name would be given during chaincode install time
    channelName = args[1]


	// Query chaincode_example02
	f := "queryCount"
	queryArgs := toChaincodeArgs(f, "ee")
        //return shim.Success(queryArgs)
       // ab := string(queryArgs[0]) + string(queryArgs[1]);
       // return shim.Success([]byte(ab))


	response := stub.InvokeChaincode(chaincodeName, queryArgs, channelName)
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", response.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}
	
	return shim.Success([]byte(string(response.Payload)))
}