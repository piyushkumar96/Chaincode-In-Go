
package main

import (
    
    "encoding/json"
    "fmt"
    "strconv"
    "strings"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

//Chaincode for Hospital
type WhoChaincode struct {   
}

func toChaincodeArgs(args ...string) [][]byte {
	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}
	return bargs
}


func (t *WhoChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

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

//###############           FUNCTION Invoke (For Invokation of data inside Ledger)   ############### 
func (t *WhoChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

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

//###############           FUNCTION addHospital (For adding Hospital in Ledger )   ###############
func (t *WhoChaincode) addHospital(stub shim.ChaincodeStubInterface, args []string) pb.Response { 

    shim.Success([]byte("######################  Entering in addHospital Method  ######################"))

    if len(args) != 9 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of the arguments,  Expecting 9 Arguments @@@@@@@@@@@@@@@@@@@")
    }
    
    countBytes, err := stub.GetState("count")
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@ count doesnot exists @@@@@@@@@@@@@@@@@@@")                        
    }
    count1, _ := strconv.Atoi(string(countBytes))
    count1++;
    
    transplant_count := "0";
	donor_count := "0";
	recepient_count := "0";

    countryId := args[8];
    hospitalId := countryId + "-H"+strconv.Itoa(count1);
    countryName := args[0];
    hospitalName := args[1];
    regSince := args[2];
    address := args[3];
    zipCode := args[4];
    city := args[5];
    state := args[6];
    phoneNo := args[7]    
    
    hospitalInfo := `{ "hospitalId" : "` + hospitalId  + `" , "countryName" : "` + countryName + `" , "hospitalName" : "` + hospitalName + `" , "regSince" : "` + regSince + `" , "address" : "` + address + `" , "zipCode" : "` + zipCode + `" , "city" : "` + city + `" , "state" : "` + state + `" , "phoneNo" : "` + phoneNo + `" }`
    
    err1 := stub.PutState(string(hospitalId), []byte(hospitalInfo))
    if err1 != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to add Hospital in Ledger  @@@@@@@@@@@@@@@@@@@")
    }
    err2 := stub.PutState("count", []byte(strconv.Itoa(count1)))
    if err2 != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to update count of Hospital in Ledger  @@@@@@@@@@@@@@@@@@@")
    }
    err3 := stub.PutState(hospitalId+"T", []byte(transplant_count))
    if err3 != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to add count of Transplant in Ledger  @@@@@@@@@@@@@@@@@@@")
    }
    err4 := stub.PutState(hospitalId+"D", []byte(donor_count))
    if err4 != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to add count of Donor in Ledger  @@@@@@@@@@@@@@@@@@@")
    }
    err5 := stub.PutState(hospitalId+"R", []byte(recepient_count))
    if err5 != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to add count of Recepient in Ledger  @@@@@@@@@@@@@@@@@@@")
    }
return shim.Success([]byte("######################  Add Hospital Successfully in Ledger  ######################"))

}

//###############           FUNCTION updateHospital (For updating Details of  Hospital in Ledger )   ###############
func (t *WhoChaincode) updateHospital(stub shim.ChaincodeStubInterface, args []string) pb.Response { 

    shim.Success([]byte("######################  Entering in updateHospital Method  ######################"))

    if len(args) != 9 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of the arguments,  Expecting 9 Arguments @@@@@@@@@@@@@@@@@@@")
    }
    
    hospitalId := args[8];
    countryName := args[0];
    hospitalName := args[1];
    regSince := args[2];
    address := args[3];
    zipCode := args[4];
    city := args[5];
    state := args[6];
    phoneNo := args[7]    
    
    hospitalInfo := `{ "hospitalId" : "` + hospitalId  + `" , "countryName" : "` + countryName + `" , "hospitalName" : "` + hospitalName + `" , "regSince" : "` + regSince + `" , "address" : "` + address +  `" , "zipCode" : "` + zipCode + `" , "city" : "` + city + `" , "state" : "` + state + `" , "phoneNo" : "` + phoneNo + `" }`
    
    err := stub.PutState(hospitalId, []byte(hospitalInfo))
    if err != nil {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to Update Hospital in Ledger  @@@@@@@@@@@@@@@@@@@")
    }
    
return shim.Success([]byte("######################  Update Hospital Successfully in Ledger  ######################"))

}

