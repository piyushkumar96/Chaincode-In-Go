package main

//==============================================================================================================================
//	 Common chain code error messages
//==============================================================================================================================
const (
	ArgumentErrorMessage = "Insufficient number of arguments"
	PutErrorMessage  = "Error in inserting data into ledger"
	GetStateErrorMessage  = "Unable to retrieve key value"
	DeleteStateErrorMessage  = "Unable to delete key value"
	MarshalErrorMessage  = "Unable to marshal json data"
	UnmarshalErrorMessage  = "Unable to unmarshal data"
	AuthorizationErrorMessage = "This user is not authorize to perform this operation on the chain"
)
