package main

// Define the ContractorInfo structure, with following properties.
type Contractor struct {
	FamilyName				    string					`json:"family_name" validate:"required"`
	FirstName				    string					`json:"first_name" validate:"required"`
	MiddleName				    string					`json:"middle_name"`                 
	Gender				        string					`json:"gender" validate:"required"`
	DOB				    		string					`json:"dob" validate:"required"`
	Nationality				    string					`json:"nationality" validate:"required"`
	PersonalEmail				string					`json:"personal_email" validate:"required"`
	OfficialEmail				string					`json:"official_email" validate:"required"`
	MobileNumber				string					`json:"mobile_number" validate:"required"`
	OfficialTelephone			string					`json:"official_telephone" validate:"required"`
	PresentAddress				string					`json:"present_address" validate:"required"`
	PermanentAddress			string					`json:"permanent_address" validate:"required"`
	RegistrationDate			string					`json:"registration_date" validate:"required"`
	Languages					[]Languages	    		`json:"languages"`
	EmployementHistory			[]EmployementHistory	`json:"employement_history"`
}

type Languages struct {
	Name						string					`json:"name"`
	Read	    				bool					`json:"read"`
	Write	    				bool					`json:"write"`
	Speak	    				bool					`json:"speak"`
	Fluency	    				string					`json:"fluency"`
}

type EmployementHistory struct {
	Title    					string						`json:"title"`
	Company	    				string						`json:"company"`
	Country	        			string						`json:"country"`
	Grade	        			string						`json:"grade"`
	ExperienceLocation  		string						`json:"experience_location"`
	Description     			string						`json:"description"`
	ExpertiseArea     			map[string]interface{} 		`json:"expertise_area"`
	From     					string						`json:"from"`
	To    						string						`json:"to"`
	Experience          		int							`json:"experience"`
}

type SearchStr struct {
	ExpertiseArea     			ExpertiseArea	    		`json:"expertise_area"`
	Languages	    			[]OtherField				`json:"languages"`
	ExperienceLocation	    	[]OtherField				`json:"experience_location"`
	Experience  				[]ExperienceField			`json:"experience"`
	FirstName           		[]OtherField              	`json:"first_name"`
	PersonalEmail       		[]OtherField               	`json:"personal_email"`
	MobileNumber        		[]OtherField              	`json:"mobile_number"`
}

type ExpertiseArea struct {
	PovertyReduction     		            ExpertiseField   	`json:"poverty_reduction"`
	GenderEqualityAndEmpowermentOfWomen 	ExpertiseField		`json:"gender_equality_and_empowerment_of_women"`
	YouthEmpowerment  		                ExpertiseField		`json:"youth_empowerment"`
	Health          						ExpertiseField      `json:"health"`
}

// subcategory of expertise area
type ExpertiseField struct {
	SubCategory	    						[]string	    	`json:"sub_category"`
	Weightage	      						float64				`json:"weightage"`
}

type OtherField struct {
	Name	   								string	    		`json:"name"`
	Weightage		    					float64				`json:"weightage"`
}

type ExperienceField struct {
	Name	   								int	    			`json:"name"`
	Weightage		    					float64				`json:"weightage"`
}