//###############           FUNCTION queryHospital (For querying Details of Hospital from Ledger)   ###############
func (t *WhoChaincode) queryHospital(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
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

//###############           FUNCTION addDonorOrRecepient (For adding Donor or Recepient in Ledger )   ###############
func (t *WhoChaincode) addDonorOrRecepient(stub shim.ChaincodeStubInterface, args []string) pb.Response { 

    shim.Success([]byte("######################  Entering in addDonorOrRecepient Method  ######################"))

    if len(args) != 31 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of the arguments,  Expecting 31 Arguments @@@@@@@@@@@@@@@@@@@")
    }
    
    //PersonalInfo 
        typeOfUser := args[0];
        firstName := args[1];
        middleName := args[2];
        lastName := args[3];
        relation := args[4];
        dateOfBirth := args[5];
        age := args[6];
        gender := args[7]; 
        personalMobileNo := args[8]
        bloodGroup := args[9]
        emailIdPersonal := args[10]
        //endPersonalinfo
//IdentificationInfo
        identificationType := args[11];
        identificationCard := args[12];
        //endIdentificationInfo
//AddressInfo
        addressLine1 := args[13];
        addressLine2 := args[14];
        district := args[15];
        state := args[16];
        country := args[17];
        pinCode := args[18];
        phoneNumber := args[19];
        //endAddressInfo
//emergencyContactInfo
        fullNameEmergencyConct := args[20];
        mobileNoEmergencyConct := args[21];
        emailIdEmergencyContct := args[22];
        relationEmergencyConct := args[23];
        fullAddressEmergencyConct := args[24];
        districtEmergencyConct := args[25];
        stateEmergencyConct := args[26];
        countryEmergencyConct := args[27];
        pincodeEmergencyConct := args[28];
        //endEmergencyContactlInfo
//organforDonationOrAcceptance
        organStatus := args[29];
        hospitalId := args[30];
        //endorganforDonationOrAcceptance   


    //adding Recepient
    if strings.ToLower(typeOfUser) == "recepient" {
        recepientBytes, err := stub.GetState("recepient")
        if err != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@ country recepient count doesnot exists @@@@@@@@@@@@@@@@@@@")                        
        }
        recepient1, _ := strconv.Atoi(string(recepientBytes))
        recepient1++;

        recepientHospitalBytes, err1 := stub.GetState(hospitalId+"R")
        if err1 != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@ hospital recepient count doesnot exists @@@@@@@@@@@@@@@@@@@")                        
        }
        recepientHospital1, _ := strconv.Atoi(string(recepientHospitalBytes))
        recepientHospital1++;
        recepientId := args[30]+"-R"+ strconv.Itoa(recepientHospital1)
        
        recepientInfo := `{ "recepientId" : "` + recepientId  + `" , "typeOfUser" : "` + typeOfUser + `" , "firstName" : "` + firstName + `" , "middleName" : "` + middleName + `" , "lastName" : "` + lastName + `" , "relation" : "` + relation + `" , "dateOfBirth" : "` + dateOfBirth + `" , "age" : "` + age + `" , "gender" : "` + gender + `" , "personalMobileNo" : "` + personalMobileNo + `" , "bloodGroup" : "` + bloodGroup + `" , "emailIdPersonal" : "` + emailIdPersonal + `" , "identificationType" : "` + identificationType + `" , "identificationCard" : "` + identificationCard + `" , "addressLine1" : "` + addressLine1 + `" , "addressLine2" : "` + addressLine2 + `" , "district" : "` + district + `" , "pinCode" : "` + pinCode +`" , "state" : "` + state +`" , "country" : "` + country + `" , "phoneNumber" : "` + phoneNumber + `" , "fullNameEmergencyConct" : "` + fullNameEmergencyConct + `" , "mobileNoEmergencyConct" : "` + mobileNoEmergencyConct + `" , "emailIdEmergencyContct" : "` + emailIdEmergencyContct + `" , "relationEmergencyConct" : "` + relationEmergencyConct + `" , "fullAddressEmergencyConct" : "` + fullAddressEmergencyConct + `" , "districtEmergencyConct" : "` + districtEmergencyConct + `" , "stateEmergencyConct" : "` + stateEmergencyConct + `" , "countryEmergencyConct" : "` + countryEmergencyConct + `" , "pincodeEmergencyConct" : "` + pincodeEmergencyConct + `" , "organStatus" : "` + organStatus + `" }`
    
        err2 := stub.PutState(recepientId   , []byte(recepientInfo))
        if err2 != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to add Recepient in Ledger  @@@@@@@@@@@@@@@@@@@")
        }
        err3 := stub.PutState("recepient", []byte(strconv.Itoa(recepient1)))
        if err3 != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to update country recepient count in Ledger  @@@@@@@@@@@@@@@@@@@")
        }
        err4 := stub.PutState(hospitalId+"R", []byte(strconv.Itoa(recepientHospital1)))
        if err4 != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to update hospital recepient count  in Ledger  @@@@@@@@@@@@@@@@@@@")
        }
    
      return shim.Success([]byte(recepientInfo))
    }   
    
    //adding Donor
    if strings.ToLower(typeOfUser) == "donor" {
        donorBytes, err := stub.GetState("donor")
        if err != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@ country donor count doesnot exists @@@@@@@@@@@@@@@@@@@")                        
        }
        donor1, _ := strconv.Atoi(string(donorBytes))
        donor1++;

        donorHospitalBytes, err1 := stub.GetState(hospitalId+"D")
        if err1 != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@ hospital donor count doesnot exists @@@@@@@@@@@@@@@@@@@")                        
        }
        donorHospital1, _ := strconv.Atoi(string(donorHospitalBytes))
        donorHospital1++;
        donorId := args[30]+"-D"+ strconv.Itoa(donorHospital1)
        
        donorInfo := `{ "donorId" : "` + donorId  + `" , "typeOfUser" : "` + typeOfUser + `" , "firstName" : "` + firstName + `" , "middleName" : "` + middleName + `" , "lastName" : "` + lastName + `" , "relation" : "` + relation + `" , "dateOfBirth" : "` + dateOfBirth + `" , "age" : "` + age + `" , "gender" : "` + gender + `" , "personalMobileNo" : "` + personalMobileNo + `" , "bloodGroup" : "` + bloodGroup + `" , "emailIdPersonal" : "` + emailIdPersonal + `" , "identificationType" : "` + identificationType + `" , "identificationCard" : "` + identificationCard + `" , "addressLine1" : "` + addressLine1 + `" , "addressLine2" : "` + addressLine2 + `" , "district" : "` + district + `" , "state" : "` + state +`" , "country" : "` + country + `" , "pinCode" : "` + pinCode + `" , "phoneNumber" : "` + phoneNumber + `" , "fullNameEmergencyConct" : "` + fullNameEmergencyConct + `" , "mobileNoEmergencyConct" : "` + mobileNoEmergencyConct + `" , "emailIdEmergencyContct" : "` + emailIdEmergencyContct + `" , "relationEmergencyConct" : "` + relationEmergencyConct + `" , "fullAddressEmergencyConct" : "` + fullAddressEmergencyConct + `" , "districtEmergencyConct" : "` + districtEmergencyConct + `" , "stateEmergencyConct" : "` + stateEmergencyConct + `" , "countryEmergencyConct" : "` + countryEmergencyConct + `" , "pincodeEmergencyConct" : "` + pincodeEmergencyConct + `" , "organStatus" : "` + organStatus + `" }`
    
        err2 := stub.PutState(donorId, []byte(donorInfo))
        if err2 != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to add Donor in Ledger  @@@@@@@@@@@@@@@@@@@")
        }
        err3 := stub.PutState("donor", []byte(strconv.Itoa(donor1)))
        if err3 != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to update country donor count in Ledger  @@@@@@@@@@@@@@@@@@@")
        }
        err4 := stub.PutState(hospitalId+"D", []byte(strconv.Itoa(donorHospital1)))
        if err4 != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to update hospital donor count  in Ledger  @@@@@@@@@@@@@@@@@@@")
        }
    
      return shim.Success([]byte(donorInfo))
    }
    return shim.Success([]byte(typeOfUser))
}


