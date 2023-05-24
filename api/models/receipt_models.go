package models

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// Item struct
type Item struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price,string"`
}

// Receipt struct
type Receipt struct {
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Total        float64 `json:"total,string"`
	Items        []Item  `json:"items"`
}

// ConvertJsonToReceipt -  Convert Json string to Receipt datatype
//
// Parameters:
//   - jsonStr: Json as string
//
// Return:
//   - Receipt: Receipt data if the json string is valid
//   - error: Error explaining the incompatiblity in the json string
func ConvertJsonToReceipt(jsonStr string) (Receipt, error) {
	var receipt Receipt
	err := json.Unmarshal([]byte(jsonStr), &receipt)
	if err != nil {
		return Receipt{}, err
	}

	err = validateReceipt(receipt)
	if err != nil {
		return Receipt{}, err
	}
	return receipt, nil
}

// validateitem -  Validates Item data
//
// Parameters:
//   - item: item data
//
// Return:
//   - error: Error explaining which validation rule is not followed
func validateItem(item Item) error {
	//validation rule for description name
	nameValidationRegex := regexp.MustCompile(`^[\w\s\-]+$`)
	if item.ShortDescription == "" {
		return fmt.Errorf("item name is required")
	} else if !nameValidationRegex.MatchString(item.ShortDescription) {
		return fmt.Errorf("item name should adhere to ^[\\w\\s\\-]+$")
	}

	return nil
}

// validateReceipt -  Validates Receipt data
//
// Parameters:
//   - receipt: Receipt data
//
// Return:
//   - error: Error explaining which validation rule is not followed
func validateReceipt(receipt Receipt) error {
	//validation rule for Retailer name
	nameValidationRegex := regexp.MustCompile(`^\S+$`)
	if receipt.Retailer == "" {
		return fmt.Errorf("receipt name is required")
	} else if !nameValidationRegex.MatchString(receipt.PurchaseDate) {
		return fmt.Errorf("receipt name should adhere to ^\\S+$")
	}

	//validation rule for Date
	dateValidationRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if receipt.PurchaseDate == "" {
		return fmt.Errorf("purchase date is required")
	} else if !dateValidationRegex.MatchString(receipt.PurchaseDate) {
		return fmt.Errorf("purchase date format should be in the format yyyy-mm-dd")
	}

	// validation rule for Time
	timeValidationRegex := regexp.MustCompile(`^\d{2}:\d{2}$`)
	if receipt.PurchaseTime == "" {
		return fmt.Errorf("purchase time is required")
	} else if !timeValidationRegex.MatchString(receipt.PurchaseTime) {
		return fmt.Errorf("purchase time format should be in the format hh:mm")
	}

	//validation rule for items
	noOfItems := len(receipt.Items)
	if noOfItems == 0 {
		return fmt.Errorf("no of items should not be 0")
	}

	for index := 0; index < len(receipt.Items); index++ {
		item := receipt.Items[index]
		err := validateItem(item)

		if err != nil {
			return fmt.Errorf(err.Error())
		}
	}

	return nil
}
