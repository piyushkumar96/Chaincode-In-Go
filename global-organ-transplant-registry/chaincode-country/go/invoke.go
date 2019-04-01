package main

import (
	"encoding/json"
    "time"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//###############           FUNCTION addCountry (For adding Country in Ledger )   ###############
func addCountry(stub shim.ChaincodeStubInterface, args []string) pb.Response { 

    shim.Success([]byte("######################  Entering in addCountry Method  ######################"))

    if len(args) != 6 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of the arguments,  Expecting 6 Arguments @@@@@@@@@@@@@@@@@@@")
    }
    

    
    current_time := time.Now().Format("02/01/2006")
    countryName := args[0]
    address := args[1];
    zipcode := args[2];
    city := args[3];
    phoneNo := args[4];
    countryId := args[5];     
    countryInfo := `{ "countryId" : "` + countryId  + `" , "countryName" : "` + countryName + `" , "address" : "` + address + `" , "zipcode" : "` + zipcode + `" , "city" : "` + city + `" , "phoneNo" : "` + phoneNo + `" , "regSince" : "` + current_time + `" }`
    
    err := stub.PutState(countryId, []byte(countryInfo))
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to add addCountry in Ledger  @@@@@@@@@@@@@@@@@@@")
    }
    
return shim.Success([]byte("######################  Add country Successfully in Ledger  ######################"))

}


//###############           FUNCTION updateCountry (For updating details  of country in Ledger)   ###############
func updateCountry(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in updateCountry Method  ######################"))

    if len(args) != 6 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of the arguments,  Expecting 6 Arguments @@@@@@@@@@@@@@@@@@@")
    }
    
    countryName := args[0]
    address := args[1];
    zipcode := args[2];
    city := args[3];
    phoneNo := args[4];
    countryId := args[5];   

    countryBytes, _ := stub.GetState(countryId)
    if countryBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate country  @@@@@@@@@@@@@@@@@@@")
    }
    
    var raw map[string]string
    err :=  json.Unmarshal(countryBytes, &raw)
    if err != nil {
		return shim.Error("@@@@@@@@@@@@@@@@@@@  Could not unmarshal Country with ID  @@@@@@@@@@@@@@@@@@@")
	}
    raw["countryName"] = countryName
    raw["address"] = address
    raw["zipcode"] = zipcode
    raw["city"] = city
    raw["phoneNo"] = phoneNo
    raw["countryId"] = countryId   
        

    countrybytes,_ := json.Marshal(raw)
    err1 := stub.PutState(countryId, []byte(countrybytes))
    if err1 != nil{
        return shim.Error(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@  Failed to update the details of Country @@@@@@@@@@@@@@@@@@@"))
    }
 return shim.Success(countrybytes)
}
