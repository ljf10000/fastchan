package main

/*

1. udp chan protocol
1.1 ping c==>s
{
	"version":1,
	"mac":"MAC",
	"cmd":"ping"
}

1.2 pong s==>c
{
	"version":1,
	"mac":"MAC",
	"cmd":"pong"
}

1.3 request s==>c
{
	"version":1,
	"mac":"MAC",
	"seq":SEQ,
	"ack":ACK,
	"cmd":"request",
	"command":"COMMAND",
	"mode":"syn/asyn/ack",
	"timeout":TIMEOUT(ms),
	"url":"URL"
}

1.4 response s==>c
{
	"version":1,
	"mac":"MAC",
	"seq":SEQ,
	"ack":ACK,
	"cmd":"response",
	"stdout":"STDOUT",
	"stderr":"STDERR",
	"errno":ERRNO
}

2. tcp service protocol
   http service protocol(http://lms4.autelan.com/v1)
2.1 list
2.1.1 request
{
	"version":1,
	"cmd":"list",
	"select":"online/offline/all",
	"mac":"MAC",
	"mask":"MASK",
}
2.1.2 response
{
	"version":1,
	"results": [
		{
			"mac":"MAC",
			"onlineTime":"TIME",
		},
		{
			"mac":"MAC",
			"onlineTime":"TIME",
		}
	]
}

2.2 command
2.2.1 request
{
	"version":1,
	"cmd":"command",
	"devs":[
		"mac1",
		"macX"
	],
	"command":"COMMAND",
	"timeout":TIMEOUT,
}

2.2.2 response
{
	"version":1,
	"results":[
		{
			"mac":"MAC",
			"stdout":"STDOUT",
			"stderr":"STDERR",
			"errno":ERRNO	
		}
	]
}

*/
