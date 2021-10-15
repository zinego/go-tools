package main

import (
	"github.com/zinego/go-tools/utils/log"
)

func init() {
	log.Init()
}

func main() {
	dir := "../image/img"
	fname := saveToRespository(dir)
	upload(dir, fname)
}
