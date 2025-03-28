package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

var (
	receiptStore = make(map[string]int)
	mutex        = sync.Mutex{}
)

func processReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	points := calculatePoints(receipt)

	mutex.Lock()
	receiptStore[id] = points
	mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func getPoints(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/receipts/")
	id = strings.TrimSuffix(id, "/points")

	mutex.Lock()
	points, exists := receiptStore[id]
	mutex.Unlock()

	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"points": points})
}

func calculatePoints(receipt Receipt) int {
	points := 0

	alphaNumRegex := regexp.MustCompile(`[a-zA-Z0-9]`)
	points += len(alphaNumRegex.FindAllString(receipt.Retailer, -1))

	totalFloat, _ := strconv.ParseFloat(receipt.Total, 64)
	if math.Mod(totalFloat, 1.0) == 0 {
		points += 50
	}

	if math.Mod(totalFloat, 0.25) == 0 {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5

	for _, item := range receipt.Items {
		trimmedDesc := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDesc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	dateParts := strings.Split(receipt.PurchaseDate, "-")
	day, _ := strconv.Atoi(dateParts[2])
	if day%2 != 0 {
		points += 6
	}

	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}

func main() {
	http.HandleFunc("/receipts/process", processReceipt)
	http.HandleFunc("/receipts/", getPoints)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
