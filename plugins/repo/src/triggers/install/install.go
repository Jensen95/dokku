package main

import (
	"fmt"

	"github.com/dokku/dokku/plugins/common"
)

// runs the install step for the repo plugin
func main() {
	if err := common.PropertySetup("repo"); err != nil {
		common.LogFail(fmt.Sprintf("Unable to install the repo plugin: %s", err.Error()))
	}
}
