
package main

import (
    
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)


//###############           FUNCTION addOrder (For adding details of replace Order request  in Ledger)   ###############
func addOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response { 

    shim.Success([]byte("######################  Entering in addOrder Method  ######################"))

    if len(args) != 4 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of the arguments,  Expecting 4 Arguments @@@@@@@@@@@@@@@@@@@")
    }

    var Orderid = args[0]
    Orderid = Orderid[7:len(Orderid)] 
        
    OrderInfo := `{ "ordId" : "` + Orderid  + `" , "custName" : "` + args[1] + `" , "ordDate" : "` + args[3] + `" , "recDate" : "NA" ,"status" : "0" ,
                    "manufacturer" : { "quantity" : "` + args[2] + `" , "drfc" : "` + args[3] + `" , "dnts":"NA" } ,
                    "shipper" : { "dcbs" : "NA" , "dntl" : "NA" } ,
                    "logistic" : { "dcbl" : "NA" , "ddtm" : "NA" }
                  }`
    
    err := stub.PutState(Orderid, []byte(OrderInfo))
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to add Order in Ledger  @@@@@@@@@@@@@@@@@@@")
    }
    
return shim.Success([]byte("######################  Add Order Successfully in Ledger  ######################"))

}

//###############           FUNCTION updateDate (For updating detials of Order in Ledger)   ###############
func updateDate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in updateDate Method  ######################"))

    var Orderid = args[0]
    var key = args[1]
    var value = args[2]
    var err error 
    
    if len(args) != 3{
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of arguments, Expecting 3 Arguments  @@@@@@@@@@@@@@@@@@@")
    }

    orderBytes, err := stub.GetState(Orderid)
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ Order doesnot exists @@@@@@@@@@@@@@@@@@@")                        
    }

    var raw map[string]interface{}
    err =  json.Unmarshal(orderBytes, &raw)
    if err != nil {
		return shim.Error("@@@@@@@@@@@@@@@@@@@  Could not unmarshal Order with ID  @@@@@@@@@@@@@@@@@@@")
	}

    if key == "dnts"{
            raw["manufacturer"].(map[string]interface{})["dnts"] = value
            raw["status"] = 1
            OrdBytes,_ := json.Marshal(raw)
            err = stub.PutState(Orderid, []byte(OrdBytes))
            if err != nil{
                return shim.Error(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@  Failed to update the value dnts (Date of notification to Shipper)  @@@@@@@@@@@@@@@@@@@"))
            }
    } else if key == "dcbs"{
            raw["shipper"].(map[string]interface{})["dcbs"] = value
            raw["status"] = 2
            OrdBytes,_ := json.Marshal(raw)
            err = stub.PutState(Orderid, []byte(OrdBytes))
            if err != nil{
                return shim.Error(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@  Failed to update the value dcbs  (date of Confirmation by logistic to manufacturer)  @@@@@@@@@@@@@@@@@@@"))
            }
    } else if key == "dntl"{
            raw["shipper"].(map[string]interface{})["dntl"]=value
            raw["status"] = 3
            OrdBytes,_ := json.Marshal(raw)
            err = stub.PutState(Orderid, []byte(OrdBytes))
            if err != nil{
                return shim.Error(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@  Failed to update the value dntl (date of notification to logistic) @@@@@@@@@@@@@@@@@@@"))
            }
    } else if key == "dcbl"{
            raw["logistic"].(map[string]interface{})["dcbl"]=value
            raw["status"] = 4
            OrdBytes,_ := json.Marshal(raw)
            err = stub.PutState(Orderid, []byte(OrdBytes))
            if err != nil{
                return shim.Error(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@  Failed to update the value dcbl (date of Confirmation by logistic to Shipper)  @@@@@@@@@@@@@@@@@@@"))
            }
    } else if key == "ddtc"{
            raw["logistic"].(map[string]interface{})["ddtc"]=value
            raw["status"] = 5
            OrdBytes,_ := json.Marshal(raw)
            err = stub.PutState(Orderid, []byte(OrdBytes))
            if err != nil{
                return shim.Error(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@  Failed to update the value ddtm (date of dispatch to client)  @@@@@@@@@@@@@@@@@@@"))
            }
    } else if key == "recDate"{
            raw["recDate"]=value
            raw["status"] = 6
            OrdBytes,_ := json.Marshal(raw)
            err = stub.PutState(Orderid, []byte(OrdBytes))
            if err != nil{
                return shim.Error(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@  Failed to update the value ddtm (date of recieve of product back by client)  @@@@@@@@@@@@@@@@@@@"))
            }
    } else {
            return shim.Error(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@  Invalid Field Name for Updation  @@@@@@@@@@@@@@@@@@@"))
    }
    
return shim.Success([]byte("###################### update date Successfully  ######################"))

}