
package main

import (
    
    "encoding/json"
    "time"
    "fmt"
    "bytes"
    "strconv"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)


//###############           FUNCTION queryStatus (For querying status of country from Ledger)   ###############
func queryStatus(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in queryStatus Method  ######################"))
    
   if len(args) != 1 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of the arguments,  Expecting 1 Arguments @@@@@@@@@@@@@@@@@@@")
    }
    

    var countryid = args[0]
    countryBytes, _ := stub.GetState(countryid)
    if countryBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate country  @@@@@@@@@@@@@@@@@@@")
    }
    
    var raw map[string]string
    err :=  json.Unmarshal(countryBytes, &raw)
    if err != nil {
		return shim.Error("@@@@@@@@@@@@@@@@@@@  Could not unmarshal Country with ID  @@@@@@@@@@@@@@@@@@@")
	}
    
 return shim.Success([]byte(raw["status"]))
}

//###############           FUNCTION queryCountry  (For getting detials of single country in Ledger) ###############
func queryCountry(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    if len(args) != 1{
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of arguments Expecting 1 argument  @@@@@@@@@@@@@@@@@@@")
    }

    var countryId = args[0]
    countryBytes, _ := stub.GetState(countryId)
    if countryBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate country  @@@@@@@@@@@@@@@@@@@")
    }
    str := "["+string(countryBytes)+"]"
       
return shim.Success([]byte(str))

}


//###############           FUNCTION queryAllCountries (For getting details of all countries in Ledger)   ###############
func queryAllCountries(stub shim.ChaincodeStubInterface) pb.Response {

    countBytes, err := stub.GetState("count")
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ count doesnot exists @@@@@@@@@@@@@@@@@@@")                        
    }
    count1, _ := strconv.Atoi(string(countBytes))
    
    str := "[ "

    var i int
    b := false 
    
    k:=0;
    if count1 < 3 {
         k=0;
    } else {
         k = count1-3;
    }


    for i=count1; i>k; i-- {
        countryBytes, _ := stub.GetState(strconv.Itoa(i))
        if countryBytes == nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate country  @@@@@@@@@@@@@@@@@@@")
        }
        
        if b == true {
            str = str + ","
        }
        str = str + string(countryBytes) 
        b = true 
    } 

    str = str + " ]"
	return shim.Success([]byte(str))    
}


//###############           FUNCTION recentRegisteredCountries (For getting detials of first  three  countries in Ledger)   ###############
func recentRegisteredCountries(stub shim.ChaincodeStubInterface) pb.Response {

    countBytes, err := stub.GetState("count")
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ count doesnot exists @@@@@@@@@@@@@@@@@@@")                        
    }
    count1, _ := strconv.Atoi(string(countBytes))
    
    var startKey string
    if count1 < 3 {
        startKey = "1"
    } else {
        startKey = strconv.Itoa(count1-3)
    }
	
	endKey := strconv.Itoa(count1)

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
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("recentRegisteredCountries:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}


//###############           FUNCTION queryPorts (For querying detials of Ports from Ledger)   ###############
func queryPorts(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in queryStatus Method  ######################"))

    var countryid = args[0]
    countryBytes, _ := stub.GetState(countryid)
    if countryBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate country  @@@@@@@@@@@@@@@@@@@")
    }
    
    var raw map[string]string
    err :=  json.Unmarshal(countryBytes, &raw)
    if err != nil {
		return shim.Error("@@@@@@@@@@@@@@@@@@@  Could not unmarshal Country with ID  @@@@@@@@@@@@@@@@@@@")
	}
    str := `{ "peer0port": `+ raw["peer0port"] + `, "peer1port": `+ raw["peer1port"] + `, "reqport": `+ raw["reqport"] +` }`
 return shim.Success([]byte(str))
}


//###############           FUNCTION queryCount (For querying detials of Ports from Ledger)   ###############
func queryCount(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in queryCount Method  ######################"))

    var queryc = args[0]
    querycBytes, _ := stub.GetState(queryc)
    if querycBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate country  @@@@@@@@@@@@@@@@@@@")
    }
    
 return shim.Success([]byte(string(querycBytes)))
}