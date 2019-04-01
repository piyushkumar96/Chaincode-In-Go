package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//###############           FUNCTION queryCountry  (For getting detials of single country in Ledger) ###############
func queryCountry(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    if len(args) != 1{
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of arguments Expecting 1 argument  @@@@@@@@@@@@@@@@@@@")
    }

    countryId := args[0]
    countryBytes, _ := stub.GetState(countryId)
    if countryBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate country  @@@@@@@@@@@@@@@@@@@")
    }
    str := "["+string(countryBytes)+"]"
       
return shim.Success([]byte(str))

}
