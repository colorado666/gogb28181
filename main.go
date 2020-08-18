package main

import (
	"github.com/abiosoft/ishell"
	"gogb28181/transport"
	"gogb28181/tu"
)

func main() {
	// create new shell.
	// by default, new shell includes 'exit', 'help' and 'clear' commands.
	shell := ishell.New()

	// display welcome info.
	shell.Println("gb28181 test")

	// simulate an authentication
	shell.AddCmd(&ishell.Cmd{
		Name: "tcps",
		Help: "tcp server",
		Func: func(c *ishell.Context) {
			transport.RunServerTCP()
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "tcpc",
		Help: "tcp client",
		Func: func(c *ishell.Context) {
			transport.RunClientTCP()
		},
	})
	// simulate an authentication
	shell.AddCmd(&ishell.Cmd{
		Name: "udps",
		Help: "udp server",
		Func: func(c *ishell.Context) {
			transport.RunServerUDP()
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "udpc",
		Help: "udp client",
		Func: func(c *ishell.Context) {
			transport.RunClientUDP()
		},
	})

	//=======================================
	shell.AddCmd(&ishell.Cmd{
		Name: "uas",
		Help: "sip user agent server",
		Func: func(c *ishell.Context) {
			tu.RunServer()
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "uac",
		Help: "sip user agent client",
		Func: func(c *ishell.Context) {
			tu.RunClient()
		},
	})

	//=======================================
	// run shell
	shell.Run()
}
