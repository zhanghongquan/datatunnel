package storage

/*
引入storage的主要目的是为了尽量保证数据不丢失，即便client因为某些原因暂时丢失了和broker之间的连接，那么在连接丢失的时候的即便是QOS为0的消息也不会丢失。
在存储消息的时候， 每一条消息都会按照 timestamp / topic 的方式
*/

import "github.com/zhanghongquan/gorocksdb"

type DataStorage struct {
	db *gorocksdb.DB
}
