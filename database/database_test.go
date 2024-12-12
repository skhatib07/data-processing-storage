package database_test

import (
	"testing"

	"github.com/skhatib07/data-processing-storage/database"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryDB(t *testing.T) {
	inmemoryDB := database.NewInMemoryDB()

	t.Run("Get non-existent key", func(t *testing.T) {
		// should return null, because A doesn’t exist in the DB yet
		val := inmemoryDB.Get("A")
		assert.Nil(t, val, "Expected nil for key 'A'")
	})

	t.Run("Put outside transaction", func(t *testing.T) {
		// should throw an error because a transaction is not in progress
		err := inmemoryDB.Put("A", 5)
		assert.EqualError(t, err, "ERROR: put operation not allowed outside of a transaction", "Expected error when Put is called outside of a transaction")
	})

	t.Run("Begin Transaction and Put", func(t *testing.T) {
		// starts a new transaction
		inmemoryDB.BeginTransaction()

		// set’s value of A to 5, but its not committed yet
		err := inmemoryDB.Put("A", 5)
		assert.NoError(t, err, "Unexpected error when calling Put within transaction")

		// should return null, because updates to A are not committed yet
		val := inmemoryDB.Get("A")
		assert.Nil(t, val, "Expected nil for key 'A' before commit")

		// update A’s value to 6 within the transaction
		err = inmemoryDB.Put("A", 6)
		assert.NoError(t, err, "Unexpected error when updating 'A' within transaction")

		// commits the open transaction
		err = inmemoryDB.Commit()
		assert.NoError(t, err, "Unexpected error during commit")

		// should return 6, that was the last value of A to be committed
		val = inmemoryDB.Get("A")
		assert.NotNil(t, val, "Expected non-nil value for key 'A' after commit")
		assert.Equal(t, 6, *val, "Expected value 6 for key 'A' after commit")
	})

	t.Run("Commit without transaction", func(t *testing.T) {
		// throws an error, because there is no open transaction
		err := inmemoryDB.Commit()
		assert.EqualError(t, err, "ERROR: no transaction in progress", "Expected error for Commit without transaction")
	})

	t.Run("Rollback without transaction", func(t *testing.T) {
		// throws an error because there is no ongoing transaction
		err := inmemoryDB.Rollback()
		assert.EqualError(t, err, "ERROR: no transaction in progress", "Expected error for Rollback without transaction")
	})

	t.Run("Get non-existent key B", func(t *testing.T) {
		// should return null because B does not exist in the database
		val := inmemoryDB.Get("B")
		assert.Nil(t, val, "Expected nil for key 'B'")
	})

	t.Run("Transaction with Rollback", func(t *testing.T) {
		// starts a new transaction
		inmemoryDB.BeginTransaction()

		// Set key B’s value to 10 within the transaction
		err := inmemoryDB.Put("B", 10)
		assert.NoError(t, err, "Unexpected error when calling Put within transaction")

		// Rollback the transaction - revert any changes made to B
		err = inmemoryDB.Rollback()
		assert.NoError(t, err, "Unexpected error during Rollback")

		// Should return null because changes to B were rolled back
		val := inmemoryDB.Get("B")
		assert.Nil(t, val, "Expected nil for key 'B' after rollback")
	})
}
