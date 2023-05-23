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

// TODO: Should this conversion function also present in the same file?
func ConvertJsonToRecept(jsonStr string) (Receipt, error) {
	// TODO: Is this the way to do?
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

func validateReceipt(receipt Receipt) error {
	//validation rule for date
	dateValidationRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if receipt.PurchaseDate == "" {
		return fmt.Errorf("Purchase date is required")
	} else if !dateValidationRegex.MatchString(receipt.PurchaseDate) {
		return fmt.Errorf("Purchase date format should be in the format yyyy-mm-dd")
	}

	// validation rule for time
	timeValidationRegex := regexp.MustCompile(`^\d{2}:\d{2}$`)
	if receipt.PurchaseTime == "" {
		return fmt.Errorf("Purchase date is required")
	} else if !timeValidationRegex.MatchString(receipt.PurchaseTime) {
		return fmt.Errorf("Purchase date format should be in the format yyyy-mm-dd")
	}

	return nil
}