//###############           FUNCTION updateDonorOrRecepient (For updating Donor or Recepient in Ledger )   ###############
func (t *WhoChaincode) updateDonorOrRecepient(stub shim.ChaincodeStubInterface, args []string) pb.Response { 

    shim.Success([]byte("######################  Entering in updateDonorOrRecepient Method  ######################"))

    if len(args) != 31 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of the arguments,  Expecting 31 Arguments @@@@@@@@@@@@@@@@@@@")
    }
    
    //PersonalInfo 
        typeOfUser := args[0];
        firstName := args[1];
        middleName := args[2];
        lastName := args[3];
        relation := args[4];
        dateOfBirth := args[5];
        age := args[6];
        gender := args[7]; 
        personalMobileNo := args[8]
        bloodGroup := args[9]
        emailIdPersonal := args[10]
        //endPersonalinfo
//IdentificationInfo
        identificationType := args[11];
        identificationCard := args[12];
        //endIdentificationInfo
//AddressInfo
        addressLine1 := args[13];
        addressLine2 := args[14];
        district := args[15];
        state := args[16];
        country := args[17];
        pinCode := args[18];
        phoneNumber := args[19];
        //endAddressInfo
//emergencyContactInfo
        fullNameEmergencyConct := args[20];
        mobileNoEmergencyConct := args[21];
        emailIdEmergencyContct := args[22];
        relationEmergencyConct := args[23];
        fullAddressEmergencyConct := args[24];
        districtEmergencyConct := args[25];
        stateEmergencyConct := args[26];
        countryEmergencyConct := args[27];
        pincodeEmergencyConct := args[28];
        //endEmergencyContactlInfo
