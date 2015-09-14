package main

import (
	"fmt"
	. "strconv"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":" + Itoa(int(conf.HttpPort)), nil)
	if nil!=err {
		fmt.Println("ListenAndServe", err)
	}
}

func init() {
	loadConfig()
	devManagerInit()
	http.HandleFunc("/v1", fastChan)
}
