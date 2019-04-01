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

//###############           FUNCTION addCountry (For adding Country in Ledger )   ###############
func addCountry(stub shim.ChaincodeStubInterface, args []string) pb.Response { 

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
func updateStatus(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
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

//###############           FUNCTION updatePorts (For updating Ports of peers  of country in Ledger)   ###############
func updatePorts(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
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


//###############           FUNCTION addCountry (For adding Country in Ledger )   ###############
func addCount(stub shim.ChaincodeStubInterface, args []string) pb.Response { 


    err1 := stub.PutState("rr", []byte("6666"))
    if err1 != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to add addCount in Ledger  @@@@@@@@@@@@@@@@@@@")
    }
return shim.Success([]byte("######################  Add count Successfully in Ledger  ######################"))

}



