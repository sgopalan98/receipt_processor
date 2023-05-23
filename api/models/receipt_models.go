package models

import (
	"encoding/json"
	"fmt"
)

// Item struct
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// Receipt struct
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total        string `json:"total"`
	Items        []Item `json:"items"`
}

// TODO: Is this conversion function also present in the same file?
func ConvertJsonToRecept(jsonStr string) Receipt {
	// TODO: Is this the way to do?
	var receipt Receipt
	err := json.Unmarshal([]byte(jsonStr), &receipt)
	if err != nil {
		fmt.Println("Error: ", err)
		return Receipt{}
	}

	return receipt
}
