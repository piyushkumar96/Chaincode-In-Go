package main

import (
    "encoding/json"
    "fmt"
    "strconv"
    "bytes"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

type NewChaincode struct {   
}

                                // Global Variable to store No. Of Employees in company

type Employee struct {
        empid string `json:"empid"`
        firstname string `json:"firstname"`
        lastname string `json:"lastname"`
        age string `json:"age"`
        designation string `json:"designation"`
        amount string `json:"amount"`
        mobile string `json:"mobile"`
        email string `json:"email"`
}

//###############           FUNCTION Init (For Initiation of chaincode)   ############### 
func (s *NewChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

    _, args := stub.GetFunctionAndParameters()

    err2 := stub.PutState(args[0],[]byte(args[1]))
    if err2 != nil {
        return shim.Error(err2.Error())
    }

    return shim.Success([]byte("###################### Successfully Init of Ledger  ######################"))
}

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

//###############           FUNCTION noofemployees (For getting No. of Employees  in Ledger)   ###############
func (s *NewChaincode) noofemployees(stub shim.ChaincodeStubInterface) pb.Response {
    
    noofempbytes, err := stub.GetState("noofemployees")
    if err != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@ Failed to get NoOfEmployees Value from Ledger @@@@@@@@@@@@@@@@@@@")
    }

    return shim.Success(noofempbytes)
}

//###############           FUNCTION addMember (For adding details of Employee  in Ledger)   ###############
func (s *NewChaincode) addMember(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    if len(args) != 8 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ Incorrect no. of the arguments @@@@@@@@@@@@@@@@@@@")
    }
    
    noofempbytes, err1 := stub.GetState("noofemployees")
    if err1 != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@ Failed to get NoOfEmployees Value from Ledger @@@@@@@@@@@@@@@@@@@")
    }

    NoOfEmployees, _ := strconv.Atoi(string(noofempbytes))
    NoOfEmployees++;

    var employee = Employee{empid: args[0], firstname: args[1], lastname: args[2], age: args[3], designation: args[4], amount: args[5], mobile: args[6], email: args[7]}
        
    employeebytes,_ := json.Marshal(employee)
        
    err2 := stub.PutState(string(NoOfEmployees),employeebytes)
    if err2 != nil {
        return shim.Error(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@ Failed to record employee Details: %s @@@@@@@@@@@@@@@@@@@", args[0]))
    }
  
    
    err3 := stub.PutState(noofemployees,[]byte(strconv.Itoa(NoOfEmployees)))
    if err3 != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ Failed to Increment the value of NoOfEmployees by 1 as New Member Joined The Company @@@@@@@@@@@@@@@@@@@")
    }

return shim.Success([]byte("######################  Add Employee Successfully in Ledger  ######################"))
}


