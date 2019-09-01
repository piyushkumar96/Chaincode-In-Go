
package main

import (
    
    "encoding/json"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)


//###############           FUNCTION createTicket    ###############
func (r *RMAChaincode) createTicket(stub shim.ChaincodeStubInterface, args []string) pb.Response { 

    shim.Success([]byte("######################  Entering in createTicket Method  ######################"))

    if len(args) != 1 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of the arguments,  Expecting 1 Arguments @@@@@@@@@@@@@@@@@@@")
    }

    var rmaNo = args[0]

    _, err := stub.GetState(rmaNo)
    if err == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ Ticket No already exists @@@@@@@@@@@@@@@@@@@")                        
    }
        
    RMAInfo := `{ "rmaNo" : "` + rmaNo  + `" ,
                    "messages" : [
                                    { "type" : "3b11" , "message" : "" , "receiveddate":"" } ,
                                    { "type" : "3b3sir" , "message" : "" , "receiveddate":"" },
                                    { "type" : "3b3eta" , "message" : "" , "receiveddate":"" },
                                    { "type" : "3b13" , "message" : "" , "receiveddate":"" },
                                    { "type" : "3b3pod" , "message" : "" , "receiveddate":"" }
                                ]
                  }`
    
    err = stub.PutState(rmaNo, []byte(RMAInfo))
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to Create Ticket in Ledger  @@@@@@@@@@@@@@@@@@@")
    }
    
return shim.Success([]byte("######################  Ticket Created Successfully in Ledger  ######################"))

}

//###############           FUNCTION updates3b11 (For updating details of RMA ticket in Ledger)   ###############
func (r *RMAChaincode) updates3b11(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in updates3b11 Method  ######################"))

    if len(args) != 3{
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of arguments, Expecting 3 Arguments  @@@@@@@@@@@@@@@@@@@")
    }

    var rmaNo = args[0]
    var message = args[1]
    var receiveddate = args[2]

    var err error 
    

    rmaBytes, err := stub.GetState(rmaNo)
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ RMA Ticket Number doesnot exists @@@@@@@@@@@@@@@@@@@")                        
    }

    var raw map[string]interface{}
    err =  json.Unmarshal(rmaBytes, &raw)
    if err != nil {
		return shim.Error("@@@@@@@@@@@@@@@@@@@  Could not unmarshal RMA with given number  @@@@@@@@@@@@@@@@@@@")
	}
            messages := raw["messages"]

            //return shim.Success([]byte(messages))
           for i := 0; i < len(messages.([]interface{})); i++ {
                if (messages.([]interface{})[i]).(map[string]interface {})["type"] == "3b11" {
                    (messages.([]interface{})[i]).(map[string]interface {})["message"] = message
                    (messages.([]interface{})[i]).(map[string]interface {})["receiveddate"] = receiveddate
                }
            }
            raw["messages"] = messages
            RMABytes,_ := json.Marshal(raw)
            err = stub.PutState(rmaNo, []byte(RMABytes))
            if err != nil{
                return shim.Error("@@@@@@@@@@@@@@@@@@@  Failed to update the message and receivedate  @@@@@@@@@@@@@@@@@@@")
            } 
    
return shim.Success([]byte("###################### update Successfully  ######################"))

}

//###############           FUNCTION updates3bsir (For updating details of RMA ticket in Ledger)   ###############
func (r *RMAChaincode) updates3b3sir(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in updates3b3sir Method  ######################"))

    if len(args) != 3{
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of arguments, Expecting 3 Arguments  @@@@@@@@@@@@@@@@@@@")
    }

    var rmaNo = args[0]
    var message = args[1]
    var receiveddate = args[2]

    var err error 
    

    rmaBytes, err := stub.GetState(rmaNo)
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ RMA Ticket Number doesnot exists @@@@@@@@@@@@@@@@@@@")                        
    }

    var raw map[string]interface{}
    err =  json.Unmarshal(rmaBytes, &raw)
    if err != nil {
		return shim.Error("@@@@@@@@@@@@@@@@@@@  Could not unmarshal RMA with given number  @@@@@@@@@@@@@@@@@@@")
	}
           messages := raw["messages"]
            
           for i := 0; i < len(messages.([]interface{})); i++ {
                if (messages.([]interface{})[i]).(map[string]interface {})["type"] == "3b3sir" {
                    (messages.([]interface{})[i]).(map[string]interface {})["message"] = message
                    (messages.([]interface{})[i]).(map[string]interface {})["receiveddate"] = receiveddate
                }
            }
            
            raw["messages"] = messages
            RMABytes,_ := json.Marshal(raw)
            err = stub.PutState(rmaNo, []byte(RMABytes))
            if err != nil{
                return shim.Error("@@@@@@@@@@@@@@@@@@@  Failed to update the message and receivedate  @@@@@@@@@@@@@@@@@@@")
            }
    
return shim.Success([]byte("###################### update Successfully  ######################"))

}

