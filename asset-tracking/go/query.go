package main

import (
    
    "fmt"
    "bytes"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

//###############           FUNCTION queryDetails  (For getting detials of single Order in Ledger) ###############
func (t *TrackChaincode) queryDetails(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    if len(args) != 1{
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of arguments Expecting 1 argument  @@@@@@@@@@@@@@@@@@@")
    }

    var Orderid = args[0]
    Orderid = Orderid[4:len(Orderid)]
    
    ordBytes, _ := stub.GetState(Orderid)
    if ordBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate Order  @@@@@@@@@@@@@@@@@@@")
    }
    str := "["+string(ordBytes)+"]"
       
return shim.Success([]byte(str))

}

//###############           FUNCTION queryAllOrder (For getting detials of all Order in Ledger)   ###############
func (s *TrackChaincode) queryAllOrder(stub shim.ChaincodeStubInterface) pb.Response {

	startKey := "111"
	endKey := "9999"

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add comma before array members,suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllOrder:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

//###############           FUNCTION main   ############### 
func main() {
	// Create a new Smart Contract
	err := shim.Start(new(TrackChaincode))
	if err != nil {
		fmt.Printf("@@@@@@@@@@@@@@@@@@@  Error creating new TrackChaincode : %s  @@@@@@@@@@@@@@@@@@@", err)
	}
} 
