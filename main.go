package main

import (
	"github.com/MaximChernomorov/challenge-test/cmd"
	_ "net/http/pprof"
	"runtime"
)

func main() {
	runtime.SetCPUProfileRate(100)

	cmd.Execute()
}
