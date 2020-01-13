package deckdb

import (
	"log"
	"time"
	"fmt"

	"github.com/boltdb/bolt"
)

/*
func OpenDb(file string) {
   db, err := bolt.Open(file, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


   db.Update(func(tx *bolt.Tx) error {
      root, err := tx.CreateBucketIfNotExists([]byte("DB"))
      if err != nil {
         return fmt.Errorf("could not create root bucket: %v", err)
      }

      _, err = root.CreateBucketIfNotExists([]byte("CARDS"))
      if err != nil {
         return fmt.Errorf("could not create root bucket: %v", err)
      }

      return nil
   })
}
*/

func OpenDb(file string) {
   db, err := bolt.Open(file, 0600, &bolt.Options{Timeout: 1 * time.Second})
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    key := []byte("hello")
    value := []byte("Hello World!")

    // store some data
    err = db.Update(func(tx *bolt.Tx) error {
        bucket, err := tx.CreateBucketIfNotExists([]byte("world"))
        if err != nil {
            return err
        }

        err = bucket.Put(key, value)
        if err != nil {
            return err
        }
        return nil
    })

    if err != nil {
        log.Fatal(err)
    }

    // retrieve the data
    err = db.View(func(tx *bolt.Tx) error {
        bucket := tx.Bucket([]byte("world"))
        if bucket == nil {
            return fmt.Errorf("Bucket %q not found!", []byte("world"))
        }

        val := bucket.Get(key)
        fmt.Println(string(val))

        return nil
    })

    if err != nil {
        log.Fatal(err)
    }
 }
