package main

import (
	"flag"
	"os"
	"strings"

	"github.com/dokku/dokku/plugins/common"
	"github.com/otiai10/copy"
)

func copyDirectory(sourceBasePath string, sourceFolder string, destinationPath string) {
	if sourceFolder != "" {
		sourcePath := strings.Join([]string{sourceBasePath, sourceFolder}, "/")
		if stat, err := os.Stat(sourcePath); err == nil && stat.IsDir() {
			copy.Copy(sourcePath, destinationPath)
		}
	}
}

// copies files into place for a given repo
func main() {
	flag.Parse()
	appName := flag.Arg(0)
	tmpWorkDir := flag.Arg(1)

	appRoot := strings.Join([]string{common.MustGetEnv("DOKKU_ROOT"), appName}, "/")

	hostCopyFolder := common.PropertyGet("repo", appName, "host-copy-folder")
	containerCopyFolder := common.PropertyGet("repo", appName, "container-copy-folder")

	copyDirectory(tmpWorkDir, hostCopyFolder, appRoot)
	copyDirectory(tmpWorkDir, containerCopyFolder, tmpWorkDir)
}
