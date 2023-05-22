package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

//TODO: Is this the way to do?
func getSampleJsonPaths() []string {
	samples := []string{"samples/simple-receipt.json", "samples/morning-receipt.json"}
	return samples
}

// TODO: Is this the way to do?
func getSampleReceipts() []Receipt {
	samples := []Receipt{
		{
			Retailer:     "Walgreens",
			PurchaseDate: "2022-01-02",
			PurchaseTime: "08:13",
			Total:        "2.65",
			Items: []Item{
				{
					ShortDescription: "Pepsi - 12-oz",
					Price:            "1.25",
				},
				{
					ShortDescription: "Dasani",
					Price:            "1.40",
				},
			},
		},
		{
			Retailer:     "Target",
			PurchaseDate: "2022-01-02",
			PurchaseTime: "13:13",
			Total:        "1.25",
			Items: []Item{
				{
					ShortDescription: "Pepsi - 12-oz",
					Price:            "1.25",
				},
			},
		},
	}
	return samples
}

// TODO: Compute manually
func getSamplePoints() []int {
	points := []int{14, 31}
	return points
}

// TODO: None of the falses explain what the difference is.
func compareItems(itemA, itemB Item) bool {
	if itemA.ShortDescription != itemB.ShortDescription {
		return false
	}
	if itemA.Price != itemB.Price {
		return false
	}
	return true
}

// TODO: None of the falses return what is the difference in the receipts
func compareReceipts(receiptA, receiptB Receipt) bool {
	if receiptA.Retailer != receiptB.Retailer {
		return false
	}
	if receiptA.PurchaseDate != receiptB.PurchaseDate {
		return false
	}
	if receiptA.PurchaseTime != receiptB.PurchaseTime {
		return false
	}
	if len(receiptA.Items) != len(receiptB.Items) {
		return false
	}

	for index := 0; index < len(receiptA.Items); index++ {
		itemA := receiptA.Items[index]
		itemB := receiptB.Items[index]
		if !compareItems(itemA, itemB) {
			return false
		}
	}
	return true
}

func TestConvertJson(t *testing.T) {
	fmt.Printf("STARTING\n")
	sampleJsonPaths := getSampleJsonPaths()
	sampleReceipts := getSampleReceipts()

	for index := 0; index < len(sampleJsonPaths); index++ {
		sampleJsonPath := sampleJsonPaths[index]

		// Read the file contents as bytes
		fileBytes, err := ioutil.ReadFile(sampleJsonPath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		// Convert bytes to string
		sampleJson := string(fileBytes)

		expectedReceipt := sampleReceipts[index]

		actualReceipt := convertJson(sampleJson)
		comparison := compareReceipts(actualReceipt, expectedReceipt)
		if !comparison {
			t.Errorf("NOT EQUAL \n")
		}
	}
}

func TestComputePoints(t *testing.T) {
	sampleJsonReceipts := getSampleReceipts()
	samplePoints := getSamplePoints()
	sampleJsonPaths := getSampleJsonPaths()

	// TODO: Is this the for loop way?
	for index := 0; index < len(sampleJsonReceipts); index++ {
		sampleJsonReceipt := sampleJsonReceipts[index]
		expectedPoints := samplePoints[index]

		actualPoints := computePoints(sampleJsonReceipt)

		if expectedPoints != actualPoints {
			t.Errorf("For %s, computed points = %d, expected points = %d \n", sampleJsonPaths[index], actualPoints, expectedPoints)
		}
	}
}
