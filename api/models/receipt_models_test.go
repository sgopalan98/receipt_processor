package models

import (
	"fmt"
	"io/ioutil"
	"testing"
)

// mock data for validation
func getValidationTestMockData() ([]Receipt, []error) {
	var receipts []Receipt
	var errors []error

	// Date not in the format yyyy-mm-dd
	dateNotInFormat := Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: "202-01-02",
		PurchaseTime: "08:13",
		Total:        2.65,
		Items: []Item{
			{
				ShortDescription: "Pepsi - 12-oz",
				Price:            1.25,
			},
		},
	}
	receipts = append(receipts, dateNotInFormat)
	errors = append(errors, fmt.Errorf("purchase date format should be in the format yyyy-mm-dd"))

	// Date in format yyyy-mm-dd
	dateInFormat := Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "08:13",
		Total:        2.65,
		Items: []Item{
			{
				ShortDescription: "Pepsi - 12-oz",
				Price:            1.25,
			},
		},
	}
	receipts = append(receipts, dateInFormat)
	errors = append(errors, nil)

	// Date not in the format yyyy-mm-dd
	timeNotInFormat := Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "8:13",
		Total:        2.65,
		Items: []Item{
			{
				ShortDescription: "Pepsi - 12-oz",
				Price:            1.25,
			},
		},
	}
	receipts = append(receipts, timeNotInFormat)
	errors = append(errors, fmt.Errorf("purchase time format should be in the format hh:mm"))

	// Date in format yyyy-mm-dd
	timeInFormat := Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "08:13",
		Total:        2.65,
		Items: []Item{
			{
				ShortDescription: "Pepsi - 12-oz",
				Price:            1.25,
			},
		},
	}
	receipts = append(receipts, timeInFormat)
	errors = append(errors, nil)

	return receipts, errors
}

// mock data for json convertion
func getConvertJsonMockData() ([]string, []Receipt) {
	var jsonPaths []string
	var receipts []Receipt

	// simple json
	simpleJsonPath := "samples/simple.json"

	simpleJsonReceipt := Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "13:13",
		Total:        1.25,
		Items: []Item{
			{
				ShortDescription: "Pepsi - 12-oz",
				Price:            1.25,
			},
		},
	}

	jsonPaths = append(jsonPaths, simpleJsonPath)
	receipts = append(receipts, simpleJsonReceipt)
	// complex json

	complexJsonPath := "samples/complex.json"

	complexJsonReceipt := Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Total:        35.35,
		Items: []Item{
			{
				ShortDescription: "Mountain Dew 12PK",
				Price:            6.49,
			},
			{
				ShortDescription: "Emils Cheese Pizza",
				Price:            12.25,
			},
			{
				ShortDescription: "Knorr Creamy Chicken",
				Price:            1.26,
			},
			{
				ShortDescription: "Doritos Nacho Cheese",
				Price:            3.35,
			},
			{
				ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
				Price:            12.00,
			},
		},
	}

	jsonPaths = append(jsonPaths, complexJsonPath)
	receipts = append(receipts, complexJsonReceipt)

	return jsonPaths, receipts
}

func compareItems(itemA, itemB Item) error {
	if itemA.ShortDescription != itemB.ShortDescription {
		return fmt.Errorf("item shortdescription not same; Actual: %s, Expected: %s\n", itemA.ShortDescription, itemB.ShortDescription)
	}
	if itemA.Price != itemB.Price {
		return fmt.Errorf("item shortdescription not same; Actual: %f, Expected: %f\n", itemA.Price, itemB.Price)
	}
	return nil
}

func compareReceipts(receiptA, receiptB Receipt) error {
	if receiptA.Retailer != receiptB.Retailer {
		return fmt.Errorf("Retailer name not same; Actual: %s, Expected: %s\n", receiptA.Retailer, receiptB.Retailer)
	}
	if receiptA.PurchaseDate != receiptB.PurchaseDate {
		return fmt.Errorf("Purchase date not same; Actual: %s, Expected: %s\n", receiptA.PurchaseDate, receiptB.PurchaseDate)
	}
	if receiptA.PurchaseTime != receiptB.PurchaseTime {
		return fmt.Errorf("Purchase time not same; Actual: %s, Expected: %s\n", receiptA.PurchaseTime, receiptB.PurchaseTime)
	}
	if receiptA.Total != receiptB.Total {
		return fmt.Errorf("Purchase total not same; Actual: %f, Expected: %f\n", receiptA.Total, receiptB.Total)
	}
	if len(receiptA.Items) != len(receiptB.Items) {
		return fmt.Errorf("No of items not same\n")
	}

	for index := 0; index < len(receiptA.Items); index++ {
		itemA := receiptA.Items[index]
		itemB := receiptB.Items[index]
		err := compareItems(itemA, itemB)
		if err != nil {
			return fmt.Errorf("Item %d don't match with error: %s\n", index, err)
		}
	}
	return nil
}

// TODO: Think about if you actually need this - because the code is pretty much standard.
func TestConvertJson(t *testing.T) {

	jsonPaths, receipts := getConvertJsonMockData()

	for index := 0; index < len(jsonPaths); index++ {
		sampleJsonPath := jsonPaths[index]

		// Read the file contents as bytes
		fileBytes, err := ioutil.ReadFile(sampleJsonPath)
		if err != nil {
			t.Errorf("Failed to read: %s", sampleJsonPath)
			continue
		}

		// Convert bytes to string
		sampleJson := string(fileBytes)
		actualReceipt, err := ConvertJsonToReceipt(sampleJson)
		if err != nil {
			t.Errorf("Failed to convert %s\n", err)
			continue
		}

		expectedReceipt := receipts[index]

		// Compare actual and expected
		err = compareReceipts(actualReceipt, expectedReceipt)
		if err != nil {
			t.Errorf("Comparison failed for %s with error: %s", sampleJsonPath, err)
			continue
		}
	}
}

func TestValidationRules(t *testing.T) {
	inputs, expectedOutputs := getValidationTestMockData()

	for index := 0; index < len(inputs); index++ {
		input := inputs[index]
		expectedOutput := expectedOutputs[index]
		actualOutput := validateReceipt(input)

		if actualOutput != nil && expectedOutput == nil {
			t.Errorf("For input %d, expected: nil actual: %s \n", index, actualOutput)
			continue
		} else if actualOutput == nil && expectedOutput != nil {
			t.Errorf("For input %d, expected: %s actual: nil \n", index, expectedOutput)
			continue
		} else if actualOutput != nil && expectedOutput != nil && actualOutput.Error() != expectedOutput.Error() {
			t.Errorf("For input %d, expected: %s actual: %s \n", index, expectedOutput, actualOutput)
			continue
		}
	}
}
