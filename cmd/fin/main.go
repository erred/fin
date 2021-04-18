package main

import (
	"flag"
	"fmt"
	"os"

	"go.seankhliao.com/fin/run"
	"go.seankhliao.com/fin/run/importcsv"
	"go.seankhliao.com/fin/run/summary"
)

func main() {
	o := run.NewOptions(flag.CommandLine)
	flag.Parse()

	subCommands := map[string]run.Cmd{
		"importcsv": importcsv.Run,
		"summary":   summary.Run,
	}

	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "no command given")
		os.Exit(1)
	}

	cmd, ok := subCommands[flag.Arg(0)]
	if !ok {
		fmt.Fprintln(os.Stderr, "unknown command")
		os.Exit(1)
	}
	err := cmd(o, flag.Args())
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
