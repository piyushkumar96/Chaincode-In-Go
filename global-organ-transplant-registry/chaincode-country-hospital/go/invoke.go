package main

import (
    "bytes"
    "strconv"
    "strings"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

//###############           FUNCTION addHospital (For adding Hospital in Ledger )   ###############
func  addHospital(stub shim.ChaincodeStubInterface, args []string) pb.Response { 

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
func updateHospital(stub shim.ChaincodeStubInterface, args []string) pb.Response { 

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


//###############           FUNCTION addDonorOrRecepient (For adding Donor or Recepient in Ledger )   ###############
func addDonorOrRecepient(stub shim.ChaincodeStubInterface, args []string) pb.Response { 

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
func updateDonorOrRecepient(stub shim.ChaincodeStubInterface, args []string) pb.Response { 

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
func addTransplant(stub shim.ChaincodeStubInterface, args []string) pb.Response { 

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
func updateTransplant(stub shim.ChaincodeStubInterface, args []string) pb.Response { 

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
