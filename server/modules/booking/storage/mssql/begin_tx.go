package mssql

import "gorm.io/gorm"

func (s *sqlStorage) Begin() *gorm.DB {
	return s.db.Begin() // Start a new transaction
}
