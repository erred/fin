//go:build generate
// +build generate

package main

//go:generate protoc --go_out=. --go_opt=paths=source_relative proto/fin/fin.proto
