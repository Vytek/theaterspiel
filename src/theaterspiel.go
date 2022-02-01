package main

import (
	"strings"

	"github.com/abiosoft/ishell/v2"
	"moul.io/banner"
)

const VERSION = "Version 0.1 Alfa"

func main() {
	// create new shell.
	// by default, new shell includes 'exit', 'help' and 'clear' commands.
	shell := ishell.New()

	// display welcome info.
	shell.Println(banner.Inline("theaterspiel"))
	shell.Println(VERSION)

	// register a function for "greet" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "greet",
		Help: "greet user",
		Func: func(c *ishell.Context) {
			c.Println("Hello", strings.Join(c.Args, " "))
		},
	})

	// run shell
	shell.Run()
}
