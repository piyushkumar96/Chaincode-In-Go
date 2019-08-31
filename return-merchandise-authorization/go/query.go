package main

import (
    
    "fmt"
    "bytes"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

//###############           FUNCTION getAllMessagesById  (For getting detials of single RMA Ticket in Ledger) ###############
func (r *RMAChaincode) getAllMessagesById(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    if len(args) != 1{
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of arguments Expecting 1 argument  @@@@@@@@@@@@@@@@@@@")
    }

    var rmaNo = args[0]
    
    rmaBytes, _ := stub.GetState(rmaNo)
    if rmaBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate RMA Ticket  @@@@@@@@@@@@@@@@@@@")
    }
       
return shim.Success(rmaBytes)

}

//###############           FUNCTION getAllMessages (For getting detials of all RMA Ticket in Ledger)   ###############
func (r *RMAChaincode) getAllMessages(stub shim.ChaincodeStubInterface) pb.Response {

	queryString := fmt.Sprintf("{\"selector\":{}}")
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error("Unable to retrieve all RMA Ticket from ledger."+err.Error())
	}
	return shim.Success(queryResults)
}



//###############           FUNCTION  getQueryResultForQueryString Retrieve Records based on query string    ßßß###############
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return buffer.Bytes(), nil
}