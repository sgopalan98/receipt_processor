package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/sgoplan98/receipt_processor/api/models"
)

// TODO: Is this the right way to do? When is this created?
var receiptPointsMapping = make(map[string]int)

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
			// TODO: Error handling
			price, _ := strconv.ParseFloat(item.Price, 64)
			//TODO: Problem statement - wrong? Round != Ceil. Notify
			roundedPoints := int(math.Ceil(price * 0.2))
			points += roundedPoints
		}
	}

	return points
}

func datePoints(receipt models.Receipt) int {
	// TODO: I'm blindly using 2 index here. How else should I do it?
	// TODO: do you need string. then, convert to int?
	dateString := strings.Split(receipt.PurchaseDate, "-")[2]
	// TODO: Handle error?
	date, _ := strconv.Atoi(dateString)

	if date%2 == 1 {
		return 6
	}
	return 0
}

func timePoints(receipt models.Receipt) int {
	timeString := receipt.PurchaseTime
	// TODO: 0 and 1 hardcoded.
	hourString := strings.Split(timeString, ":")[0]
	minuteString := strings.Split(timeString, ":")[1]

	// TODO: Error handling
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
	// TODO: Check if you need to make functions for every point calculcation
	// TODO: Check the datatype of total points
	totalPoints := 0

	// Add a point for each of the alphanum character in retailer name
	retailerNamePoints := computeRetailerNamePoints(receipt.Retailer)
	totalPoints += retailerNamePoints

	// Add 50 points if Receipt total is round
	// TODO: Check if any other type is possible?
	// TODO: Check if I can convert the types when JSON is parsed?
	receiptTotal, _ := strconv.ParseFloat(receipt.Total, 64)
	isRound := math.Trunc(receiptTotal) == receiptTotal
	if isRound {
		totalPoints += 50
	}

	// Add 25 points if total is a multiple of 25 cents
	// TODO: Check if this is variable naming is correct
	is25Multiple := math.Trunc(4*receiptTotal) == 4*receiptTotal
	if is25Multiple {
		totalPoints += 25
	}

	// Add 5 points for every two items in the receipt
	// TODO: name better
	// TODO: Should I calculate length or can I use len() directly?
	receiptItemsLength := len(receipt.Items)
	numberOf2s := receiptItemsLength / 2
	totalPoints += numberOf2s * 5

	// Add points based off of description
	// TODO: TODOs inside the function
	totalPoints += descriptionPoints(receipt)

	// Add points based off of date
	// TODO: do you need a function for this?
	totalPoints += datePoints(receipt)

	// Add points based off of time
	// TODO: do you need a function for this?
	totalPoints += timePoints(receipt)

	return totalPoints
}

// TODO: All errors should be handled
func ReceiptsProcessHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("receipts process handler called")
	if r.Method != "POST" {
		http.NotFound(w, r)
	}
	reqBodyBytes, err := ioutil.ReadAll(r.Body)

	// TODO: Error handling done to be better
	if err != nil {
		fmt.Println("Error reading from request body:", err)
		return
	}

	// Convert bytes to string
	receiptJson := string(reqBodyBytes)
	receipt := models.ConvertJsonToRecept(receiptJson)
	points := computePoints(receipt)

	// Generate a UUID
	uuidObj, _ := uuid.NewRandom()
	// Convert the UUID to an alphanumeric string
	id := uuidObj.String()

	receiptPointsMapping[id] = points

	// Generate response and send response
	response := models.ProcessResponse{
		Id: id,
	}

	// Convert Response to JSON
	jsonData, _ := json.Marshal(response)

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write the JSON data to the response body
	fmt.Fprint(w, string(jsonData))
}

// TODO: all errors should be handled
func ReceiptsPointsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("receipts points handler called")
	if r.Method != "GET" {
		http.NotFound(w, r)
		// TODO: Do you need return?
		return
	}

	// Get the path from the request URL
	path := r.URL.Path

	// Extract the id from the path
	parts := strings.Split(path, "/")

	// Checking if URL is correct
	if parts[1] != "receipts" {
		http.NotFound(w, r)
		return
	}

	if parts[3] != "points" {
		http.NotFound(w, r)
		return
	}
	id := parts[2]

	points, ok := receiptPointsMapping[id]

	if !ok {
		fmt.Printf("ID %s not found\n", id)
		http.NotFound(w, r)
		return
	}

	pointsReponse := models.PointsResponse{
		Points: strconv.Itoa(points),
	}

	// Convert Response to JSON
	jsonData, _ := json.Marshal(pointsReponse)

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write the JSON data to the response body
	fmt.Fprint(w, string(jsonData))
}
