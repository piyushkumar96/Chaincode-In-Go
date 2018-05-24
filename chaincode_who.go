
package main

import (
    
    "encoding/json"
    "strings"
    "time"
    "fmt"
    "bytes"
    "strconv"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

//Chaincode for WHO
type WhoChaincode struct {   
}

func toChaincodeArgs(args ...string) [][]byte {
	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}
	return bargs
}

func (t *WhoChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

    _, args := stub.GetFunctionAndParameters()
    var count = args[0]
    var value = args[1]
    err := stub.PutState(count,[]byte(value))
    if err != nil {
        return shim.Error(err.Error())
    }
     
    return shim.Success([]byte("###################### Init Successfull ######################"))
}

//###############           FUNCTION Invoke (For Invokation of data inside Ledger)   ############### 
func (t *WhoChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

    function, args := stub.GetFunctionAndParameters()

    if len(args) == 0 {
         return shim.Error("@@@@@@@@@@@@@@@@@@@ Expecting Arguments @@@@@@@@@@@@@@@@@@@")
    }
    
    if function == "addCountry" {
         return t.addCountry(stub, args)
    } else if function == "queryAllCountries" {
         return t.queryAllCountries(stub)
    } else if function == "queryCountry" {
         return t.queryCountry(stub, args)
    } else if function == "recentRegisteredCountries" {
         return t.recentRegisteredCountries(stub)
    } else if function == "updateStatus" {
         return t.updateStatus(stub, args)
    } else if function == "queryStatus" {
         return t.queryStatus(stub, args)
    } else if function == "updatePorts" {
         return t.updatePorts(stub, args)
    } else if function == "queryPorts" {
         return t.queryPorts(stub, args)
    } else if function == "queryCount" {    
         return t.queryCount(stub, args)
    }  else if function == "addCount" {
         return t.addCount(stub, args)
    }  else if function == "queryAllCountryCount" {
         return t.queryAllCountryCount(stub)
    } else {
         return shim.Error("@@@@@@@@@@@@@@@@@@@  Function called doesnot exits @@@@@@@@@@@@@@@@@@@")
    }
        
}

//###############           FUNCTION addCountry (For adding Country in Ledger )   ###############
func (t *WhoChaincode) addCountry(stub shim.ChaincodeStubInterface, args []string) pb.Response { 

    shim.Success([]byte("######################  Entering in addCountry Method  ######################"))

    if len(args) != 1 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of the arguments,  Expecting 1 Arguments @@@@@@@@@@@@@@@@@@@")
    }
    

    countBytes, err := stub.GetState("count")
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ count doesnot exists @@@@@@@@@@@@@@@@@@@")                        
    }
    count1, _ := strconv.Atoi(string(countBytes))
    count1++;
    
    current_time := time.Now().Format("02/01/2006")
    countryName := args[0]
    channelName1 := "channelwho"+strings.ToLower(countryName);
    channelName2 := "channel"+strings.ToLower(countryName);      
    countryInfo := `{ "countryId" : "` + strconv.Itoa(count1) + `" , "countryName" : "` + countryName + `" , "regSince" : "` + current_time + `" , "channelName1" : "` + channelName1 + `" , "channelName2" : "` + channelName2 +`" ,"status" : "inactive", "peer0port" : "", "peer1port" : "", "reqport" : "" }`
    
    err1 := stub.PutState(strconv.Itoa(count1), []byte(countryInfo))
    if err1 != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to add addCountry in Ledger  @@@@@@@@@@@@@@@@@@@")
    }
    err2 := stub.PutState("count", []byte(strconv.Itoa(count1)))
    if err2 != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to update count of country in Ledger  @@@@@@@@@@@@@@@@@@@")
    }
return shim.Success([]byte("######################  Add country Successfully in Ledger  ######################"))

}


//###############           FUNCTION updateStatus (For updating status  of country in Ledger)   ###############
func (t *WhoChaincode) updateStatus(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in updateStatus Method  ######################"))
    
    if len(args) != 2 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of the arguments,  Expecting 2 Arguments @@@@@@@@@@@@@@@@@@@")
    }
    
    var countryid = args[0]
    var status = args[1]
    countryBytes, _ := stub.GetState(countryid)
    if countryBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate country  @@@@@@@@@@@@@@@@@@@")
    }
    
    var raw map[string]string
    err :=  json.Unmarshal(countryBytes, &raw)
    if err != nil {
		return shim.Error("@@@@@@@@@@@@@@@@@@@  Could not unmarshal Country with ID  @@@@@@@@@@@@@@@@@@@")
	}
    raw["status"] = status
    countrybytes,_ := json.Marshal(raw)
    err1 := stub.PutState(countryid, []byte(countrybytes))
    if err1 != nil{
        return shim.Error(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@  Failed to update the status of Country @@@@@@@@@@@@@@@@@@@"))
    }
 return shim.Success(countrybytes)
}

