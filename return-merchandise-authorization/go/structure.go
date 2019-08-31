package main

// Order Structure 
type RMA struct {
	rmaNo                   string              `json:"rmaNo"`             // Unique Ticket Number
	messages				[]interface{}				`json:"messages"`			// messages
}


/*// Message Structure 
type Message struct {
	mc_3b11                		MC_3B11              	`json:"mc_3b11"`  
	mc_3b3sir                	MC_3B3SIR                `json:"mc_3b3sir"`  
	mc_3b3eta                	MC_3B3ETA               `json:"mc_3b3eta"`  
	mc_3b13                		MC_3B13              	`json:"mc_3b13"`  
	mc_3b3pod                	MC_3B3POD               `json:"mc_3b3pod"`         
	
}

// Message Code 3B11                                                 
type MC_3B11 struct {
	_type                      string              ` default:"3b11" json:"_type"`              
	message                   string              `json:"message"`             				 
	receiveddate			  string			  `json:"receiveddate"`              			
}

// Message Code 3B3SIR                                                 
type MC_3B3SIR struct {
	_type                      string              ` default:"3b3sir" json:"_type"`              
	message                   string              `json:"message"`             				 
	receiveddate			  string			  `json:"receiveddate"`              			
}

// Message Code 3B3ETA                                                 
type MC_3B3ETA struct {
	_type                      string              ` default:"3b3eta" json:"_type"`              
	message                   string              `json:"message"`             				 
	receiveddate			  string			  `json:"receiveddate"`              			
}

// Message Code 3B13                                                 
type MC_3B13 struct {
	_type                      string              ` default:"3b13" json:"_type"`              
	message                   string              `json:"message"`             				 
	receiveddate			  string			  `json:"receiveddate"`              			
}

// Message Code 3B3POD                                                 
type MC_3B3POD struct {
	_type                      string              ` default:"3b3pod" json:"_type"`              
	message                   string              `json:"message"`             				 
	receiveddate			  string			  `json:"receiveddate"`              			
}
*/