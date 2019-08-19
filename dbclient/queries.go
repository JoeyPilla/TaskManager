package dbclient

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
)

// func ListTodos() error {
// 	db, err := openDB()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("You have the following tasks:")
// 	err = db.View(func(tx *bolt.Tx) error {

// 		count := 1
// 		b.ForEach(func(k, v []byte) error {
// 			t := Entry{}
// 			json.Unmarshal(v, &t)
// 			fmt.Printf("%d. %s\n", count, t.Task)
// 			count++
// 			return nil
// 		})
// 		return nil
// 	})
// 	return err
// }

func (bc *BoltClient) ListTodos() {
	bc.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("TodosBucket"))
		count := 1
		return b.ForEach(func(k, v []byte) error {
			t := Todo{}
			json.Unmarshal(v, &t)
			fmt.Printf("%d. %s\n", count, t.Task)
			count++
			return nil
		})
	})
}
