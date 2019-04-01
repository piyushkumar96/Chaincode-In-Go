
package main

import (
    
	"encoding/json"
	"bytes"
    "fmt"
    "strconv"
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