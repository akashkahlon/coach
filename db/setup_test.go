package db

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)


func TestInitDB_Success(t *testing.T) {
	db, err := InitDB()

	assert.NoError(t, err)
	assert.NotNil(t, db)
}