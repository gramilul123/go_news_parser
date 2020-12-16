package db

import (
	"testing"
)

func TestDBConnection(t *testing.T) {
	dbConn := GetDB()

	if dbConn == nil {
		t.Error("DB Connection error")
	}
}
