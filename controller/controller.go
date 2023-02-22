package controller

import (
	"Project3Go/models"
	"Project3Go/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func HowMuchTillPayday(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	params, present := query["pay_day"]
	if !present || len(params) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		badResp, err := json.Marshal("Bad request")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Write(badResp)
		return
	}
	payDay, err := strconv.Atoi(params[0])
	if err != nil || payDay <= 0 || payDay > 31 {
		invalidPayDay, err := json.Marshal("Invalid pay day")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		http.Error(w, string(invalidPayDay), http.StatusBadRequest)
		return
	}

	result := service.TillSalary(payDay, time.Now().Month())
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)

}

func PayDayListDates(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	components := strings.Split(path, "/")
	if components[4] != "list-dates" || len(components) > 5 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	day := components[3]

	// Print the value to the response writer
	//fmt.Fprintf(w, "The value of 'day' is %s", day)
	payDay, err := strconv.Atoi(day)
	if err != nil || payDay <= 0 || payDay > 31 {
		http.Error(w, "Invalid pay day", http.StatusBadRequest)
		return
	}
	//fmt.Fprintf(w, "The value of 'payday' is %s", payDay)
	now := time.Now()
	_, month, _ := now.Date()
	currentMonth := int(month)
	//fmt.Printf("current month: %v", currentMonth)
	//if currentMonth == 12 && payDay>= currentDay{
	//	return
	//}

	result := service.TillSalary(payDay, month)

	var response models.PayDayResponse
	var lastMonth int

	if payDay >= time.Now().Day() {
		lastMonth = 12
	} else {
		lastMonth = 11
	}

	for currentMonth <= lastMonth {
		newMonth2 := now.AddDate(0, currentMonth-1, 0)
		response.PayDays = append(response.PayDays, result.NextDate)

		result = service.TillSalary(payDay, newMonth2.Month())

		currentMonth++
	}

	//response := make([]time.Time, 0, 12-currentMonth+1)
	//for currentMonth <= 12 {
	//	nextPayday := time.Date(year, time.Month(currentMonth), payDay, 0, 0, 0, 0, now.Location())
	//	response = append(response, nextPayday)
	//	currentMonth++
	//}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
