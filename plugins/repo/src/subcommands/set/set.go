package main

import (
	"flag"

	"github.com/dokku/dokku/plugins/common"
	"github.com/dokku/dokku/plugins/repo"
)

// set or clear a repo property for an app
func main() {
	flag.Parse()
	appName := flag.Arg(1)
	property := flag.Arg(2)
	value := flag.Arg(3)

	common.CommandPropertySet("repo", appName, property, value, repo.DefaultProperties)
}
