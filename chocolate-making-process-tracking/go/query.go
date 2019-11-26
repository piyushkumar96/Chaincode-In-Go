package main

import (
    
    "fmt"
	"strconv"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

//###############           FUNCTION getCocoaBeanBagById  (For getting details of single Cocoa Bean Bag from Ledger) ###############
func (r *ChocoChaincode) getCocoaBeanBagById(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    if len(args) != 1{
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of arguments Expecting 1 argument  @@@@@@@@@@@@@@@@@@@")
    }
    
    var beanBagId = args[0]
    
    ccbBytes, _ := stub.GetState(beanBagId)
    if ccbBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate Cocoa Bean Bag with given Id  @@@@@@@@@@@@@@@@@@@")
    }
       
	return shim.Success(ccbBytes)
}

//###############           FUNCTION getChocolateBarById  (For getting details of single Chocolate Bar from Ledger) ###############
func (r *ChocoChaincode) getChocolateBarById(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    if len(args) != 1{
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of arguments Expecting 1 argument  @@@@@@@@@@@@@@@@@@@")
    }
    
    var barId = args[0]
    
    cbBytes, _ := stub.GetState(barId)
    if cbBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate Chocolate Bar with given Id  @@@@@@@@@@@@@@@@@@@")
    }
       
	return shim.Success(cbBytes)
}

//###############           FUNCTION getCBCountryCount  (For getting Choco Bars country wise count from Ledger) ###############
func (r *ChocoChaincode) getCBCountryCount(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    if len(args) != 1{
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of arguments Expecting 1 argument  @@@@@@@@@@@@@@@@@@@")
    }

    var countryCode = args[0]
    
    // making the query string 
    queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"country\": \""+ countryCode + "\"}, {\"isConsumed\":true}]}}")
    
    // getting the count of records 
    recordCount, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error("Unable to retrieve Cocoa Bean Bags from ledger."+err.Error())
	}

	numberOfChocoBars := recordCount * 100 
	return shim.Success([]byte(strconv.Itoa(numberOfChocoBars))) 
}

//###############           FUNCTION  getQueryResultForQueryString (Retrieve Number of Records based on query string)    ###############
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) (int, error) {

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return 0, err
	}
	defer resultsIterator.Close()

	count := 0
	for resultsIterator.HasNext() {
		resultsIterator.Next()
		count++;
	}

	return count, nil
}