//organforDonationOrAcceptance
        organStatus := args[29];
        updateId := args[30];
        //endorganforDonationOrAcceptance   


    //adding Recepient
    if strings.ToLower(typeOfUser) == "recepient" {
        
        recepientInfo := `{ "recepientId" : "` + updateId  + `" , "typeOfUser" : "` + typeOfUser + `" , "firstName" : "` + firstName + `" , "middleName" : "` + middleName + `" , "lastName" : "` + lastName + `" , "relation" : "` + relation + `" , "dateOfBirth" : "` + dateOfBirth + `" , "age" : "` + age + `" , "gender" : "` + gender + `" , "personalMobileNo" : "` + personalMobileNo + `" , "bloodGroup" : "` + bloodGroup + `" , "emailIdPersonal" : "` + emailIdPersonal + `" , "identificationType" : "` + identificationType + `" , "identificationCard" : "` + identificationCard + `" , "addressLine1" : "` + addressLine1 + `" , "addressLine2" : "` + addressLine2 + `" , "district" : "` + district + `" , "state" : "` + state +`" , "country" : "` + country + `" , "pinCode" : "` + pinCode + `" , "phoneNumber" : "` + phoneNumber + `" , "fullNameEmergencyConct" : "` + fullNameEmergencyConct + `" , "mobileNoEmergencyConct" : "` + mobileNoEmergencyConct + `" , "emailIdEmergencyContct" : "` + emailIdEmergencyContct + `" , "relationEmergencyConct" : "` + relationEmergencyConct + `" , "fullAddressEmergencyConct" : "` + fullAddressEmergencyConct + `" , "districtEmergencyConct" : "` + districtEmergencyConct + `" , "stateEmergencyConct" : "` + stateEmergencyConct + `" , "countryEmergencyConct" : "` + countryEmergencyConct + `" , "pincodeEmergencyConct" : "` + pincodeEmergencyConct + `" , "organStatus" : "` + organStatus + `" }`
    
        err := stub.PutState(updateId, []byte(recepientInfo))
        if err != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to update Recepient in Ledger  @@@@@@@@@@@@@@@@@@@")
        }
        
        
      return shim.Success([]byte(recepientInfo))
    }   
    
    //adding Donor
    if strings.ToLower(typeOfUser) == "donor" {
        
        donorInfo := `{ "donorId" : "` + updateId  + `" , "typeOfUser" : "` + typeOfUser + `" , "firstName" : "` + firstName + `" , "middleName" : "` + middleName + `" , "lastName" : "` + lastName + `" , "relation" : "` + relation + `" , "dateOfBirth" : "` + dateOfBirth + `" , "age" : "` + age + `" , "gender" : "` + gender + `" , "personalMobileNo" : "` + personalMobileNo + `" , "bloodGroup" : "` + bloodGroup + `" , "emailIdPersonal" : "` + emailIdPersonal + `" , "identificationType" : "` + identificationType + `" , "identificationCard" : "` + identificationCard + `" , "addressLine1" : "` + addressLine1 + `" , "addressLine2" : "` + addressLine2 + `" , "district" : "` + district + `" , "state" : "` + state +`" , "country" : "` + country + `" , "pinCode" : "` + pinCode + `" , "phoneNumber" : "` + phoneNumber + `" , "fullNameEmergencyConct" : "` + fullNameEmergencyConct + `" , "mobileNoEmergencyConct" : "` + mobileNoEmergencyConct + `" , "emailIdEmergencyContct" : "` + emailIdEmergencyContct + `" , "relationEmergencyConct" : "` + relationEmergencyConct + `" , "fullAddressEmergencyConct" : "` + fullAddressEmergencyConct + `" , "districtEmergencyConct" : "` + districtEmergencyConct + `" , "stateEmergencyConct" : "` + stateEmergencyConct + `" , "countryEmergencyConct" : "` + countryEmergencyConct + `" , "pincodeEmergencyConct" : "` + pincodeEmergencyConct + `" , "organStatus" : "` + organStatus + `" }`
    
        err1 := stub.PutState(updateId, []byte(donorInfo))
        if err1 != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to update details of Donor in Ledger  @@@@@@@@@@@@@@@@@@@")
        }
      return shim.Success([]byte(donorInfo))
    }

    return shim.Success([]byte(typeOfUser))
}


