package shortdb

import (
	"log"
	"os"

	"go.etcd.io/bbolt"
)

const storePath string = "ls.db"
const dbFileMode os.FileMode = 0600

var bucketName []byte = []byte("links")

type ShortStore struct {
	conn *bbolt.DB
	path string
}

/*
type Options struct {
	BoltOptions *bbolt.Options
}
*/

/* Notes:
DB.View() and DB.Update() are wrappers for BD.Begin and handle starting the transaction, executing a function and safely
closes the transaction

Only a single read/write operation can be run at a time
There can be as many reads as the server is capable of
*/

// DBSetup initializes the database and buckets
// The BD file is created within the same directory as main.go
// TODO: take another look at this function and possibly refactor
func DBInit() (*ShortStore, error) {

	db, err := bbolt.Open(storePath, dbFileMode, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	store := &ShortStore{db, storePath}
	// not sure about this block
	if !store.conn.IsReadOnly() {
		if err := store.bucketInit(); err != nil {
			store.conn.Close()
		}
	}

	return store, nil
}

// Init Bucket
func (s *ShortStore) bucketInit() error {
	tx, err := s.conn.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.CreateBucketIfNotExists(bucketName); err != nil {
		return err
	}

	return tx.Commit()
}

// Add a key/value pair to the bucket - read/write operation
func (s *ShortStore) Put(sLink []byte, fullLink []byte) error {
	tx, err := s.conn.Begin(true)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	bucket := tx.Bucket(bucketName)
	if err := bucket.Put(sLink, fullLink); err != nil {
		return err
	}
	return tx.Commit()
}

// Remove a key/value pair from the bucket - read/write op
func (s *ShortStore) Delete(sLink []byte) ([]byte, error) {
	value := []byte("")
	err := s.conn.View(func(tx *bbolt.Tx) error {
		tx.Bucket(bucketName).Delete(sLink)
		return nil
	})
	return value, err
}

// Get takes in a short link and returns the URL stored as a value
func (s *ShortStore) Get(sLink []byte) (string, error) {
	tx, err := s.conn.Begin(false) // Begin takes a bool to mark the connection as writable. False is non-writable transaction
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	val := tx.Bucket(bucketName).Get(sLink)

	return string(val), nil

}

// Check if the DB already exists in the dir - could later be used in the DBinit() func to avoid extra work
func (s *ShortStore) DBexists() bool {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if storePath == file.Name() {
			return true
		}
	}
	return false
}

// Function to check if a key is in the bucket
func (s *ShortStore) DBKeyExists(sLink []byte) bool {
	tx, err := s.conn.Begin(false)
	if err != nil {
		log.Fatal(err)
	}

	k := tx.Bucket(bucketName).Get(sLink)
	if k != nil {
		return k != nil
	}
	return false
}
