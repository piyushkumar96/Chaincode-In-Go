
package main

import (
    
    "encoding/json"
    "time"
    "fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

//Chaincode for Country
type WhoChaincode struct {   
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
    } else if function == "queryCountry" {
         return t.queryCountry(stub, args)
    } else if function == "updateCountry" {
         return t.updateCountry(stub, args)
    } else {
         return shim.Error("@@@@@@@@@@@@@@@@@@@  Function called doesnot exits @@@@@@@@@@@@@@@@@@@")
    }
        
}

//###############           FUNCTION addCountry (For adding Country in Ledger )   ###############
func (t *WhoChaincode) addCountry(stub shim.ChaincodeStubInterface, args []string) pb.Response { 

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
func (t *WhoChaincode) updateCountry(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
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

//###############           FUNCTION queryCountry  (For getting detials of single country in Ledger) ###############
func (t *WhoChaincode) queryCountry(stub shim.ChaincodeStubInterface, args []string) pb.Response {

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


//###############           FUNCTION main   ############### 
func main() {
	// Create a new Smart Contract
	err := shim.Start(new(WhoChaincode))
	if err != nil {
		fmt.Printf("@@@@@@@@@@@@@@@@@@@  Error creating new WhoChaincode : %s  @@@@@@@@@@@@@@@@@@@", err)
	}
} 



