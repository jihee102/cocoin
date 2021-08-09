package db

import (
	"github.com/boltdb/bolt"
	"github.com/jihee102/explorer/utils"
)

var db *bolt.DB

const (
	dbName       = "cocoin.db"
	dataBucket   = "data"
	blocksBucket = "blocks"
)

func DB() *bolt.DB {
	if db == nil {
		//init db
		dbPointer, err := bolt.Open(dbName, 6500, nil)
		db = dbPointer
		utils.HandleErr(err)
		err = db.Update(func(t *bolt.Tx) error {
			_, err := t.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_, err = t.CreateBucketIfNotExists([]byte(blocksBucket))
			return err
		})
		utils.HandleErr(err)

	}
	return db
}