//###############           FUNCTION addTransplant (For adding Donor or Recepient in Ledger )   ###############
func (t *WhoChaincode) addTransplant(stub shim.ChaincodeStubInterface, args []string) pb.Response { 

    shim.Success([]byte("######################  Entering in addTransplant Method  ######################"))

    if len(args) != 18 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of the arguments,  Expecting 18 Arguments @@@@@@@@@@@@@@@@@@@")
    }
        typeOfOrgan := args[0];
        dateOfTransplant := args[1];
        HospitalName := args[2];
        city := args[3];
        status := args[4];
//Donor Info
        donorName := args[5];
        donordateOfBirth := args[6];
        donorage := args[7];
        donormobile := args[8]; 
        donorfullAddress := args[9];
        donorcity := args[10];
//Recepient Info
        recepientName := args[11];
        recepdateOfBirth := args[12];
        recepage := args[13];
        recepmobile := args[14]; 
        recepfullAddress := args[15];
        recepcity := args[16]; 
        hospitalId := args[17]; 
     
        transplantBytes, err := stub.GetState("transplant")
        
          

        if err != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@ country transplant count doesnot exists @@@@@@@@@@@@@@@@@@@")                        
        }
        transplant1, _ := strconv.Atoi(string(transplantBytes))
        transplant1++;

        transplantHospitalBytes, err1 := stub.GetState(hospitalId+"T")
        if err1 != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@ hospital transplant count doesnot exists @@@@@@@@@@@@@@@@@@@")                        
        }
        transplantHospital1, _ := strconv.Atoi(string(transplantHospitalBytes))
        transplantHospital1++;
        transplant1Id := hospitalId +"-T"+ strconv.Itoa(transplantHospital1)
      

        transplantInfo := `{ "transplantId" : "` + transplant1Id  + `" , "typeOfOrgan" : "` + typeOfOrgan + `" , "dateOfTransplant" : "` + dateOfTransplant + `" , "HospitalName" : "` + HospitalName + `" , "city" : "` + city + `" , "status" : "` + status + `" , "donorName" : "` + donorName + `" , "donordateOfBirth" : "` + donordateOfBirth + `" , "donorage" : "` + donorage + `" , "donormobile" : "` + donormobile + `" , "donorfullAddress" : "` + donorfullAddress + `" , "donorcity" : "` + donorcity + `" , "recepientName" : "` + recepientName + `" , "recepdateOfBirth" : "` + recepdateOfBirth + `" , "recepage" : "` + recepage + `" , "recepmobile" : "` + recepmobile + `" , "recepfullAddress" : "` + recepfullAddress + `" , "recepcity" : "` + recepcity + `" }`
    
        err2 := stub.PutState(transplant1Id, []byte(transplantInfo))
        if err2 != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to add Transplant in Ledger  @@@@@@@@@@@@@@@@@@@")
        }
        err3 := stub.PutState("transplant", []byte(strconv.Itoa(transplant1)))
        if err3 != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to update country transplant count in Ledger  @@@@@@@@@@@@@@@@@@@")
        }
        err4 := stub.PutState(hospitalId+"T", []byte(strconv.Itoa(transplantHospital1)))
        if err4 != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to update hospital transplant count  in Ledger  @@@@@@@@@@@@@@@@@@@")
        }
    
      return shim.Success([]byte(transplantInfo))

}

