package service

import (
	"Project3Go/models"
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

func TestTillSalary(t *testing.T) {
	testCases := []struct {
		name       string
		payDay     int
		currentDay int
		month      time.Month
		expected   models.NextPayDayResponse
	}{
		{
			name:       "Test1 - When the payday is bigger than current day and payday is bigger than last day of month",
			payDay:     31,
			currentDay: 24,
			month:      time.Month(2),
			expected: models.NextPayDayResponse{
				DaysUntil: 4,
				NextDate:  time.Date(time.Now().Year(), time.Month(2), 28, 0, 0, 0, 0, time.Now().Location()).Format("02-01-2006"),
			},
		},
		{
			name:       "Test2 - When the payday is bigger than current day and payday is smaller than last day of month",
			payDay:     27,
			currentDay: 24,
			month:      time.Month(2),
			expected: models.NextPayDayResponse{
				DaysUntil: 3,
				NextDate:  time.Date(time.Now().Year(), time.Month(2), 27, 0, 0, 0, 0, time.Now().Location()).Format("02-01-2006"),
			},
		},
		{
			name:       "Test3 - When the payday is smaller than current day and payday is smaller than last day of month",
			payDay:     12,
			currentDay: 24,
			month:      time.Month(2),
			expected: models.NextPayDayResponse{
				DaysUntil: 16,
				NextDate:  time.Date(time.Now().Year(), time.Month(3), 12, 0, 0, 0, 0, time.Now().Location()).Format("02-01-2006"),
			},
		},
		{
			name:       "Test4 - When the payday is equal to the current day",
			payDay:     15,
			currentDay: 15,
			month:      time.Month(2),
			expected: models.NextPayDayResponse{
				DaysUntil: 0,
				NextDate:  time.Date(time.Now().Year(), time.Month(2), 15, 0, 0, 0, 0, time.Now().Location()).Format("02-01-2006"),
			},
		},
		{
			name:       "Test5 - When we are in last month of year and paydays is in first month of next year",
			payDay:     15,
			currentDay: 20,
			month:      time.Month(12),
			expected: models.NextPayDayResponse{
				DaysUntil: 26,
				NextDate:  time.Date(time.Now().Year()+1, time.Month(1), 15, 0, 0, 0, 0, time.Now().Location()).Format("02-01-2006"),
			},
		},
	}

	for _, tests := range testCases {
		actual := NextPayDay(tests.payDay, tests.currentDay, tests.month)
		if diff := cmp.Diff(actual, tests.expected); diff != "" {
			t.Errorf("NextPayDay(%v, %v) = %v; expected %v", tests.payDay, tests.month, actual, tests.expected)
		}
	}
}

func TestPayDayList(t *testing.T) {
	testCases := []struct {
		name       string
		payDay     int
		currentDay int
		month      time.Month
		expected   models.PayDayListResponse
	}{
		{
			name:       "Test1 - When the payday is bigger than current day and payday is bigger than last day of month",
			payDay:     31,
			currentDay: 25,
			month:      time.Month(2),
			expected: models.PayDayListResponse{
				PayDays: []string{
					"28-02-2023", "31-03-2023", "30-04-2023", "31-05-2023", "30-06-2023", "31-07-2023",
					"31-08-2023", "30-09-2023", "31-10-2023", "30-11-2023", "31-12-2023",
				},
			},
		},
		{
			name:       "Test2 - When the payday is bigger than current day and payday is smaller than last day of month",
			payDay:     27,
			currentDay: 25,
			month:      time.Month(4),
			expected: models.PayDayListResponse{
				PayDays: []string{
					"27-04-2023", "27-05-2023", "27-06-2023", "27-07-2023", "27-08-2023",
					"27-09-2023", "27-10-2023", "27-11-2023", "27-12-2023",
				},
			},
		},
		{
			name:       "Test3 - When is the last month and the payday already passed this month, no more paydays this year",
			payDay:     15,
			currentDay: 25,
			month:      time.Month(12),
			expected: models.PayDayListResponse{
				PayDays: nil,
			},
		},
		{
			name:       "Test4 - When payday is equal to current day",
			payDay:     15,
			currentDay: 15,
			month:      time.Month(8),
			expected: models.PayDayListResponse{
				PayDays: []string{
					"15-08-2023", "15-09-2023", "15-10-2023", "15-11-2023", "15-12-2023",
				},
			},
		},
		{
			name:       "Test4 - When payday is smaller than current day",
			payDay:     15,
			currentDay: 19,
			month:      time.Month(10),
			expected: models.PayDayListResponse{
				PayDays: []string{
					"15-11-2023", "15-12-2023",
				},
			},
		},
	}

	for _, tests := range testCases {
		actual := PayDayList(tests.payDay, tests.currentDay, tests.month)
		if diff := cmp.Diff(actual, tests.expected); diff != "" {
			t.Errorf("\nPayDayList(%v, %v, %v) = %v; \nexpected %v", tests.payDay, tests.currentDay, tests.month,
				actual, tests.expected)
		}
	}
}
