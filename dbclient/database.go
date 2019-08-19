package dbclient

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDB() {
	var err error
	bc.boltDB, err = bolt.Open("test.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("TodosBucket"))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}
