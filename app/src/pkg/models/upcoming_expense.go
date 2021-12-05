package models

import "time"

type UpcomingExpense struct {
	Id          int
	HouseholdId int
	Name        string
	Amount      float64
	Deadline    time.Time
}
