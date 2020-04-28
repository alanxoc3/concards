package deckdb

import (
	"log"
	"time"
	"fmt"
   "encoding/json"
	"github.com/boltdb/bolt"
	"github.com/alanxoc3/concards/card"
)

func dbInit(db *bolt.DB) (err error) {
   err = db.Update(func(tx *bolt.Tx) error {
      b, err := tx.CreateBucketIfNotExists([]byte("cards"))
      if err != nil {
         return err
      }

      if c, err := card.New(
         map[string]bool{},
         []string{"THIS_IS_A_QUESTION"},
         [][]string{},
         [][]string{},
         []string{}); err != nil {
         return err
      } else {
         insertCard(b, c)
      }

      return nil
   })

   return
}

func insertGroups(cardBucket *bolt.Bucket, groups *card.Groups) {
   // Reset the groups to be new values.
   if err := cardBucket.DeleteBucket([]byte("@>")); err != nil {
      // The bucket either does not exist, or is not a bucket.
   }

   gb, err := cardBucket.CreateBucket([]byte("@>"))
   if err != nil {
      // The key already exists.
      panic(err)
   }



   err := cardBucket.DeleteBucket("@>")
   cardBucket.DeleteBucket("@>")


   cardBucket.Put([]byte("@q")
   sum := c.Hash()
   cardBucket, err := b.CreateBucketIfNotExists(sum[:])
   if err != nil {
      panic(err)
   }

   cardBucket.Put([]byte("@q"), []byte(c.Question))
   answers, _ := json.Marshal(c.Answers)
   cardBucket.Put([]byte("@a"), answers)
}

func insertCard(b *bolt.Bucket, c *card.Card) {
   sum := c.Hash()
   cardBucket, err := b.CreateBucketIfNotExists(sum[:])
   if err != nil {
      panic(err)
   }

   cardBucket.Put([]byte("@q"), []byte(c.Question))
   answers, _ := json.Marshal(c.Answers)
   cardBucket.Put([]byte("@a"), answers)
}

func printBucket(b *bolt.Bucket) {
   c := b.Cursor()

   for k, v := c.First(); k != nil; k, v = c.Next() {
      nestedBucket := b.Bucket(k)
      if nestedBucket != nil {
         printBucket(nestedBucket)
      } else {
         fmt.Printf("key: \"%s\", value \"%s\".\n", k, v)
      }
   }
}

func OpenDb(file string) {
   db, err := bolt.Open(file, 0600, &bolt.Options{Timeout: 1 * time.Second})
   if err != nil {
      panic(err)
   }
   defer db.Close()

   err = dbInit(db)
   if err != nil {
      panic(err)
   }

   // retrieve the data
   err = db.View(func(tx *bolt.Tx) error {
      bucket := tx.Bucket([]byte("cards"))
      if bucket == nil {
         return fmt.Errorf("Bucket %q not found!", []byte("world"))
      }

      printBucket(bucket)

      return nil
   })

   if err != nil {
      println("heasoentuh")
      log.Fatal(err)
   }
}
