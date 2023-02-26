package models

type NextPayDayResponse struct {
	DaysUntil int    `json:"days_until"`
	NextDate  string `json:"next_date"`
}

type PayDayListResponse struct {
	PayDays []string `json:"pay_days"`
}
