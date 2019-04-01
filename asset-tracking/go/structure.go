package main

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