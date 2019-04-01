package main

import (
	
	"fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)


//=======================================
//	Chain Code Structure Declaration
//=======================================
//Chaincode for Hospital
type WHOCountryHospitalChainCode struct {   
}

var Logger = shim.NewLogger("[WHOCountryHospitalChainCode: ]")

/*
 * The Init method is called when the Smart Contract is instantiated by the blockchain network
 */
 func (t *WHOCountryHospitalChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
    Logger.Info("Initializing the WHOCountryHospital chain code")
    _, args := stub.GetFunctionAndParameters()
    var count = args[0]
    var value = args[1]
    var recepient = args[2]
    var value2 = args[3]
    var donor = args[4]
    var value3 = args[5]
    var transplant = args[6]
    var value4 = args[7]
    err := stub.PutState(count,[]byte(value))
    if err != nil {
        return shim.Error(err.Error())
    }
    err1 := stub.PutState(recepient,[]byte(value2))
    if err1 != nil {
        return shim.Error(err1.Error())
    }
    err2 := stub.PutState(donor,[]byte(value3))
    if err2 != nil {
        return shim.Error(err2.Error())
    }
    err3 := stub.PutState(transplant,[]byte(value4))
    if err3 != nil {
        return shim.Error(err3.Error())
    }
     
    return shim.Success([]byte("###################### Init Successfull ######################"))
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
//###############           FUNCTION Invoke (For Invokation of data inside Ledger)   ############### 
func (t *WHOCountryHospitalChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

    function, args := stub.GetFunctionAndParameters()

    if len(args) == 0 {
         return shim.Error("@@@@@@@@@@@@@@@@@@@ Expecting Arguments @@@@@@@@@@@@@@@@@@@")
    }
    
    if function == "addHospital" {
         return t.addHospital(stub, args)
    } else if function == "updateHospital" {
         return t.updateHospital(stub, args)
    } else if function == "queryHospital" {
         return t.queryHospital(stub, args)
    } else if function == "addDonorOrRecepient" {
         return t.addDonorOrRecepient(stub, args)
    } else if function == "updateDonorOrRecepient" {
         return t.updateDonorOrRecepient(stub, args)
    } else if function == "addTransplant" {
         return t.addTransplant(stub, args)
    } else if function == "updateTransplant" {
         return t.updateTransplant(stub, args)
    } else if function == "queryTransplant" {
         return t.queryTransplant(stub, args)
    } else if function == "queryDonor" {
         return t.queryDonor(stub, args)
    } else if function == "queryRecepient" {
         return t.queryRecepient(stub, args)
    } else if function == "queryHospitalById" {
         return t.queryHospitalById(stub, args)
    } else if function == "queryDonorById" {
         return t.queryDonorById(stub, args)
    } else if function == "queryRecepientById" {
         return t.queryRecepientById(stub, args)
    } else if function == "queryTransplantById" {
         return t.queryTransplantById(stub, args)
    } else if function == "queryCount" {
         return t.queryCount(stub, args)
    } else if function == "queryCountryCount" {
         return t.queryCountryCount(stub)
    } else if function == "getHospitalName" {
         return t.getHospitalName(stub, args)
    } else if function == "callOtherChaincodeInvoke" {
         return t.callOtherChaincodeInvoke(stub, args)
    } else if function == "callOtherChaincodeQuery" {
         return t.callOtherChaincodeQuery(stub, args)
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
	err := shim.Start(new(WHOCountryHospitalChainCode))
	if err != nil {
		fmt.Printf("@@@@@@@@@@@@@@@@@@@  Error creating new WHOCountryHospitalChainCode : %s  @@@@@@@@@@@@@@@@@@@", err)
	}
} 