//###############           FUNCTION updates3beta (For updating details of RMA ticket in Ledger)   ###############
func (r *RMAChaincode) updates3b3eta(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in updates3beta Method  ######################"))

    if len(args) != 3{
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of arguments, Expecting 3 Arguments  @@@@@@@@@@@@@@@@@@@")
    }

    var rmaNo = args[0]
    var message = args[1]
    var receiveddate = args[2]

    var err error 
    

    rmaBytes, err := stub.GetState(rmaNo)
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ RMA Ticket Number doesnot exists @@@@@@@@@@@@@@@@@@@")                        
    }

    var raw map[string]interface{}
    err =  json.Unmarshal(rmaBytes, &raw)
    if err != nil {
		return shim.Error("@@@@@@@@@@@@@@@@@@@  Could not unmarshal RMA with given number  @@@@@@@@@@@@@@@@@@@")
	}
            messages := raw["messages"]
            for i := 0; i < len(messages.([]interface{})); i++ {
                if (messages.([]interface{})[i]).(map[string]interface {})["type"] == "3b3eta" {
                    (messages.([]interface{})[i]).(map[string]interface {})["message"] = message
                    (messages.([]interface{})[i]).(map[string]interface {})["receiveddate"] = receiveddate
                }
            }
        
            raw["messages"] = messages
            RMABytes,_ := json.Marshal(raw)
            err = stub.PutState(rmaNo, []byte(RMABytes))
            if err != nil{
                return shim.Error("@@@@@@@@@@@@@@@@@@@  Failed to update the message and receivedate  @@@@@@@@@@@@@@@@@@@")
            }
    
return shim.Success([]byte("###################### update Successfully  ######################"))

}

//###############           FUNCTION updates3b13 (For updating details of RMA ticket in Ledger)   ###############
func (r *RMAChaincode) updates3b13(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in updates3b13 Method  ######################"))

    if len(args) != 3{
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of arguments, Expecting 3 Arguments  @@@@@@@@@@@@@@@@@@@")
    }

    var rmaNo = args[0]
    var message = args[1]
    var receiveddate = args[2]

    var err error 
    

    rmaBytes, err := stub.GetState(rmaNo)
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ RMA Ticket Number doesnot exists @@@@@@@@@@@@@@@@@@@")                        
    }

    var raw map[string]interface{}
    err =  json.Unmarshal(rmaBytes, &raw)
    if err != nil {
		return shim.Error("@@@@@@@@@@@@@@@@@@@  Could not unmarshal RMA with given number  @@@@@@@@@@@@@@@@@@@")
	}
            messages := raw["messages"]
            for i := 0; i < len(messages.([]interface{})); i++ {
                if (messages.([]interface{})[i]).(map[string]interface {})["type"] == "3b13" {
                    (messages.([]interface{})[i]).(map[string]interface {})["message"] = message
                    (messages.([]interface{})[i]).(map[string]interface {})["receiveddate"] = receiveddate
                }
            }
            
            raw["messages"] = messages
            RMABytes,_ := json.Marshal(raw)
            err = stub.PutState(rmaNo, []byte(RMABytes))
            if err != nil{
                return shim.Error("@@@@@@@@@@@@@@@@@@@  Failed to update the message and receivedate  @@@@@@@@@@@@@@@@@@@")
            }
    
return shim.Success([]byte("###################### update Successfully  ######################"))

}

//###############           FUNCTION updates3b3pod (For updating details of RMA ticket in Ledger)   ###############
func (r *RMAChaincode) updates3b3pod(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in updates3b3pod Method  ######################"))

    if len(args) != 3{
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of arguments, Expecting 3 Arguments  @@@@@@@@@@@@@@@@@@@")
    }

    var rmaNo = args[0]
    var message = args[1]
    var receiveddate = args[2]

    var err error 
    

    rmaBytes, err := stub.GetState(rmaNo)
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ RMA Ticket Number doesnot exists @@@@@@@@@@@@@@@@@@@")                        
    }

    var raw map[string]interface{}
    err =  json.Unmarshal(rmaBytes, &raw)
    if err != nil {
		return shim.Error("@@@@@@@@@@@@@@@@@@@  Could not unmarshal RMA with given number  @@@@@@@@@@@@@@@@@@@")
	}
            messages := raw["messages"]

            for i := 0; i < len(messages.([]interface{})); i++ {
                if (messages.([]interface{})[i]).(map[string]interface {})["type"] == "3b3pod" {
                    (messages.([]interface{})[i]).(map[string]interface {})["message"] = message
                    (messages.([]interface{})[i]).(map[string]interface {})["receiveddate"] = receiveddate
                }
            }

            raw["messages"] = messages
            RMABytes,_ := json.Marshal(raw)
            err = stub.PutState(rmaNo, []byte(RMABytes))
            if err != nil{
                return shim.Error("@@@@@@@@@@@@@@@@@@@  Failed to update the message and receivedate  @@@@@@@@@@@@@@@@@@@")
            } 
    
return shim.Success([]byte("###################### update Successfully  ######################"))

}
