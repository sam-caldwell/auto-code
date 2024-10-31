package main

import (
	"github.com/sam-caldwell/arguments"
	"log"
)

const ipAddressRegex = `
\b(
    (25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)\.            # Matches IPv4 sections (0-255)
    ((25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)\.){2}       # Repeats for the first three sections
    (25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)              # Final IPv4 segment
\b|
    (
        ([A-Fa-f0-9]{1,4}:){7}[A-Fa-f0-9]{1,4}      # Matches full IPv6 address
      | ([A-Fa-f0-9]{1,4}:){1,7}:                   # Compressed IPv6 forms
      | ([A-Fa-f0-9]{1,4}:){1,6}:[A-Fa-f0-9]{1,4}   # Mixed IPv6 forms
    )
\b
`

func main() {

	config, err := arguments.NewParser().
		ProgramName("Example Program").
		ProgramDescription("A simple demonstration").
		Copyright("(c) 2024 Sam Caldwell.  MIT License").
		//A parameter is defined as a record name and its associated value as defined by some source...
		Parameter(
			"network.address",                 // Define the identifier used in the parser to represent the value.
			"this is the listener Ip address", // define help text used in error messages"
			arguments.Type("0.0.0.0", arguments.Pattern(ipAddressRegex)), // Define the parameter's type
			arguments.YamlFile("config.yaml", "net.ip"),                  // Define the data source(s) in
			arguments.Environment("LISTEN_IP"),                           // order of reverse precedence
			arguments.CliOption("-n"),
			arguments.CliOption("--ip"),
		).
		Parameter(
			"network.port",
			"this is the listener ip port",
			arguments.Type(uint16(8080), arguments.MinMaxInt(8000, 8888)),
			arguments.YamlFile("config.yaml", "net.port"),
			arguments.Environment("LISTEN_PORT"),
			arguments.CliOption("-p"),
			arguments.CliOption("--port"),
		).
		Parameter(
			"pi",
			"this defines pi...how irrational",
			arguments.Type(3.14, arguments.MinMaxFloat(3, 4)),
			arguments.CliOption("--pi"),
		).
		Parse() // return a resolved configuration and any error state

	if err != nil {
		log.Fatal(err)
	}
	config.Dump()
}
