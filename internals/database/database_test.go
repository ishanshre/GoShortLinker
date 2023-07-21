package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConnection(t *testing.T) {
	dsn := os.Getenv("dsn")
	db, err := NewConnection(dsn)
	assert.NoError(t, err)
	assert.NotNil(t, db)
	collection := db.GetUrlCollections()
	assert.NotNil(t, collection)
	name := collection.Name()
	assert.Equal(t, name, "urls")
}

func TestNewConnection_Failure(t *testing.T) {
	dsn := "error-dsn"
	db, err := NewConnection(dsn)
	assert.Error(t, err)
	assert.Nil(t, db)
}
