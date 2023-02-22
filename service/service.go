package service

import (
	"Project3Go/models"
	"fmt"
	"time"
)

func TillSalary(payDay int) models.SalaryResponse {
	now := time.Now()
	response := models.SalaryResponse{}
	year, month, _ := now.Date()
	lastDayOfMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
	fmt.Println(lastDayOfMonth)

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
