package main

// Cocoa Bean Bag Structure 
type CocoaBeanBag struct {
	BeanBagId                   string              `json:"beanBagId"`         	    // Unique Cocoa Bean Bag Id
	Country						string				`json:"country"`				// Country Code eg US, IN
	FarmerId					string				`json:"farmerId"`				// eg F201
	IsConsumed					bool				`json:"isConsumed"`				// eg. true/false
}

// Chocolate Bar Structure 
type ChocoBar struct {
	BarId                   	string              `json:"barId"`         	    	// Unique Choco Bar Id
	BeanBagId                   string              `json:"beanBagId"`         	    // Unique Cocoa Bean Bag Id
}
