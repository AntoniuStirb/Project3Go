package controller

import (
	"Project3Go/models"
	"Project3Go/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func PayDayListDates(w http.ResponseWriter, r *http.Request) {
	payDayStr := r.URL.Path[len("/till-sallary/pay-day/"):]
	payDay, err := strconv.Atoi(payDayStr)
	if err != nil || payDay <= 0 || payDay > 31 {
		http.Error(w, "Invalid pay day", http.StatusBadRequest)
		return
	}

	now := time.Now()
	_, month, _ := now.Date()
	currentMonth := int(month)
	//if currentMonth == 12 && payDay>= currentDay{
	//	return
	//}

	result := service.TillSalary(payDay)
	//nextPayday:= result.NextDate

	var response models.PayDayResponse
	test := response.PayDays
	for currentMonth <= 12 {
		nextPayDay := result.NextDate
		nextPayDayInt, err1 := strconv.Atoi(nextPayDay)
		fmt.Println(nextPayDayInt)
		if err1 != nil || payDay <= 0 || payDay > 31 {
			http.Error(w, "Invalid pay day", http.StatusBadRequest)
			return
		}
		result = service.TillSalary(payDay)
		test = append(test, nextPayDay)
		currentMonth++
	}

	//response := make([]time.Time, 0, 12-currentMonth+1)
	//for currentMonth <= 12 {
	//	nextPayday := time.Date(year, time.Month(currentMonth), payDay, 0, 0, 0, 0, now.Location())
	//	response = append(response, nextPayday)
	//	currentMonth++
	//}

	jsonResponse, err := json.Marshal(test)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

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

	result := service.TillSalary(payDay)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)

}
