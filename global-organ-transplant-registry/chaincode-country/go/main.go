package main

import (

    "fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

//Chaincode for Country
type WHOCountryChaincode struct {   
}

var Logger = shim.NewLogger("[WHOCountryChaincode: ]")

/*
 * The Init method is called when the Smart Contract is instantiated by the blockchain network
 */

func (t *WHOCountryChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	Logger.Info("Initializing the WHOCountryChaincode chain code")
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
func (t *WHOCountryChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

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

// ===================================================================================
// Bootstrap chaincode
// ===================================================================================
//###############           FUNCTION main   ############### 
func main() {
	// Create a new Smart Contract
	err := shim.Start(new(WHOCountryChaincode))
	if err != nil {
		fmt.Printf("@@@@@@@@@@@@@@@@@@@  Error creating new WHOCountryChaincode : %s  @@@@@@@@@@@@@@@@@@@", err)
	}
} 



