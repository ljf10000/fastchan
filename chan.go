package main

type chanBaseHeader struct {
	Version 	uint32
	Mac 		string
	Cmd 		string // ping/pong/request/response
}

type chanHeader struct {
	chanBaseHeader
	
	Seq 		uint64
	Ack 		uint64
}

type chanPing chanBaseHeader
type chanPong chanBaseHeader

type chanRequest struct {
	chanHeader
	
	Command 	string
	Mode 		string // syn/asyn/ack
	Timeout 	uint32 // ms
	Url			string
}

type chanResponse struct {
	chanHeader
	Std
}
