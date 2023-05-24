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

// TODO: Is this the right way to do? When is this created?
var receiptPointsMapping sync.Map

// Compute points based off Retailer Name
func computeRetailerNamePoints(retailerName string) int {
	count := 0
	for _, char := range retailerName {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			count++
		}
	}
	return count

}

func descriptionPoints(receipt models.Receipt) int {
	items := receipt.Items
	points := 0
	for _, item := range items {
		trimmedItemDescription := strings.TrimSpace(item.ShortDescription)
		// TODO: DO you actually need these variables or can you just use len straight forward?
		trimmedItemDescriptionLength := len(trimmedItemDescription)
		if trimmedItemDescriptionLength%3 == 0 {
			//TODO: Problem statement - wrong? Round != Ceil. Notify
			roundedPoints := int(math.Ceil(item.Price * 0.2))
			points += roundedPoints
		}
	}

	return points
}

func datePoints(receipt models.Receipt) int {
	dateString := strings.Split(receipt.PurchaseDate, "-")[2]
	date, _ := strconv.Atoi(dateString)

	if date%2 == 1 {
		return 6
	}
	return 0
}

func timePoints(receipt models.Receipt) int {
	timeString := receipt.PurchaseTime
	hourString := strings.Split(timeString, ":")[0]
	minuteString := strings.Split(timeString, ":")[1]

	hour, _ := strconv.Atoi(hourString)
	minute, _ := strconv.Atoi(minuteString)

	// TODO: Write logic better?
	if hour == 14 && minute > 0 {
		return 10
	}

	if hour == 15 {
		return 10
	}
	return 0
}

func computePoints(receipt models.Receipt) int {
	totalPoints := 0

	// Add a point for each of the alphanum character in retailer name
	retailerNamePoints := computeRetailerNamePoints(receipt.Retailer)
	totalPoints += retailerNamePoints

	// Add 50 points if Receipt total is round
	receiptTotal := receipt.Total
	isRound := math.Trunc(receiptTotal) == receiptTotal
	if isRound {
		totalPoints += 50
	}

	// Add 25 points if total is a multiple of 25 cents
	is25Multiple := math.Trunc(4*receiptTotal) == 4*receiptTotal
	if is25Multiple {
		totalPoints += 25
	}

	// Add 5 points for every two items in the receipt
	// TODO: Should I calculate length or can I use len() directly?
	receiptItemsLength := len(receipt.Items)
	numberOf2s := receiptItemsLength / 2
	totalPoints += numberOf2s * 5

	// Add points based on description
	totalPoints += descriptionPoints(receipt)

	// Add points based on date
	// TODO: do you need a function for this?
	totalPoints += datePoints(receipt)

	// Add points based on time
	// TODO: do you need a function for this?
	totalPoints += timePoints(receipt)

	return totalPoints
}

// TODO: All errors should be handled
func ReceiptsProcessHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("receipts process handler called")
	if r.Method != "POST" {
		http.NotFound(w, r)
		// TODO: Should you return after not found? Change in all if required
		return
	}
	reqBodyBytes, err := ioutil.ReadAll(r.Body)

	// TODO: Error handling done to be better - better text
	if err != nil {
		fmt.Println("Error reading from request body:", err)
		http.Error(w, "Error reading the request body", http.StatusBadRequest)
		return
	}

	// Convert bytes to string
	receiptJson := string(reqBodyBytes)
	receipt, err := models.ConvertJsonToRecept(receiptJson)
	if err != nil {
		fmt.Println("Invalid data: ", err)
		http.Error(w, "Error reading the request body", http.StatusBadRequest)
		return
	}

	points := computePoints(receipt)

	// Generate a UUID string
	uuidObj, _ := uuid.NewRandom()
	id := uuidObj.String()

	receiptPointsMapping.Store(id, strconv.Itoa(points))

	// Generate response and send response
	response := models.ProcessResponse{
		Id: id,
	}

	// Convert Response to JSON
	jsonData, _ := json.Marshal(response)

	// Set the response headers and write the json data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonData))
}

// TODO: all errors should be handled
func ReceiptsPointsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("receipts points handler called")
	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}

	// Get the path from the request URL
	path := r.URL.Path

	// Extract the id from the path
	parts := strings.Split(path, "/")

	// Checking if URL is correct
	// SHould you check for more than 4? - Error?
	if parts[1] != "receipts" {
		http.NotFound(w, r)
		return
	}
	if parts[3] != "points" {
		http.NotFound(w, r)
		return
	}
	id := parts[2]

	points, ok := receiptPointsMapping.Load(id)

	if !ok {
		fmt.Printf("ID %s not found\n", id)
		http.NotFound(w, r)
		return
	}

	pointsReponse := models.PointsResponse{
		Points: points.(string),
	}

	// Convert Response to JSON
	jsonData, _ := json.Marshal(pointsReponse)

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write the JSON data to the response body
	fmt.Fprint(w, string(jsonData))
}
