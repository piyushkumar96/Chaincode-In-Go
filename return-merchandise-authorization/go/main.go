
package main

import (
    "fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

//Chaincode for Tracking of return of product 
type RMAChaincode struct {   
}


var Logger = shim.NewLogger("[RMAChaincode: ]")

/*
 * The Init method is called when the Smart Contract is instantiated by the blockchain network
 */	
//###############           FUNCTION Init (For Initiation of chaincode)   ############### 
func (r *RMAChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
     Logger.Info("Initializing the RMAChaincode chain code")
     return shim.Success([]byte("###################### Init Successfull ######################"))
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
//###############           FUNCTION Invoke (For Invokation of data inside Ledger)   ############### 
func (r *RMAChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

     function, args := stub.GetFunctionAndParameters()
     
     if function == "createTicket" {
          return r.createTicket(stub, args)
     }else if function == "updates3b11" {
          return r.updates3b11(stub, args)
     }else if function == "updates3b3sir" {
          return r.updates3b3sir(stub, args)
     }else if function == "updates3b3eta" {
          return r.updates3b3eta(stub, args)
     }else if function == "updates3b13" {
          return r.updates3b13(stub, args)
     }else if function == "updates3b3pod" {
          return r.updates3b3pod(stub, args)
     } else if function == "getAllMessagesById" {
          return r.getAllMessagesById(stub, args)
     } else if function == "getAllMessages" {
          return r.getAllMessages(stub)
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
	err := shim.Start(new(RMAChaincode))
	if err != nil {
		fmt.Printf("@@@@@@@@@@@@@@@@@@@  Error creating new RMAChaincode : %s  @@@@@@@@@@@@@@@@@@@", err)
	}
} 



