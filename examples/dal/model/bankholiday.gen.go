// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameBankholiday = "bankholiday"

// Bankholiday mapped from table <bankholiday>
type Bankholiday struct {
	ID   int32     `gorm:"column:id;primaryKey" json:"id"`
	Name string    `gorm:"column:name;not null" json:"name"`
	Date time.Time `gorm:"column:date;not null" json:"date"`
	Year int32     `gorm:"column:year;not null" json:"year"`
}

// TableName Bankholiday's table name
func (*Bankholiday) TableName() string {
	return TableNameBankholiday
}