package main

import (
    
    "encoding/json"
    "strconv"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)



//###############           FUNCTION queryHospital (For querying Details of Hospital from Ledger)   ###############
func queryHospital(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in queryHospital Method  ######################"))

    if len(args) != 1 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of the arguments,  Expecting 1 Arguments @@@@@@@@@@@@@@@@@@@")
    }
    var countryid = args[0]

countBytes, err := stub.GetState("count")
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ count doesnot exists @@@@@@@@@@@@@@@@@@@")                        
    }
    count1, _ := strconv.Atoi(string(countBytes))
	Key := countryid+"-H"
	
    str := "[ "

    var i int
    b := false 
    for i=1; i <= count1; i++ {
        hospitalBytes, _ := stub.GetState(Key+strconv.Itoa(i))
        if hospitalBytes == nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate Hospital  @@@@@@@@@@@@@@@@@@@")
        }
        
        if b == true {
            str = str + ","
        }
        str = str + string(hospitalBytes) 
        b = true 
    } 

    str = str + " ]"
	return shim.Success([]byte(str))    
}

//###############           FUNCTION queryTransplant (For querying Details of Transplant from Ledger)   ###############
func queryTransplant(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in queryTransplant Method  ######################"))

    if len(args) != 1 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of the arguments,  Expecting 1 Arguments @@@@@@@@@@@@@@@@@@@")
    }
    var hospitalId = args[0]

    countBytes, err := stub.GetState(hospitalId+"T")
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ hospital Transplant count doesnot exists @@@@@@@@@@@@@@@@@@@")                        
    }
    count1, _ := strconv.Atoi(string(countBytes))
	Key := hospitalId+"-T"
	
    str := "[ "

    var i int
    b := false 
    for i=1; i <= count1; i++ {
        transplantBytes, _ := stub.GetState(Key+strconv.Itoa(i))
        if transplantBytes == nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate Transplant  @@@@@@@@@@@@@@@@@@@")
        }
        
        if b == true {
            str = str + ","
        }
        str = str + string(transplantBytes) 
        b = true 
    } 

    str = str + " ]"
	return shim.Success([]byte(str))  
}


//###############           FUNCTION queryDonor (For querying Details of Donor from Ledger)   ###############
func queryDonor(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in queryDonor Method  ######################"))

    if len(args) != 1 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of the arguments,  Expecting 1 Arguments @@@@@@@@@@@@@@@@@@@")
    }
    var hospitalId = args[0]

    countBytes, err := stub.GetState(hospitalId+"D")
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ hospital Donor count doesnot exists @@@@@@@@@@@@@@@@@@@")                        
    }
    count1, _ := strconv.Atoi(string(countBytes))
	
	Key := hospitalId +"-D"
	
    str := "[ "

    var i int
    b := false 
    for i=1; i <= count1; i++ {
        donorBytes, _ := stub.GetState(Key+strconv.Itoa(i))
        if donorBytes == nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate Donor  @@@@@@@@@@@@@@@@@@@")
        }
        
        if b == true {
            str = str + ","
        }
        str = str + string(donorBytes) 
        b = true 
    } 

    str = str + " ]"
	return shim.Success([]byte(str))  
}

//###############           FUNCTION queryRecepient (For querying Details of Recepient from Ledger)   ###############
func queryRecepient(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in queryRecepient Method  ######################"))

    if len(args) != 1 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of the arguments,  Expecting 1 Arguments @@@@@@@@@@@@@@@@@@@")
    }
    var hospitalId = args[0]

    countBytes, err := stub.GetState(hospitalId+"R")
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ hospital Transplant count doesnot exists @@@@@@@@@@@@@@@@@@@")                        
    }
    count1, _ := strconv.Atoi(string(countBytes))
	Key := hospitalId +"-R"
	
    str := "[ "

    var i int
    b := false 
    for i=1; i <= count1; i++ {
        recepientBytes, _ := stub.GetState(Key+strconv.Itoa(i))
        if recepientBytes == nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate Recepient  @@@@@@@@@@@@@@@@@@@")
        }
        
        if b == true {
            str = str + ","
        }
        str = str + string(recepientBytes) 
        b = true 
    } 

    str = str + " ]"
	return shim.Success([]byte(str))
}

//###############           FUNCTION queryHospitalById  (For getting detials of single Hospital in Ledger) ###############
func queryHospitalById(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in queryHospitalById Method  ######################"))

    if len(args) != 1{
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of arguments Expecting 1 argument  @@@@@@@@@@@@@@@@@@@")
    }

    var hospitalId = args[0]
    hospitalBytes, _ := stub.GetState(hospitalId)
    if hospitalBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate Hospital  @@@@@@@@@@@@@@@@@@@")
    }
    str := "["+string(hospitalBytes)+"]"
       
return shim.Success([]byte(str))

}

//###############           FUNCTION queryDonorById  (For getting detials of single Donor in Ledger) ###############
func queryDonorById(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in queryDonorById Method  ######################"))

    if len(args) != 1{
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of arguments Expecting 1 argument  @@@@@@@@@@@@@@@@@@@")
    }

    var donorId = args[0]
    donorBytes, _ := stub.GetState(donorId)
    if donorBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate Donor  @@@@@@@@@@@@@@@@@@@")
    }
    str := "["+string(donorBytes)+"]"
       
return shim.Success([]byte(str))

}

//###############           FUNCTION queryRecepientById  (For getting detials of single Recepient in Ledger) ###############
func queryRecepientById(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in queryHospitalById Method  ######################"))

    if len(args) != 1{
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of arguments Expecting 1 argument  @@@@@@@@@@@@@@@@@@@")
    }

    var recepientId = args[0]
    recepientBytes, _ := stub.GetState(recepientId)
    if recepientBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate Recepient  @@@@@@@@@@@@@@@@@@@")
    }
    str := "["+string(recepientBytes)+"]"
       
