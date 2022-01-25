package storage

import "github.com/zhanghongquan/gorocksdb"

type DataStorage struct {
	db *gorocksdb.DB
}
