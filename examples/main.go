package main

import (
	logger "log"

	"github.com/vishal1132/log"
)

func main() {
	l := log.New()
	off, err := l.Write([]byte("abcd"))
	if err != nil {
		logger.Fatal(err)
	}
	v, err := l.Read(off)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println(string(v))
}
