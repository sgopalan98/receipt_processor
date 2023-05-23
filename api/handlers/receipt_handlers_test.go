package handlers

import (
	"testing"

	"github.com/sgoplan98/receipt_processor/api/models"
)

//TODO: Is this the way to do?
func getSampleJsonPaths() []string {
	samples := []string{"samples/morning-receipt.json", "samples/simple-receipt.json", "samples/target.json", "samples/m-m.json"}
	return samples
}

// TODO: Is this the way to do?
func getSampleReceipts() []models.Receipt {
	samples := []models.Receipt{
		{
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
		{
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
		{
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
		{
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
	}
	return samples
}

// TODO: Compute manually
func getSamplePoints() []int {
	points := []int{15, 31, 28, 109}
	return points
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