//###############           FUNCTION updateTransplant (For adding Donor or Recepient in Ledger )   ###############
func (t *WhoChaincode) updateTransplant(stub shim.ChaincodeStubInterface, args []string) pb.Response { 

    shim.Success([]byte("######################  Entering in updateTransplant Method  ######################"))

    if len(args) != 18 {
        return shim.Error("@@@@@@@@@@@@@@@@@@@  Incorrect no. of the arguments,  Expecting 18 Arguments @@@@@@@@@@@@@@@@@@@")
    }
        typeOfOrgan := args[0];
        dateOfTransplant := args[1];
        HospitalName := args[2];
        city := args[3];
        status := args[4];
//Donor Info
        donorName := args[5];
        donordateOfBirth := args[6];
        donorage := args[7];
        donormobile := args[8]; 
        donorfullAddress := args[9];
        donorcity := args[10];
//Recepient Info
        recepientName := args[11];
        recepdateOfBirth := args[12];
        recepage := args[13];
        recepmobile := args[14]; 
        recepfullAddress := args[15];
        recepcity := args[16]; 
        transplantId := args[17]; 
     
       
        
        transplantInfo := `{ "transplantId" : "` + transplantId  + `" , "typeOfOrgan" : "` + typeOfOrgan + `" , "dateOfTransplant" : "` + dateOfTransplant + `" , "HospitalName" : "` + HospitalName + `" , "city" : "` + city + `" , "status" : "` + status + `" , "donorName" : "` + donorName + `" , "donordateOfBirth" : "` + donordateOfBirth + `" , "donorage" : "` + donorage + `" , "donormobile" : "` + donormobile + `" , "donorfullAddress" : "` + donorfullAddress + `" , "donorcity" : "` + donorcity + `" , "recepientName" : "` + recepientName + `" , "recepdateOfBirth" : "` + recepdateOfBirth + `" , "recepage" : "` + recepage + `" , "recepmobile" : "` + recepmobile + `" , "recepfullAddress" : "` + recepfullAddress + `" , "recepcity" : "` + recepcity + `" }`
    
        err := stub.PutState(string(transplantId), []byte(transplantInfo))
        if err != nil {
            return shim.Error("@@@@@@@@@@@@@@@@@@@  Unable to update Transplant in Ledger  @@@@@@@@@@@@@@@@@@@")
        }
        
      return shim.Success([]byte(transplantInfo))

}


//###############           FUNCTION queryTransplant (For querying Details of Transplant from Ledger)   ###############
func (t *WhoChaincode) queryTransplant(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
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
func (t *WhoChaincode) queryDonor(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
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
func (t *WhoChaincode) queryRecepient(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
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
func (t *WhoChaincode) queryHospitalById(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
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
func (t *WhoChaincode) queryDonorById(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
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
func (t *WhoChaincode) queryRecepientById(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
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
func (t *WhoChaincode) queryTransplantById(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
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
func (t *WhoChaincode) queryCount(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
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
func (t *WhoChaincode) queryCountryCount(stub shim.ChaincodeStubInterface) pb.Response {
    
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
func (t *WhoChaincode) totalTransplantCount(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
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
func (t *WhoChaincode) getHospitalName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
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


func (t *WhoChaincode) callOtherChaincodeInvoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var channelName string // Sum entity
	//var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting atleast 2")
	}

	chaincodeName := args[0] // Expecting name of the chaincode you would like to call, this name would be given during chaincode install time
    channelName = args[1]


	// Query chaincode_example02
	f := "addCount"
	queryArgs := toChaincodeArgs(f, "ee")
        //return shim.Success(queryArgs)
       // ab := string(queryArgs[0]) + string(queryArgs[1]);
       // return shim.Success([]byte(ab))


	response := stub.InvokeChaincode(chaincodeName, queryArgs, channelName)
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", response.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}
	
	return shim.Success([]byte(string(response.Payload)))
}

func (t *WhoChaincode) callOtherChaincodeQuery(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var channelName string // Sum entity
	//var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting atleast 2")
	}

	chaincodeName := args[0] // Expecting name of the chaincode you would like to call, this name would be given during chaincode install time
    channelName = args[1]


	// Query chaincode_example02
	f := "queryCount"
	queryArgs := toChaincodeArgs(f, "ee")
        //return shim.Success(queryArgs)
       // ab := string(queryArgs[0]) + string(queryArgs[1]);
       // return shim.Success([]byte(ab))


	response := stub.InvokeChaincode(chaincodeName, queryArgs, channelName)
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", response.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}
	
	return shim.Success([]byte(string(response.Payload)))
}

//###############           FUNCTION main   ############### 
func main() {
	// Create a new Smart Contract
	err := shim.Start(new(WhoChaincode))
	if err != nil {
		fmt.Printf("@@@@@@@@@@@@@@@@@@@  Error creating new WhoChaincode : %s  @@@@@@@@@@@@@@@@@@@", err)
	}
} 