return shim.Success([]byte(str))

}

//###############           FUNCTION queryTransplantById  (For getting detials of single Transplant in Ledger) ###############
func queryTransplantById(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in queryHospitalById Method  ######################"))

    if len(args) != 1{
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of arguments Expecting 1 argument  @@@@@@@@@@@@@@@@@@@")
    }

    var transplantId = args[0]
   transplantBytes, _ := stub.GetState(transplantId)
    if transplantBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate Transplant  @@@@@@@@@@@@@@@@@@@")
    }
    str := "["+string(transplantBytes)+"]"
       
return shim.Success([]byte(str))

}

//###############           FUNCTION queryCount  (For getting detials of single hospital specific count of Transplant, Donor, Recepient in Ledger) ###############
func queryCount(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in queryCount Method  ######################"))

    if len(args) != 1{
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of arguments Expecting 1 argument  @@@@@@@@@@@@@@@@@@@")
    }

    var hospitalId = args[0]

   countTransplantBytes, _ := stub.GetState(hospitalId+"T")
    if countTransplantBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to get hospital Transplant count  @@@@@@@@@@@@@@@@@@@")
    }
    // countTransplant1, _ := strconv.Atoi(string(countTransplantBytes))

    countRecepientBytes, _ := stub.GetState(hospitalId+"R")
    if countRecepientBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to get hospital Recepient count  @@@@@@@@@@@@@@@@@@@")
    }
   // countRecepient1, _ := strconv.Atoi(string(countRecepientBytes))

    countDonorBytes, _ := stub.GetState(hospitalId+"D")
    if countDonorBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to get hospital Donor count  @@@@@@@@@@@@@@@@@@@")
    }
   // countDonor1, _ := strconv.Atoi(string(countDonorBytes))

    str := `{"countT":"`+string(countTransplantBytes)+`","countR":"`+string(countRecepientBytes)+`","countD":"`+string(countDonorBytes)+`" }`
       
return shim.Success([]byte(str))

}

//###############           FUNCTION queryCountryCount  (For getting detials of single country specific count of Transplant, Donor, Recepient, Hospital in Ledger) ###############
func queryCountryCount(stub shim.ChaincodeStubInterface) pb.Response {
    
    shim.Success([]byte("######################  Entering in queryCountryCount Method  ######################"))

   countTransplantBytes, _ := stub.GetState("transplant")
    if countTransplantBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to get country Transplant count  @@@@@@@@@@@@@@@@@@@")
    }
    // countTransplant1, _ := strconv.Atoi(string(countTransplantBytes))

    countRecepientBytes, _ := stub.GetState("recepient")
    if countRecepientBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to get country Recepient count  @@@@@@@@@@@@@@@@@@@")
    }
   // countRecepient1, _ := strconv.Atoi(string(countRecepientBytes))

    countDonorBytes, _ := stub.GetState("donor")
    if countDonorBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to get country Donor count  @@@@@@@@@@@@@@@@@@@")
    }
   // countDonor1, _ := strconv.Atoi(string(countDonorBytes))

    countBytes, _ := stub.GetState("count")
    if countBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to get country Hospital count  @@@@@@@@@@@@@@@@@@@")
    }
   // count1, _ := strconv.Atoi(string(countBytes))

    str := `{"countT":"`+string(countTransplantBytes)+`","countR":"`+string(countRecepientBytes)+`","countD":"`+string(countDonorBytes)+`","count":"`+string(countBytes)+`" }`
       
return shim.Success([]byte(str))

}



//###############           FUNCTION totalTransplantCount (For getting total Transplant counts of a country from Ledger)   ###############
func totalTransplantCount(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in totalTransplantCount Method  ######################"))

    if len(args) != 1 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of the arguments,  Expecting 1 Arguments @@@@@@@@@@@@@@@@@@@")
    }
    var countryId = args[0]

    countBytes, err := stub.GetState("count")
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ hospital count doesnot exists @@@@@@@@@@@@@@@@@@@")                        
    }
    count1, _ := strconv.Atoi(string(countBytes))
	
    str := "[ "

    var i int
    b := false 
    for i=1; i <= count1; i++ {

        hospId := countryId + "-H"+strconv.Itoa(i);
        transId := countryId + "-H"+strconv.Itoa(i)+"T";
        transplantBytes, _ := stub.GetState(transId)
        if transplantBytes == nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate Transplant  @@@@@@@@@@@@@@@@@@@")
        }

        str1 :=`{ "id" : "` + hospId + `","countT" : "`+  string(transplantBytes) +`" }`
        if b == true {
            str = str + ","
        }
        
        str = str + str1
        b = true 
    } 

    str = str + " ]"
	return shim.Success([]byte(str))
}

//###############           FUNCTION getHospitalName  (For getting detials of single Hospital in Ledger) ###############
func getHospitalName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    shim.Success([]byte("######################  Entering in getHospitalName Method  ######################"))

    if len(args) != 1{
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of arguments Expecting 1 argument  @@@@@@@@@@@@@@@@@@@")
    }

    var hospitalId = args[0]
    hospitalBytes, _ := stub.GetState(hospitalId)
    if hospitalBytes == nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to locate Hospital  @@@@@@@@@@@@@@@@@@@")
    }
    var raw map[string]string
    err :=  json.Unmarshal(hospitalBytes, &raw)
    if err != nil {
		return shim.Error("@@@@@@@@@@@@@@@@@@@  Could not unmarshal Hospital with ID  @@@@@@@@@@@@@@@@@@@")
	}
    
 return shim.Success([]byte(raw["hospitalName"]))
}