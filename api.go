package main

type apiHeader struct {
	Version 	uint32
}

type apiRequestHeader struct {
	apiHeader
	
	Cmd		string // list/command
}

// list request
type apiListRequest struct {
	apiRequestHeader
	
	Select 		string // online/offline/all
	Mac 		string
	Mask 		string
}

type apiListEntry struct {
	Mac 		string
	OnlineTime	string
}

// list response
type apiListResponse struct {
	apiHeader
	
	Results []apiListEntry
}

// command request
type apiCommandRequest struct {
	apiRequestHeader
	
	Devs	[]string
	Command string
	Timeout uint32 // ms
}

type apiCommandEntry struct {
	Std
	
	Mac		string
}

// command reponse
type apiCommandResponse struct {
	apiHeader
	
	Results []apiCommandEntry
}

