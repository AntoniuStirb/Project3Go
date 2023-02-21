package models

type SalaryResponse struct {
	DaysUntil int    `json:"days_until"`
	NextDate  string `json:"next_date"`
}

type PayDayResponse struct {
	PayDays []string `json:"pay_days"`
}
