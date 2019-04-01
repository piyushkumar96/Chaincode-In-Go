
package main

import (
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

//Chaincode for Tracking of return of product 
type TrackChaincode struct {   
}


var Logger = shim.NewLogger("[TrackChaincode: ]")

/*
 * The Init method is called when the Smart Contract is instantiated by the blockchain network
 */	
//###############           FUNCTION Init (For Initiation of chaincode)   ############### 
func (t *TrackChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
     Logger.Info("Initializing the TrackChaincode chain code")
     return shim.Success([]byte("###################### Init Successfull ######################"))
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
//###############           FUNCTION Invoke (For Invokation of data inside Ledger)   ############### 
func (t *TrackChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

     function, args := stub.GetFunctionAndParameters()
 
     if len(args) == 0 {
          return shim.Error("@@@@@@@@@@@@@@@@@@@ Expecting Arguments @@@@@@@@@@@@@@@@@@@")
     }
     
     if function == "addOrder" {
          return t.addOrder(stub, args)
     } else if function == "updateDate" {
          return t.updateDate(stub, args)
     } else if function == "queryDetails" {
          return t.queryDetails(stub, args)
     } else if function == "queryAllOrder" {
          return t.queryAllOrder(stub)
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
	err := shim.Start(new(TrackChaincode))
	if err != nil {
		fmt.Printf("@@@@@@@@@@@@@@@@@@@  Error creating new TrackChaincode : %s  @@@@@@@@@@@@@@@@@@@", err)
	}
} 



