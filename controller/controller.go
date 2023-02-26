package controller

import (
	"Project3Go/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func HowMuchTillPayday(w http.ResponseWriter, r *http.Request) {
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
	daysUntilPayday := service.TillSalary(payDay, currentDay, currentMonth)

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

func PayDayListDates(w http.ResponseWriter, r *http.Request) {
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
