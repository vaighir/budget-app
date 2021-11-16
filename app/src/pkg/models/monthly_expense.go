package models

type MonthlyExpense struct {
	Id          int
	HouseholdId int
	Name        string
	Amount      float64
}