//###############           FUNCTION queryDetailsWithId  (For getting detials of single Employee using Id from Ledger) ###############
func (s *NewChaincode) queryDetailsWithId(stub shim.ChaincodeStubInterface, args string) ( []byte, bool ) {
    
    if len(args) != 1 {
        return []byte("@@@@@@@@@@@@@@@@@@@ Incorrect no. of arguments Expecting 1 argument @@@@@@@@@@@@@@@@@@@") , false
    }

    startKey := "1"
    noofempbytes, err1 := stub.GetState("noofemployees")
    if err1 != nil {
            return []byte("@@@@@@@@@@@@@@@@@@@ Failed to get NoOfEmployees Value from Ledger @@@@@@@@@@@@@@@@@@@") , false
    }

    endKey := string(noofempbytes)
    
    resultsIterator, err2 := stub.GetStateByRange(startKey, endKey)
    if err2 != nil {
	return []byte("@@@@@@@@@@@@@@@@@@@ Failed to Iterate @@@@@@@@@@@@@@@@@@@") , false
    }
    defer resultsIterator.Close()
    
    var buffer bytes.Buffer

    for resultsIterator.HasNext() {
	queryResponse, err3 := resultsIterator.Next()
	if err3 != nil {
		return []byte("@@@@@@@@@@@@@@@@@@@ Failed to Iterate @@@@@@@@@@@@@@@@@@@") , false
	}
        
        buffer.WriteString(queryResponse.Key)
    
        empBytes, err4 := stub.GetState(string(buffer.Bytes()))
        if err4 != nil {
        return []byte("@@@@@@@@@@@@@@@@@@@ Employee doesnot exists @@@@@@@@@@@@@@@@@@@") , false                       
        }

        emp := Employee{}
        err5 :=  json.Unmarshal(empBytes, &emp)
        if err5 != nil {
		     return []byte(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@  Could not unmarshal Employee with ID  %s @@@@@@@@@@@@@@@@@@@", args[0])) , false
	    }

        if emp.empid == args[0]  {
             return empBytes , true
        }
         buffer.Reset()  
      }
     return []byte(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@  Unable to get Employee with ID  %s @@@@@@@@@@@@@@@@@@@", args[0])) , false
}

//###############           FUNCTION transferMoney (For transfering Money from one Employee's Acc to Another)   ###############
func (s *NewChaincode) transferMoney(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var employeeid1, employeeid2 string 
    var balanceid1,balanceid2 int
    var transferamt int 
    var err error
    
    if len(args) != 3 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ Incorrect No. of Arguments Expecting 3 @@@@@@@@@@@@@@@@@@@")
    } 

    employeeid1 = args[0]
    employeeid2 = args[1]

    empid1bytes, bool1 := s.queryDetailsWithId(stub, employeeid1)
    //empid1bytes, err := stub.GetState(employeeid1)
    if bool1 == false {
            return shim.Error(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@ Failed to get Employee from Ledger with Id %s @@@@@@@@@@@@@@@@@@@", args[0]))
    }
    if empid1bytes == nil {
            return shim.Error(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@ Employee not Found in Ledger with ID %s @@@@@@@@@@@@@@@@@@@", args[0]))
    }
    emp1 := Employee{}
    err =  json.Unmarshal(empid1bytes, &emp1)
    if err != nil {
		return shim.Error(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@  Could not unmarshal Employee with ID  %s @@@@@@@@@@@@@@@@@@@", args[0]))
    }
    balanceid1,_ = strconv.Atoi(string(emp1.amount))

    empid2bytes, bool2 := s.queryDetailsWithId(stub, employeeid2)
    //empid2bytes, err := stub.GetState(employeeid2)
    if bool2 == false {
            return shim.Error(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@ Failed to get Employee from Ledger with Id %s @@@@@@@@@@@@@@@@@@@", args[1]))
    }
    if empid2bytes == nil {
            return shim.Error(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@ Employee not Found in Ledger with ID %s @@@@@@@@@@@@@@@@@@@", args[1]))
    }
    emp2 := Employee{}
    err =  json.Unmarshal(empid2bytes, &emp2)
    if err != nil {
		return shim.Error(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@  Could not unmarshal Employee with ID  %s @@@@@@@@@@@@@@@@@@@", args[1]))
    }
    balanceid2,_ = strconv.Atoi(string(emp2.amount))

    transferamt,err = strconv.Atoi(string(args[2]))
    if err != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@ Invalid Transaction amount, expecting a integer value @@@@@@@@@@@@@@@@@@@")
    }
    balanceid1 = balanceid1 - transferamt
    balanceid2 = balanceid2 + transferamt

    fmt.Printf("Employeeid1 balance = %d, Employeeid2 balance = %d \n",balanceid1,balanceid2)

    err = stub.PutState(employeeid1, []byte(strconv.Itoa(balanceid1)))
    if err != nil {
            return shim.Error(err.Error())
    }
    err = stub.PutState(employeeid2, []byte(strconv.Itoa(balanceid2)))
    if err != nil {
            return shim.Error(err.Error())
    }

return shim.Success([]byte("###################### Successfully Tranfer Money  ######################"))
}

func (s *NewChaincode) changeDetail(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    employeeid := args[0]
    key := args[1]
    value := args[2]

    if len(args) != 3{
        return shim.Error("@@@@@@@@@@@@@@@@@@@ Incorrect no. of arguments, required 3 args @@@@@@@@@@@@@@@@@@@")
    }

    empBytes, err1 := stub.GetState(employeeid)

    if err1 != nil {
        return shim.Error(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@ Employee doesnot exits with ID %s @@@@@@@@@@@@@@@@@@@", args[0]))
    }

    emp := Employee{}
    err2 := json.Unmarshal(empBytes, &emp)
    if err2 != nil {
		return shim.Error(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@  Could not unmarshal Employee with ID  %s @@@@@@@@@@@@@@@@@@@", args[1]))
    }
  
    if key == "empid" {
            emp.empid=value
    } else if key == "firstname" {
            emp.firstname=value
    } else if key == "lastname" {
            emp.lastname=value
    } else if key == "age" {
            emp.age=value
    } else if key == "designation" {
            emp.designation=value
    } else if key == "amount" {
            emp.amount=value
    } else if key == "email" {
            emp.email=value
    } else if key == "mobile" {
            emp.mobile=value
    }

    empBytes,_ = json.Marshal(emp)
    err3 := stub.PutState(employeeid, empBytes)
    if err3 != nil{
        return shim.Error(fmt.Sprintf("@@@@@@@@@@@@@@@@@@@ Failed to update the value @@@@@@@@@@@@@@@@@@@"))
    }

  return shim.Success([]byte("###################### Successfully Update Details  ######################"))
}


//###############           FUNCTION queryAllEmpInCompany (For getting detials of all Employees in company from Ledger)   ###############
func (s *NewChaincode) queryAllEmpInCompany(stub shim.ChaincodeStubInterface) pb.Response {

	startKey := "1"
        noofemployees := "noofemployees" 
	noofempbytes, err1 := stub.GetState(noofemployees)
        if err1 != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@ Failed to get NoOfEmployees Value from Ledger @@@@@@@@@@@@@@@@@@@") 
        }
        endKey := string(noofempbytes)

	resultsIterator, err2 := stub.GetStateByRange(startKey, endKey)
	if err2 != nil {
		return shim.Error(err2.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err3 := resultsIterator.Next()
		if err3 != nil {
			return shim.Error(err3.Error())
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

	fmt.Printf("- queryAllEmpInCompany:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *NewChaincode) delEmployee(stub shim.ChaincodeStubInterface, args []string) pb.Response { 

    if len(args) != 1  {
        return shim.Error("Incorrect no. of arguments Expecting 1 argument")
    }

    var employeeid = args[0]

    err := stub.DelState(employeeid)
    
    if err != nil {
        return shim.Error(" Unable to delete Employee Details")
    }

    return shim.Success(nil)

}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(NewChaincode))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

