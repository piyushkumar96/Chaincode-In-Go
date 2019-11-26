
package main

import (
    
    "encoding/json"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    "fmt"
    pb "github.com/hyperledger/fabric/protos/peer"
)


//###############           FUNCTION createCocoaBeanBag    ###############
func (r *ChocoChaincode) createCocoaBeanBag(stub shim.ChaincodeStubInterface, args []string) pb.Response { 

    shim.Success([]byte("######################  Entering in createCocoaBeanBag Method  ######################"))

    if len(args) != 3 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of the arguments,  Expecting 3 Arguments @@@@@@@@@@@@@@@@@@@")
    }

    var beanBagId = args[0]
    var country   = args[1]
    var farmerId  = args[2]

    // checking that Cocoa Bean bag exists with given Id or Not
    cbbBytes, err := stub.GetState(beanBagId)
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ Failed to get Cocoa Bean Bag with this Id " + err.Error() + " @@@@@@@@@@@@@@@@@@@")                        
    }else if cbbBytes != nil {
		return shim.Error("@@@@@@@@@@@@@@@@@@@ Cocoa Bean Bag already exists with this Id  @@@@@@@@@@@@@@@@@@@")
	}
    
    fmt.Println("$$$$$$$$$$$$$$$$ start creating Cocoa Bean Bag with beanBagId %s $$$$$$$$$$$$$$$$", beanBagId)

    CBeanBagInfo := `{ "beanBagId" : "` + beanBagId  + `" ,
                  "country" : "` + country  + `" ,
                  "farmerId" : "` + farmerId  + `" ,
                  "isConsumed" : false
                  }`
    
    // writing Cocoa Bean Bag to the ledger
    err = stub.PutState(beanBagId, []byte(CBeanBagInfo))
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to Cocoa Bean Bag in Ledger  @@@@@@@@@@@@@@@@@@@")
    }
    
    fmt.Println("$$$$$$$$$$$$$$$$ creating Cocoa Bean Bag with beanBagId %s $$$$$$$$$$$$$$$$", beanBagId)

    return shim.Success([]byte("######################  Cocoa Bean Bag Created Successfully in Ledger  ######################"))
}

//###############           FUNCTION createChocolateBar   ###############
func (r *ChocoChaincode) createChocolateBar(stub shim.ChaincodeStubInterface, args []string) pb.Response { 

    shim.Success([]byte("######################  Entering in createChocolateBar Method  ######################"))

    if len(args) != 2 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of the arguments,  Expecting 2 Arguments @@@@@@@@@@@@@@@@@@@")
    }

    var barId = args[0]
    var beanBagId   = args[1]

    // checking that Cocoa Bean bag exists with given Id or Not
    cbBytes, err := stub.GetState(barId)
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ Chocolate bar already exists with this Id " + err.Error() + "  @@@@@@@@@@@@@@@@@@@")                        
    }else if cbBytes != nil {
		return shim.Error("@@@@@@@@@@@@@@@@@@@ Chocolate bar already exists with this Id  @@@@@@@@@@@@@@@@@@@")
	}

    fmt.Println("$$$$$$$$$$$$$$$$ start creating Chocolate Bar with barId %s $$$$$$$$$$$$$$$$", barId)

    cbbBytes, err1 := stub.GetState(beanBagId)                  
    if err1 != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ Failed to get Chocolate Bar with this Id " + err1.Error() + " @@@@@@@@@@@@@@@@@@@")         
    }else if cbbBytes == nil {
		return shim.Error("@@@@@@@@@@@@@@@@@@@ Chocolate Bar doesnot exists with this Id  @@@@@@@@@@@@@@@@@@@")
	}

    // create an Object of struct CocoaBeanBag
    cbbData := CocoaBeanBag{}
    
    //unmarshal it aka JSON.parse()
	err = json.Unmarshal(cbbBytes, &cbbData)                    
	if err != nil {
		return shim.Error("@@@@@@@@@@@@@@@@@@@  Could not unmarshal CBB with given beanBagId  " + err.Error() + "  @@@@@@@@@@@@@@@@@@@")
	}
    
    //Updating the isConsumed for Cocoa Bean Bag
    cbbData.IsConsumed = true                                   
    
    //rewrite the Cocoa Bean Bag
    cbbJSONasBytes, _ := json.Marshal(cbbData)
	err = stub.PutState(beanBagId, cbbJSONasBytes)             
	if err != nil {
		return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to update Cocoa Bean Bag in Ledger " + err.Error() + " @@@@@@@@@@@@@@@@@@@")
	}

     //writing the Choco Bar to the ledger
    err = stub.PutState(barId, []byte(beanBagId))             
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to Create Chocolate Bar in Ledger  @@@@@@@@@@@@@@@@@@@")
    }
    
    fmt.Println("$$$$$$$$$$$$$$$$ end of creating Chocolate Bar with barId %s $$$$$$$$$$$$$$$$", barId)

    return shim.Success([]byte("######################  Chocolate Bar Created Successfully in Ledger  ######################"))
}

