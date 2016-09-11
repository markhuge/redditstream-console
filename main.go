package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/markhuge/redditstream-console/api"
)

const VERSION = "1.0.0"
const USAGE = `
redditstream thing v` + VERSION + ` - [https://github.com/markhuge/redditstream-console]

USAGE: redditstream REDDIT_POST_URL

  -h, --help Print this shit

Ctrl + c to exit
`

var help bool

func init() {
	flag.BoolVar(&help, "help", false, "halp")
	flag.BoolVar(&help, "h", false, "halp")
	flag.Parse()
}

func main() {
	// arrrrrrrrgs
	args := os.Args
	if len(args) <= 1 || help {
		fmt.Println(USAGE)
		return
	}

	uri := args[1]

	p, err := api.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	api.Print(p)
}
