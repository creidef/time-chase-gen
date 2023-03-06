// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameExpense = "expense"

// Expense mapped from table <expense>
type Expense struct {
	ID      int32     `gorm:"column:id;not null" json:"id"`
	Date    time.Time `gorm:"column:date;not null" json:"date"`
	Comment string    `gorm:"column:comment;not null" json:"comment"`
	Cost    float64   `gorm:"column:cost;not null" json:"cost"`
	Amount  int32     `gorm:"column:amount;not null" json:"amount"`
	Project int32     `gorm:"column:project;not null" json:"project"`
}

// TableName Expense's table name
func (*Expense) TableName() string {
	return TableNameExpense
}