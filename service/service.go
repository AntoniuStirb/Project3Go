package service

import (
	"Project3Go/models"
	"time"
)

func TillSalary(payDay int, month time.Month) models.SalaryResponse {
	now := time.Date(time.Now().Year(), month, time.Now().Day(), 0, 0, 0, 0, time.Now().Location())
	response := models.SalaryResponse{}
	//inputMonth := month
	lastDayOfMonth := time.Date(time.Now().Year(), month+1, 0, 0, 0, 0, 0, time.UTC).Day()
	if payDay > lastDayOfMonth { //daca data de salariu este mai mare decat ultima zi a lunii
		if now.Day() <= payDay { //salariul va intra in ultima zi a lunii
			response.DaysUntil = lastDayOfMonth - now.Day()
			response.NextDate = now.AddDate(0, 0, response.DaysUntil).Format("02-01-2006")
			return response
		} else if now.Day() > payDay { //daca data de salariu este inainte de ultima zi a lunii
			response.DaysUntil = lastDayOfMonth - now.Day() + payDay
			response.NextDate = now.AddDate(0, 0, response.DaysUntil).Format("02-01-2006")
			return response
		}
	} else {
		if now.Day() <= payDay {
			response.DaysUntil = payDay - now.Day()
			response.NextDate = now.AddDate(0, 0, response.DaysUntil).Format("02-01-2006")
			return response
		} else if now.Day() > payDay {
			response.DaysUntil = lastDayOfMonth - now.Day() + payDay
			response.NextDate = now.AddDate(0, 0, response.DaysUntil).Format("02-01-2006")
			return response
		}
	}
	return response
}
