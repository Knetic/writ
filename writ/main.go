package main

import (
	"fmt"
	"os"
	"flag"
	"writ"
)

type RunSettings struct {
	ServePath string
}

func main() {

	var settings RunSettings
	var err error

	settings, err = parseRunSettings()
	if(err != nil) {
		fatal(err, 1)
	}

	server := writ.NewServer()
	server.Listen(settings.ServePath)
}

func parseRunSettings() (RunSettings, error) {

	var ret RunSettings

	flag.StringVar(&ret.ServePath, "p", ":8080", "ip/port to serve http on")
	flag.Parse()
	return ret, nil
}

func fatal(fault error, code int) {

	fmt.Fprintf(os.Stderr, fault.Error())
	os.Exit(code)
}
