# Receipt Processor 

My submission for the Receipt Processor challenge as part of Backend Engineer application process.

For more information about the challenge, please refer to the following link: [https://github.com/fetch-rewards/receipt-processor-challenge](https://github.com/fetch-rewards/receipt-processor-challenge)

## Brief intro

This is a HTTP web-service that serves two end points:
1. /receipts/process - POST request end point that receives a JSON with the details of a receipt. End point uses this JSON, computes and stores the points for the receipt. It returns a JSON with an ID for the receipt.
2. /receipts/%id%/points - GET request end point that returns the points for the receipt with the ID passed.

## Things to install before running

* [Docker!](https://docs.docker.com/get-docker/)
* cURL (or any end point testing tool)
* Go - (if you want to run tests)

## How to build?

    docker build -t receipt-processor-image .

## How to run?

 Step 1: `docker run -p 8080:8080 receipt-processor-image`
 Step 2: Open new terminal
 Step 3: Run cURL requests (with a JSON that you want to test) 


Example curl commands: 

1. curl -X POST -H "Content-Type: application/json" -d '{
		    "retailer": "Walgreens",
		    "purchaseDate": "2022-01-02",
		    "purchaseTime": "08:13",
		    "total": "2.65",
		    "items": [
		        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
		        {"shortDescription": "Dasani", "price": "1.40"}
		    ]
		}' http://localhost:8080/receipts/process
2. Store the ID returned somewhere safe.
3. curl -X GET http://localhost:8080/receipts/ID/points (replace 'ID' with the ID stored)

## How to test?

* To run test files for the project, run

 `go test ./...`


## Tech stack

- Docker
- Golang

## Implementation

- Has a ID->Points Map data structure
- 2 end points are exposed by net/http Go module
- POST '/receipts/process' processes the receipt as JSON in the request.
	- Compute points based off the data in the JSON
	- Generates a unique ID for the Receipt
	- Stores the ID -> Points mapping in the map mentioned above
- GET '/receipts/%id%/points' returns the points for ID in the request
	- Retrieves the points associated with the ID using the map mentioned above

## Project structure & details

```
.
├── Dockerfile (contains docker commands)
├── README.md (contains README)
├── api (contains API code)
│   ├── handlers
│   │   ├── receipt_handlers.go (Handlers' endpoints)
│   │   └── receipt_handlers_test.go (Test code for all handlers)
│   └── models (contains all models used in the API)
│       ├── http_models.go (models for HTTP responses)
│       ├── receipt_models.go (models used in receipt processing)
│       ├── receipt_models_test.go
│       └── samples (sample data for test files)
│           ├── complex.json
│           └── simple.json
├── go.mod
├── go.sum
├── main.go (start point for the program)
```


### Comments 


- My first time using Go; I apologize if code looks unclean
- Some resources I found useful when learning:
	- [go.dev learning tutorials](https://go.dev/learn/)
	- [50 Shades of Go](https://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html)
	- ChatGPT (ofcourse!)
	- Please feel free to comment on anything!







