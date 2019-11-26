
package main

import (
    "fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

//Chaincode for Tracking of return of product 
type ChocoChaincode struct {   
}


var Logger = shim.NewLogger("[ChocoChaincode: ]")

/*
 * The Init method is called when the Smart Contract is instantiated by the blockchain network
 */	
//###############           FUNCTION Init (For Initiation of chaincode)   ############### 
func (r *ChocoChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
     Logger.Info("Initializing the ChocoChaincode chain code")
     return shim.Success([]byte("###################### Init Successfull ######################"))
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
//###############           FUNCTION Invoke (For Invokation of data inside Ledger)   ############### 
func (r *ChocoChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

     function, args := stub.GetFunctionAndParameters()
     
     if function == "createCocoaBeanBag" {
          return r.createCocoaBeanBag(stub, args)
     }else if function == "createChocolateBar" {
          return r.createChocolateBar(stub, args)
     }else if function == "getCocoaBeanBagById" {
          return r.getCocoaBeanBagById(stub, args)
     }else if function == "getChocolateBarById" {
          return r.getChocolateBarById(stub, args)
     }else if function == "getCBCountryCount" {
          return r.getCBCountryCount(stub, args)
     }else {
          return shim.Error("@@@@@@@@@@@@@@@@@@@  Function called doesnot exits @@@@@@@@@@@@@@@@@@@")
     }
         
 }

// ===================================================================================
// Bootstrap chaincode
// ===================================================================================
//###############           FUNCTION main   ############### 
func main() {
	// Create a new Smart Contract
	err := shim.Start(new(ChocoChaincode))
	if err != nil {
		fmt.Printf("@@@@@@@@@@@@@@@@@@@  Error creating new ChocoChaincode : %s  @@@@@@@@@@@@@@@@@@@", err)
	}
} 



