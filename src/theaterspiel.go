package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/abiosoft/ishell/v2"
	"moul.io/banner"
)

const VERSION = "Version 0.2 Alfa"

//https://go.dev/play/p/Q-Ufgrw3vZL
func DDHHMMZ() string {
	current_time := time.Now().UTC()
	return fmt.Sprintf("%d%02d%02dZ\n", current_time.Day(), current_time.Hour(), current_time.Minute())
}

func DDHHMMZmmmYY() string {
	current_time := time.Now().UTC()
	return fmt.Sprintf(current_time.Format("021504ZJan06\n"))
}

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

	shell.AddCmd(&ishell.Cmd{
		Name: "datetime",
		Help: "Zulu datetime",
		Func: func(c *ishell.Context) {
			c.Println("Zulu time:", DDHHMMZ())
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "ldatetime",
		Help: "Zulu long datetime",
		Func: func(c *ishell.Context) {
			c.Println("Zulu time:", DDHHMMZmmmYY())
		},
	})

	// run shell
	shell.Run()
}
