
package main

import (
    
    "encoding/json"
    "fmt"
    "bytes"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

//Chaincode for Tracking of return of product 
type TrackChaincode struct {   
}
// Order Structure 
type Order struct {
        ordId                   string              `json:"ordId"`             // Unique, In our case it starts From 111
        custName                string              `json:"custName"`
        ordDate                 string              `json:"ordDate"`           // date of order by client 
	    recDate			        string		        `json:"recDate"`           // item recieved by client
        status                  string              `json:"status"`            // there are 7 status  values from  [0-6] 
        manufacturer            Manufacturer        `json:"manufacturer"`      // Object of Manufacturer
        shipper                 Shipper             `json:"shipper"`           // Object of Shipper
        logistic                Logistic            `json:"logistic"`          // Object of Logistic
}

// Manufacturer Structure 
type Manufacturer struct {
        quantity                string              `json:"quantity"`          //Quantity of product 
        drfc                    string              `json:"drfc"`              //Date of replace request from client
        dnts                    string              `json:"dnts"`              //Date of notification to Shipper
}

// Shipper Structure                                                           // Shipper means Partner
type Shipper struct {
        dcbs                    string              `json:"dcbs"`              // date of Confirmation by logistic to manufacturer
        dntl                    string              `json:"dntl"`              // date of notification to logistic  
}

// Logistic Structure
type Logistic struct {                                                         // Logistic means the carrier
        
        dcbl                    string              `json:"dcbl"`              // date of Confirmation by logistic to Shipper
        ddtc                    string              `json:"ddtc"`              // date of dispatch to manufacturer
}

//###############           FUNCTION Init (For Initiation of chaincode)   ############### 
func (t *TrackChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

            return shim.Success([]byte("###################### Init Successfull ######################"))
}

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

//###############           FUNCTION addOrder (For adding details of replace Order request  in Ledger)   ###############
func (t *TrackChaincode) addOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response { 

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
func (t *TrackChaincode) updateDate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
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

//###############           FUNCTION queryDetails  (For getting detials of single Order in Ledger) ###############
func (t *TrackChaincode) queryDetails(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    if len(args) != 1{
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of arguments Expecting 1 argument  @@@@@@@@@@@@@@@@@@@")
    }

    var Orderid = args[0]
    Orderid = Orderid[4:len(Orderid)]
    
    ordBytes, _ := stub.GetState(Orderid)
    if ordBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate Order  @@@@@@@@@@@@@@@@@@@")
    }
    str := "["+string(ordBytes)+"]"
       
return shim.Success([]byte(str))

}

//###############           FUNCTION queryAllOrder (For getting detials of all Order in Ledger)   ###############
func (s *TrackChaincode) queryAllOrder(stub shim.ChaincodeStubInterface) pb.Response {

	startKey := "111"
	endKey := "9999"

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add comma before array members,suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllOrder:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

//###############           FUNCTION main   ############### 
func main() {
	// Create a new Smart Contract
	err := shim.Start(new(TrackChaincode))
	if err != nil {
		fmt.Printf("@@@@@@@@@@@@@@@@@@@  Error creating new TrackChaincode : %s  @@@@@@@@@@@@@@@@@@@", err)
	}
} 


