package main

import (

    "fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

type NewChaincode struct {   
}

var Logger = shim.NewLogger("[NewChaincode: ]")

/*
 * The Init method is called when the Smart Contract is instantiated by the blockchain network
 */

//###############           FUNCTION Init (For Initiation of chaincode)   ############### 
func (s *NewChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
     Logger.Info("Initializing the NewChaincode chain code")
 _, args := stub.GetFunctionAndParameters()

 err2 := stub.PutState(args[0],[]byte(args[1]))
 if err2 != nil {
     return shim.Error(err2.Error())
 }

 return shim.Success([]byte("###################### Successfully Init of Ledger  ######################"))
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
//###############           FUNCTION Invoke (For Invokation of data inside Ledger)   ############### 
func (s *NewChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    
     function, args := stub.GetFunctionAndParameters()
     
     if len(args) == 0 {
          return shim.Error("@@@@@@@@@@@@@@@@@@@ Expecting Arguments @@@@@@@@@@@@@@@@@@@")
     }
 
     if function == "changeDetail"{
             return s.changeDetail(stub, args)
     }else if function == "addMember" {
         return s.addMember(stub, args)
     }else if function == "transferMoney"{
             return s.transferMoney(stub,args)
     }else if function == "queryAllEmpInCompany"{
             return s.queryAllEmpInCompany(stub)
     }else if function == "delEmployee"{
             return s.delEmployee(stub,args)
     }else if function == "noofemployees"{
             return s.noofemployees(stub)
     }
 
     return shim.Error("@@@@@@@@@@@@@@@@@@@ Invalid Fucntion Name Invoked @@@@@@@@@@@@@@@@@@@")
 } 

// ===================================================================================
// Bootstrap chaincode
// ===================================================================================
//###############           FUNCTION main   ############### 
func main() {
	// Create a new Smart Contract
	err := shim.Start(new(NewChaincode))
	if err != nil {
		fmt.Printf("@@@@@@@@@@@@@@@@@@@  Error creating new NewChaincode : %s  @@@@@@@@@@@@@@@@@@@", err)
	}
} 



