package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type SalaryResponse struct {
	DaysUntil int    `json:"days_until"`
	NextDate  string `json:"next_date"`
}

func howMuchTillPayday(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	params, present := query["pay_day"]
	if !present || len(params) == 0 {
		fmt.Println("filters not present")
		return
	}
	payDay, err := strconv.Atoi(params[0])
	if err != nil || payDay <= 0 || payDay > 31 {
		http.Error(w, "Invalid pay day", http.StatusBadRequest)
		return
	}

	result := tillSalary(payDay)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(result)
	if err != nil {
		log.Printf("Error happened in JSON marshal. Err: %s", err)
	}
	w.WriteHeader(200)
	w.Write(jsonResp)

}

func tillSalary(payDay int) SalaryResponse {
	now := time.Now()
	response := SalaryResponse{}
	year, month, _ := now.Date()
	lastDayOfMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
	fmt.Println(lastDayOfMonth)

	//if payDay> lastDayOfMonth{
	//	payDay = lastDayOfMonth
	//}
	if payDay > lastDayOfMonth {
		if now.Day() <= payDay {
			response.DaysUntil = lastDayOfMonth - now.Day()
			response.NextDate = now.AddDate(0, 0, response.DaysUntil).Format("02-01-2006")
			return response
		} else if now.Day() > payDay {
			response.DaysUntil = lastDayOfMonth - now.Day() + payDay
			response.NextDate = now.AddDate(0, 0, response.DaysUntil).Format("02-01-2006")
			return response
		}
	} else {
		if now.Day() <= payDay {
			response.DaysUntil = payDay - now.Day()
			response.NextDate = now.AddDate(0, 0, lastDayOfMonth).Format("02-01-2006")
			return response
		} else if now.Day() > payDay {
			response.DaysUntil = lastDayOfMonth - now.Day() + payDay
			response.NextDate = now.AddDate(0, 0, response.DaysUntil).Format("02-01-2006")
			return response
		}
	}
	return response
}

func main() {
	http.HandleFunc("/till-salary/how-much", howMuchTillPayday)
	//http.HandleFunc("/till-sallary/pay-day/", payDayListDates)
	http.ListenAndServe(":8080", nil)
}
