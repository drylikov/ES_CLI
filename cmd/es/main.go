
package main

import (
	"log"

	"github.com/segmentio/go-env"
	
)

var version = "0.0.1"

var usage = `
  Usage:
    es nodes memory
    es -h | --help
    es --version

  Options:
    -h, --help      Output help information
    -v, --version   Output program version
`

func main() {
	args, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	client := es.New(env.MustGet("ES_ADDR"))

	switch {
	case args["nodes"].(bool) && args["memory"].(bool):
		nodesMemory(client)
	}
}