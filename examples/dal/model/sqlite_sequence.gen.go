// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSqliteSequence = "sqlite_sequence"

// SqliteSequence mapped from table <sqlite_sequence>
type SqliteSequence struct {
	Name string `gorm:"column:name;type:" json:"name"`
	Seq  string `gorm:"column:seq;type:" json:"seq"`
}

// TableName SqliteSequence's table name
func (*SqliteSequence) TableName() string {
	return TableNameSqliteSequence
}