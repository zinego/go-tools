package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/skanehira/clipboard-image/v2"
	"github.com/zinego/go-tools/utils/log"
)

func saveToRespository(dir string) (fname string) {
	imgd, err := clipboard.Read()
	if err != nil {
		log.Errorf("read clipboard image failed: %v", err)
		panic(err)
	}
	body, err := ioutil.ReadAll(imgd)
	if err != nil {
		log.Errorf("read all clipboard image failed: %v", err)
		panic(err)
	}
	fname = fmt.Sprintf("%s/%s.png", dir, time.Now().Format("2006_01_02T15_04_05"))
	f, err := os.Create(fname)
	if err != nil {
		log.Errorf("create image failed: %v", err)
		panic(err)
	}
	_, err = f.Write(body)
	if err != nil {
		log.Errorf("write image failed: %v", err)
		panic(err)
	}
	return fname
}
