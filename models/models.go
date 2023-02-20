package models

type SalaryDate struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type SalaryResponse struct {
	DaysUntil int        `json:"days_until"`
	NextDate  SalaryDate `json:"next_date"`
}

type PayDayResponse struct {
	PayDays []SalaryDate `json:"pay_days"`
}
