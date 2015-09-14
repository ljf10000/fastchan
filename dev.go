package main

import (
	. "asdf"
	"net"
)

const (
	cacheCount 		= 32
	
	devsCount		= 256
	devIdleTimeout  = 10*1000 // ms
)

type dev struct {
	mac 		[MacSize]byte
	seq 		uint64
	ack 		uint64
	
	peer 		net.Addr
	conn 		*net.PacketConn
	
	ping 		uint64
	pong 		uint64
	idle 		uint32
	
	cache 		map[uint64/* seq */]*apiCommandRequest
}

type devChan struct {
	
}

type devManager struct {
	devs 	map[[MacSize]byte]*dev
	ch 	 	chan devChan
	idle 	uint
	idx 	uint
}

func (me *devManager) init() {
	me.devs = make(map[[MacSize]byte]*dev)
	me.ch	= make(chan devChan, conf.DevsChanSize)
	me.idle = conf.DevIdleTimeout
}

// go it
func (me *devManager) run() {
	
}

var devManagers []devManager

func devManagerInit() {
	devManagers = make([]devManager, conf.DevsCount)
	
	for i:=uint(0); i<conf.DevsCount; i++ {
		dm := &devManagers[i]
		
		dm.idx = i
		dm.init()
		
		go dm.run()
	}
}
