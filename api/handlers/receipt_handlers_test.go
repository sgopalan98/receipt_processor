package handlers

import (
	"testing"

	"github.com/sgoplan98/receipt_processor/api/models"
)

var MockData = []struct {
	Name        string
	ReceiptData models.Receipt
	Points      int
}{
	{
		Name: "morning-receipt",
		ReceiptData: models.Receipt{
			Retailer:     "Walgreens",
			PurchaseDate: "2022-01-02",
			PurchaseTime: "08:13",
			Total:        2.65,
			Items: []models.Item{
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
		Points: 15,
	},
	{
		Name: "simple-receipt",
		ReceiptData: models.Receipt{
			Retailer:     "Target",
			PurchaseDate: "2022-01-02",
			PurchaseTime: "13:13",
			Total:        1.25,
			Items: []models.Item{
				{
					ShortDescription: "Pepsi - 12-oz",
					Price:            1.25,
				},
			},
		},
		Points: 31,
	},
	{
		Name: "target",
		ReceiptData: models.Receipt{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Total:        35.35,
			Items: []models.Item{
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
		Points: 28,
	},
	{
		Name: "m-m",
		ReceiptData: models.Receipt{
			Retailer:     "M&M Corner Market",
			PurchaseDate: "2022-03-20",
			PurchaseTime: "14:33",
			Total:        9.00,
			Items: []models.Item{
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
		Points: 109,
	},
}

func TestComputePoints(t *testing.T) {
	mockData := MockData
	for index := 0; index < len(mockData); index++ {
		sampleReceipt := mockData[index].ReceiptData
		expectedPoints := mockData[index].Points
		name := mockData[index].Name

		actualPoints := computePoints(sampleReceipt)

		if expectedPoints != actualPoints {
			t.Errorf("For %s, computed points = %d, expected points = %d \n", name, actualPoints, expectedPoints)
		}
	}
}
