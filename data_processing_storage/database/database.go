package database

import (
	"errors"
)

type InMemoryDB struct {
	mainState        map[string]int
	transactionState map[string]int
	inTransaction    bool
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		mainState:        make(map[string]int),
		transactionState: nil,
		inTransaction:    false,
	}
}

// Start transaction
func (db *InMemoryDB) BeginTransaction() error {
	if db.inTransaction {
		return errors.New("ERROR: A transaction is already in progress")
	}
	db.inTransaction = true
	db.transactionState = make(map[string]int)
	return nil
}

// Add key to transaction state... MUST BE IN A TRANSACTION
func (db *InMemoryDB) Put(key string, value int) error {
	if !db.inTransaction {
		return errors.New("ERROR: put operation not allowed outside of a transaction")
	}
	db.transactionState[key] = value
	return nil
}

// Gets the value of key stored to database. Does not return value if key is in transaction state
func (db *InMemoryDB) Get(key string) *int {
	if db.inTransaction {
		if _, exists := db.transactionState[key]; exists {
			return nil
		}
	}
	if value, exists := db.mainState[key]; exists {
		return &value
	}
	return nil
}

// Commits all changes in transaction state to main database and ends the transaction
func (db *InMemoryDB) Commit() error {
	if !db.inTransaction {
		return errors.New("ERROR: no transaction in progress")
	}
	for key, value := range db.transactionState {
		db.mainState[key] = value
	}
	db.inTransaction = false
	db.transactionState = nil
	return nil
}

// Rollbacks all changes within transaction state and ends the transaction
func (db *InMemoryDB) Rollback() error {
	if !db.inTransaction {
		return errors.New("ERROR: no transaction in progress")
	}
	db.inTransaction = false
	db.transactionState = nil
	return nil
}
