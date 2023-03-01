package service

import (
	"Project3Go/models"
	"time"
)

// NextPayDay takes three inputs: the user's pay day, the current day of the month, and the month of interest.
// It calculates the number of days until the next payday based on the current day and the pay day.
// Then, it returns a models.NextPayDayResponse struct that contains the number of days until the next payday and the date of the next payday.
// The function works by creating a time.Time object that represents the current date,
// then using that object to calculate the number of days until the next payday.
// Finally, it formats the date of the next payday as a string and returns the response.
// If the payDay is bigger than last day of current month, the payDay will take place on the last day of month.
func NextPayDay(payDay int, currentDay int, month time.Month) models.NextPayDayResponse {
	now := time.Date(time.Now().Year(), month, currentDay, 0, 0, 0, 0, time.Now().Location())
	response := models.NextPayDayResponse{}
	lastDayOfMonth := time.Date(time.Now().Year(), month+1, 0, 0, 0, 0, 0, time.Now().Location()).Day()

	if payDay > lastDayOfMonth { //if payday is bigger than last day of current month
		response.DaysUntil = lastDayOfMonth - currentDay
	} else { //if payday is smaller than last day of month
		if now.Day() <= payDay { //if pay day did not pass this month
			response.DaysUntil = payDay - currentDay
		} else if now.Day() > payDay { //if pay day passed this month
			response.DaysUntil = lastDayOfMonth - currentDay + payDay
		}
	}
	response.NextDate = now.AddDate(0, 0, response.DaysUntil).Format("02-01-2006")
	return response
}

// PayDayList calculates a list of pay days for the remaining months of the year based on the user's pay day and the current day of the month.
// It uses a loop to iterate over the remaining months of the year and appends the date of each pay day to a
// response of type models.PayDayListResponse The function calls the NextPayDay function to calculate the date for each month.
// Finally, the function returns the response struct containing the list of pay days.
func PayDayList(payDay int, currentDay int, month time.Month) models.PayDayListResponse {
	now := time.Now()
	currentMonth := int(month)
	result := NextPayDay(payDay, currentDay, month)
	var response models.PayDayListResponse
	var lastMonth int
	if payDay >= currentDay {
		lastMonth = 12
	} else {
		lastMonth = 11
	}

	previousMonth := currentMonth - 1

	for currentMonth <= lastMonth {
		newMonth := now.AddDate(0, currentMonth-previousMonth, 0)
		response.PayDays = append(response.PayDays, result.NextDate)
		result = NextPayDay(payDay, currentDay, newMonth.Month())
		currentMonth++
	}
	return response
}
