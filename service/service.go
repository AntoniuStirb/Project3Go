package service

import (
	"Project3Go/models"
	"time"
)

func TillSalary(payDay int, currentDay int, month time.Month) models.NextPayDayResponse {
	now := time.Date(time.Now().Year(), month, currentDay, 0, 0, 0, 0, time.Now().Location())
	response := models.NextPayDayResponse{}
	lastDayOfMonth := time.Date(time.Now().Year(), month+1, 0, 0, 0, 0, 0, time.UTC).Day()

	//if payDay <= lastDayOfMonth {
	//	payDayTime := time.Date(time.Now().Year(), month, payDay, 0, 0, 0, 0, time.Now().Location())
	//	if isWeekend(payDayTime) {
	//		fmt.Printf("AAAA: %v\n", payDayTime)
	//		payDayTime = previousFriday(payDayTime)
	//	}
	//	payDay = payDayTime.Day()
	//} else {
	//	payDayTime := time.Date(time.Now().Year(), month, lastDayOfMonth, 0, 0, 0, 0, time.Now().Location())
	//	if isWeekend(payDayTime) {
	//		fmt.Printf("BBBBB: %v\n", payDayTime)
	//		payDayTime = previousFriday(payDayTime)
	//	}
	//	payDay = payDayTime.Day()
	//}

	if payDay > lastDayOfMonth { //daca data de salariu este mai mare decat ultima zi a lunii
		if now.Day() <= payDay { //daca data de salariu nu a trecut inca in luna curenta
			response.DaysUntil = lastDayOfMonth - currentDay
		} else if now.Day() > payDay { //daca data de salariu a trecut luna aceasta
			response.DaysUntil = lastDayOfMonth - currentDay + payDay
		}
	} else { //daca data de salariu nu este mai mare decat ultima zi a lunii
		if now.Day() <= payDay { //daca data de salariu nu a trecut inca in luna curenta
			response.DaysUntil = payDay - currentDay
		} else if now.Day() > payDay { //daca data de salariu a trecut luna aceasta
			response.DaysUntil = lastDayOfMonth - currentDay + payDay
		}
	}
	response.NextDate = now.AddDate(0, 0, response.DaysUntil).Format("02-01-2006")
	return response
}

func PayDayList(payDay int, currentDay int, month time.Month) models.PayDayListResponse {
	now := time.Now()
	//_, month, currentDay := now.Date()
	currentMonth := int(month)
	result := TillSalary(payDay, currentDay, month)
	var response models.PayDayListResponse
	var lastMonth int

	if payDay >= currentDay {
		lastMonth = 12
	} else {
		lastMonth = 11
	}

	for currentMonth <= lastMonth {
		newMonth2 := now.AddDate(0, currentMonth-1, 0)
		response.PayDays = append(response.PayDays, result.NextDate)

		result = TillSalary(payDay, currentDay, newMonth2.Month())

		currentMonth++
	}
	return response
}

//func isWeekend(date time.Time) bool {
//	return date.Weekday() == time.Saturday || date.Weekday() == time.Sunday
//}
//
//func previousFriday(date time.Time) time.Time {
//	daysSinceFriday := int(date.Weekday() - time.Friday)
//	if daysSinceFriday < 0 {
//		daysSinceFriday += 7
//	}
//	fridayDate := date.AddDate(0, 0, -daysSinceFriday)
//	return fridayDate
//}
