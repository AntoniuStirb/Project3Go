package main

import (
	"Project3Go/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/till-salary/how-much", controller.HowMuchTillPayday)
	http.HandleFunc("/till-salary/pay-day/", controller.PayDayListDates)
	http.ListenAndServe(":8080", nil)
}
