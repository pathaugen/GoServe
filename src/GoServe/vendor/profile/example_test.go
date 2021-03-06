package profile_test

import (
	"flag"
	"os"
	
	// Switch to Vendoring for testing (requires vendor/)
	//"github.com/pkg/profile"
	"profile"
)

func ExampleStart() {
	// start a simple CPU profile and register
	// a defer to Stop (flush) the profiling data.
	defer profile.Start().Stop()
}

func ExampleCPUProfile() {
	// CPU profiling is the default profiling mode, but you can specify it
	// explicitly for completeness.
	defer profile.Start(profile.CPUProfile).Stop()
}

func ExampleMemProfile() {
	// use memory profiling, rather than the default cpu profiling.
	defer profile.Start(profile.MemProfile).Stop()
}

func ExampleMemProfileRate() {
	// use memory profiling with custom rate.
	defer profile.Start(profile.MemProfileRate(2048)).Stop()
}

func ExampleProfilePath() {
	// set the location that the profile will be written to
	defer profile.Start(profile.ProfilePath(os.Getenv("HOME")))
}

func ExampleNoShutdownHook() {
	// disable the automatic shutdown hook.
	defer profile.Start(profile.NoShutdownHook).Stop()
}

func ExampleStart_withFlags() {
	// use the flags package to selectively enable profiling.
	mode := flag.String("profile.mode", "", "enable profiling mode, one of [cpu, mem, block]")
	flag.Parse()
	switch *mode {
	case "cpu":
		defer profile.Start(profile.CPUProfile).Stop()
	case "mem":
		defer profile.Start(profile.MemProfile).Stop()
	case "block":
		defer profile.Start(profile.BlockProfile).Stop()
	default:
		// do nothing
	}
}
