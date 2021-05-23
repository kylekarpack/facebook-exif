package main

import (
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
			dir := args["dir"].Value
			run(dir, false)
		})

	// configure info command
	commando.
		Register("info").
		SetShortDescription("perform a 'dry run' and return info without modifying files").
		SetDescription("This command displays more information about the contents of the directory like number of photos and metadata found").
		AddArgument("dir", "directory to photos", "./photos_and_videos"). // default `./photos_and_videos`
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			dir := args["dir"].Value
			run(dir, true)
		})

	// parse command-line arguments
	commando.Parse(nil)

}
