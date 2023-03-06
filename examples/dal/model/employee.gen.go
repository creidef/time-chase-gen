// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameEmployee = "employee"

// Employee mapped from table <employee>
type Employee struct {
	ID           int32     `gorm:"column:id;not null" json:"id"`
	Firstname    string    `gorm:"column:firstname;not null" json:"firstname"`
	Lastname     string    `gorm:"column:lastname;not null" json:"lastname"`
	Birthday     time.Time `gorm:"column:birthday;not null" json:"birthday"`
	Workload     int32     `gorm:"column:workload;not null" json:"workload"`
	Workhours    int32     `gorm:"column:workhours;not null" json:"workhours"`
	Holidays     int32     `gorm:"column:holidays;not null" json:"holidays"`
	Holidaysused int32     `gorm:"column:holidaysused;not null" json:"holidaysused"`
	StartedAt    time.Time `gorm:"column:started_at;not null" json:"started_at"`
	CreatedAt    time.Time `gorm:"column:created_at;not null" json:"created_at"`
	Isactive     bool      `gorm:"column:isactive;not null" json:"isactive"`
	Nationality  int32     `gorm:"column:nationality;not null" json:"nationality"`
	Ahvnr        int32     `gorm:"column:ahvnr;not null" json:"ahvnr"`
}

// TableName Employee's table name
func (*Employee) TableName() string {
	return TableNameEmployee
}
