package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/sgoplan98/receipt_processor/api/models"
)

var receiptPointsMap sync.Map

// computeRetailerNamePoints -  Compute points based on Retailer Name
//
// Parameters:
//   - retailerName: Name of the retailer in the receipt.
//
// Return:
//   - Points awarded based on retailer name.
func computeRetailerNamePoints(retailerName string) int {
	count := 0
	// Add a point for each of the alphanum character in retailer name
	for _, char := range retailerName {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			count++
		}
	}
	return count
}

// computeDescriptionPoints -  Compute points based on Item descriptions
//
// Parameters:
//   - items: items in the receipt which have descriptions
//
// Return:
//   - Points awarded based on item descriptions.
func computeDescriptionPoints(items []models.Item) int {
	points := 0
	for _, item := range items {
		trimmedItemDescription := strings.TrimSpace(item.ShortDescription)
		trimmedItemDescriptionLength := len(trimmedItemDescription)
		if trimmedItemDescriptionLength%3 == 0 {
			//TODO: Problem statement - wrong? Round != Ceil. Notify
			roundedPoints := int(math.Ceil(item.Price * 0.2))
			points += roundedPoints
		}
	}

	return points
}

// computeDateTimePoints -  Compute points based on date time
//
// Parameters:
//   - Receipt: the receipt data which has date & time
//
// Return:
//   - Points awarded based on date and time.
func computeDateTimePoints(receipt models.Receipt) int {
	dateTimePoints := 0
	dateString := strings.Split(receipt.PurchaseDate, "-")[2]
	date, _ := strconv.Atoi(dateString)

	// Add 6 points if the day in the purchase date is odd
	if date%2 == 1 {
		dateTimePoints += 6
	}

	timeString := receipt.PurchaseTime
	hourString := strings.Split(timeString, ":")[0]
	minuteString := strings.Split(timeString, ":")[1]

	hour, _ := strconv.Atoi(hourString)
	minute, _ := strconv.Atoi(minuteString)

	// Add 10 points if time of purchase is after 2pm and before 4pm
	if (hour == 14 && minute > 0) || hour == 15 {
		dateTimePoints += 10
	}

	return dateTimePoints
}

// computeTotalCostPoints -  Compute points based on total cost
//
// Parameters:
//   - receiptTotal: the total cost for the receipt
//
// Return:
//   - Points awarded based on the total cost of the receipt
func computeTotalCostPoints(receiptTotal float64) int {
	totalCostPoints := 0

	// Add 50 points if Receipt total is round
	isRound := math.Trunc(receiptTotal) == receiptTotal
	if isRound {
		totalCostPoints += 50
	}

	// Add 25 points if total is a multiple of 25 cents
	is25Multiple := math.Trunc(4*receiptTotal) == 4*receiptTotal
	if is25Multiple {
		totalCostPoints += 25
	}

	return totalCostPoints
}

// computeTotalCostPoints -  Compute total points for the receipt
//
// Parameters:
//   - receipt : Receipt data
//
// Return:
//   - Points awarded based on the data in receipt
func computePoints(receipt models.Receipt) int {
	totalPoints := 0

	// Add points based on retailter name
	totalPoints += computeRetailerNamePoints(receipt.Retailer)

	// Add points based off total cost in receipt
	totalPoints += computeTotalCostPoints(receipt.Total)

	// Add 5 points for every two items in the receipt
	totalPoints += (len(receipt.Items) / 2) * 5

	// Add points based on description
	totalPoints += computeDescriptionPoints(receipt.Items)

	// Add points based on date and time
	totalPoints += computeDateTimePoints(receipt)

	return totalPoints
}

// ReceiptsProcessHandler -  Handler to process the data of a receipt and generate unique ID for the receipt
//
// Parameters:
//   - w : http.ResponseWriter
//   - r : *http.Request
//
// Return:
//   - ID generated for the receipt
func ReceiptsProcessHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("receipts process handler called")
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	reqBodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading from request body:", err)
		http.Error(w, "Error reading the request body", http.StatusBadRequest)
		return
	}

	// Convert bytes to string
	receiptJson := string(reqBodyBytes)
	receipt, err := models.ConvertJsonToReceipt(receiptJson)
	if err != nil {
		fmt.Println("Invalid data: ", err)
		http.Error(w, "Error in JSON format", http.StatusBadRequest)
		return
	}

	points := computePoints(receipt)

	// Generate a UUID string - ID for the receipt
	uuidObj, _ := uuid.NewRandom()
	id := uuidObj.String()
	receiptPointsMap.Store(id, strconv.Itoa(points))

	// Generate response and send response
	response := models.ProcessResponse{
		Id: id,
	}
	jsonData, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonData))
}

// ReceiptsProcessHandler -  Handler to get the data for the ID present in the URL
//
// Parameters:
//   - w : http.ResponseWriter
//   - r : *http.Request
//
// Return:
//   - Points stored for the Receipt ID
func ReceiptsPointsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("receipts points handler called")
	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}

	// Get the path from the request URL
	path := r.URL.Path
	// Extract the id from the path & doing validation
	parts := strings.Split(path, "/")
	if parts[1] != "receipts" {
		http.NotFound(w, r)
		return
	}
	if parts[3] != "points" {
		http.NotFound(w, r)
		return
	}
	id := parts[2]

	points, ok := receiptPointsMap.Load(id)
	if !ok {
		fmt.Printf("ID %s not found\n", id)
		http.NotFound(w, r)
		return
	}

	// Generate response and send response
	pointsReponse := models.PointsResponse{
		Points: points.(string),
	}
	jsonData, _ := json.Marshal(pointsReponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonData))
}
