package models

import (
	"fmt"
	"io/ioutil"
	"testing"
)

//TODO: Is this the way to do?
func getSampleJsonPaths() []string {
	samples := []string{"samples/morning-receipt.json", "samples/simple-receipt.json", "samples/target.json", "samples/m-m.json"}
	return samples
}

// TODO: Is this the way to do?
func getSampleReceipts() []Receipt {
	samples := []Receipt{
		{
			Retailer:     "Walgreens",
			PurchaseDate: "2022-01-02",
			PurchaseTime: "08:13",
			Total:        2.65,
			Items: []Item{
				{
					ShortDescription: "Pepsi - 12-oz",
					Price:            1.25,
				},
				{
					ShortDescription: "Dasani",
					Price:            1.40,
				},
			},
		},
		{
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
		{
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
		{
			Retailer:     "M&M Corner Market",
			PurchaseDate: "2022-03-20",
			PurchaseTime: "14:33",
			Total:        9.00,
			Items: []Item{
				{
					ShortDescription: "Gatorade",
					Price:            2.25,
				},
				{
					ShortDescription: "Gatorade",
					Price:            2.25,
				},
				{
					ShortDescription: "Gatorade",
					Price:            2.25,
				},
				{
					ShortDescription: "Gatorade",
					Price:            2.25,
				},
			},
		},
	}
	return samples
}

// TODO: None of the falses explain what the difference is.
func compareItems(itemA, itemB Item) error {
	if itemA.ShortDescription != itemB.ShortDescription {
		return fmt.Errorf("item shortdescription not same; Actual: %s, Expected: %s\n", itemA.ShortDescription, itemB.ShortDescription)
	}
	if itemA.Price != itemB.Price {
		return fmt.Errorf("item shortdescription not same; Actual: %f, Expected: %f\n", itemA.Price, itemB.Price)
	}
	return nil
}

// TODO: None of the falses return what is the difference in the receipts
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
	sampleJsonPaths := getSampleJsonPaths()
	sampleReceipts := getSampleReceipts()

	for index := 0; index < len(sampleJsonPaths); index++ {
		sampleJsonPath := sampleJsonPaths[index]

		// Read the file contents as bytes
		fileBytes, err := ioutil.ReadFile(sampleJsonPath)
		if err != nil {
			t.Errorf("Failed to read: %s", sampleJsonPath)
			continue
		}

		// Convert bytes to string
		sampleJson := string(fileBytes)
		actualReceipt, err := ConvertJsonToRecept(sampleJson)
		if err != nil {
			t.Errorf("Failed to convert %s\n", err)
			continue
		}

		expectedReceipt := sampleReceipts[index]

		// Compare actual and expected
		err = compareReceipts(actualReceipt, expectedReceipt)
		if err != nil {
			t.Errorf("Comparison failed for %s with error: %s", sampleJsonPath, err)
			continue
		}
	}
}
