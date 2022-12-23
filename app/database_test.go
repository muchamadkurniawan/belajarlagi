package app

import "testing"
import _ "github.com/go-sql-driver/mysql"

func TestNewDB(t *testing.T) {
	db := NewDB()
	db.Begin()
}
