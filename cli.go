package main

import (
	"fmt"

	"github.com/thatisuday/commando"
)

func main() {

	// configure commando
	commando.
		SetExecutableName("fix-fb-meta").
		SetVersion("1.0.0").
		SetDescription("This tool combines JSON metadata with downloaded photos from Facebook to restore information that was stripped during upload")

	// configure the root command
	commando.
		Register(nil).
		AddArgument("dir", "directory to photos", "./photos_and_videos"). // default `./photos_and_videos`
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Printf("Printing options of the `root` command...\n\n")

			// print arguments
			for k, v := range args {
				fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			}

			// print flags
			for k, v := range flags {
				fmt.Printf("flag -> %v: %v(%T)\n", k, v.Value, v.Value)
			}
		})

	// configure info command
	commando.
		Register("info").
		SetShortDescription("displays detailed information of a directory").
		SetDescription("This command displays more information about the contents of the directory like size, permission and ownership, etc.").
		AddArgument("dir", "directory to photos", "./photos_and_videos"). // default `./photos_and_videos`
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Printf("Printing options of the `info` command...\n\n")

			// print arguments
			for k, v := range args {
				fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			}

			// print flags
			for k, v := range flags {
				fmt.Printf("flag -> %v: %v(%T)\n", k, v.Value, v.Value)
			}
		})

	// parse command-line arguments
	commando.Parse(nil)

}
