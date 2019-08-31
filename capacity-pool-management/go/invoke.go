package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// ===============================================
// newContractorRegistration() - registeration of new contarctors
// ===============================================
func newContractorRegistration(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//Check the correct no. of arguments in an array
	if len(args) < 2 {
		str := fmt.Sprintf(ArgumentErrorMessage)
		return shim.Error(str)
	}
	
	uuid    := args[0]
	regInfo := args[1]

    //Check same uuid exists in a ledger or not
	uuidAsBytes, err := stub.GetState(uuid)
	if uuidAsBytes != nil || err != nil {
		return shim.Error("A UUID already exists in the system. Please provide new UUID.")
	}
	
	//Storing new contractor detail corresponding to its UUID
	err = stub.PutState(uuid, []byte(regInfo))
	if err != nil {
		return shim.Error(PutErrorMessage)

	}

	return shim.Success([]byte("Contratcor Registered successfully."))
}


// ===============================================
// updateContractorDetails() - update contractors details
// ===============================================
func updateContractorDetails(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//Check the correct no. of arguments in an array
	if len(args) < 2 {
		str := fmt.Sprintf(ArgumentErrorMessage)
		return shim.Error(str)
	}
	
	uuid    := args[0]
	updatedInfo := args[1]

    //Check uuid exists in a ledger or not
	uuidAsBytes, err := stub.GetState(uuid)
	if uuidAsBytes == nil || err == nil {
		return shim.Error("The contractor with this uuid doesnot exists in the system.")
	}
	
	//Updating the contractor detail corresponding to its UUID
	err = stub.PutState(uuid, []byte(updatedInfo))
	if err != nil {
		return shim.Error(PutErrorMessage)

	}

	return shim.Success([]byte("Contratcor Details are successfully updated in ledger."))
}


// ===============================================================
// deleteContractor() - Delete extisting Contractor from the world state
// ===============================================================
func deleteContractor(stub shim.ChaincodeStubInterface, UUID string) pb.Response {

	err := stub.DelState(UUID)
	if err != nil {
		return shim.Error(DeleteStateErrorMessage)
	}

	return shim.Success([]byte("Contractor deleted successfully from the world state. "))
}
