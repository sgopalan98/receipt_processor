package handlers

import (
	"testing"

	"github.com/sgoplan98/receipt_processor/api/models"
)

var MockComputePointsData = []struct {
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

var MockComputeRetailerNamePointsData = []struct {
	RetailerName string
	Points       int
}{
	{
		RetailerName: "",
		Points:       0,
	},
	{
		RetailerName: "    a",
		Points:       1,
	},
	{
		RetailerName: "abc",
		Points:       3,
	},
	{
		RetailerName: "123",
		Points:       3,
	},
	{
		RetailerName: "abc123",
		Points:       6,
	},
}

var MockComputeDescriptionPointsData = []struct {
	items  []models.Item
	Points int
}{
	{
		items: []models.Item{
			{
				ShortDescription: "",
				Price:            3.00,
			},
		},
		Points: 1,
	},
	{
		items: []models.Item{
			{
				ShortDescription: "A",
				Price:            3.00,
			},
		},
		Points: 0,
	},
	{
		items: []models.Item{
			{
				ShortDescription: "ABC",
				Price:            3.00,
			},
		},
		Points: 1,
	},
	{
		items: []models.Item{
			{
				ShortDescription: "    A",
				Price:            3.00,
			},
		},
		Points: 0,
	},
}

var MockComputeTotalCostPointsData = []struct {
	TotalCost float64
	Points    int
}{
	{
		TotalCost: 0,
		Points:    75,
	},
	{
		TotalCost: 0.10,
		Points:    0,
	},
	{
		TotalCost: 2,
		Points:    75,
	},
	{
		TotalCost: 0.25,
		Points:    25,
	},
	{
		TotalCost: 0.75,
		Points:    25,
	},
}

var MockComputeDateTimePointsData = []struct {
	ReceiptData models.Receipt
	Points      int
}{
	{
		ReceiptData: models.Receipt{
			Retailer:     "Target",
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
		Points: 0,
	},
	{
		ReceiptData: models.Receipt{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "14:00",
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
		Points: 6,
	},
	{
		ReceiptData: models.Receipt{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "14:13",
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
		Points: 16,
	},
	{
		ReceiptData: models.Receipt{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "15:13",
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
		Points: 16,
	},
}

func TestComputePoints(t *testing.T) {
	mockData := MockComputePointsData
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

func TestComputeRetailerNamePoints(t *testing.T) {
	mockData := MockComputeRetailerNamePointsData
	for index := 0; index < len(mockData); index++ {
		sampleName := mockData[index].RetailerName
		expectedPoints := mockData[index].Points

		actualPoints := computeRetailerNamePoints(sampleName)

		if expectedPoints != actualPoints {
			t.Errorf("For %s, computed points = %d, expected points = %d \n", sampleName, actualPoints, expectedPoints)
		}
	}
}

func TestComputeDescriptionNamePoints(t *testing.T) {
	mockData := MockComputeDescriptionPointsData
	for index := 0; index < len(mockData); index++ {
		sampleItems := mockData[index].items
		expectedPoints := mockData[index].Points

		actualPoints := computeDescriptionPoints(sampleItems)

		if expectedPoints != actualPoints {
			t.Errorf("For %d, computed points = %d, expected points = %d \n", index, actualPoints, expectedPoints)
		}
	}
}

func TestComputeTotalCostPoints(t *testing.T) {
	mockData := MockComputeTotalCostPointsData
	for index := 0; index < len(mockData); index++ {
		sampleTotal := mockData[index].TotalCost
		expectedPoints := mockData[index].Points

		actualPoints := computeTotalCostPoints(sampleTotal)

		if expectedPoints != actualPoints {
			t.Errorf("For %f, computed points = %d, expected points = %d \n", sampleTotal, actualPoints, expectedPoints)
		}
	}
}

func TestComputeDateTimePoints(t *testing.T) {
	mockData := MockComputeDateTimePointsData
	for index := 0; index < len(mockData); index++ {
		receipt := mockData[index].ReceiptData
		expectedPoints := mockData[index].Points

		actualPoints := computeDateTimePoints(receipt)

		if expectedPoints != actualPoints {
			t.Errorf("For %d, computed points = %d, expected points = %d \n", index, actualPoints, expectedPoints)
		}
	}
}
