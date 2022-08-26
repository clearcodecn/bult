package main

import (
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var (
		ba *bolt.Bucket
		bb *bolt.Bucket
	)

	db.Update(func(tx *bolt.Tx) error {
		ba, _ = tx.CreateBucket([]byte("a"))
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		bb, _ = tx.CreateBucket([]byte("a"))
		return nil
	})

	err = ba.Put([]byte("k1"), []byte("v1"))
	if err != nil {
		panic(err)
	}
	err = bb.Put([]byte("k1"), []byte("v1"))
	if err != nil {
		panic(err)
	}
}
