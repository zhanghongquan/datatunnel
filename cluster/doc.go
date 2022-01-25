//go:generate protoc --proto_path=../protos/ --go_out=./ --go_opt=paths=source_relative ../protos/gossip.proto
package cluster
