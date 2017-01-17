package main

import (
	"fmt"
	"os"
	
	// Switch to Vendoring
	//"GoServe/external/cli-master"
	"cli-master"
	//"github.com/miketheprogrammer/go-thrust/lib/spawn"
	
	// Switch to Vendoring
	//"GoServe/external/go-thrust-master/lib/spawn"
	"go-thrust-master/lib/spawn"
)

func install(c *cli.Context) {
	spawn.SetBaseDirectory("") // Default to usr.homedir.
	tp := spawn.NewThrustProvisioner()
	if err := tp.Provision(); err != nil {
		panic(err)
	}
	fmt.Println("Thrust installed")
}

func main() {
	app := cli.NewApp()
	app.Name = "go-thrust"
	app.Usage = "Tools to developp Thrust applications in Go"

	app.Commands = []cli.Command{
		{
			Name:      "install",
			ShortName: "i",
			Usage:     "Install Thrust",
			Action: func(c *cli.Context) {
				install(c)
			},
		},
	}

	app.Run(os.Args)
}
