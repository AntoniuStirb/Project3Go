package controller

import (
	"Project3Go/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// HowMuchTillPayday is an HTTP request handler which accepts only GET requests that calculates the number of days until the next
// payday and the date based on the user's pay day and the current day of the month. It parses the pay_day query parameter from the URL,
// validates it, and calls the NextPayDay function from the service package to calculate the number of days until the next payday.
// Finally, it writes the response of type models.NextPayDayResponse as JSON to the HTTP response writer.
func HowMuchTillPayday(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Parse the query parameters.
	payDayStr := r.URL.Query().Get("pay_day")
	if payDayStr == "" {
		http.Error(w, "Missing pay_day parameter", http.StatusBadRequest)
		return
	}

	// Parse the payDay parameter.
	payDay, err := strconv.Atoi(payDayStr)
	if err != nil || payDay <= 0 || payDay > 31 {
		http.Error(w, "invalid pay_day parameter", http.StatusBadRequest)
		return
	}

	// Calculate the number of days until the next payday.
	currentMonth := time.Now().Month()
	currentDay := time.Now().Day()
	daysUntilPayday := service.NextPayDay(payDay, currentDay, currentMonth)

	// Write the response as JSON.
	jsonResponse, err := json.Marshal(daysUntilPayday)
	if err != nil {
		http.Error(w, "failed to encode response as JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// PayDayListDates is a handler for a HTTP endpoint that expects a URL path with the format
// "/payday/{day}/list-dates", where {day} is an integer representing the day of the month when the user gets the payday.
// The function only accepts GET requests, and returns a JSON response with a list of upcoming paydays based on the current day
// current month, and the user's pay day. If the URL path or pay day parameter are invalid,
// the function returns an appropriate HTTP error status code.
func PayDayListDates(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	path := r.URL.Path
	components := strings.Split(path, "/")
	if components[4] != "list-dates" || len(components) != 5 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	payDayStr := components[3]

	payDay, err := strconv.Atoi(payDayStr)
	if err != nil || payDay <= 0 || payDay > 31 {
		http.Error(w, "Invalid pay day", http.StatusBadRequest)
		return
	}

	currentMonth := time.Now().Month()
	currentDay := time.Now().Day()
	result := service.PayDayList(payDay, currentDay, currentMonth)

	jsonResponse, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
