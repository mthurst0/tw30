package main

import (
	"github.com/mthurst0/tw30/internal/appenv"

	"github.com/mthurst0/tw30/cmd"
)

var (
	BuildTime  string
	CommitHash string
	GoVersion  string
	GitTag     string
)

func main() {
	appenv.SetVersionInfo(appenv.VersionInfo{
		Version:   GitTag,
		Commit:    CommitHash,
		BuildTime: BuildTime,
		GoVersion: GoVersion,
	})
	cmd.Execute()
}
