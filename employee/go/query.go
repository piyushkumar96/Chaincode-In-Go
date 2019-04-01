package main

import (
    "encoding/json"
    "fmt"
    "strconv"
    "bytes"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

//###############           FUNCTION noofemployees (For getting No. of Employees  in Ledger)   ###############
func (s *NewChaincode) noofemployees(stub shim.ChaincodeStubInterface) pb.Response {
    
    noofempbytes, err := stub.GetState("noofemployees")
    if err != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@ Failed to get NoOfEmployees Value from Ledger @@@@@@@@@@@@@@@@@@@")
    }

    return shim.Success(noofempbytes)
}

//###############           FUNCTION queryDetailsWithId  (For getting detials of single Employee using Id from Ledger) ###############
func (s *NewChaincode) queryDetailsWithId(stub shim.ChaincodeStubInterface, args string) ( []byte, bool ) {
    
    if len(args) != 1 {
        return []byte("@@@@@@@@@@@@@@@@@@@ Incorrect no. of arguments Expecting 1 argument @@@@@@@@@@@@@@@@@@@") , false
    }

    startKey := "1"
    noofempbytes, err1 := stub.GetState("noofemployees")
    if err1 != nil {
            return []byte("@@@@@@@@@@@@@@@@@@@ Failed to get NoOfEmployees Value from Ledger @@@@@@@@@@@@@@@@@@@") , false
    }

    endKey := string(noofempbytes)
    
    resultsIterator, err2 := stub.GetStateByRange(startKey, endKey)
    if err2 != nil {
	return []byte("@@@@@@@@@@@@@@@@@@@ Failed to Iterate @@@@@@@@@@@@@@@@@@@") , false
    }
    defer resultsIterator.Close()
    
    var buffer bytes.Buffer

    for resultsIterator.HasNext() {
	queryResponse, err3 := resultsIterator.Next()
	if err3 != nil {
		return []byte("@@@@@@@@@@@@@@@@@@@ Failed to Iterate @@@@@@@@@@@@@@@@@@@") , false
	}
        
        buffer.WriteString(queryResponse.Key)
    
        empBytes, err4 := stub.GetState(string(buffer.Bytes()))
        if err4 != nil {
        return []byte("@@@@@@@@@@@@@@@@@@@ Employee doesnot exists @@@@@@@@@@@@@@@@@@@") , false                       
        }

        emp := Employee{}
        err5 :=  json.Unmarshal(empBytes, &emp)
        if err5 != nil {
		     return []byte(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@  Could not unmarshal Employee with ID  %s @@@@@@@@@@@@@@@@@@@", args[0])) , false
	    }

        if emp.empid == args[0]  {
             return empBytes , true
        }
         buffer.Reset()  
      }
     return []byte(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@  Unable to get Employee with ID  %s @@@@@@@@@@@@@@@@@@@", args[0])) , false
}


//###############           FUNCTION queryAllEmpInCompany (For getting detials of all Employees in company from Ledger)   ###############
func (s *NewChaincode) queryAllEmpInCompany(stub shim.ChaincodeStubInterface) pb.Response {

	startKey := "1"
        noofemployees := "noofemployees" 
	noofempbytes, err1 := stub.GetState(noofemployees)
        if err1 != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@ Failed to get NoOfEmployees Value from Ledger @@@@@@@@@@@@@@@@@@@") 
        }
        endKey := string(noofempbytes)

	resultsIterator, err2 := stub.GetStateByRange(startKey, endKey)
	if err2 != nil {
		return shim.Error(err2.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err3 := resultsIterator.Next()
		if err3 != nil {
			return shim.Error(err3.Error())
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

	fmt.Printf("- queryAllEmpInCompany:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}