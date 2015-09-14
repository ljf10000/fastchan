package main

import (
	. "asdf"
	"fmt"
	"bytes"
	"os"
	"encoding/json"
)

const fileConfig = "config.json"

type config struct {
	ChanPort		uint
	HttpPort		uint
	
	DevIdleTimeout 	uint
	
	DevsCount		uint
	DevsChanSize	uint
}

var conf = &config{
	ChanPort:9527,
	HttpPort:7259,
	
	DevIdleTimeout:10*1000,
	
	DevsCount:256,
	DevsChanSize:128,
}

func (me *config) Load(file string) {
	f, err := os.Open(file)
	if nil != err {
		fmt.Println("no found", file)

		return // no config file, just return
	}
	defer f.Close()

	b := new(bytes.Buffer)
	_, err = b.ReadFrom(f)
	if nil != err {
		panic(err)
	}

	Log.Info("load default:%+v" + Crlf, me)
	json.Unmarshal(b.Bytes(), me)
	Log.Info("load config(%s):%+v" + Crlf, file, me)
}

func loadConfig() {
	conf.Load(fileConfig)
}