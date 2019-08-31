package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// ============================================================================================================================
// getContractorBySearchingCriteria() - Retrieve all contractors based on searching criteria
// ============================================================================================================================
func getContractorBySearchingCriteria(stub shim.ChaincodeStubInterface , args []string) pb.Response {
	
	var srstr = args[0];
	var searchstr SearchStr ;

	err := json.Unmarshal([]byte(srstr), &searchstr)
	if err != nil {
		return shim.Error("@@@@@@@@@@@@@@@@@@@  Could not unmarshal Search String  @@@@@@@@@@@@@@@@@@@")
	}
	 
	var first_name []string;
    for j := 0; j < len(searchstr.FirstName) ; j++ {
        first_name = append(first_name, searchstr.FirstName[j].Name)
    }
    var languages []string;
    for j := 0; j < len(searchstr.Languages) ; j++ {
        languages = append(languages, searchstr.Languages[j].Name)
    }
    var experience_location []string;
    for j := 0; j < len(searchstr.ExperienceLocation) ; j++ {
        experience_location = append(experience_location, searchstr.ExperienceLocation[j].Name)
    }
    var experience []int;
    for j := 0; j < len(searchstr.Experience) ; j++ {
        experience = append(experience, searchstr.Experience[j].Name)
    }
    var personal_email []string;
    for j := 0; j < len(searchstr.PersonalEmail) ; j++ {
        personal_email = append(personal_email, searchstr.PersonalEmail[j].Name)
    }
    var mobile_number []string;
    for j := 0; j < len(searchstr.MobileNumber) ; j++ {
        mobile_number = append(mobile_number, searchstr.MobileNumber[j].Name)
	}
	

	poverty_reduction := searchstr.ExpertiseArea.PovertyReduction.SubCategory
	poverty_reductionBytes, err := json.MarshalIndent(poverty_reduction, "", "  ")
	if err != nil {
	  //Handle the error
	}
	var poverty_reductionStr interface{} = string(poverty_reductionBytes)
	if ( len(poverty_reduction) == 0 ) { 
		poverty_reductionStr = []string{}
	}
	
	gender_equality := searchstr.ExpertiseArea.GenderEqualityAndEmpowermentOfWomen.SubCategory
	gender_equalityBytes, err := json.MarshalIndent(gender_equality, "", "  ")
	if err != nil {
	  //Handle the error
	}
	var gender_equalityStr interface{} = string(gender_equalityBytes)
	if ( len(gender_equality) == 0 ) { 
		gender_equalityStr = []string{}
	}

	gender_equality := searchstr.ExpertiseArea.GenderEqualityAndEmpowermentOfWomen.SubCategory
	gender_equalityBytes, err := json.MarshalIndent(gender_equality, "", "  ")
	if err != nil {
	  //Handle the error
	}
	var gender_equalityStr interface{} = string(gender_equalityBytes)
	if ( len(gender_equality) == 0 ) { 
		gender_equalityStr = []string{}
	}	

	first_nameBytes, err := json.MarshalIndent(first_name, "", "")
	if err != nil {
		// Handle the error
	}
	var first_nameStr interface{} = string(first_nameBytes)
	if ( len(first_name) == 0 ) { 
		first_nameStr = []string{}
	}

	languagesBytes, err := json.MarshalIndent(languages, "", "")
	if err != nil {
		// Handle the error
	}
	var languagesStr interface{} = string(languagesBytes)
	if ( len(languages) == 0 ) { 
		languagesStr = []string{}
	}
    
	experience_locationBytes, err := json.MarshalIndent(experience_location, "", "")
	if err != nil {
		// Handle the error
	}
    var experience_locationStr interface{} = string(experience_locationBytes)
	if ( len(experience_location) == 0 ) { 
		experience_locationStr = []string{}
	}

	// experienceStr, err := json.MarshalIndent(experience, "", "")
	// if err != nil {
	// 	// Handle the error
	// }
		
	personal_emailBytes, err := json.MarshalIndent(personal_email, "", "")
	if err != nil {
		// Handle the error
	}
	var personal_emailStr interface{} = string(personal_emailBytes)
	if ( len(personal_email) == 0 ) { 
		personal_emailStr = []string{}
	}

	mobile_numberBytes, err := json.MarshalIndent(mobile_number, "", "  ")
	if err != nil {
		// Handle the error
	}
    var mobile_numberStr interface{} = string(mobile_numberBytes)
	if ( len(mobile_number) == 0 ) { 
		mobile_numberStr = []string{}
	}
	
	queryString := fmt.Sprintf("{\"selector\":{\"$or\":[{\"EmployementHistory\":{\"$elemMatch\":{\"$or\":[{\"ExpertiseArea\":{\"$or\":[{\"PovertyReduction\":{\"$in\": %v }},{\"GenderEqualityAndEmpowermentOfWomen\":{\"$in\": %v }}]}},{\"ExperienceLocation\":{\"$in\": %v }},{\"Experience\":{\"gte\": %v }}]}}},{\"Languages\":{\"$elemMatch\":{\"Name\":{\"$in\": %v }}}},{\"FirstName\":{\"$in\": %v }},{\"PersonalEmail\":{\"$in\": %v }},{\"MobileNumber\":{\"$in\": %v }}]}}",poverty_reductionStr, gender_equalityStr, experience_locationStr, experience[0], languagesStr, first_nameStr, personal_emailStr, mobile_numberStr)
    
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error("Unable to retrieve contractors records from ledger."+err.Error())
	}
	return shim.Success(queryResults)
}

// ============================================================================================================================
// getQueryResultForQueryString() - Retrieve Records based on query string
// ============================================================================================================================
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return buffer.Bytes(), nil
}	