//###############           FUNCTION queryStatus (For querying status of country from Ledger)   ###############
func (t *WhoChaincode) queryStatus(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
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
func (t *WhoChaincode) queryCountry(stub shim.ChaincodeStubInterface, args []string) pb.Response {

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
func (s *WhoChaincode) queryAllCountries(stub shim.ChaincodeStubInterface) pb.Response {

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
func (s *WhoChaincode) recentRegisteredCountries(stub shim.ChaincodeStubInterface) pb.Response {

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

//###############           FUNCTION updatePorts (For updating Ports of peers  of country in Ledger)   ###############
func (t *WhoChaincode) updatePorts(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in updateStatus Method  ######################"))
    if len(args) != 4 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of the arguments,  Expecting 4 Arguments @@@@@@@@@@@@@@@@@@@")
    }
    
    var countryid = args[0]
    var port1 = args[1]
    var port2 = args[2]
    var port3 = args[3]

    countryBytes, _ := stub.GetState(countryid)
    if countryBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate country  @@@@@@@@@@@@@@@@@@@")
    }
    
    var raw map[string]string
    err :=  json.Unmarshal(countryBytes, &raw)
    if err != nil {
		return shim.Error("@@@@@@@@@@@@@@@@@@@  Could not unmarshal Country with ID  @@@@@@@@@@@@@@@@@@@")
	}
    raw["peer0port"] = port1
     raw["peer1port"] = port2
      raw["reqport"] = port3
    countrybytes,_ := json.Marshal(raw)
    err1 := stub.PutState(countryid, []byte(countrybytes))
    if err1 != nil{
        return shim.Error(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@  Failed to update the status of Country @@@@@@@@@@@@@@@@@@@"))
    }
 return shim.Success(countrybytes)
}

//###############           FUNCTION queryPorts (For querying detials of Ports from Ledger)   ###############
func (t *WhoChaincode) queryPorts(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
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


//###############           FUNCTION addCountry (For adding Country in Ledger )   ###############
func (t *WhoChaincode) addCount(stub shim.ChaincodeStubInterface, args []string) pb.Response { 


    err1 := stub.PutState("rr", []byte("6666"))
    if err1 != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to add addCount in Ledger  @@@@@@@@@@@@@@@@@@@")
    }
return shim.Success([]byte("######################  Add count Successfully in Ledger  ######################"))

}


//###############           FUNCTION queryCount (For querying detials of Ports from Ledger)   ###############
func (t *WhoChaincode) queryCount(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in queryCount Method  ######################"))

    var queryc = args[0]
    querycBytes, _ := stub.GetState(queryc)
    if querycBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate country  @@@@@@@@@@@@@@@@@@@")
    }
    
 return shim.Success([]byte(string(querycBytes)))
}


//###############           FUNCTION queryAllCountryCount (For querying all country counts  from Ledger)   ###############
func (t *WhoChaincode) queryAllCountryCount(stub shim.ChaincodeStubInterface) pb.Response {
    
   countBytes, err := stub.GetState("count")
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ count doesnot exists @@@@@@@@@@@@@@@@@@@")                        
    }
    count1, _ := strconv.Atoi(string(countBytes)) 
   
   str := "[ "

   var i int
    b := false 
    for i=1; i <= count1; i++ { 
      countryBytes, _ := stub.GetState(strconv.Itoa(i))
       
        if countryBytes == nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate country  @@@@@@@@@@@@@@@@@@@")
        }
        var raw map[string]string
        err1 :=  json.Unmarshal(countryBytes, &raw)
        if err1 != nil {
		    return shim.Error("@@@@@@@@@@@@@@@@@@@  Could not unmarshal Country with ID  @@@@@@@@@@@@@@@@@@@")
	    }
       
        chaincodeName := raw["channelName2"] 
	    channelName := raw["channelName2"]
    
       
	
	f := "queryCountryCount"
	queryArgs := toChaincodeArgs(f,"a")

	//   if chaincode being invoked is on the same channel,
	//   then channel defaults to the current channel and args[2] can be "".
	//   If the chaincode being called is on a different channel,
	//   then you must specify the channel name in args[2]

	response := stub.InvokeChaincode(chaincodeName, queryArgs, channelName)
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", response.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}
        
        if b == true {
            str = str + ","
        }
        str = str + string(response.Payload)
        b = true 
    } 
    str = str + " ]"
	return shim.Success([]byte(str))   

}


//###############           FUNCTION main   ############### 
func main() {
	// Create a new Smart Contract
	err := shim.Start(new(WhoChaincode))
	if err != nil {
		fmt.Printf("@@@@@@@@@@@@@@@@@@@  Error creating new WhoChaincode : %s  @@@@@@@@@@@@@@@@@@@", err)
	}
} 



