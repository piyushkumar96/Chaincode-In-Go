package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


//=======================================
//	Chain Code Structure Declaration
//=======================================
type ChainCode struct {

}

var Logger = shim.NewLogger("[UNDPChainCode: ]")

/*
 * The Init method is called when the Smart Contract is instantiated by the blockchain network
 */
func (t *ChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	Logger.Info("Initializing the chain code")

	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (t *ChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()

	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "newContractorRegistration" {
		return newContractorRegistration(stub, args)
	} else if function == "updateContractorDetails" {
	    return updateContractorDetails(stub, args)
	} else if function == "deleteContractor" {
	    return deleteContractor(stub, args[0])
	} else if function == "getContractorBySearchingCriteria" {
	    return getContractorBySearchingCriteria(stub, args)
	} else {			
		return shim.Error("Invalid function call")
	}
}

// ===================================================================================
// Bootstrap chaincode
// ===================================================================================
func main() {
	err := shim.Start(new(ChainCode))
	if err != nil {
		fmt.Printf("Error starting in chaincode: %s", err)
	}
}