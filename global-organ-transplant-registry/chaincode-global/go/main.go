
package main

import (
    
    "encoding/json"
    "fmt"
    "bytes"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)


//=======================================
//	Chain Code Structure Declaration
//=======================================
type WHOChaincode struct {   
}

var Logger = shim.NewLogger("[WHOChaincode: ]")

/*
 * The Init method is called when the Smart Contract is instantiated by the blockchain network
 */
 func (t *WHOChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

     _, args := stub.GetFunctionAndParameters()
     var count = args[0]
     var value = args[1]
     err := stub.PutState(count,[]byte(value))
     if err != nil {
         return shim.Error(err.Error())
     }
      
     return shim.Success([]byte("###################### Init Successfull ######################"))
 }


/*
 * The Invoke method is called as a result of an application request to run the Smart Contract
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
//###############           FUNCTION Invoke (For Invokation of data inside Ledger)   ############### 
//###############           FUNCTION Invoke (For Invokation of data inside Ledger)   ############### 
func (t *WHOChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

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

// ===================================================================================
// Bootstrap chaincode
// ===================================================================================
//###############           FUNCTION main   ############### 
func main() {
	// Create a new Smart Contract
	err := shim.Start(new(WHOChaincode))
	if err != nil {
		fmt.Printf("@@@@@@@@@@@@@@@@@@@  Error creating new WHOChaincode : %s  @@@@@@@@@@@@@@@@@@@", err)
	}
} 