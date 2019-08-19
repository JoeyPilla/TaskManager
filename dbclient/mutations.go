package dbclient

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

func (bc *BoltClient) AddTodo(task string, date time.Time) {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("TodosBucket"))
		id, err := b.NextSequence()
		check(err)
		todo := Todo{Id: int(id), Task: task}
		todoJSON, err := json.Marshal(todo)
		check(err)
		err = b.Put([]byte(date.Format(time.RFC3339)), todoJSON)
		fmt.Printf(`Added "%s" to your task list`+"\n", task)
		return err
	})
}

func (bc *BoltClient) RemoveTodo(item string) {
	key, value, err := findItem(bc.boltDB, item)
	if err != nil {
		fmt.Println(err)
	} else {
		bc.boltDB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("TodosBucket"))
			err := b.Delete(key)
			check(err)
			todo := Todo{}
			json.Unmarshal(value, &todo)
			fmt.Printf(`You have completed the "%s" task.`+"\n", todo.Task)
			return nil
		})
	}
}

func findItem(db *bolt.DB, item string) ([]byte, []byte, error) {
	key := []byte{}
	value := []byte{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("TodosBucket"))
		counter := 1
		return b.ForEach(func(k, v []byte) error {
			if strconv.Itoa(counter) == item {
				key = k
				value = v
				return nil
			}
			return fmt.Errorf("Could not find: %s", item)
		})
	})
	return key, value, err
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
