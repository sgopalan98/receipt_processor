package models

import (
	"fmt"
	"io/ioutil"
	"testing"
)

// mock data for validation
var validationTestMockData = []struct {
	ReceiptData Receipt
	Error       error
}{
	{
		ReceiptData: Receipt{
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
		},
		Error: fmt.Errorf("purchase date format should be in the format yyyy-mm-dd"),
	},
	{
		ReceiptData: Receipt{
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
		},
		Error: nil,
	},
	{
		ReceiptData: Receipt{
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
		},
		Error: fmt.Errorf("purchase time format should be in the format hh:mm"),
	},
	{
		ReceiptData: Receipt{
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
		},
		Error: nil,
	},
}

// mock data for json convertion
var convertJsonMockData = []struct {
	JsonPath    string
	ReceiptData Receipt
}{
	{
		JsonPath: "samples/simple.json",
		ReceiptData: Receipt{
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
		},
	},
	{
		JsonPath: "samples/complex.json",
		ReceiptData: Receipt{
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
		},
	},
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

func TestConvertJson(t *testing.T) {

	mockData := convertJsonMockData

	for index := 0; index < len(mockData); index++ {
		sampleJsonPath := mockData[index].JsonPath

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

		expectedReceipt := mockData[index].ReceiptData

		// Compare actual and expected
		err = compareReceipts(actualReceipt, expectedReceipt)
		if err != nil {
			t.Errorf("Comparison failed for %s with error: %s", sampleJsonPath, err)
			continue
		}
	}
}

func TestValidationRules(t *testing.T) {
	mockData := validationTestMockData

	for index := 0; index < len(mockData); index++ {
		input := mockData[index].ReceiptData
		expectedOutput := mockData[index].Error
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
